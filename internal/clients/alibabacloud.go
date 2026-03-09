/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/crossplane-contrib/provider-alibabacloud/internal/version"
	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/crossplane-contrib/provider-alibabacloud/apis/v1beta1"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal alicloud credentials as JSON"

	// OIDC error messages
	errReadOIDCToken      = "cannot read OIDC token from file"
	errAssumeRoleWithOIDC = "cannot assume role with OIDC"

	// OIDC constants
	defaultTokenPath       = "/var/run/secrets/upbound.io/provider/token"
	defaultSessionName     = "crossplane-oidc-session"
	defaultRegion          = "cn-hangzhou"
	defaultDurationSeconds = 3600 // 1 hour
)

// credentialsSourceWebIdentity is the source value for OIDC/WebIdentity authentication.
// xpv1.CredentialsSource is a string type, so we define the constant here.
const credentialsSourceWebIdentity v1.CredentialsSource = "WebIdentity"

var (
	oidcCredentialCaches = &OIDCCredentialCacheMap{
		caches: make(map[string]*OIDCCredentialCache),
	}
)

// OIDCCredentialCacheMap manages credential caches per ProviderConfig
type OIDCCredentialCacheMap struct {
	caches map[string]*OIDCCredentialCache
	mu     sync.RWMutex
}

// getCache retrieves or creates a cache for a specific ProviderConfig
func (m *OIDCCredentialCacheMap) getCache(configName string) *OIDCCredentialCache {
	m.mu.RLock()
	cache, exists := m.caches[configName]
	m.mu.RUnlock()

	if exists {
		return cache
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	// Double-check after acquiring write lock
	if existingCache, exists := m.caches[configName]; exists {
		return existingCache
	}

	newCache := &OIDCCredentialCache{
		creds: make(map[string]string),
	}
	m.caches[configName] = newCache
	return newCache
}

// OIDCCredentialCache caches temporary credentials obtained via OIDC token exchange
type OIDCCredentialCache struct {
	creds      map[string]string
	expiration time.Time
	mu         sync.RWMutex
}

// isValid checks if the cached credentials are still valid and not about to expire.
// Returns true if credentials are valid for at least 5 minutes.
func (c *OIDCCredentialCache) isValid() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	// Check if credentials are about to expire soon (less than 5 minutes)
	return time.Now().Add(5 * time.Minute).Before(c.expiration)
}

// getCredentials returns a copy of the cached credentials
func (c *OIDCCredentialCache) getCredentials() map[string]string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	creds := make(map[string]string)
	for k, v := range c.creds {
		creds[k] = v
	}
	return creds
}

// setCredentials updates the cached credentials
func (c *OIDCCredentialCache) setCredentials(creds map[string]string, expiration time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.creds = creds
	c.expiration = expiration
}

// readOIDCToken reads the OIDC token from the file specified by PSAT_TOKEN_PATH env var.
// Falls back to defaultTokenPath if the env var is not set.
func readOIDCToken() (string, error) {
	tokenPath := os.Getenv("PSAT_TOKEN_PATH")
	if tokenPath == "" {
		tokenPath = defaultTokenPath
	}

	// #nosec G304 -- tokenPath is controlled by PSAT_TOKEN_PATH env var or default constant
	data, err := os.ReadFile(tokenPath)
	if err != nil {
		return "", errors.Wrap(err, errReadOIDCToken)
	}

	return string(data), nil
}

// assumeRoleWithOIDC exchanges the OIDC token for temporary STS credentials
func assumeRoleWithOIDC(token, roleArn, providerArn, region, sessionName string) (map[string]string, time.Time, error) {
	stsClient, err := sts.NewClientWithAccessKey(region, "", "")
	if err != nil {
		return nil, time.Time{}, errors.Wrap(err, errAssumeRoleWithOIDC)
	}

	request := sts.CreateAssumeRoleWithOIDCRequest()
	request.Scheme = "HTTPS"
	request.RoleArn = roleArn
	request.OIDCProviderArn = providerArn
	request.OIDCToken = token
	request.RoleSessionName = sessionName
	request.DurationSeconds = requests.NewInteger(defaultDurationSeconds)

	response, err := stsClient.AssumeRoleWithOIDC(request)
	if err != nil {
		return nil, time.Time{}, errors.Wrap(err, errAssumeRoleWithOIDC)
	}

	creds := map[string]string{
		"access_key":     response.Credentials.AccessKeyId,
		"secret_key":     response.Credentials.AccessKeySecret,
		"security_token": response.Credentials.SecurityToken,
		"region":         region,
	}

	expiration, err := time.Parse(time.RFC3339, response.Credentials.Expiration)
	if err != nil {
		expiration = time.Now().Add(time.Duration(defaultDurationSeconds-300) * time.Second)
	}

	return creds, expiration, nil
}

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, c client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(c, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		creds, err := extractAndUnmarshalCredentials(ctx, c, configRef)
		if err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		region, err := getRegion(mg, creds)
		if err != nil {
			return ps, errors.Wrap(err, "cannot get region")
		}

		// Set credentials in Terraform provider configuration.
		ps.Configuration = map[string]any{
			"region": region,
		}
		if v, ok := creds["access_key"]; ok {
			ps.Configuration["access_key"] = v
		}
		if v, ok := creds["secret_key"]; ok {
			ps.Configuration["secret_key"] = v
		}
		if v, ok := creds["security_token"]; ok {
			ps.Configuration["security_token"] = v
		}
		ps.Configuration["configuration_source"] = getUserAgent()
		return ps, nil
	}
}

