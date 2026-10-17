// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"testing"
)

// The offline setup exists to make the diff server independent of the host it
// runs on. Unlike the AWS provider, the Alibaba Cloud SDK has no implicit
// credential chain once static keys are configured, so static placeholder
// keys are sufficient to guarantee that. What still needs pinning are the two
// provider fields whose lazy resolution makes network calls: account_id
// (resolved via an STS GetCallerIdentity call) and account_type (probed via a
// BssOpenApi QueryAvailableInstances call).

func TestOfflineCredentialsAreStatic(t *testing.T) {
	// An empty key would send the alicloud provider down its fallback path -
	// profile files, credentials_uri, ECS role name - any of which can read
	// the host's environment or make network calls.
	if offlineAccessKeyID == "" || offlineSecretAccessKey == "" {
		t.Error("the offline placeholder credentials must not be empty, or the alicloud provider falls back to its credential discovery path and the setup stops being offline")
	}
}

func TestOfflineSetupPinsNetworkResolvingFields(t *testing.T) {
	// The diff server runs without credentials, so it cannot ask Alibaba
	// Cloud for the account ID or the account type. The provider resolves
	// both lazily over the network the first time a resource needs them,
	// which an offline plan must never trigger. Both have to be pinned in
	// the configuration.
	if DefaultOfflineAccountID == "" {
		t.Error("the offline account ID placeholder must not be empty, or the provider resolves it with an STS GetCallerIdentity call")
	}
	if offlineAccountType == "" {
		t.Error("the offline account type placeholder must not be empty, or the provider probes BssOpenApi for it")
	}
}
