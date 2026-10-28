#!/usr/bin/env bash
# Tests hack/cloud-cleanup.sh against a stub `aliyun` on PATH, so no call ever
# reaches Alibaba Cloud. Run it from anywhere: hack/cloud-cleanup-test.sh
set -euo pipefail

SCRIPT="$(cd "$(dirname "$0")" && pwd)/cloud-cleanup.sh"
WORK="$(mktemp -d)"
trap 'rm -rf "${WORK}"' EXIT
mkdir -p "${WORK}/bin"

# The stub logs every call and answers by API action (the second argument,
# such as DeleteVpc, or rm for `aliyun oss rm`), from files in $STUB_DIR:
#   creds                 "bad" fails every call, as with a rotated key
#   fail/<action>         fails that action with the error code in the file
#   fail-once/<action>    fails only its first call
#   fail-deletes          fails every delete
#   response/<action>     printed as the response; response/<action>.p<N>
#                         answers page N of a paged call
#   format                "json" prints errors as the CLI's AI agent mode
#                         does; otherwise as text, as in CI, with every line
#                         colored so that parsing has to strip the colors
# Anything else succeeds with an empty JSON object, or an empty policy list
# for ListPoliciesForRole.
cat >"${WORK}/bin/aliyun" <<'STUB'
#!/usr/bin/env bash
set -euo pipefail
echo "$*" >>"${STUB_DIR}/calls"
fail() {
  if [ "$(cat "${STUB_DIR}/format" 2>/dev/null)" = json ]; then
    echo "{\"schemaVersion\":\"v1\",\"message\":\"stub error\",\"error_code\":\"$1\",\"status_code\":400}" >&2
  else
    printf '\033[1;31m%s\033[0m\n' "ERROR: SDK.ServerError" "ErrorCode: $1" "Recommend: https://example.invalid" "Message: stub error" >&2
  fi
  exit 1
}
[ "$(cat "${STUB_DIR}/creds" 2>/dev/null)" = bad ] && fail InvalidAccessKeyId.NotFound
action="$2"
if [ -f "${STUB_DIR}/fail-once/${action}" ] && [ ! -f "${STUB_DIR}/failed-once-${action}" ]; then
  touch "${STUB_DIR}/failed-once-${action}"
  fail "$(cat "${STUB_DIR}/fail-once/${action}")"
fi
[ -f "${STUB_DIR}/fail/${action}" ] && fail "$(cat "${STUB_DIR}/fail/${action}")"
if [ -f "${STUB_DIR}/fail-deletes" ]; then
  case "${action}" in Delete* | rm) fail StubDeleteFailed ;; esac
fi
page=""
args=("$@")
for ((i = 0; i < ${#args[@]} - 1; i++)); do
  [ "${args[i]}" = --PageNumber ] && page="${args[i + 1]}"
done
for response in "${STUB_DIR}/response/${action}.p${page}" "${STUB_DIR}/response/${action}"; do
  if [ -f "${response}" ]; then
    cat "${response}"
    exit 0
  fi
done
case "${action}" in
  GetCallerIdentity) echo '{"Arn":"acs:ram::1:user/stub"}' ;;
  ListPoliciesForRole) echo '{"Policies":{"Policy":[]}}' ;;
  ls) ;;
  *) echo '{}' ;;
esac
STUB
chmod +x "${WORK}/bin/aliyun"

failures=0
case_dir=""

# start <name> resets the stub's state for a new case.
start() {
  case_dir="${WORK}/$1"
  mkdir -p "${case_dir}/fail" "${case_dir}/fail-once" "${case_dir}/response"
  : >"${case_dir}/calls"
  echo "--- $1"
}

# respond <action> <json> sets the stub's response to an action.
respond() { echo "$2" >"${case_dir}/response/$1"; }
fail_with() { echo "$2" >"${case_dir}/fail/$1"; }

