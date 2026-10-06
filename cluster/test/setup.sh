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

# The provider's CLI/WorkspaceStore reconciliation only checks on a pending
# async terraform apply once per poll interval, and can take two full cycles
# to notice completion. At the 10m default that alone can eat half of
# chainsaw's --default-timeout budget before a resource's external-name is
# even recorded. Shorten it for this e2e run only; production installs keep
# the binary's own default.
echo "Shortening the provider's poll interval for faster e2e feedback..."
for rtc in $(${KUBECTL} get deploymentruntimeconfig -o jsonpath='{.items[*].metadata.name}'); do
  ${KUBECTL} patch deploymentruntimeconfig "${rtc}" --type merge -p '{"spec":{"deploymentTemplate":{"spec":{"template":{"spec":{"containers":[{"name":"package-runtime","args":["--debug","--poll=1m"]}]}}}}}}'
done

echo "Waiting for the provider to roll out with the shortened poll interval..."
${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m
${KUBECTL} -n upbound-system wait --for=condition=Available deployment --all --timeout=5m

echo "Creating a default provider config..."
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

echo "Creating a login profile secret for RAM tests"

${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m
${KUBECTL} -n upbound-system wait --for=condition=Available deployment --all --timeout=5m
