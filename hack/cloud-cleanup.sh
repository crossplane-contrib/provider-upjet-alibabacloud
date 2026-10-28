#!/usr/bin/env bash
# Lists and deletes Alibaba Cloud resources that end-to-end test runs left
# behind, for example when a run was cancelled before its cleanup finished.
# Driven by .github/workflows/cloud-cleanup.yaml; it expects an `aliyun` CLI
# that is already configured with credentials, and bash 4 or later. To run it
# locally against a named CLI profile, set ALIBABA_CLOUD_PROFILE=<profile>.
#
#   cloud-cleanup.sh inventory <region>...
#       Writes a Markdown inventory of the resource types the examples create.
#
#   CONFIRM=delete cloud-cleanup.sh delete <targets-file>
#       Deletes exactly the listed resources, children before parents, and
#       retries the ones that are still blocked by a dependency.
#
#   cloud-cleanup.sh types
#       Lists the resource types a target can name.
#
# Each target line is "<region> <id>" when the id has one of the prefixes that
# `types` lists, or "<region> <type>:<id-or-name>" otherwise. Blank lines and
# lines starting with # are ignored. Use any region for the global types. An
# id or name starts with a letter or digit and has only letters, digits, ".",
# "_" and "-": every type's names fit, and nothing else can reach a command,
# such as an OSS object path or an empty name.
#
# The output never shows the account ID, which some resource names contain,
# such as SLS's audit projects: it reads "<account>" instead, because the
# workflow publishes the output in a public repository.
set -euo pipefail

if ((BASH_VERSINFO[0] < 4)); then
  echo "$0 needs bash 4 or later; this is bash ${BASH_VERSION}" >&2
  exit 2
fi

# Passes over the remaining targets, and the wait between them, while
# dependents finish deleting.
PASSES="${PASSES:-6}"
PASS_WAIT="${PASS_WAIT:-60}"

# The resource types a target can name: type, delete rank (lower ranks are
# deleted first, so children go before what contains them), and the id
# prefix that identifies the type on its own ("-" when the type must be
# named, because its ids are names or share a prefix with another type: ALB
# and NLB listeners are both "lsn-", their server groups both "sgp-"). A
# PrivateLink endpoint goes before its service, which goes before the load
# balancer behind it.
TYPES='
alb-rule              1 rule-
alb-ascript           1 as-
privatelink-endpoint  1 ep-
alb-listener          2 -
nlb-listener          2 -
privatelink-service   2 epsrv-
alb-servergroup       3 -
nlb-servergroup       3 -
alb                   3 alb-
nlb                   3 nlb-
slb                   3 lb-
ecs                   4 i-
rds                   4 rm-
rds-readonly          3 rr-
polardb               4 pc-
tair                  4 r-
ackone                4 -
kafka                 4 -
sls                   4 -
mns                   4 -
fc                    4 -
oss                   4 -
cms                   4 -
ram                   4 -
sg                    5 sg-
vswitch               6 vsw-
vpc                   7 vpc-
'

types() {
  echo "| Type | Id prefix | Example target |"
  echo "| --- | --- | --- |"
  local resource_type prefix
  while read -r resource_type _ prefix; do
    [ -n "${resource_type}" ] || continue
    if [ "${prefix}" = - ]; then
      echo "| ${resource_type} | (name the type) | \`us-west-1 ${resource_type}:<id-or-name>\` |"
    else
      echo "| ${resource_type} | \`${prefix}\` | \`us-west-1 ${prefix}...\` |"
    fi
  done <<<"${TYPES}"
}

# type_of <target> prints the type of an "<id>" or "<type>:<id>" target, or
# fails when the type is unknown.
type_of() {
  local target="$1"
  if [[ "${target}" == *:* ]]; then
    awk -v t="${target%%:*}" '$1 == t { print $1; found = 1 } END { exit !found }' <<<"${TYPES}"
  else
    awk -v p="${target%%-*}-" '$3 != "-" && $3 == p { print $1; found = 1 } END { exit !found }' <<<"${TYPES}"
  fi
}