# run <args>... runs the script against the stub, saving output and status.
run() {
  set +e
  STUB_DIR="${case_dir}" PATH="${WORK}/bin:${PATH}" PASSES="${PASSES:-3}" PASS_WAIT=0 CONFIRM="${CONFIRM-delete}" \
    bash "${SCRIPT}" "$@" >"${case_dir}/out" 2>&1
  echo $? >"${case_dir}/status"
  set -e
}

targets() {
  printf '%s\n' "$@" >"${case_dir}/targets"
  run delete "${case_dir}/targets"
}

check() {
  local description="$1"
  shift
  if "$@"; then
    echo "ok: ${description}"
  else
    echo "FAIL: ${description}"
    sed 's/^/    /' "${case_dir}/out"
    failures=$((failures + 1))
  fi
}
status_is() { [ "$(cat "${case_dir}/status")" = "$1" ]; }
output_has() { grep -qF -- "$1" "${case_dir}/out"; }
output_lacks() { ! grep -qF -- "$1" "${case_dir}/out"; }
called() { grep -qF -- "$1" "${case_dir}/calls"; }
not_called() { ! grep -qF -- "$1" "${case_dir}/calls"; }
first_line_of() { grep -nF -- "$1" "${case_dir}/calls" | head -1 | cut -d: -f1; }
before() { [ "$(first_line_of "$1")" -lt "$(first_line_of "$2")" ]; }

for format in text json; do
  start "bad-credentials-${format}"
  echo "${format}" >"${case_dir}/format"
  echo bad >"${case_dir}/creds"
  targets "us-west-1 vpc-a"
  check "a bad key fails the run (${format} errors)" status_is 1
  check "a bad key is reported with its code (${format} errors)" \
    output_has "cannot authenticate to Alibaba Cloud: InvalidAccessKeyId.NotFound: stub error"
  check "a bad key deletes nothing (${format} errors)" not_called "DeleteVpc"
done

start bad-credentials-inventory
echo bad >"${case_dir}/creds"
run inventory us-west-1
check "inventory with a bad key fails" status_is 1

start unknown-named-type
targets "us-west-1 kafak:x"
check "a mistyped type is rejected" status_is 1
check "the rejection names the target" output_has 'unknown resource type in "kafak:x"'
check "nothing is deleted for a rejected target" not_called "Delete"

start missing-comma
targets "us-west-1 vpc-a us-west-1 vpc-b"
check "a line with two targets is rejected" status_is 1
check "the rejection names the line" output_has 'target "us-west-1 vpc-a us-west-1 vpc-b" needs exactly a region and an id'
check "nothing is deleted for a line with two targets" not_called "Delete"

start unknown-prefix
targets "us-west-1 foo-1"
check "an id without a known prefix is rejected" status_is 1

start shared-prefixes
targets "us-west-1 lsn-1"
check "a bare listener id is rejected, since ALB and NLB share the prefix" status_is 1

start prefixes
targets "us-west-1 epsrv-1" "us-west-1 ep-1" "us-west-1 rule-1" "us-west-1 r-1" "us-west-1 rr-1" "us-west-1 rm-1" \
  "us-west-1 alb-listener:lsn-1" "us-west-1 nlb-servergroup:sgp-1"
check "all targets deleted" status_is 0
check "epsrv- is an endpoint service" called "privatelink DeleteVpcEndpointService --RegionId us-west-1 --ServiceId epsrv-1"
check "ep- is an endpoint" called "privatelink DeleteVpcEndpoint --RegionId us-west-1 --EndpointId ep-1"
check "rule- is an ALB rule, not Tair" called "alb DeleteRule --region us-west-1 --RuleId rule-1"
check "r- is Tair" called "r-kvstore DeleteInstance --region us-west-1 --InstanceId r-1"
check "rr- is an RDS read-only instance" called "rds DeleteDBInstance --region us-west-1 --DBInstanceId rr-1"
check "a named ALB listener goes to ALB" called "alb DeleteListener --region us-west-1 --ListenerId lsn-1"
check "a named NLB server group goes to NLB" called "nlb DeleteServerGroup --RegionId us-west-1 --ServerGroupId sgp-1"

