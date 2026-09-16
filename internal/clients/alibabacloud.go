/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/crossplane-contrib/provider-alibabacloud/internal/version"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/crossplane-contrib/provider-alibabacloud/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig      = "no providerConfigRef provided"
	errGetProviderConfig     = "cannot get referenced ProviderConfig"
	errTrackUsage            = "cannot track ProviderConfig usage"
	errExtractCredentials    = "cannot extract credentials"
	errInvalidProviderConfig = "invalid ProviderConfig"
	errUnmarshalCredentials  = "cannot unmarshal alicloud credentials as JSON"
)

var providerCredentialKeys = []string{
	"access_key",
	"secret_key",
	"security_token",
}

var providerSpecCredentialKeys = map[string]string{
	"assume_role":           "spec.assumeRole",
	"assume_role_with_oidc": "spec.assumeRoleWithOIDC",
}

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(tfProvider *schema.Provider) terraform.SetupFn {
	return func(ctx context.Context, c client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(c, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		pc, creds, err := extractAndUnmarshalCredentials(ctx, c, configRef)
		if err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}
		if validateErr := validateProviderConfig(pc, creds); validateErr != nil {
			return ps, errors.Wrap(validateErr, errInvalidProviderConfig)
		}

		region, err := getRegion(mg, creds)
		if err != nil {
			return ps, errors.Wrap(err, "cannot get region")
		}

		ps.Configuration = buildProviderConfiguration(region, creds, pc)
		ps.Configuration["configuration_source"] = getUserAgent()
		return ps, errors.Wrap(configureNoForkAlibabaCloudClient(ctx, &ps, *tfProvider), "failed to configure the no-fork AlibabaCloud client")
	}
}

// configureNoForkAlibabaCloudClient configures the Terraform provider with the
// credentials resolved from the ProviderConfig and hands the resulting provider
// meta to upjet, which passes it to the resource CRUD functions.
//
// Please be aware that this implementation relies on the schema.Provider
// parameter p being a non-pointer. This is because the Terraform plugin SDK
// normally configures the provider only once, and using a pointer argument here
// would cause race conditions between resources referring to different
// ProviderConfigs.
func configureNoForkAlibabaCloudClient(ctx context.Context, ps *terraform.Setup, p schema.Provider) error {
	diag := p.Configure(context.WithoutCancel(ctx), &tfsdk.ResourceConfig{
		Config: ps.Configuration,
	})
	if diag != nil && diag.HasError() {
		return errors.Errorf("failed to configure the provider: %v", diag)
	}
	ps.Meta = p.Meta()
	return nil
}

func buildProviderConfiguration(region string, creds map[string]any, pc *v1beta1.ProviderConfig) terraform.ProviderConfiguration {
	cfg := terraform.ProviderConfiguration{
		"region": region,
	}
	for _, key := range providerCredentialKeys {
		v, ok := creds[key]
		if !ok {
			continue
		}
		cfg[key] = v
	}
	if pc == nil {
		return cfg
	}
	if pc.Spec.AssumeRole != nil {
		cfg["assume_role"] = []any{assumeRoleConfiguration(pc.Spec.AssumeRole)}
	}
	if pc.Spec.AssumeRoleWithOIDC != nil {
		cfg["assume_role_with_oidc"] = []any{assumeRoleWithOIDCConfiguration(pc.Spec.AssumeRoleWithOIDC)}
	}
	return cfg
}

func assumeRoleConfiguration(assumeRole *v1beta1.AssumeRoleOptions) map[string]any {
	cfg := map[string]any{
		"role_arn": assumeRole.RoleARN,
	}
	if assumeRole.SessionName != nil {
		cfg["session_name"] = *assumeRole.SessionName
	}
	if assumeRole.Policy != nil {
		cfg["policy"] = *assumeRole.Policy
	}
	if assumeRole.SessionExpiration != nil {
		cfg["session_expiration"] = *assumeRole.SessionExpiration
	}
	if assumeRole.ExternalID != nil {
		cfg["external_id"] = *assumeRole.ExternalID
	}
	return cfg
}