rank_of() {
  awk -v t="$1" '$1 == t { print $2 }' <<<"${TYPES}"
}

# The CLI prints errors as text ("ErrorCode: ..." and "Message: ..." lines,
# in color) unless it runs in its AI agent mode, where they are JSON. Both are
# parsed, so the script does not depend on which mode the CLI picks.
ESC=$'\033'

strip_color() { sed -E "s/${ESC}\[[0-9;]*m//g" <<<"$1"; }

# error_code <output> prints the API error code, or nothing for a local error.
error_code() {
  local plain code
  plain=$(strip_color "$1")
  code=$(jq -r '.error_code // empty' <<<"${plain}" 2>/dev/null) || code=""
  [ -n "${code}" ] || code=$(sed -nE 's/^ErrorCode: *([^[:space:]]+).*/\1/p' <<<"${plain}" | head -1)
  echo "${code}"
}

# describe_error <output> prints an API error as "Code: message", or the first
# line of anything else, such as a local CLI error.
describe_error() {
  local plain code message
  plain=$(strip_color "$1")
  code=$(error_code "${plain}")
  message=$(jq -r '.message // empty' <<<"${plain}" 2>/dev/null) || message=""
  [ -n "${message}" ] || message=$(sed -nE 's/^Message: *//p' <<<"${plain}" | head -1)
  [ -n "${message}" ] || message=$(grep -m1 . <<<"${plain}" || true)
  [ -z "${code}" ] || message="${code}: ${message}"
  head -1 <<<"${message}" | tr -d '`|' | cut -c1-200
}

# check_credentials fails unless the CLI can authenticate, and sets
# ACCOUNT_ID. Without it, a rotated key would fail every call in the same way
# and hide what happened. Run locally, it also says which identity is in use;
# not in CI, where the output is public.
check_credentials() {
  local out
  if ! out=$(aliyun sts GetCallerIdentity 2>&1); then
    echo "::error::cannot authenticate to Alibaba Cloud: $(describe_error "${out}")"
    return 1
  fi
  if ! ACCOUNT_ID=$(jq -er '.AccountId | strings | select(test("^[0-9]+$"))' <<<"${out}" 2>/dev/null); then
    echo "::error::cannot read the account ID from GetCallerIdentity, so it cannot be kept out of the output"
    return 1
  fi
  if [ "${GITHUB_ACTIONS:-}" != true ]; then
    echo "Using $(jq -r '.Arn // "an unknown identity"' <<<"${out}" 2>/dev/null)" >&2
  fi
}

# cells turns an array of values into a Markdown table row.
CELLS='map(if . == null then "" else tostring end | gsub("[|]"; "\\|") | gsub("\n"; " ")) | "| " + join(" | ") + " |"'

# section <title> <target> <header> <list-path> <columns> <aliyun args>...
# Lists one resource type and renders it as a Markdown table, followed by the
# form of a delete target for it. list-path must
# exist in the response: a renamed or mistyped key is reported, never shown as
# an empty list. A failed call is reported too, instead of aborting, so one
# missing permission does not hide the rest of the inventory.
section() {
  local title="$1" target="$2" header="$3" path="$4" columns="$5"
  shift 5
  local out rows shown total
  echo "#### ${title}"
  echo
  if ! out=$(aliyun "$@" 2>&1); then
    echo "Could not list: \`$(describe_error "${out}")\`"
    echo
    return
  fi
  if ! rows=$(jq -r "(${path} // error(\"missing\"))[] | ${columns} | ${CELLS}" <<<"${out}" 2>/dev/null); then
    echo "Unexpected response: no \`${path}\` in it. Top-level keys: \`$(jq -r 'keys | join(", ")' <<<"${out}" 2>/dev/null || echo "not JSON")\`."
    echo
    return
  fi
  if [ -z "${rows}" ]; then
    echo "None."
    echo
    return
  fi
  echo "| ${header} |"
  echo "|$(sed -E 's/[^|]+/ --- /g' <<<"${header}")|"
  echo "${rows}"
  echo
  echo "Target: \`${target}\`"
  echo
  shown=$(wc -l <<<"${rows}" | tr -d ' ')
  total=$(jq -r '.TotalCount // .totalCount // .total // .Total // .Data.Total // .TotalRecordCount // empty' <<<"${out}" 2>/dev/null || true)
  if [ -n "${total}" ] && [ "${total}" -gt "${shown}" ]; then
    echo "Showing ${shown} of ${total}."
    echo
  elif jq -e '((.NextToken // .nextToken // "") != "") or (.IsTruncated == true)' <<<"${out}" >/dev/null 2>&1; then
    echo "Showing the first ${shown}; there are more."
    echo
  fi
}