start order-and-retry
echo DependencyViolation.SecurityGroup >"${case_dir}/fail-once/DeleteVpc"
targets "us-west-1 vpc-a" "us-west-1 sg-a" "us-west-1 alb-a" "us-west-1 alb-listener:lsn-a" \
  "us-west-1 epsrv-a" "us-west-1 ep-a"
check "the run succeeds" status_is 0
check "the summary counts every delete" output_has "Delete accepted for 6, not found 0, failed 0."
check "the blocked VPC was retried" \
  output_has "pass 1: \`vpc-a\` (vpc, us-west-1) not deleted yet: \`DependencyViolation.SecurityGroup: stub error"
check "endpoint before endpoint service" before DeleteVpcEndpoint DeleteVpcEndpointService
check "endpoint service before load balancer" before DeleteVpcEndpointService "alb DeleteLoadBalancer"
check "listener before load balancer" before "alb DeleteListener" "alb DeleteLoadBalancer"
check "load balancer before security group" before "alb DeleteLoadBalancer" DeleteSecurityGroup
check "security group before VPC" before DeleteSecurityGroup "vpc DeleteVpc "

start already-gone
fail_with DeleteVpc InvalidVpcId.NotFound
respond DescribeVpcs '{"Vpcs":{"Vpc":[]}}'
targets "us-west-1 vpc-gone"
check "a missing resource is not a failure" status_is 0
check "a missing resource is reported as such" output_has "\`vpc-gone\` (vpc) does not exist in us-west-1; nothing to delete"
check "a run that deleted nothing warns" output_has "::warning::nothing was deleted"

start check-fails
fail_with DeleteVpc InternalError
fail_with DescribeVpcs Throttling
targets "us-west-1 vpc-a"
check "a failed delete whose check also fails is a failure" status_is 1
check "it is never reported as gone" output_has "Delete accepted for 0, not found 0, failed 1."
check "the summary lists the target that was not deleted" output_has "- \`us-west-1 vpc:vpc-a\`"

start still-there
fail_with DeleteVpc IncorrectStatus
respond DescribeVpcs '{"Vpcs":{"Vpc":[{"VpcId":"vpc-a"}]}}'
targets "us-west-1 vpc-a"
check "a failed delete of an existing resource is a failure" status_is 1

start missing-key-is-not-gone
fail_with DeleteVpc InvalidVpcId.NotFound
respond DescribeVpcs '{"VPCs":{"Vpc":[]}}'
targets "us-west-1 vpc-a"
check "an existence check with a renamed key fails the target" status_is 1

start ram-role-gone
fail_with DeleteRole EntityNotExist.Role
fail_with GetRole EntityNotExist.Role
respond ListPoliciesForRole '{"Policies":{"Policy":[]}}'
targets "us-west-1 ram:test-role-1"
check "a role that is already gone is found absent from text errors" status_is 0
check "it is reported as absent" output_has "\`test-role-1\` (ram) does not exist in us-west-1"

start ram-policies-unreadable
fail_with ListPoliciesForRole Forbidden.RAM
respond GetRole '{"Role":{"RoleName":"test-role-1"}}'
targets "us-west-1 ram:test-role-1"
check "a role whose policies cannot be listed is not deleted" not_called "DeleteRole"
check "the listing error is the one reported" output_has "not deleted yet: \`Forbidden.RAM: stub error"

# Only names that every type's own rules allow reach a command.
for bad in 'oss:bkt/logs' 'oss:' 'cms:' 'sls:a"b' '-x' 'ram:-x'; do
  start "bad-name-${bad}"
  targets "us-west-1 ${bad}"
  check "\"${bad}\" is rejected" status_is 1
  check "\"${bad}\" is rejected with the rule" output_has "is not a valid id or name"
  check "nothing is called for \"${bad}\"" not_called "Delete"
  check "nothing is removed for \"${bad}\"" not_called "oss rm"
done

