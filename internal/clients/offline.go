// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"

	"github.com/crossplane/crossplane-runtime/v2/pkg/errors"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Placeholder credentials and identifiers the offline setup configures the
// Alibaba Cloud Terraform provider with.
//
// They are what makes the setup credential-free: the alicloud provider has no
// implicit credential chain once static keys are set, so handing it a static
// key pair keeps the diff server from both requiring and using the host's
// credentials. A key left empty would instead send the provider down its
// fallback path - profile files, credentials_uri, ECS role name - any of which
// can read the host's environment or make network calls.
const (
	offlineAccessKeyID     = "OFFLINEDIFFSRVRPLACEHLD" //nolint:gosec // a placeholder, never used to sign a request
	offlineSecretAccessKey = "offlineDiffServerPlaceholderSecretKey000"
	offlineSecurityToken   = "offlineDiffServerPlaceholderToken0000"

	// DefaultOfflineAccountID stands in for the Alibaba Cloud account ID when
	// the caller does not supply one. The diff server cannot ask STS for the
	// real account, but the alicloud provider resolves the account ID lazily
	// via a GetCallerIdentity call the first time a resource needs it, and
	// several resources template the account ID into their Terraform IDs.
	// Setting account_id in the provider configuration short-circuits that
	// call; the placeholder keeps those IDs syntactically valid.
	DefaultOfflineAccountID = "0000000000000000"

	// offlineAccountType pins the provider's account type so the client does
	// not probe BssOpenApi for it the first time a resource consults the
	// billing API. "Domestic" is the provider's own default when probing
	// fails, so it is a safe placeholder.
	offlineAccountType = "Domestic"
)

const (
	errOfflineNoProviderConfig   = "the managed resource does not reference a provider config"
	errOfflineRegion             = "cannot get the region of the managed resource"
	errOfflineConfigureProvider  = "cannot configure the Alibaba Cloud provider offline"
	errOfflineNotLegacyManaged   = "resource is not a legacy (cluster-scoped) managed resource"
	errOfflineNilTerraformSchema = "terraform provider cannot be nil"
)

// OfflineTerraformSetupBuilder returns a Terraform setup that configures the
// Alibaba Cloud Terraform provider without credentials and without contacting
// Alibaba Cloud.
//
// Use it in place of TerraformSetupBuilder in the diff server only: the
// resulting terraform.Setup carries a provider meta built from placeholder
// credentials, which is enough for the provider's schema and CustomizeDiff
// logic. The ProviderConfig referenced by the managed resource is still
// fetched - through the diff server's in-memory client, so no cluster access -
// but its credentials are ignored: only the region it influences matters for
// a diff, and that comes from the resource itself (spec.forProvider.region)
// with the ProviderConfig's credential region as the fallback.
//
// AssumeRole and AssumeRoleWithOIDC configurations are dropped: assuming a
// role requires an STS call, which is not possible offline.
//
// Unlike the AWS provider, there is no single egress-guard hook that covers
// every service client the alicloud provider can build, so a resource whose
// CustomizeDiff genuinely calls Alibaba Cloud fails with the connection or
// authentication error of the placeholder credentials rather than a clean
// DIFF_COMPUTATION_NOT_SUPPORTED. The account_id and account_type
// placeholders keep the two known construction-time probes (GetCallerIdentity
// and BssOpenApi QueryAvailableInstances) from firing; everything beyond that
// is left to the request failing on its own.
//
// As of the pinned provider (v1.294.1), only one CustomizeDiff in the whole
// alicloud provider makes a cloud call - alicloud_mse_cluster's
// QueryClusterInfo - and it is not among the resources this provider exposes.
// The nine CustomizeDiffs on exposed resources (alicloud_instance,
// alicloud_alb_server_group, alicloud_ecs_launch_template,
// alicloud_nlb_load_balancer, and the unexposed remainder) are pure
// computation over the diff, so the limitation is currently theoretical.
// Re-audit when the upstream provider pin changes.
func OfflineTerraformSetupBuilder(tfProvider *schema.Provider) terraform.SetupFn {
	return func(ctx context.Context, c client.Client, mg xpresource.Managed) (terraform.Setup, error) {
		lmg, ok := mg.(xpresource.LegacyManaged) //nolint:staticcheck // cluster-scoped MRs are legacy by definition
		if !ok {
			return terraform.Setup{}, errors.New(errOfflineNotLegacyManaged)
		}
		configRef := lmg.GetProviderConfigReference()
		if configRef == nil {
			return terraform.Setup{}, errors.New(errOfflineNoProviderConfig)
		}

		// Resolve the ProviderConfig only to honor the region its
		// credentials declare; the credentials themselves are discarded.
		pc, creds, err := extractAndUnmarshalCredentials(ctx, c, configRef)
		if err != nil {
			// The diff server's object store may legitimately lack the
			// Secret backing the ProviderConfig. Fall back to an empty
			// credential map so the region is taken from the resource.
			creds = map[string]any{}
		}
		_ = pc

		region, err := getRegion(mg, creds)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, errOfflineRegion)
		}

		ps := terraform.Setup{
			Configuration: terraform.ProviderConfiguration{
				"region":         region,
				"access_key":     offlineAccessKeyID,
				"secret_key":     offlineSecretAccessKey,
				"security_token": offlineSecurityToken,
				// Setting these two keeps the provider from making their
				// lazy network calls; see the constants above.
				"account_id":   DefaultOfflineAccountID,
				"account_type": offlineAccountType,
			},
		}
		if tfProvider == nil {
			return terraform.Setup{}, errors.New(errOfflineNilTerraformSchema)
		}
		// Configuring with static placeholder keys builds the client
		// purely from configuration: no profile files, no metadata
		// service, no STS. The provider is passed by value for the same
		// per-request isolation reason as the online setup; see
		// configureNoForkAlibabaCloudClient.
		return ps, errors.Wrap(configureNoForkAlibabaCloudClient(ctx, &ps, *tfProvider), errOfflineConfigureProvider)
	}
}
