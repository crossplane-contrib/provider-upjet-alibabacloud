#!/usr/bin/env bash
# Checks that an endpoint service keeps the same NLB attached in two zones
# (#212): Terraform provider 2.0.0 collapsed the two attachments into one.
#
# resource[].resourceId takes no reference, so the manifest cannot name the
# NLB it creates. Once everything is ready, this hook:
#   1. adds both attachments to the existing service (the update path);
#   2. creates a second service with both attachments from the start (the
#      create path, as reported), checks it, and deletes it.
# Each check waits for status.atProvider.resource to list both zones.
set -euo pipefail

KUBECTL="${KUBECTL:-kubectl}"
SERVICE=vpcendpointservice.privatelink.alibabacloud.crossplane.io
ZONES='["us-west-1a","us-west-1b"]'
# Well inside the chainsaw exec budget this hook shares with the assert step.
WAIT=600

nlb=$("${KUBECTL}" get loadbalancer.nlb.alibabacloud.crossplane.io nlb-zones -o jsonpath='{.metadata.annotations.crossplane\.io/external-name}')
[ -n "${nlb}" ] || {
  echo "post-assert-hook: the NLB has no external name" >&2
  exit 1
}
resources=$(jq -nc --arg id "${nlb}" --argjson zones "${ZONES}" '[$zones[] | {resourceId: $id, resourceType: "nlb", zoneId: .}]')

# wait_for_zones <name> waits until the service reports both zones attached.
wait_for_zones() {
  local name="$1" deadline=$((SECONDS + WAIT)) got
  while :; do
    got=$("${KUBECTL}" get "${SERVICE}" "${name}" -o json |
      jq -c --arg id "${nlb}" '[.status.atProvider.resource[]? | select(.resourceId == $id) | .zoneId] | sort')
    if [ "${got}" = "${ZONES}" ]; then
      echo "post-assert-hook: ${name} has ${nlb} attached in ${got}"
      return 0
    fi
    if ((SECONDS > deadline)); then
      echo "post-assert-hook: ${name} has ${nlb} attached in ${got}, want ${ZONES}" >&2
      "${KUBECTL}" get "${SERVICE}" "${name}" -o yaml >&2 || true
      return 1
    fi
    sleep 15
  done
}

echo "post-assert-hook: attaching ${nlb} in both zones to the existing service"
"${KUBECTL}" patch "${SERVICE}" nlb-zones --type=merge -p "{\"spec\":{\"forProvider\":{\"resource\":${resources}}}}"
wait_for_zones nlb-zones

echo "post-assert-hook: creating a service with ${nlb} in both zones"
"${KUBECTL}" apply -f - <<EOF
apiVersion: privatelink.alibabacloud.crossplane.io/v1alpha1
kind: VPCEndpointService
metadata:
  name: nlb-zones-create
spec:
  forProvider:
    region: us-west-1
    autoAcceptConnection: false
    serviceDescription: crossplane-example-nlb-zones-create
    serviceResourceType: nlb
    resource: ${resources}
EOF
"${KUBECTL}" wait "${SERVICE}" nlb-zones-create --for=condition=Ready --timeout="${WAIT}s"
wait_for_zones nlb-zones-create
"${KUBECTL}" delete "${SERVICE}" nlb-zones-create --timeout="${WAIT}s"
