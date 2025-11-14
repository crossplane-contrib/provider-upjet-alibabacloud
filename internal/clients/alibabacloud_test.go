/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestOIDCCredentialCache(t *testing.T) {
	tests := map[string]struct {
		creds      map[string]string
		expiration time.Time
		wantValid  bool
	}{
		"ValidCredentials": {
			creds: map[string]string{
				"access_key":     "test-access-key",
				"secret_key":     "test-secret-key",
				"security_token": "test-security-token",
			},
			expiration: time.Now().Add(1 * time.Hour),
			wantValid:  true,
		},
		"ExpiredCredentials": {
			creds: map[string]string{
				"access_key":     "test-access-key",
				"secret_key":     "test-secret-key",
				"security_token": "test-security-token",
			},
			expiration: time.Now().Add(-1 * time.Hour),
			wantValid:  false,
		},
		"EmptyCredentials": {
			creds:      map[string]string{},
			expiration: time.Now().Add(1 * time.Hour),
			wantValid:  true,
		},
		"AboutToExpire": {
			creds: map[string]string{
				"access_key":     "test-access-key",
				"secret_key":     "test-secret-key",
				"security_token": "test-security-token",
			},
			expiration: time.Now().Add(3 * time.Minute),
			wantValid:  false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			cache := &OIDCCredentialCache{
				creds: make(map[string]string),
			}
			cache.setCredentials(tc.creds, tc.expiration)

			// Test isValid
			gotValid := cache.isValid()
			if gotValid != tc.wantValid {
				t.Errorf("isValid() = %v, want %v", gotValid, tc.wantValid)
			}

			// Test getCredentials returns a copy
			gotCreds := cache.getCredentials()
			if diff := cmp.Diff(tc.creds, gotCreds); diff != "" {
				t.Errorf("getCredentials() mismatch (-want +got):\n%s", diff)
			}

			// Verify it's a copy by modifying the returned map
			gotCreds["new_key"] = "new_value"
			// Get credentials again to verify original is unchanged
			gotCreds2 := cache.getCredentials()
			if _, exists := gotCreds2["new_key"]; exists {
				t.Error("getCredentials() should return a copy, but returned the original map")
			}
		})
	}
}

func TestOIDCCredentialCacheMap(t *testing.T) {
	cacheMap := &OIDCCredentialCacheMap{
		caches: make(map[string]*OIDCCredentialCache),
	}

	// Test getting cache for different ProviderConfigs
	cache1 := cacheMap.getCache("config1")
	cache2 := cacheMap.getCache("config2")
	cache1Again := cacheMap.getCache("config1")

	// Verify cache1 and cache1Again are the same instance
	if cache1 != cache1Again {
		t.Error("getCache should return the same cache instance for the same config name")
	}

	// Verify cache1 and cache2 are different instances
	if cache1 == cache2 {
		t.Error("getCache should return different cache instances for different config names")
	}

	// Set credentials in cache1 and verify they don't affect cache2
	creds1 := map[string]string{"access_key": "key1"}
	cache1.setCredentials(creds1, time.Now().Add(1*time.Hour))

	cache1Creds := cache1.getCredentials()
	cache2Creds := cache2.getCredentials()

	if diff := cmp.Diff(creds1, cache1Creds); diff != "" {
		t.Errorf("cache1 credentials mismatch (-want +got):\n%s", diff)
	}

	if len(cache2Creds) != 0 {
		t.Errorf("cache2 should be empty, got %v", cache2Creds)
	}
}

func TestOIDCCredentialCacheExpirationChecking(t *testing.T) {
	tests := map[string]struct {
		expiration time.Time
		wantValid  bool
	}{
		"ValidForMoreThan5Minutes": {
			expiration: time.Now().Add(10 * time.Minute),
			wantValid:  true,
		},
		"ExpiresInLessThan5Minutes": {
			expiration: time.Now().Add(3 * time.Minute),
			wantValid:  false,
		},
		"ExpiresInExactly5Minutes": {
			expiration: time.Now().Add(5 * time.Minute),
			wantValid:  false, // Should be false because we check for "before" 5 minutes
		},
		"AlreadyExpired": {
			expiration: time.Now().Add(-1 * time.Minute),
			wantValid:  false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			cache := &OIDCCredentialCache{}
			creds := map[string]string{
				"access_key":     "test-access-key",
				"secret_key":     "test-secret-key",
				"security_token": "test-security-token",
			}
			cache.setCredentials(creds, tc.expiration)

			gotValid := cache.isValid()
			if gotValid != tc.wantValid {
				t.Errorf("isValid() = %v, want %v", gotValid, tc.wantValid)
			}
		})
	}
}

func TestOIDCCredentialCacheThreadSafety(t *testing.T) {
	cache := &OIDCCredentialCache{}
	creds := map[string]string{
		"access_key":     "test-access-key",
		"secret_key":     "test-secret-key",
		"security_token": "test-security-token",
	}

	// Test concurrent access
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			cache.setCredentials(creds, time.Now().Add(1*time.Hour))
			cache.isValid()
			cache.getCredentials()
			done <- true
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}
}

func TestConstants(t *testing.T) {
	// Test that our constants are defined
	tests := []struct {
		name     string
		expected string
		actual   string
	}{
		{
			name:     "defaultTokenPath",
			expected: "/var/run/secrets/upbound.io/provider/token",
			actual:   defaultTokenPath,
		},
		{
			name:     "defaultSessionName",
			expected: "crossplane-oidc-session",
			actual:   defaultSessionName,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.actual != tc.expected {
				t.Errorf("Expected %s = %q, got %q", tc.name, tc.expected, tc.actual)
			}
		})
	}
}