inventory() {
  echo "## Alibaba Cloud inventory"
  echo
  echo "Generated $(date -u +%Y-%m-%dT%H:%MZ)."
  echo
  local region
  for region in "$@"; do
    [ -n "${region}" ] || continue
    echo "### ${region}"
    echo
    section "${region} VPCs" "${region} <ID>" \
      "ID | Name | Status | Created" \
      '.Vpcs.Vpc' '[.VpcId, .VpcName, .Status, .CreationTime]' \
      vpc DescribeVpcs --RegionId "${region}" --PageSize 50
    section "${region} vSwitches" "${region} <ID>" \
      "ID | Name | VPC | Zone | Created" \
      '.VSwitches.VSwitch' '[.VSwitchId, .VSwitchName, .VpcId, .ZoneId, .CreationTime]' \
      vpc DescribeVSwitches --RegionId "${region}" --PageSize 50
    section "${region} security groups" "${region} <ID>" \
      "ID | Name | Description | VPC | Service managed | Created" \
      '.SecurityGroups.SecurityGroup' '[.SecurityGroupId, .SecurityGroupName, .Description, .VpcId, .ServiceManaged, .CreationTime]' \
      ecs DescribeSecurityGroups --RegionId "${region}" --PageSize 50
    section "${region} ECS instances" "${region} <ID>" \
      "ID | Name | Status | VPC | Created" \
      '.Instances.Instance' '[.InstanceId, .InstanceName, .Status, .VpcAttributes.VpcId, .CreationTime]' \
      ecs DescribeInstances --RegionId "${region}" --PageSize 100
    section "${region} ALB load balancers" "${region} <ID>" \
      "ID | Name | Status | VPC | Deletion protection | Created" \
      '.LoadBalancers' '[.LoadBalancerId, .LoadBalancerName, .LoadBalancerStatus, .VpcId, .DeletionProtectionConfig.Enabled, .CreateTime]' \
      alb ListLoadBalancers --region "${region}" --MaxResults 100
    section "${region} ALB listeners" "${region} alb-listener:<ID>" \
      "ID | Load balancer | Protocol | Port" \
      '.Listeners' '[.ListenerId, .LoadBalancerId, .ListenerProtocol, .ListenerPort]' \
      alb ListListeners --region "${region}" --MaxResults 100
    section "${region} ALB rules" "${region} <ID>" \
      "ID | Listener | Name" \
      '.Rules' '[.RuleId, .ListenerId, .RuleName]' \
      alb ListRules --region "${region}" --MaxResults 100
    section "${region} ALB server groups" "${region} alb-servergroup:<ID>" \
      "ID | Name | VPC" \
      '.ServerGroups' '[.ServerGroupId, .ServerGroupName, .VpcId]' \
      alb ListServerGroups --region "${region}" --MaxResults 100
    section "${region} ALB AScripts" "${region} <ID>" \
      "ID | Name | Listener" \
      '.AScripts' '[.AScriptId, .AScriptName, .ListenerId]' \
      alb ListAScripts --region "${region}" --MaxResults 100
    section "${region} NLB load balancers" "${region} <ID>" \
      "ID | Name | Status | VPC | Created" \
      '.LoadBalancers' '[.LoadBalancerId, .LoadBalancerName, .LoadBalancerStatus, .VpcId, .CreateTime]' \
      nlb ListLoadBalancers --RegionId "${region}" --MaxResults 100
    section "${region} NLB listeners" "${region} nlb-listener:<ID>" \
      "ID | Load balancer | Protocol | Port" \
      '.Listeners' '[.ListenerId, .LoadBalancerId, .ListenerProtocol, .ListenerPort]' \
      nlb ListListeners --RegionId "${region}" --MaxResults 100
    section "${region} NLB server groups" "${region} nlb-servergroup:<ID>" \
      "ID | Name | VPC" \
      '.ServerGroups' '[.ServerGroupId, .ServerGroupName, .VpcId]' \
      nlb ListServerGroups --RegionId "${region}" --MaxResults 100
    section "${region} SLB load balancers" "${region} <ID>" \
      "ID | Name | Status | VPC | Created" \
      '.LoadBalancers.LoadBalancer' '[.LoadBalancerId, .LoadBalancerName, .LoadBalancerStatus, .VpcId, .CreateTime]' \
      slb DescribeLoadBalancers --RegionId "${region}" --PageSize 100
    section "${region} PrivateLink endpoint services" "${region} <ID>" \
      "ID | Name | Status | Created" \
      '.Services' '[.ServiceId, .ServiceName, .ServiceStatus, .CreateTime]' \
      privatelink ListVpcEndpointServices --RegionId "${region}" --MaxResults 100
    section "${region} PrivateLink endpoints" "${region} <ID>" \
      "ID | Name | Status | VPC | Created" \
      '.Endpoints' '[.EndpointId, .EndpointName, .EndpointStatus, .VpcId, .CreateTime]' \
      privatelink ListVpcEndpoints --RegionId "${region}" --MaxResults 100
    section "${region} RDS instances" "${region} <ID>" \
      "ID | Description | Engine | Status | VPC | Created" \
      '.Items.DBInstance' '[.DBInstanceId, .DBInstanceDescription, .Engine, .DBInstanceStatus, .VpcId, .CreateTime]' \
      rds DescribeDBInstances --RegionId "${region}" --PageSize 100
    section "${region} PolarDB clusters" "${region} <ID>" \
      "ID | Description | Type | Status | VPC | Created" \
      '.Items.DBCluster' '[.DBClusterId, .DBClusterDescription, .DBType, .DBClusterStatus, .VpcId, .CreateTime]' \
      polardb DescribeDBClusters --RegionId "${region}" --PageSize 100
    section "${region} Tair (Redis) instances" "${region} <ID>" \
      "ID | Name | Status | VPC | Created" \
      '.Instances.KVStoreInstance' '[.InstanceId, .InstanceName, .InstanceStatus, .VpcId, .CreateTime]' \
      r-kvstore DescribeInstances --RegionId "${region}" --PageSize 50
    section "${region} ACK One hub clusters" "${region} ackone:<ID>" \
      "ID | Name | State | VPC | Created" \
      '.Clusters' '[.ClusterInfo.ClusterId, .ClusterInfo.Name, .ClusterInfo.State, .Network.VpcId, .ClusterInfo.CreationTime]' \
      adcp DescribeHubClusters --region "${region}"
    section "${region} Kafka instances" "${region} kafka:<ID>" \
      "ID | Name | Status | Deploy type | VPC | Created" \
      '.InstanceList.InstanceVO' '[.InstanceId, .Name, .ServiceStatus, .DeployType, .VpcId, .CreateTime]' \
      alikafka GetInstanceList --RegionId "${region}"
    section "${region} SLS projects" "${region} sls:<name>" \
      "Name | Status | Created" \
      '.projects' '[.projectName, .status, .createTime]' \
      sls ListProject --region "${region}" --size 100
    section "${region} MNS queues" "${region} mns:<name>" \
      "Name | Created" \
      '.Data.PageData' '[.QueueName, .CreateTime]' \
      mns-open ListQueue --region "${region}" --PageSize 50
    section "${region} FC functions" "${region} fc:<name>" \
      "Name | Runtime | Created" \
      '.functions' '[.functionName, .runtime, .createdTime]' \
      fc ListFunctions --region "${region}" --limit 100
  done
  echo "### Global"
  echo
  echo "#### OSS buckets"
  echo
  local buckets
  if buckets=$(aliyun oss ls -s 2>&1); then
    if grep '^oss://' <<<"${buckets}" | sed 's/^/- /' | grep .; then
      echo
      echo "Target: \`us-west-1 oss:<name>\`, the name without oss://"
    else
      echo "None."
    fi
  else
    echo "Could not list: \`$(describe_error "${buckets}")\`"
  fi
  echo
  section "RAM roles" "us-west-1 ram:<name>" \
    "Name | Description | Created" \
    '.Roles.Role' '[.RoleName, .Description, .CreateDate]' \
    ram ListRoles --MaxItems 100
  section "CloudMonitor contact groups" "us-west-1 cms:<name>" \
    "Name" \
    '.ContactGroups.ContactGroup' '[.]' \
    cms DescribeContactGroupList --PageSize 100
}

