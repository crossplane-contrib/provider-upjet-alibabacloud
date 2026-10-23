#!/usr/bin/env bash
# Lists and deletes Alibaba Cloud resources that end-to-end test runs left
# behind, for example when a run was cancelled before its cleanup finished.
# Driven by .github/workflows/cloud-cleanup.yaml; it expects an `aliyun` CLI
# that is already configured with credentials.
#
#   cloud-cleanup.sh inventory <region>...
#       Writes a Markdown inventory of the resource types the examples create
#       to stdout, and every raw API response to $RAW_DIR.
#
#   cloud-cleanup.sh delete <targets-file>
#       Deletes exactly the listed resources, children before parents, and
#       retries the ones that are still blocked by a dependency.
#
# Each target line is "<region> <id>" for ids with a recognizable prefix
# (vpc-, vsw-, sg-, i-, alb-, lsn-, rule-, sgp-, as-, epsrv-), or
# "<region> <type>:<name>" otherwise, where type is one of ackone, kafka, sls,
# mns, fc, oss, cms or ram. Blank lines and lines starting with # are ignored.
# Use any region for the global types (cms, ram).
set -euo pipefail

RAW_DIR="${RAW_DIR:-_output/cloud-cleanup}"
# Passes over the remaining targets, and the wait between them, while
# dependents finish deleting.
PASSES="${PASSES:-6}"
PASS_WAIT="${PASS_WAIT:-60}"

# Errors that mean the resource is already gone.
GONE_PATTERN='NotFound|NotExist|not exist|not found|NoSuchBucket|EntityNotExist'

# first_error <output> prints the most useful line of an API error.
first_error() {
  { grep -m1 -E 'Code|Message|ERROR' <<<"$1" || head -1 <<<"$1"; } | tr -d '`' | cut -c1-200
}

# section <title> <header> <jq-rows> <aliyun args>...
# Calls the API, keeps the raw response, and renders the rows the jq filter
# extracts as a Markdown table. A failed call is reported instead of aborting,
# so one missing permission does not hide the rest of the inventory.
section() {
  local title="$1" header="$2" rows="$3"
  shift 3
  local file out table
  file="${RAW_DIR}/$(tr ' /' '__' <<<"${title}").json"
  echo "#### ${title}"
  echo
  if ! out=$(aliyun "$@" 2>&1); then
    echo "Could not list: \`$(first_error "${out}")\`"
    echo
    return
  fi
  echo "${out}" >"${file}"
  if ! table=$(jq -r "${rows} | map(tostring) | \"| \" + join(\" | \") + \" |\"" <<<"${out}" 2>/dev/null); then
    echo "Unexpected response shape; see \`$(basename "${file}")\` in the workflow artifact."
    echo
    return
  fi
  if [ -z "${table}" ]; then
    echo "None."
  else
    echo "| ${header} |"
    echo "|$(sed -E 's/[^|]+/ --- /g' <<<"${header}")|"
    echo "${table}"
  fi
  echo
}

