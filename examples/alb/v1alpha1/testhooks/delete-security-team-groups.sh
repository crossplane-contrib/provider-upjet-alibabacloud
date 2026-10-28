#!/usr/bin/env bash
# Deletes the security groups that Alibaba Cloud's Security Team adds to the
# test account's VPCs, so that uptest can delete the VPCs.
#
# A few minutes after some ALB examples create their VPC, a security group
# named CloudSecurityTeamDefaultSecurityGroupName ("Created by Alibaba Cloud
# Security Team") appears in it. The provider never created it and does not
# know about it, so the VPC's delete fails with a dependency error until
# someone removes the group by hand, and the run hangs in its delete step
# (runs 39956574647 and 40002095956). Not every ALB VPC gets one, and what
# triggers it is not known, so this hook runs for every ALB example's VPC and
# does nothing when there is no such group.
#
# It runs as a pre-delete hook on the VPC, inside the same script that issues
# every other delete. A failing hook would stop the deletes after it and leak
# resources, so this hook never fails: on any problem it says what went wrong
# and exits 0, and the VPC is left as it would have been without the hook.
#
# It deletes only groups with exactly that name, in VPCs that this test
# cluster manages, and never one that the cluster manages itself. It needs
# the Alibaba Cloud CLI; on Linux it installs the pinned version when the CLI
# is missing. Its output is public in CI, so it prints ids and error codes,
# never API responses.
set -uo pipefail

KUBECTL="${KUBECTL:-kubectl}"
GROUP_NAME="CloudSecurityTeamDefaultSecurityGroupName"
REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../../.." && pwd)"
WORK="$(mktemp -d)"
trap 'rm -rf "${WORK}"' EXIT
CONFIG="${WORK}/config.json"

say() { echo "pre-delete-hook: $*"; }

# error_code <output> prints the API error code of a failed CLI call, from
# either of the CLI's error formats, or "unknown".
error_code() {
  local code
  code=$(sed -E $'s/\033\\[[0-9;]*m//g' <<<"$1" | sed -nE 's/^ErrorCode: *([^[:space:]]+).*/\1/p' | head -1)
  [ -n "${code}" ] || code=$(jq -r '.error_code // empty' <<<"$1" 2>/dev/null)
  echo "${code:-unknown}"
}

ensure_cli() {
  command -v aliyun >/dev/null && return 0
  if [ "$(uname -s)" != Linux ]; then
    say "the aliyun CLI is not installed"
    return 1
  fi
  say "installing the aliyun CLI"
  "${REPO_ROOT}/hack/install-aliyun-cli.sh" >/dev/null || {
    say "could not install the aliyun CLI"
    return 1
  }
}

# write_config writes a CLI profile with the provider's credentials, so the
# hook never reads or changes the user's own CLI configuration.
write_config() {
  local credentials
  credentials=$("${KUBECTL}" -n crossplane-system get secret provider-secret \
    -o jsonpath='{.data.credentials}' | base64 -d) || {
    say "could not read the provider's credentials"
    return 1
  }
  jq -e '{current: "hook", meta_path: "", profiles: [{name: "hook", mode: "AK",
      access_key_id: (.access_key | strings), access_key_secret: (.secret_key | strings),
      region_id: "us-west-1", output_format: "json", language: "en"}]}' \
    <<<"${credentials}" >"${CONFIG}" 2>/dev/null || {
    say "the provider's credentials have no access_key and secret_key"
    return 1
  }
}

main() {
  local vpcs managed region vpc_id out groups group
  ensure_cli || return 0
  write_config || return 0

  vpcs=$("${KUBECTL}" get vpcs.vpc.alibabacloud.crossplane.io -o json |
    jq -r '.items[] | [.spec.forProvider.region, (.metadata.annotations["crossplane.io/external-name"] // .status.atProvider.id // "")] | @tsv') || {
    say "could not list the VPCs"
    return 0
  }
  managed=$("${KUBECTL}" get securitygroups.ecs.alibabacloud.crossplane.io -o json 2>/dev/null |
    jq -r '.items[].metadata.annotations["crossplane.io/external-name"] // empty') || managed=""

  while IFS=$'\t' read -r region vpc_id; do
    [[ "${vpc_id}" == vpc-* && -n "${region}" ]] || continue
    if ! out=$(aliyun ecs DescribeSecurityGroups --config-path "${CONFIG}" \
      --RegionId "${region}" --VpcId "${vpc_id}" --SecurityGroupName "${GROUP_NAME}" --PageSize 50 2>&1); then
      say "could not list the security groups in ${vpc_id}: $(error_code "${out}")"
      continue
    fi
    groups=$(jq -r --arg name "${GROUP_NAME}" \
      '.SecurityGroups.SecurityGroup[] | select(.SecurityGroupName == $name) | .SecurityGroupId' <<<"${out}" 2>/dev/null) || {
      say "unexpected response listing the security groups in ${vpc_id}"
      continue
    }
    if [ -z "${groups}" ]; then
      say "no ${GROUP_NAME} in ${vpc_id}"
      continue
    fi
    for group in ${groups}; do
      if grep -qxF -- "${group}" <<<"${managed}"; then
        say "${group} in ${vpc_id} is managed by this test; leaving it"
        continue
      fi
      if out=$(aliyun ecs DeleteSecurityGroup --config-path "${CONFIG}" \
        --RegionId "${region}" --SecurityGroupId "${group}" 2>&1); then
        say "deleted ${group} (${GROUP_NAME}) in ${vpc_id}"
      else
        say "could not delete ${group} in ${vpc_id}: $(error_code "${out}")"
      fi
    done
  done <<<"${vpcs}"
}

main
exit 0