# count <id> <list-path> <aliyun args>... prints how many items the call
# returns at list-path, which refers to the id as $id. It fails, rather than
# printing 0, when the call fails or the path is missing, so a broken check
# never reads as "already gone".
count() {
  local id="$1" path="$2" out
  shift 2
  out=$(aliyun "$@" 2>&1) || return 1
  jq --arg id "${id}" "(${path} // error(\"missing\")) | length" <<<"${out}" 2>/dev/null
}

# remaining <region> <type> <id> prints how many resources with that id still
# exist in the region (0 or 1), or fails when it cannot tell. Ids reach jq and
# grep as values, never as part of a program or pattern.
remaining() {
  local region="$1" resource_type="$2" id="$3"
  # shellcheck disable=SC2016 # $id in the list paths is jq's, not the shell's.
  case "${resource_type}" in
    vpc) count "${id}" '.Vpcs.Vpc' vpc DescribeVpcs --RegionId "${region}" --VpcId "${id}" ;;
    vswitch) count "${id}" '.VSwitches.VSwitch' vpc DescribeVSwitches --RegionId "${region}" --VSwitchId "${id}" ;;
    sg) count "${id}" '.SecurityGroups.SecurityGroup' ecs DescribeSecurityGroups --RegionId "${region}" --SecurityGroupId "${id}" ;;
    ecs) count "${id}" '.Instances.Instance' ecs DescribeInstances --RegionId "${region}" --InstanceIds "$(jq -nc --arg id "${id}" '[$id]')" ;;
    alb) count "${id}" '.LoadBalancers' alb ListLoadBalancers --region "${region}" --LoadBalancerIds.1 "${id}" ;;
    alb-listener) count "${id}" '.Listeners' alb ListListeners --region "${region}" --ListenerIds.1 "${id}" ;;
    alb-rule) count "${id}" '.Rules' alb ListRules --region "${region}" --RuleIds.1 "${id}" ;;
    alb-servergroup) count "${id}" '.ServerGroups' alb ListServerGroups --region "${region}" --ServerGroupIds.1 "${id}" ;;
    alb-ascript) count "${id}" '.AScripts' alb ListAScripts --region "${region}" --AScriptIds.1 "${id}" ;;
    nlb) count "${id}" '.LoadBalancers' nlb ListLoadBalancers --RegionId "${region}" --LoadBalancerIds.1 "${id}" ;;
    nlb-listener) count "${id}" '.Listeners' nlb ListListeners --RegionId "${region}" --ListenerIds.1 "${id}" ;;
    nlb-servergroup) count "${id}" '.ServerGroups' nlb ListServerGroups --RegionId "${region}" --ServerGroupIds.1 "${id}" ;;
    slb) count "${id}" '.LoadBalancers.LoadBalancer' slb DescribeLoadBalancers --RegionId "${region}" --LoadBalancerId "${id}" ;;
    privatelink-service) count "${id}" '.Services' privatelink ListVpcEndpointServices --RegionId "${region}" --ServiceId "${id}" ;;
    privatelink-endpoint) count "${id}" '.Endpoints' privatelink ListVpcEndpoints --RegionId "${region}" --EndpointId "${id}" ;;
    rds | rds-readonly) count "${id}" '.Items.DBInstance' rds DescribeDBInstances --RegionId "${region}" --DBInstanceId "${id}" ;;
    polardb) count "${id}" '.Items.DBCluster' polardb DescribeDBClusters --RegionId "${region}" --DBClusterIds "${id}" ;;
    tair) count "${id}" '.Instances.KVStoreInstance' r-kvstore DescribeInstances --RegionId "${region}" --InstanceIds "${id}" ;;
    ackone) count "${id}" '[(.Clusters // error("missing"))[] | select(.ClusterInfo.ClusterId == $id)]' adcp DescribeHubClusters --region "${region}" ;;
    kafka) count "${id}" '.InstanceList.InstanceVO' alikafka GetInstanceList --RegionId "${region}" --InstanceId.1 "${id}" ;;
    sls) count "${id}" '[(.projects // error("missing"))[] | select(.projectName == $id)]' sls ListProject --region "${region}" --projectName "${id}" ;;
    mns) count "${id}" '[(.Data.PageData // error("missing"))[] | select(.QueueName == $id)]' mns-open ListQueue --region "${region}" --QueueName "${id}" ;;
    fc) count "${id}" '[(.functions // error("missing"))[] | select(.functionName == $id)]' fc ListFunctions --region "${region}" --prefix "${id}" ;;
    cms)
      # The list cannot be filtered by name, so read every page.
      local page=1 out
      while :; do
        out=$(aliyun cms DescribeContactGroupList --PageSize 100 --PageNumber "${page}" 2>&1) || return 1
        jq -e --arg id "${id}" '(.ContactGroups.ContactGroup // error("missing")) | index($id) != null' <<<"${out}" >/dev/null 2>&1 &&
          { echo 1; return; }
        [ "$(jq '.ContactGroups.ContactGroup | length' <<<"${out}" 2>/dev/null || echo error)" = 100 ] ||
          break
        page=$((page + 1))
      done
      # Distinguish "not on any page" from a broken response.
      jq -e '.ContactGroups.ContactGroup | arrays' <<<"${out}" >/dev/null 2>&1 || return 1
      echo 0
      ;;
    oss)
      local buckets
      buckets=$(aliyun oss ls -s 2>&1) || return 1
      grep -cxF "oss://${id}" <<<"${buckets}" || true
      ;;
    ram)
      local out
      if out=$(aliyun ram GetRole --RoleName "${id}" 2>&1); then
        echo 1
      elif [ "$(error_code "${out}")" = "EntityNotExist.Role" ]; then
        echo 0
      else
        return 1
      fi
      ;;
    *)
      echo "no existence check for type ${resource_type}" >&2
      return 1
      ;;
  esac
}

