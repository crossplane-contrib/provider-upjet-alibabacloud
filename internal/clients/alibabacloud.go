/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	v1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/v2/pkg/fieldpath"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/crossplane-contrib/provider-alibabacloud/internal/version"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tfsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	clusterv1beta1 "github.com/crossplane-contrib/provider-alibabacloud/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/crossplane-contrib/provider-alibabacloud/apis/namespaced/v1beta1"
)

const (
	// error messages
	errNoProviderConfig      = "no providerConfigRef provided"
	errNotManaged            = "resource is neither a legacy nor a modern managed resource"
	errUnknownPCKind         = "referenced provider config kind is not supported"
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

		pcSpec, creds, err := resolveProviderConfig(ctx, c, mg)
		if err != nil {
			return ps, err
		}
		if validateErr := validateProviderConfig(pcSpec, creds); validateErr != nil {
			return ps, errors.Wrap(validateErr, errInvalidProviderConfig)
		}

		region, err := getRegion(mg, creds)
		if err != nil {
			return ps, errors.Wrap(err, "cannot get region")
		}

		ps.Configuration = buildProviderConfiguration(region, creds, pcSpec)
		ps.Configuration["configuration_source"] = getUserAgent()
		return ps, errors.Wrap(configureNoForkAlibabaCloudClient(ctx, &ps, *tfProvider), "failed to configure the no-fork AlibabaCloud client")
	}
}

// resolveProviderConfig resolves the ProviderConfig a managed resource
// references, tracks its usage, and returns its spec normalised to the modern
// (namespaced API group) type so that everything downstream operates on one
// type regardless of which of the three provider config kinds was referenced.
//
// crossplane-runtime v2 splits resource.Managed into LegacyManaged, whose
// providerConfigRef is untyped and always names a ProviderConfig in the legacy
// API group, and ModernManaged, whose reference is typed and may name either a
// namespaced ProviderConfig or a cluster-scoped ClusterProviderConfig in the
// modern group. Legacy MRs must only resolve legacy configs and modern MRs only
// modern ones, so the two paths are kept deliberately separate.
func resolveProviderConfig(ctx context.Context, c client.Client, mg resource.Managed) (*namespacedv1beta1.ProviderConfigSpec, map[string]any, error) {
	switch managed := mg.(type) {
	case resource.LegacyManaged: //nolint:staticcheck // cluster-scoped MRs are legacy by definition
		return resolveLegacy(ctx, c, managed)
	case resource.ModernManaged:
		return resolveModern(ctx, c, managed)
	default:
		return nil, nil, errors.New(errNotManaged)
	}
}

func resolveLegacy(ctx context.Context, c client.Client, mg resource.LegacyManaged) (*namespacedv1beta1.ProviderConfigSpec, map[string]any, error) { //nolint:staticcheck // see resolveProviderConfig
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, nil, errors.New(errNoProviderConfig)
	}

	t := resource.NewLegacyProviderConfigUsageTracker(c, &clusterv1beta1.ProviderConfigUsage{}) //nolint:staticcheck // legacy PCU carries an untyped reference
	if err := t.Track(ctx, mg); err != nil {
		return nil, nil, errors.Wrap(err, errTrackUsage)
	}

	pc := &clusterv1beta1.ProviderConfig{}
	if err := c.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, nil, errors.Wrap(err, errGetProviderConfig)
	}

	spec := legacySpecToModern(&pc.Spec)
	creds, err := extractCredentials(ctx, c, spec, "")
	return spec, creds, err
}

func resolveModern(ctx context.Context, c client.Client, mg resource.ModernManaged) (*namespacedv1beta1.ProviderConfigSpec, map[string]any, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, nil, errors.New(errNoProviderConfig)
	}

	t := resource.NewProviderConfigUsageTracker(c, &namespacedv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, nil, errors.Wrap(err, errTrackUsage)
	}

	var spec *namespacedv1beta1.ProviderConfigSpec
	switch configRef.Kind {
	case namespacedv1beta1.ProviderConfigKind:
		pc := &namespacedv1beta1.ProviderConfig{}
		// A namespaced ProviderConfig is always read from the MR's own
		// namespace; there is no cross-namespace reference.
		if err := c.Get(ctx, types.NamespacedName{Name: configRef.Name, Namespace: mg.GetNamespace()}, pc); err != nil {
			return nil, nil, errors.Wrap(err, errGetProviderConfig)
		}
		spec = &pc.Spec
	case namespacedv1beta1.ClusterProviderConfigKind:
		pc := &namespacedv1beta1.ClusterProviderConfig{}
		if err := c.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return nil, nil, errors.Wrap(err, errGetProviderConfig)
		}
		spec = &pc.Spec
	default:
		return nil, nil, errors.Errorf("%s: %q", errUnknownPCKind, configRef.Kind)
	}

	// Secret references on a modern ProviderConfig are local: they resolve in
	// the referring MR's namespace, not in whatever namespace the manifest
	// happens to carry.
	creds, err := extractCredentials(ctx, c, spec, mg.GetNamespace())
	return spec, creds, err
}