inventory() {
  mkdir -p "${RAW_DIR}"
  echo "## Alibaba Cloud inventory"
  echo
  echo "Generated $(date -u +%Y-%m-%dT%H:%MZ). Each list shows at most one page (50-100 items)."
  echo
  local r
  for r in "$@"; do
    echo "### ${r}"
    echo
    section "${r} VPCs" "ID | Name | Status | Created" \
      '.Vpcs.Vpc[]? | [.VpcId, .VpcName, .Status, .CreationTime]' \
      vpc DescribeVpcs --RegionId "${r}" --PageSize 50
    section "${r} vSwitches" "ID | Name | VPC | Zone | Created" \
      '.VSwitches.VSwitch[]? | [.VSwitchId, .VSwitchName, .VpcId, .ZoneId, .CreationTime]' \
      vpc DescribeVSwitches --RegionId "${r}" --PageSize 50
    section "${r} security groups" "ID | Name | Description | VPC | Service managed | Created" \
      '.SecurityGroups.SecurityGroup[]? | [.SecurityGroupId, .SecurityGroupName, .Description, .VpcId, .ServiceManaged, .CreationTime]' \
      ecs DescribeSecurityGroups --RegionId "${r}" --PageSize 50
    section "${r} ECS instances" "ID | Name | Status | VPC | Created" \
      '.Instances.Instance[]? | [.InstanceId, .InstanceName, .Status, .VpcAttributes.VpcId, .CreationTime]' \
      ecs DescribeInstances --RegionId "${r}" --PageSize 100
    section "${r} ALB load balancers" "ID | Name | Status | VPC | Deletion protection | Created" \
      '.LoadBalancers[]? | [.LoadBalancerId, .LoadBalancerName, .LoadBalancerStatus, .VpcId, .DeletionProtectionConfig.Enabled, .CreateTime]' \
      alb ListLoadBalancers --region "${r}" --MaxResults 100
    section "${r} ALB listeners" "ID | Load balancer | Protocol | Port" \
      '.Listeners[]? | [.ListenerId, .LoadBalancerId, .ListenerProtocol, .ListenerPort]' \
      alb ListListeners --region "${r}" --MaxResults 100
    section "${r} ALB rules" "ID | Listener | Name" \
      '.Rules[]? | [.RuleId, .ListenerId, .RuleName]' \
      alb ListRules --region "${r}" --MaxResults 100
    section "${r} ALB server groups" "ID | Name | VPC" \
      '.ServerGroups[]? | [.ServerGroupId, .ServerGroupName, .VpcId]' \
      alb ListServerGroups --region "${r}" --MaxResults 100
    section "${r} ALB AScripts" "ID | Name | Listener" \
      '.AScripts[]? | [.AScriptId, .AScriptName, .ListenerId]' \
      alb ListAScripts --region "${r}" --MaxResults 100
    section "${r} PrivateLink endpoint services" "ID | Name | Status | Created" \
      '.Services[]? | [.ServiceId, .ServiceName, .ServiceStatus, .CreateTime]' \
      privatelink ListVpcEndpointServices --RegionId "${r}" --MaxResults 100
    section "${r} ACK One hub clusters" "ID | Name | State | VPC | Created" \
      '.Clusters[]? | [.ClusterInfo.ClusterId, .ClusterInfo.Name, .ClusterInfo.State, .Network.VpcId, .ClusterInfo.CreationTime]' \
      adcp DescribeHubClusters --region "${r}"
    section "${r} Kafka instances" "ID | Name | Status | Deploy type | VPC | Created" \
      '.InstanceList.InstanceVO[]? | [.InstanceId, .Name, .ServiceStatus, .DeployType, .VpcId, .CreateTime]' \
      alikafka GetInstanceList --RegionId "${r}"
    section "${r} SLS projects" "Name | Status | Created" \
      '.projects[]? | [.projectName, .status, .createTime]' \
      sls ListProject --region "${r}" --size 100
    section "${r} MNS queues" "Name | Created" \
      '.Data.PageData[]? | [.QueueName, .CreateTime]' \
      mns-open ListQueue --region "${r}" --PageSize 50
    section "${r} FC functions" "Name | Runtime | Created" \
      '.functions[]? | [.functionName, .runtime, .createdTime]' \
      fc ListFunctions --region "${r}" --limit 100
  done
  echo "### Global"
  echo
  echo "#### OSS buckets"
  echo
  local buckets
  if buckets=$(aliyun oss ls -s 2>&1); then
    grep '^oss://' <<<"${buckets}" | sed 's/^/- /' || echo "None."
  else
    echo "Could not list: \`$(first_error "${buckets}")\`"
  fi
  echo
  section "RAM roles" "Name | Description | Created" \
    '.Roles.Role[]? | [.RoleName, .Description, .CreateDate]' \
    ram ListRoles --MaxItems 100
  section "CloudMonitor contact groups" "Name" \
    '.ContactGroups.ContactGroup[]? | [.]' \
    cms DescribeContactGroupList --PageSize 100
}

# type_of <id> prints the resource type for an id or "<type>:<name>" target.
type_of() {
  case "$1" in
    *:*) echo "${1%%:*}" ;;
    vpc-*) echo vpc ;;
    vsw-*) echo vswitch ;;
    sg-*) echo sg ;;
    i-*) echo ecs ;;
    alb-*) echo alb ;;
    lsn-*) echo listener ;;
    rule-*) echo rule ;;
    sgp-*) echo servergroup ;;
    as-*) echo ascript ;;
    epsrv-*) echo privatelink ;;
    *) echo unknown ;;
  esac
}

# rank <type> orders deletes so that children go before what contains them.
rank() {
  case "$1" in
    rule | ascript) echo 1 ;;
    listener) echo 2 ;;
    alb | servergroup) echo 3 ;;
    ecs | ackone | kafka | privatelink | sls | mns | fc | oss | cms | ram) echo 4 ;;
    sg) echo 5 ;;
    vswitch) echo 6 ;;
    vpc) echo 7 ;;
    *) echo 9 ;;
  esac
}