# prepare <description> <aliyun args>... runs a step that makes a delete
# possible, such as turning off deletion protection. A failure is reported but
# does not stop the delete, whose own error then has the context it needs.
prepare() {
  local description="$1" out
  shift
  if ! out=$(aliyun "$@" 2>&1); then
    echo "(could not ${description}: $(describe_error "${out}")) "
  fi
}

# delete_one <region> <type> <id> deletes one resource. On success it prints
# any warnings from the preparation steps; on failure, those and the API's
# error.
delete_one() {
  local region="$1" resource_type="$2" id="$3" result=""
  case "${resource_type}" in
    alb-rule) aliyun alb DeleteRule --region "${region}" --RuleId "${id}" ;;
    alb-ascript) aliyun alb DeleteAScripts --region "${region}" --AScriptIds.1 "${id}" ;;
    alb-listener) aliyun alb DeleteListener --region "${region}" --ListenerId "${id}" ;;
    alb-servergroup) aliyun alb DeleteServerGroup --region "${region}" --ServerGroupId "${id}" ;;
    alb)
      result=$(prepare "turn off deletion protection" alb DisableDeletionProtection --region "${region}" --ResourceId "${id}")
      aliyun alb DeleteLoadBalancer --region "${region}" --LoadBalancerId "${id}"
      ;;
    nlb-listener) aliyun nlb DeleteListener --RegionId "${region}" --ListenerId "${id}" ;;
    nlb-servergroup) aliyun nlb DeleteServerGroup --RegionId "${region}" --ServerGroupId "${id}" ;;
    nlb) aliyun nlb DeleteLoadBalancer --RegionId "${region}" --LoadBalancerId "${id}" ;;
    slb)
      result=$(prepare "turn off deletion protection" slb SetLoadBalancerDeleteProtection --RegionId "${region}" --LoadBalancerId "${id}" --DeleteProtection off)
      aliyun slb DeleteLoadBalancer --RegionId "${region}" --LoadBalancerId "${id}"
      ;;
    privatelink-endpoint) aliyun privatelink DeleteVpcEndpoint --RegionId "${region}" --EndpointId "${id}" ;;
    privatelink-service) aliyun privatelink DeleteVpcEndpointService --RegionId "${region}" --ServiceId "${id}" ;;
    ecs) aliyun ecs DeleteInstance --region "${region}" --InstanceId "${id}" --Force true ;;
    rds | rds-readonly)
      result=$(prepare "turn off deletion protection" rds ModifyDBInstanceDeletionProtection --region "${region}" --DBInstanceId "${id}" --DeletionProtection false)
      aliyun rds DeleteDBInstance --region "${region}" --DBInstanceId "${id}"
      ;;
    polardb)
      result=$(prepare "turn off deletion protection" polardb ModifyDBClusterDeletion --region "${region}" --DBClusterId "${id}" --Protection false)
      aliyun polardb DeleteDBCluster --region "${region}" --DBClusterId "${id}"
      ;;
    tair) aliyun r-kvstore DeleteInstance --region "${region}" --InstanceId "${id}" ;;
    ackone)
      # DeleteHubCluster refuses while ArgoCD is enabled, which is the default.
      result=$(prepare "turn off ArgoCD and deletion protection" adcp UpdateHubClusterFeature --region "${region}" \
        --ClusterId "${id}" --ArgoCDEnabled false --DeletionProtection false)
      aliyun adcp DeleteHubCluster --region "${region}" --ClusterId "${id}"
      ;;
    kafka)
      # Released instances stay listed until they are deleted as well.
      result=$(prepare "release the instance" alikafka ReleaseInstance --RegionId "${region}" --InstanceId "${id}" --ForceDeleteInstance true)
      aliyun alikafka DeleteInstance --RegionId "${region}" --InstanceId "${id}"
      ;;
    sls) aliyun sls DeleteProject --region "${region}" --project "${id}" ;;
    mns) aliyun mns-open DeleteQueue --region "${region}" --QueueName "${id}" ;;
    fc) aliyun fc DeleteFunction --region "${region}" --functionName "${id}" ;;
    oss) aliyun oss rm "oss://${id}" --region "${region}" -b -a -r -f --all-versions ;;
    cms) aliyun cms DeleteContactGroup --ContactGroupName "${id}" ;;
    ram)
      # A role with policies attached cannot be deleted. If listing them
      # fails, that error is the one to report, not DeleteRole's.
      local policies attached name policy_type
      if policies=$(aliyun ram ListPoliciesForRole --RoleName "${id}") &&
        attached=$(jq -er '(.Policies.Policy // error("missing"))[] | "\(.PolicyName) \(.PolicyType)"' <<<"${policies}" ||
          [ $? -eq 4 ]); then
        while read -r name policy_type; do
          [ -n "${name}" ] || continue
          result+=$(prepare "detach policy ${name}" ram DetachPolicyFromRole --RoleName "${id}" --PolicyName "${name}" --PolicyType "${policy_type}")
        done <<<"${attached}"
        aliyun ram DeleteRole --RoleName "${id}"
      else
        false
      fi
      ;;
    sg) aliyun ecs DeleteSecurityGroup --RegionId "${region}" --SecurityGroupId "${id}" ;;
    vswitch) aliyun vpc DeleteVSwitch --RegionId "${region}" --VSwitchId "${id}" ;;
    vpc) aliyun vpc DeleteVpc --RegionId "${region}" --VpcId "${id}" ;;
    *)
      echo "no delete for type ${resource_type}" >&2
      return 1
      ;;
  esac >"${ERR_FILE}" 2>&1 || {
    echo "${result}$(describe_error "$(cat "${ERR_FILE}")")"
    return 1
  }
  printf '%s' "${result}"
}

