#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"
echo "Creating cloud credential secret..."
${KUBECTL} -n upbound-system create secret generic provider-secret --from-literal=credentials="${UPTEST_CLOUD_CREDENTIALS}" --dry-run=client -o yaml | ${KUBECTL} apply -f -

echo "Creating the value secret for the OOS SecretParameter test..."
${KUBECTL} -n upbound-system create secret generic example-secret --from-literal=example-key="oos-secret-value" --dry-run=client -o yaml | ${KUBECTL} apply -f -

echo "Creating the certificate secrets for the SSL Certificates Service tests..."
# The Certificate examples read the certificate body and the private key from
# Secrets. Generate a throwaway self-signed pair per example; CAS only needs a
# well-formed certificate that matches its key.
CERT_DIR="$(mktemp -d)"
create_certificate_secrets() {
  local name="$1" domain="$2"
  openssl req -x509 -newkey rsa:2048 -nodes -days 365 \
    -keyout "${CERT_DIR}/${name}.key" -out "${CERT_DIR}/${name}.crt" \
    -subj "/CN=${domain}" -addext "subjectAltName=DNS:${domain}" 2>/dev/null
  ${KUBECTL} -n upbound-system create secret generic "${name}-cert" \
    --from-file=cert="${CERT_DIR}/${name}.crt" --dry-run=client -o yaml | ${KUBECTL} apply -f -
  ${KUBECTL} -n upbound-system create secret generic "${name}-key" \
    --from-file=key="${CERT_DIR}/${name}.key" --dry-run=client -o yaml | ${KUBECTL} apply -f -
}
create_certificate_secrets example-certificate example.crossplane.io
create_certificate_secrets crossplane-example-default default.crossplane.io
create_certificate_secrets crossplane-example-additional additional.crossplane.io

echo "Waiting until provider is healthy..."
${KUBECTL} wait provider.pkg --all --for condition=Healthy --timeout 5m

echo "Waiting for all pods to come online..."
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