// legacySpecToModern converts a legacy cluster-scoped ProviderConfig spec to
// the modern type. The two are field-for-field identical; they are distinct Go
// types only because they belong to different API groups.
func legacySpecToModern(in *clusterv1beta1.ProviderConfigSpec) *namespacedv1beta1.ProviderConfigSpec {
	out := &namespacedv1beta1.ProviderConfigSpec{
		Credentials: namespacedv1beta1.ProviderCredentials{
			Source:                    in.Credentials.Source,
			CommonCredentialSelectors: in.Credentials.CommonCredentialSelectors,
		},
	}
	if in.AssumeRole != nil {
		out.AssumeRole = &namespacedv1beta1.AssumeRoleOptions{
			RoleARN:           in.AssumeRole.RoleARN,
			SessionName:       in.AssumeRole.SessionName,
			Policy:            in.AssumeRole.Policy,
			SessionExpiration: in.AssumeRole.SessionExpiration,
			ExternalID:        in.AssumeRole.ExternalID,
		}
	}
	if in.AssumeRoleWithOIDC != nil {
		out.AssumeRoleWithOIDC = &namespacedv1beta1.AssumeRoleWithOIDCOptions{
			RoleARN:           in.AssumeRoleWithOIDC.RoleARN,
			OIDCProviderARN:   in.AssumeRoleWithOIDC.OIDCProviderARN,
			OIDCTokenFile:     in.AssumeRoleWithOIDC.OIDCTokenFile,
			RoleSessionName:   in.AssumeRoleWithOIDC.RoleSessionName,
			Policy:            in.AssumeRoleWithOIDC.Policy,
			SessionExpiration: in.AssumeRoleWithOIDC.SessionExpiration,
		}
	}
	return out
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

func buildProviderConfiguration(region string, creds map[string]any, pc *namespacedv1beta1.ProviderConfigSpec) terraform.ProviderConfiguration {
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
	if pc.AssumeRole != nil {
		cfg["assume_role"] = []any{assumeRoleConfiguration(pc.AssumeRole)}
	}
	if pc.AssumeRoleWithOIDC != nil {
		cfg["assume_role_with_oidc"] = []any{assumeRoleWithOIDCConfiguration(pc.AssumeRoleWithOIDC)}
	}
	return cfg
}

func assumeRoleConfiguration(assumeRole *namespacedv1beta1.AssumeRoleOptions) map[string]any {
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

func assumeRoleWithOIDCConfiguration(assumeRole *namespacedv1beta1.AssumeRoleWithOIDCOptions) map[string]any {
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

func validateProviderConfig(pc *namespacedv1beta1.ProviderConfigSpec, creds map[string]any) error {
	for key, field := range providerSpecCredentialKeys {
		if _, ok := creds[key]; ok {
			return errors.Errorf("credentials JSON must not contain %q; use %s", key, field)
		}
	}
	if pc == nil {
		return nil
	}
	if ar := pc.AssumeRole; ar != nil {
		if err := validateAssumeRole(ar); err != nil {
			return err
		}
	}
	if ar := pc.AssumeRoleWithOIDC; ar != nil {
		if err := validateAssumeRoleWithOIDC(ar); err != nil {
			return err
		}
	}
	return nil
}

func validateAssumeRole(ar *namespacedv1beta1.AssumeRoleOptions) error {
	if strings.TrimSpace(ar.RoleARN) == "" {
		return errors.New("spec.assumeRole.roleARN is required")
	}
	return validateSessionExpiration("spec.assumeRole.sessionExpiration", ar.SessionExpiration, 3600)
}

func validateAssumeRoleWithOIDC(ar *namespacedv1beta1.AssumeRoleWithOIDCOptions) error {
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

// extractCredentials reads and unmarshals the credentials a ProviderConfig spec
// points at. When namespace is non-empty the secret is read from there,
// overriding whatever the spec carries: secret references on a modern
// ProviderConfig are local to the referring managed resource.
func extractCredentials(ctx context.Context, c client.Client, spec *namespacedv1beta1.ProviderConfigSpec, namespace string) (map[string]any, error) {
	creds := map[string]any{}

	if spec.Credentials.Source == v1.CredentialsSourceInjectedIdentity ||
		spec.Credentials.Source == v1.CredentialsSourceNone {
		return creds, nil
	}

	selectors := spec.Credentials.CommonCredentialSelectors
	if namespace != "" && selectors.SecretRef != nil {
		local := *selectors.SecretRef
		local.Namespace = namespace
		selectors.SecretRef = &local
	}

	data, err := resource.CommonCredentialExtractor(ctx, spec.Credentials.Source, c, selectors)
	if err != nil {
		return creds, errors.Wrap(err, errExtractCredentials)
	}
	if err = json.Unmarshal(data, &creds); err != nil {
		return creds, errors.Wrap(err, errUnmarshalCredentials)
	}
	return creds, nil
}

func getUserAgent() string {
	// user agent formats as "crossplane/<CROSSPLANE_VERSION> <PROJECT_NAME>/<PROJECT_VERSION>"
	return fmt.Sprintf("crossplane/%s provider-upjet-alibabacloud/%s", version.CrossplaneVersion, version.ProviderVersion)
}