func assumeRoleWithOIDCConfiguration(assumeRole *v1beta1.AssumeRoleWithOIDCOptions) map[string]any {
	cfg := map[string]any{
		"role_arn":          assumeRole.RoleARN,
		"oidc_provider_arn": assumeRole.OIDCProviderARN,
		"oidc_token_file":   assumeRole.OIDCTokenFile,
	}
	if assumeRole.RoleSessionName != nil {
		cfg["role_session_name"] = *assumeRole.RoleSessionName
	}
	if assumeRole.Policy != nil {
		cfg["policy"] = *assumeRole.Policy
	}
	if assumeRole.SessionExpiration != nil {
		cfg["session_expiration"] = *assumeRole.SessionExpiration
	}
	return cfg
}

func validateProviderConfig(pc *v1beta1.ProviderConfig, creds map[string]any) error {
	for key, field := range providerSpecCredentialKeys {
		if _, ok := creds[key]; ok {
			return errors.Errorf("credentials JSON must not contain %q; use %s", key, field)
		}
	}
	if pc == nil {
		return nil
	}
	if ar := pc.Spec.AssumeRole; ar != nil {
		if err := validateAssumeRole(ar); err != nil {
			return err
		}
	}
	if ar := pc.Spec.AssumeRoleWithOIDC; ar != nil {
		if err := validateAssumeRoleWithOIDC(ar); err != nil {
			return err
		}
	}
	return nil
}

func validateAssumeRole(ar *v1beta1.AssumeRoleOptions) error {
	if strings.TrimSpace(ar.RoleARN) == "" {
		return errors.New("spec.assumeRole.roleARN is required")
	}
	return validateSessionExpiration("spec.assumeRole.sessionExpiration", ar.SessionExpiration, 3600)
}

func validateAssumeRoleWithOIDC(ar *v1beta1.AssumeRoleWithOIDCOptions) error {
	if strings.TrimSpace(ar.RoleARN) == "" {
		return errors.New("spec.assumeRoleWithOIDC.roleARN is required")
	}
	if strings.TrimSpace(ar.OIDCProviderARN) == "" {
		return errors.New("spec.assumeRoleWithOIDC.oidcProviderARN is required")
	}
	if strings.TrimSpace(ar.OIDCTokenFile) == "" {
		return errors.New("spec.assumeRoleWithOIDC.oidcTokenFile is required")
	}
	return validateSessionExpiration("spec.assumeRoleWithOIDC.sessionExpiration", ar.SessionExpiration, 43200)
}

func validateSessionExpiration(field string, v *int, max int) error {
	if v == nil {
		return nil
	}
	if *v < 900 || *v > max {
		return errors.Errorf("%s must be between 900 and %d", field, max)
	}
	return nil
}

func getRegion(obj runtime.Object, creds map[string]any) (string, error) {
	fromMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return "", errors.Wrap(err, "cannot convert to unstructured")
	}
	credsRegion := stringCredential(creds, "region")
	if credsRegion == "" {
		// region_id is used as a fallback for old version
		credsRegion = stringCredential(creds, "region_id")
	}
	r, err := fieldpath.Pave(fromMap).GetString("spec.forProvider.region")
	if fieldpath.IsNotFound(err) {
		// Region is not required for all resources, e.g. resource in "ram" group.
		return credsRegion, nil
	}
	return r, err
}

func stringCredential(creds map[string]any, key string) string {
	v, ok := creds[key].(string)
	if !ok {
		return ""
	}
	return v
}

func extractAndUnmarshalCredentials(ctx context.Context, c client.Client, configRef *v1.Reference) (*v1beta1.ProviderConfig, map[string]any, error) {
	pc := &v1beta1.ProviderConfig{}
	creds := map[string]any{}
	if err := c.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return pc, creds, errors.Wrap(err, errGetProviderConfig)
	}

	if pc.Spec.Credentials.Source == v1.CredentialsSourceInjectedIdentity ||
		pc.Spec.Credentials.Source == v1.CredentialsSourceNone {
		return pc, creds, nil
	}

	data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, c, pc.Spec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return pc, creds, errors.Wrap(err, errExtractCredentials)
	}
	if err = json.Unmarshal(data, &creds); err != nil {
		return pc, creds, errors.Wrap(err, errUnmarshalCredentials)
	}
	return pc, creds, nil
}
func getUserAgent() string {
	// user agent formats as "crossplane/<CROSSPLANE_VERSION> <PROJECT_NAME>/<PROJECT_VERSION>"
	return fmt.Sprintf("crossplane/%s provider-upjet-alibabacloud/%s", version.CrossplaneVersion, version.ProviderVersion)
}
