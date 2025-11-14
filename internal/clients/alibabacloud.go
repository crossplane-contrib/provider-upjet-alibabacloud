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
	errOIDCTokenPathNotSet = "PSAT_TOKEN_PATH environment variable not set"
	errReadOIDCToken       = "cannot read OIDC token from file"
	errAssumeRoleWithOIDC  = "cannot assume role with OIDC"
	errInvalidOIDCConfig   = "invalid OIDC configuration: missing required fields"
	errMarshalCredentials  = "cannot marshal credentials"
	errParseExpiration     = "cannot parse expiration timestamp"
	errFetchCredentials    = "cannot fetch new credentials"
	errValidateCredentials = "cannot validate existing credentials"

	// OIDC constants
	defaultTokenPath       = "/var/run/secrets/upbound.io/provider/token"
	defaultSessionName     = "crossplane-oidc-session"
	defaultDurationSeconds = 3600 // 1 hour
	minimumValidityMinutes = 5    // Minimum validity period in minutes
)

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
	if cache, exists := m.caches[configName]; exists {
		return cache
	}

	cache = &OIDCCredentialCache{
		creds: make(map[string]string),
	}
	m.caches[configName] = cache
	return cache
}

type OIDCCredentialCache struct {
	creds      map[string]string
	expiration time.Time
	mu         sync.RWMutex
}

// isValid checks if the cached credentials are still valid and not about to expire
// Returns true if credentials are valid for at least 5 minutes
func (c *OIDCCredentialCache) isValid() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	// Check if credentials are about to expire soon (less than 5 minutes)
	return time.Now().Add(5 * time.Minute).Before(c.expiration)
}

// getCredentials returns the cached credentials
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

// readOIDCToken reads the OIDC token from the file specified by PSAT_TOKEN_PATH
func readOIDCToken() (string, error) {
	tokenPath := os.Getenv("PSAT_TOKEN_PATH")
	if tokenPath == "" {
		tokenPath = defaultTokenPath
	}

	data, err := os.ReadFile(tokenPath)
	if err != nil {
		return "", errors.Wrap(err, errReadOIDCToken)
	}

	return string(data), nil
}

// assumeRoleWithOIDC exchanges the OIDC token for temporary credentials
func assumeRoleWithOIDC(token, roleArn, providerArn, region string) (map[string]string, time.Time, error) {
	client, err := sts.NewClientWithAccessKey(region, "", "")
	if err != nil {
		return nil, time.Time{}, errors.Wrap(err, errAssumeRoleWithOIDC)
	}

	request := sts.CreateAssumeRoleWithOIDCRequest()
	request.Scheme = "HTTPS"
	request.RoleArn = roleArn
	request.OIDCProviderArn = providerArn
	request.OIDCToken = token
	request.RoleSessionName = defaultSessionName
	request.DurationSeconds = requests.NewInteger(defaultDurationSeconds)

	response, err := client.AssumeRoleWithOIDC(request)
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

	// Check if OIDC authentication is configured
	if pc.Spec.Credentials.Region != "" && pc.Spec.Credentials.ProviderARN != "" && pc.Spec.Credentials.RoleARN != "" {
		// Validate that the source is set appropriately for OIDC
		if pc.Spec.Credentials.Source != v1.CredentialsSourceInjectedIdentity {
			return creds, errors.Errorf("invalid credentials source %q for OIDC authentication, must be %q",
				pc.Spec.Credentials.Source, v1.CredentialsSourceInjectedIdentity)
		}

		// Get or create cache for this ProviderConfig
		cache := oidcCredentialCaches.getCache(configRef.Name)

		// Use OIDC authentication with per-ProviderConfig caching
		if cache.isValid() {
			// Return cached credentials
			return cache.getCredentials(), nil
		}

		// Read OIDC token and assume role
		token, err := readOIDCToken()
		if err != nil {
			return creds, errors.Wrapf(err, "failed to read OIDC token for ProviderConfig %q", configRef.Name)
		}

		creds, expiration, err := assumeRoleWithOIDC(token, pc.Spec.Credentials.RoleARN, pc.Spec.Credentials.ProviderARN, pc.Spec.Credentials.Region)
		if err != nil {
			return creds, errors.Wrapf(err, "OIDC authentication failed for ProviderConfig %q with role %q",
				configRef.Name, pc.Spec.Credentials.RoleARN)
		}

		// Cache the credentials
		cache.setCredentials(creds, expiration)
		return creds, nil
	}

	// Use standard credential extraction
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
