// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"sync"
	"time"

	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

// defaultMetaTTL bounds how long a configured client is reused. It is kept
// below the shortest session the provider config validation allows (900s), so
// even without a configured session_expiration an assumed role's STS token is
// replaced well before it can expire.
const defaultMetaTTL = 10 * time.Minute

// providerMetaCache reuses the configured Terraform provider meta -- the
// upstream *connectivity.AliyunClient -- across reconciles that share the same
// provider configuration.
//
// Without it every reconcile configures a fresh client, and the upstream
// client's endpoint cache dies with it. Most products resolve their endpoint
// through Alibaba's Location service, which lives only in mainland China
// (location.aliyuncs.com, then location-readonly.aliyuncs.com) and retries for
// up to two minutes. Many reads wrap their API call, endpoint lookup included,
// in a one-minute retry (DescribeVpcAttribute, for one), so whenever Location
// is slow to reach -- routinely from outside mainland China -- the read times
// out on every reconcile and the resource can never be observed or deleted.
// Reusing the client means the lookup, and its fallback, happen once per
// product per configuration per TTL.
//
// Upstream shares one client across concurrent resource operations, so sharing
// it across reconciles is safe.
type providerMetaCache struct {
	mu      sync.Mutex
	entries map[string]metaCacheEntry
	now     func() time.Time
}

type metaCacheEntry struct {
	meta    any
	expires time.Time
}

func newProviderMetaCache() *providerMetaCache {
	return &providerMetaCache{entries: map[string]metaCacheEntry{}, now: time.Now}
}

var metaCache = newProviderMetaCache()

// configure sets ps.Meta from the cache when an unexpired client exists for
// ps.Configuration, and otherwise configures p and caches the result. Failed
// configurations are not cached.
func (c *providerMetaCache) configure(ctx context.Context, ps *terraform.Setup, p schema.Provider) error {
	key, err := metaCacheKey(ps.Configuration)
	if err != nil {
		return err
	}
	now := c.now()
	c.mu.Lock()
	e, ok := c.entries[key]
	c.mu.Unlock()
	if ok && now.Before(e.expires) {
		ps.Meta = e.meta
		return nil
	}
	if err := configureNoForkAlibabaCloudClient(ctx, ps, p); err != nil {
		return err
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, e := range c.entries {
		if !now.Before(e.expires) {
			delete(c.entries, k)
		}
	}
	c.entries[key] = metaCacheEntry{meta: ps.Meta, expires: now.Add(metaTTL(ps.Configuration))}
	return nil
}

// metaCacheKey hashes the full provider configuration, credentials included, so
// a rotated secret or a changed ProviderConfig gets a new client at once.
// encoding/json writes map keys in sorted order, so equal configurations hash
// equally.
func metaCacheKey(cfg terraform.ProviderConfiguration) (string, error) {
	b, err := json.Marshal(cfg)
	if err != nil {
		return "", errors.Wrap(err, "cannot marshal the provider configuration")
	}
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:]), nil
}

// metaTTL returns defaultMetaTTL, or half the assumed role's session_expiration
// when that is shorter.
func metaTTL(cfg terraform.ProviderConfiguration) time.Duration {
	ttl := defaultMetaTTL
	for _, k := range []string{"assume_role", "assume_role_with_oidc"} {
		blocks, ok := cfg[k].([]any)
		if !ok || len(blocks) == 0 {
			continue
		}
		block, ok := blocks[0].(map[string]any)
		if !ok {
			continue
		}
		if exp, ok := block["session_expiration"].(int); ok && exp > 0 {
			if half := time.Duration(exp) * time.Second / 2; half < ttl {
				ttl = half
			}
		}
	}
	return ttl
}