delete() {
  local targets=() failed=() line region target extra resource_type id pass result left
  local accepted=0 absent=0
  while IFS= read -r line || [ -n "${line}" ]; do
    line="$(sed -E 's/#.*//; s/^[[:space:]]+//; s/[[:space:]]+$//' <<<"${line}")"
    [ -z "${line}" ] && continue
    region="" target="" extra=""
    read -r region target extra <<<"${line}"
    if [ -z "${target}" ] || [ -n "${extra}" ]; then
      echo "::error::target \"${line}\" needs exactly a region and an id"
      return 1
    fi
    if ! [[ "${target#*:}" =~ ^[A-Za-z0-9][A-Za-z0-9._-]*$ ]]; then
      echo "::error::\"${target}\" is not a valid id or name: it must start with a letter or digit, and have only letters, digits, \".\", \"_\" and \"-\""
      return 1
    fi
    if ! resource_type="$(type_of "${target}")"; then
      echo "::error::unknown resource type in \"${target}\"; run \"$0 types\" for the list"
      return 1
    fi
    targets+=("$(rank_of "${resource_type}") ${region} ${resource_type} ${target#*:}")
  done <"$1"
  if [ "${#targets[@]}" -eq 0 ]; then
    echo "::error::no targets given"
    return 1
  fi
  if [ "${CONFIRM:-}" != delete ]; then
    echo "::error::set CONFIRM=delete to delete these ${#targets[@]} targets"
    return 1
  fi
  mapfile -t targets < <(printf '%s\n' "${targets[@]}" | sort -n -k1,1)
  ERR_FILE="$(mktemp)"
  trap 'rm -f "${ERR_FILE}"' EXIT

  echo "## Alibaba Cloud cleanup"
  echo
  for ((pass = 1; pass <= PASSES; pass++)); do
    failed=()
    for line in "${targets[@]}"; do
      read -r _ region resource_type id <<<"${line}"
      if result=$(delete_one "${region}" "${resource_type}" "${id}"); then
        accepted=$((accepted + 1))
        echo "- delete accepted for \`${id}\` (${resource_type}, ${region}) ${result}"
        continue
      fi
      if left=$(remaining "${region}" "${resource_type}" "${id}") && [ "${left}" = 0 ]; then
        absent=$((absent + 1))
        echo "- \`${id}\` (${resource_type}) does not exist in ${region}; nothing to delete"
      else
        failed+=("${line}")
        echo "- pass ${pass}: \`${id}\` (${resource_type}, ${region}) not deleted yet: \`${result}\`"
      fi
    done
    [ "${#failed[@]}" -eq 0 ] && break
    targets=("${failed[@]}")
    if ((pass < PASSES)); then
      sleep "${PASS_WAIT}"
    fi
  done
  echo
  echo "**Delete accepted for ${accepted}, not found ${absent}, failed ${#failed[@]}.**"
  if [ "${accepted}" -gt 0 ]; then
    echo
    echo "Some types, such as load balancers, databases and clusters, finish deleting in the background. The next inventory shows whether they are gone."
  fi
  if [ "${#failed[@]}" -gt 0 ]; then
    echo
    echo "Not deleted:"
    for line in "${failed[@]}"; do
      read -r _ region resource_type id <<<"${line}"
      echo "- \`${region} ${resource_type}:${id}\`"
    done
    return 1
  fi
  if [ "${accepted}" -eq 0 ]; then
    echo "::warning::nothing was deleted: none of the targets exist in the regions given"
  fi
}

# redact replaces the account ID in its input with "<account>".
redact() { sed "s/${ACCOUNT_ID}/<account>/g"; }

case "${1:-}" in
  inventory)
    shift
    check_credentials
    inventory "$@" | redact
    ;;
  delete)
    targets_file="${2:?usage: $0 delete <targets-file>}"
    check_credentials
    delete "${targets_file}" | redact
    ;;
  types)
    types
    ;;
  *)
    echo "usage: $0 inventory <region>... | delete <targets-file> | types" >&2
    exit 2
    ;;
esac