# delete_one <region> <type> <target> deletes one resource; it fails with the
# API's error output when the delete was refused.
delete_one() {
  local r="$1" t="$2" id="${3#*:}"
  case "${t}" in
    rule) aliyun alb DeleteRule --region "${r}" --RuleId "${id}" ;;
    ascript) aliyun alb DeleteAScripts --region "${r}" --AScriptIds.1 "${id}" ;;
    listener) aliyun alb DeleteListener --region "${r}" --ListenerId "${id}" ;;
    servergroup) aliyun alb DeleteServerGroup --region "${r}" --ServerGroupId "${id}" ;;
    alb)
      aliyun alb DisableDeletionProtection --region "${r}" --ResourceId "${id}" >/dev/null 2>&1 || true
      aliyun alb DeleteLoadBalancer --region "${r}" --LoadBalancerId "${id}"
      ;;
    ecs) aliyun ecs DeleteInstance --region "${r}" --InstanceId "${id}" --Force true ;;
    sg) aliyun ecs DeleteSecurityGroup --RegionId "${r}" --SecurityGroupId "${id}" ;;
    vswitch) aliyun vpc DeleteVSwitch --RegionId "${r}" --VSwitchId "${id}" ;;
    vpc) aliyun vpc DeleteVpc --RegionId "${r}" --VpcId "${id}" ;;
    privatelink) aliyun privatelink DeleteVpcEndpointService --RegionId "${r}" --ServiceId "${id}" ;;
    ackone)
      # DeleteHubCluster refuses while ArgoCD is enabled, which is the default.
      aliyun adcp UpdateHubClusterFeature --region "${r}" --ClusterId "${id}" \
        --ArgoCDEnabled false --DeletionProtection false >/dev/null 2>&1 || true
      aliyun adcp DeleteHubCluster --region "${r}" --ClusterId "${id}"
      ;;
    kafka)
      # Released instances stay listed until they are deleted as well.
      aliyun alikafka ReleaseInstance --RegionId "${r}" --InstanceId "${id}" --ForceDeleteInstance true >/dev/null 2>&1 || true
      aliyun alikafka DeleteInstance --RegionId "${r}" --InstanceId "${id}"
      ;;
    sls) aliyun sls DeleteProject --region "${r}" --project "${id}" ;;
    mns) aliyun mns-open DeleteQueue --region "${r}" --QueueName "${id}" ;;
    fc) aliyun fc DeleteFunction --region "${r}" --functionName "${id}" ;;
    oss) aliyun oss rm "oss://${id}" --region "${r}" -b -a -r -f ;;
    cms) aliyun cms DeleteContactGroup --ContactGroupName "${id}" ;;
    ram)
      local policies name ptype
      policies=$(aliyun ram ListPoliciesForRole --RoleName "${id}")
      while read -r name ptype; do
        [ -n "${name}" ] || continue
        aliyun ram DetachPolicyFromRole --RoleName "${id}" --PolicyName "${name}" --PolicyType "${ptype}" >/dev/null
      done < <(jq -r '.Policies.Policy[]? | "\(.PolicyName) \(.PolicyType)"' <<<"${policies}")
      aliyun ram DeleteRole --RoleName "${id}"
      ;;
    *)
      echo "unknown resource type for $3"
      return 1
      ;;
  esac
}

delete() {
  local targets=() remaining=() line r id t pass out
  while IFS= read -r line || [ -n "${line}" ]; do
    line="$(sed -E 's/#.*//; s/^[[:space:]]+//; s/[[:space:]]+$//' <<<"${line}")"
    [ -z "${line}" ] && continue
    r="" id=""
    read -r r id <<<"${line}"
    if [ -z "${id}" ]; then
      echo "::error::target \"${line}\" needs a region and an id"
      return 1
    fi
    t="$(type_of "${id}")"
    if [ "${t}" = unknown ]; then
      echo "::error::cannot tell the type of \"${id}\"; write it as <type>:<name>"
      return 1
    fi
    targets+=("$(rank "${t}") ${r} ${t} ${id}")
  done <"$1"
  if [ "${#targets[@]}" -eq 0 ]; then
    echo "::error::no targets given"
    return 1
  fi
  mapfile -t targets < <(printf '%s\n' "${targets[@]}" | sort -n -k1,1)

  echo "## Alibaba Cloud cleanup"
  echo
  for ((pass = 1; pass <= PASSES; pass++)); do
    remaining=()
    for line in "${targets[@]}"; do
      read -r _ r t id <<<"${line}"
      if out=$(delete_one "${r}" "${t}" "${id}" 2>&1); then
        echo "- deleted \`${id}\` (${t}, ${r})"
      elif grep -qE "${GONE_PATTERN}" <<<"${out}"; then
        echo "- \`${id}\` (${t}, ${r}) was already gone"
      else
        remaining+=("${line}")
        echo "- pass ${pass}: \`${id}\` (${t}, ${r}) not deleted yet: \`$(first_error "${out}")\`"
      fi
    done
    [ "${#remaining[@]}" -eq 0 ] && break
    targets=("${remaining[@]}")
    if ((pass < PASSES)); then
      sleep "${PASS_WAIT}"
    fi
  done
  echo
  if [ "${#remaining[@]}" -gt 0 ]; then
    echo "**${#remaining[@]} target(s) could not be deleted;** see the errors above."
    return 1
  fi
  echo "All targets are gone."
}

case "${1:-}" in
  inventory)
    shift
    inventory "$@"
    ;;
  delete)
    delete "${2:?usage: $0 delete <targets-file>}"
    ;;
  *)
    echo "usage: $0 inventory <region>... | delete <targets-file>" >&2
    exit 2
    ;;
esac