start not-confirmed
CONFIRM="" targets "us-west-1 vpc-a"
check "delete without CONFIRM=delete fails" status_is 1
check "it says how to confirm" output_has "set CONFIRM=delete to delete these 1 targets"
check "it deletes nothing" not_called "DeleteVpc"

# A dot in a bucket name is not a pattern.
start dotted-bucket
fail_with rm InternalError
respond ls 'oss://aXb'
targets "us-west-1 oss:a.b"
check "a dot in a bucket name matches only a dot" status_is 0

start cms-paging
fail_with DeleteContactGroup InternalError
respond DescribeContactGroupList.p1 "$(jq -nc '{ContactGroups: {ContactGroup: [range(100) | "group-\(.)"]}}')"
respond DescribeContactGroupList.p2 '{"ContactGroups":{"ContactGroup":["tf-example"]}}'
targets "us-west-1 cms:tf-example"
check "a contact group on page two is found" status_is 1
check "page two was read" called "cms DescribeContactGroupList --PageSize 100 --PageNumber 2"

start preparation-failure
fail_with UpdateHubClusterFeature Forbidden.RAM
targets "cn-beijing ackone:c123"
check "the delete still runs" called "adcp DeleteHubCluster --region cn-beijing --ClusterId c123"
check "the failed preparation is reported" output_has "could not turn off ArgoCD and deletion protection: Forbidden.RAM: stub error"

# Every type needs a delete and an existence check. The types list is checked
# against the table in the script, so the loop cannot pass by running zero
# times.
start types-list
mapfile -t type_list < <(bash "${SCRIPT}" types | awk -F'|' 'NR > 2 { gsub(/ /, "", $2); print $2 }')
type_rows=$(awk "/^TYPES='/ { in_table = 1; next } in_table && /^'/ { exit } in_table && NF" "${SCRIPT}" | wc -l | tr -d " ")
check "types lists types" [ "${#type_list[@]}" -gt 0 ]
check "types lists every type in the table (${#type_list[@]} of ${type_rows})" [ "${#type_list[@]}" -eq "${type_rows}" ]
for resource_type in "${type_list[@]}"; do
  start "type-${resource_type}"
  targets "us-west-1 ${resource_type}:x-1"
  check "${resource_type} has a delete" output_has "Delete accepted for 1, not found 0, failed 0."
  start "type-${resource_type}-check"
  touch "${case_dir}/fail-deletes"
  PASSES=1 targets "us-west-1 ${resource_type}:x-1"
  check "${resource_type} has an existence check" output_lacks "no existence check"
done

start inventory
respond DescribeVpcs '{"TotalCount": 5, "Vpcs": {"Vpc": [{"VpcId": "vpc-1", "VpcName": "a|b"}, {"VpcId": "vpc-2", "VpcName": null}]}}'
respond DescribeVSwitches '{"VSwitchSet": {}}'
fail_with DescribeSecurityGroups Forbidden.RAM
respond DescribeDBInstances '{"TotalRecordCount": 3, "Items": {"DBInstance": [{"DBInstanceId": "rm-1"}]}}'
run inventory us-west-1 "" cn-beijing
check "inventory succeeds" status_is 0
check "truncation is shown" output_has "Showing 2 of 5."
check "truncation is shown from TotalRecordCount" output_has "Showing 1 of 3."
check "a pipe in a cell is escaped" output_has '| vpc-1 | a\|b |'
check "a null cell is empty" output_has '| vpc-2 |  |'
check "a section shows a bare target for a type with its own prefix" output_has "Target: \`us-west-1 <ID>\`"
check "a renamed key is reported" output_has "Unexpected response: no \`.VSwitches.VSwitch\` in it. Top-level keys: \`VSwitchSet\`."
check "a failed list is reported" output_has "Could not list: \`Forbidden.RAM: stub error\`"
check "an empty region is skipped" [ "$(grep -c '^### ' "${case_dir}/out")" -eq 3 ]

echo
if [ "${failures}" -gt 0 ]; then
  echo "${failures} check(s) failed"
  exit 1
fi
echo "all checks passed"
