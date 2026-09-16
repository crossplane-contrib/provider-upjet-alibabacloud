#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"
echo "Creating cloud credential secret..."
${KUBECTL} -n upbound-system create secret generic provider-secret --from-literal=credentials="${UPTEST_CLOUD_CREDENTIALS}" --dry-run=client -o yaml | ${KUBECTL} apply -f -

echo "Creating the value secret for the OOS SecretParameter test..."
${KUBECTL} -n upbound-system create secret generic example-secret --from-literal=example-key="oos-secret-value" --dry-run=client -o yaml | ${KUBECTL} apply -f -

echo "Waiting until provider is healthy..."
${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m

echo "Waiting for all pods to come online..."
${KUBECTL} -n upbound-system wait --for=condition=Available deployment --all --timeout=5m

echo "Creating a default provider config for the legacy cluster-scoped API group..."
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: alibabacloud.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      namespace: upbound-system
      key: credentials
EOF

# Namespaced managed resources resolve kind: ClusterProviderConfig, name: default
# when spec.providerConfigRef is omitted, so the modern API group needs its own
# default credential for the namespaced examples to reconcile.
echo "Creating a default cluster provider config for the modern API group..."
cat <<EOF | ${KUBECTL} apply -f -
apiVersion: alibabacloud.m.crossplane.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: provider-secret
      namespace: upbound-system
      key: credentials
EOF

echo "Creating a login profile secret for RAM tests"

${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m
${KUBECTL} -n upbound-system wait --for=condition=Available deployment --all --timeout=5m