func getRegion(obj runtime.Object, creds map[string]string) (string, error) {
	fromMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return "", errors.Wrap(err, "cannot convert to unstructured")
	}
	credsRegion := creds["region"]
	if credsRegion == "" {
		// region_id is used as a fallback for old version
		credsRegion = creds["region_id"]
	}
	r, err := fieldpath.Pave(fromMap).GetString("spec.forProvider.region")
	if fieldpath.IsNotFound(err) {
		// Region is not required for all resources, e.g. resource in "ram" group.
		return credsRegion, nil
	}
	return r, err
}

func extractAndUnmarshalCredentials(ctx context.Context, c client.Client, configRef *v1.Reference) (map[string]string, error) {
	pc := &v1beta1.ProviderConfig{}
	creds := map[string]string{}
	if err := c.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return creds, errors.Wrap(err, errGetProviderConfig)
	}

	// Check if OIDC/WebIdentity authentication is configured
	if isOIDCConfigured(pc) {
		return extractOIDCCredentials(configRef, pc)
	}

	// Use standard credential extraction
	return extractStandardCredentials(ctx, c, pc)
}

// isOIDCConfigured returns true when the ProviderConfig is set up for OIDC/WebIdentity auth
func isOIDCConfigured(pc *v1beta1.ProviderConfig) bool {
	return pc.Spec.Credentials.Source == credentialsSourceWebIdentity &&
		pc.Spec.Credentials.OIDC != nil &&
		pc.Spec.Credentials.OIDC.RoleARN != "" &&
		pc.Spec.Credentials.OIDC.ProviderARN != ""
}

// extractOIDCCredentials handles OIDC/WebIdentity-based credential extraction with caching
func extractOIDCCredentials(configRef *v1.Reference, pc *v1beta1.ProviderConfig) (map[string]string, error) {
	creds := map[string]string{}

	opts := pc.Spec.Credentials.OIDC

	// Resolve region: use the configured value or fall back to the default STS region
	region := opts.Region
	if region == "" {
		region = defaultRegion
	}

	// Resolve session name
	sessionName := opts.RoleSessionName
	if sessionName == "" {
		sessionName = defaultSessionName
	}

	// Get or create cache for this ProviderConfig
	cache := oidcCredentialCaches.getCache(configRef.Name)

	// Return cached credentials if they are still valid
	if cache.isValid() {
		return cache.getCredentials(), nil
	}

	// Read the OIDC token from disk
	token, err := readOIDCToken()
	if err != nil {
		return creds, errors.Wrapf(err, "failed to read OIDC token for ProviderConfig %q", configRef.Name)
	}

	// Exchange the OIDC token for temporary STS credentials
	oidcCreds, expiration, err := assumeRoleWithOIDC(token, opts.RoleARN, opts.ProviderARN, region, sessionName)
	if err != nil {
		return creds, errors.Wrapf(err, "OIDC authentication failed for ProviderConfig %q with role %q",
			configRef.Name, opts.RoleARN)
	}

	// Cache the credentials
	cache.setCredentials(oidcCreds, expiration)
	return oidcCreds, nil
}

// extractStandardCredentials handles standard credential extraction from secrets or other sources
func extractStandardCredentials(ctx context.Context, c client.Client, pc *v1beta1.ProviderConfig) (map[string]string, error) {
	creds := map[string]string{}
	data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, c, pc.Spec.Credentials.CommonCredentialSelectors)
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
