#!/usr/bin/env bash
# Tests the checks in hack/upgrade-test.sh against a stub `kubectl` on PATH
# that answers from fixture files, so nothing reaches a cluster or the cloud.
# Run it from the repository root: hack/upgrade-test-test.sh
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
WORK="$(mktemp -d)"
trap 'rm -rf "${WORK}"' EXIT
mkdir -p "${WORK}/bin"

# The stub answers `get managed` from $STUB_DIR/managed.json and `get events`
# from $STUB_DIR/events.json. It records Provider applies and gets in
# $STUB_DIR/calls, and reports each applied Provider healthy at its package.
cat >"${WORK}/bin/kubectl" <<'STUB'
#!/usr/bin/env bash
case "$*" in
  "get managed -A -o json") cat "${STUB_DIR}/managed.json" ;;
  "get events -A -o json") cat "${STUB_DIR}/events.json" ;;
  "apply -f -")
    manifest=$(cat)
    name=$(sed -n 's/^  name: //p' <<<"${manifest}")
    sed -n 's/^  package: //p' <<<"${manifest}" >"${STUB_DIR}/package-${name}"
    printf '%s\n' "${manifest}" >"${STUB_DIR}/manifest-${name}"
    echo "apply ${name}" >>"${STUB_DIR}/calls"
    ;;
  "get provider.pkg "*" -o json")
    echo "get $3" >>"${STUB_DIR}/calls"
    jq -n --rawfile id "${STUB_DIR}/package-$3" '{status: {currentIdentifier: ($id | rtrimstr("\n")),
      conditions: [{type: "Installed", status: "True"}, {type: "Healthy", status: "True"}]}}'
    ;;
  *) echo "stub kubectl: unexpected call: $*" >&2; exit 1 ;;
esac
STUB
chmod +x "${WORK}/bin/kubectl"
export PATH="${WORK}/bin:${PATH}" KUBECTL=kubectl

cd "${ROOT}"
# The script is checked on its own, so shellcheck need not follow it.
# shellcheck disable=SC1091
source hack/upgrade-test.sh
set +e

checks=0
problems=0
check() {
  local description="$1"
  shift
  checks=$((checks + 1))
  if "$@"; then
    echo "ok: ${description}"
  else
    echo "FAIL: ${description}"
    printf '%s' "${SUMMARY}" | sed 's/^/    /'
    problems=$((problems + 1))
  fi
}

# start <name> resets the script's state and the stub's fixtures.
start() {
  export STUB_DIR="${WORK}/$1"
  mkdir -p "${STUB_DIR}"
  echo '{"items": []}' >"${STUB_DIR}/events.json"
  failures=0
  SUMMARY=""
  echo "--- $1"
}

# mr <uid> <name> <external-name> <ready> <synced> [<for-provider-json>]
mr() {
  jq -n --arg uid "$1" --arg name "$2" --arg ext "$3" --arg ready "$4" --arg synced "$5" \
    --argjson fp "${6:-{\"region\": \"us-west-1\"\}}" '{
      kind: "VPC", metadata: {uid: $uid, name: $name, annotations: {"crossplane.io/external-name": $ext}},
      spec: {forProvider: $fp},
      status: {conditions: [{type: "Ready", status: $ready}, {type: "Synced", status: $synced}]}}'
}

managed() { jq -s '{items: .}' >"${STUB_DIR}/managed.json"; }

# event <uid> <type> <reason> <time>
event() {
  jq -n --arg uid "$1" --arg type "$2" --arg reason "$3" --arg time "$4" '{
    involvedObject: {uid: $uid, kind: "VPC", name: "a"}, type: $type, reason: $reason,
    lastTimestamp: $time, message: "stub"}'
}

events() { jq -s '{items: .}' >"${STUB_DIR}/events.json"; }

summary_has() { grep -qF -- "$1" <<<"${SUMMARY}"; }
summary_lacks() { ! grep -qF -- "$1" <<<"${SUMMARY}"; }

start unchanged
{ mr u1 a vpc-1 True True; mr u2 b vpc-2 True True; } | managed
snapshot "${STUB_DIR}/before"
snapshot "${STUB_DIR}/after"
compare "${STUB_DIR}/before" "${STUB_DIR}/after"
check "an unchanged resource passes" [ "${failures}" -eq 0 ]
check "the first resource is compared too, not skipped" summary_has "| VPC/a | same | True | True |  |"
check "all_ready accepts Ready and Synced resources" all_ready "${STUB_DIR}/after"

start gone-and-renamed
{ mr u1 a vpc-1 True True; mr u2 b vpc-2 True True; } | managed
snapshot "${STUB_DIR}/before"
mr u2 b vpc-other True True | managed
snapshot "${STUB_DIR}/after"
compare "${STUB_DIR}/before" "${STUB_DIR}/after"
check "both problems fail" [ "${failures}" -eq 2 ]
check "a missing resource is reported" summary_has "VPC/a is gone after the upgrade"
check "a new external name is reported" summary_has "VPC/b has a different external name"

start not-synced
mr u1 a vpc-1 True True | managed
snapshot "${STUB_DIR}/before"
mr u1 a vpc-1 True False | managed
snapshot "${STUB_DIR}/after"
compare "${STUB_DIR}/before" "${STUB_DIR}/after"
check "a resource that is not Synced fails" summary_has "VPC/a is not Ready and Synced after the upgrade (Ready True, Synced False)"
check "all_ready rejects it" bash -c "! jq -e 'length > 0 and all(.ready == \"True\" and .synced == \"True\")' '${STUB_DIR}/after' >/dev/null"

start for-provider-changed
mr u1 a vpc-1 True True '{"region": "us-west-1", "runtime": {"name": "containerd"}}' | managed
snapshot "${STUB_DIR}/before"
mr u1 a vpc-1 True True '{"region": "us-west-1"}' | managed
snapshot "${STUB_DIR}/after"
compare "${STUB_DIR}/before" "${STUB_DIR}/after"
check "a pruned field does not fail" [ "${failures}" -eq 0 ]
check "a pruned field is reported by name" summary_has "spec.forProvider changed: runtime"

start added
mr u1 a vpc-1 True True | managed
snapshot "${STUB_DIR}/before"
{ mr u1 a vpc-1 True True; mr u9 z vpc-9 True True; } | managed
snapshot "${STUB_DIR}/after"
compare "${STUB_DIR}/before" "${STUB_DIR}/after"
check "a new resource is reported" summary_has "New managed resources appeared after the upgrade: VPC/z"

start events
echo '["u1", "u2"]' >"${STUB_DIR}/uids"
{
  event u1 Normal CreatedExternalResource 2026-11-25T10:05:00Z
  event u2 Normal UpdatedExternalResource 2026-11-25T10:06:00Z
  event u1 Warning CannotObserveExternalResource 2026-11-25T10:07:00Z
  event u1 Normal DeletedExternalResource 2026-11-25T09:00:00Z
  event other Normal DeletedExternalResource 2026-11-25T10:08:00Z
} | events
check_events 2026-11-25T10:00:00Z "${STUB_DIR}/uids"
check "only the create after the upgrade fails" [ "${failures}" -eq 1 ]
check "the create is reported" summary_has "VPC/a: CreatedExternalResource after the upgrade"
check "an update is reported" summary_has "updated the cloud resource after the upgrade"
check "a warning is reported" summary_has "warning CannotObserveExternalResource"
check "a delete before the upgrade is ignored" summary_lacks "DeletedExternalResource"

start packages
cat >"${STUB_DIR}/example.yaml" <<'EOF'
apiVersion: vpc.alibabacloud.m.crossplane.io/v1alpha1
kind: VPC
---
apiVersion: ecs.alibabacloud.crossplane.io/v1alpha1
kind: SecurityGroup
metadata:
  annotations:
    uptest.upbound.io/timeout: "3600"
---
apiVersion: v1
kind: Secret
---
apiVersion: alibabacloud.crossplane.io/v1beta1
kind: ProviderConfig
EOF
check "both scopes map to their service, and other kinds are skipped" \
  [ "$(packages "${STUB_DIR}/example.yaml" | tr '\n' ' ')" = "provider-family-alibabacloud provider-alibabacloud-ecs provider-alibabacloud-vpc " ]
check "the longest uptest timeout wins" [ "$(ready_timeout "${STUB_DIR}/example.yaml")" = 3600 ]
check "the timeout is at least 2400s" [ "$(ready_timeout examples/vpc/v1alpha1/vpc.yaml)" = 2400 ]

start examples
check "a curated example is accepted" bash -c "source hack/upgrade-test.sh; check_examples examples/vpc/v1alpha1/vpc.yaml"
check "a generated example is rejected" bash -c "source hack/upgrade-test.sh; ! (check_examples examples-generated/vpc/v1alpha1/vpc.yaml) 2>/dev/null"
mkdir -p "${ROOT}/examples/.upgrade-test-fixture"
trap 'rm -rf "${WORK}" "${ROOT}/examples/.upgrade-test-fixture"' EXIT
# shellcheck disable=SC2016 # the ${...} is uptest's syntax, not the shell's.
printf 'spec:\n  # ${ignored} in a comment\n  name: ${Rand.RFC1123Subdomain}\n' >"${ROOT}/examples/.upgrade-test-fixture/x.yaml"
check "an example with an uptest substitution is rejected" \
  bash -c "source hack/upgrade-test.sh; ! (check_examples examples/.upgrade-test-fixture/x.yaml) 2>/dev/null"
for version in v1.4.0 v1.5.0-rc.0 v2.0.0-beta.0; do
  check "${version} is a valid version" bash -c "source hack/upgrade-test.sh; check_version ${version}"
done
for version in 1.4.0 v1.4 latest 'v1.4.0;x'; do
  check "\"${version}\" is rejected" bash -c "source hack/upgrade-test.sh; ! (check_version '${version}') 2>/dev/null"
done

# Providers are named as Crossplane names dependencies, and the family must be
# healthy before any service is applied, wherever it is in the list.
family=crossplane-contrib-provider-family-alibabacloud
ack=crossplane-contrib-provider-alibabacloud-ack
vpc=crossplane-contrib-provider-alibabacloud-vpc
for order in "family first" "family last"; do
  start "install, ${order}"
  if [ "${order}" = "family first" ]; then
    install v1.5.0 provider-family-alibabacloud provider-alibabacloud-ack provider-alibabacloud-vpc >/dev/null
  else
    install v1.5.0 provider-alibabacloud-ack provider-alibabacloud-vpc provider-family-alibabacloud >/dev/null
  fi
  check "the family is healthy before the services are applied" \
    [ "$(head -3 "${STUB_DIR}/calls" | tr '\n' ' ')" = "apply ${family} get ${family} apply ${ack} " ]
  check "every package is applied once and waited for" [ "$(sort "${STUB_DIR}/calls" | tr '\n' ' ')" = \
    "apply ${ack} apply ${vpc} apply ${family} get ${ack} get ${vpc} get ${family} " ]
done

start "provider names"
check "the family is named as Crossplane names its dependency" [ "${family}" = crossplane-contrib-provider-family-alibabacloud ]
check "another registry gives the same name" \
  [ "$(PACKAGE_ORG=xpkg.upbound.io/crossplane-contrib provider_name provider-alibabacloud-ack)" = "${ack}" ]
check "the package is the one the name was made from" \
  [ "$(cat "${WORK}/install, family last/package-${vpc}")" = xpkg.crossplane.io/crossplane-contrib/provider-alibabacloud-vpc:v1.5.0 ]

# DEBUG points every Provider at the debug runtime config, and only then.
for debug in true false; do
  start "debug-${debug}"
  DEBUG="${debug}" install v1.5.0 provider-family-alibabacloud provider-alibabacloud-ack >/dev/null
  if [ "${debug}" = true ]; then
    check "with DEBUG, the providers use the debug runtime config" \
      grep -qx "    name: ${DEBUG_RUNTIME_CONFIG}" "${STUB_DIR}/manifest-${family}" "${STUB_DIR}/manifest-${ack}"
  else
    check "without DEBUG, the providers keep the default runtime config" \
      bash -c "! grep -q runtimeConfigRef '${STUB_DIR}/manifest-${family}' '${STUB_DIR}/manifest-${ack}'"
  fi
done

start summary
for case in "0 Passed" "3 stopped early (exit status 3)"; do
  status="${case%% *}" want="${case#* }"
  check "exit status ${status} is reported as \"${want}\"" bash -c "
    source hack/upgrade-test.sh; unset GITHUB_STEP_SUMMARY
    FROM=v1.4.0 TO=v1.5.0 OUTPUT_DIR='${WORK}/summary-${status}'; mkdir -p \"\${OUTPUT_DIR}\"
    (trap write_summary EXIT; exit ${status}) >/dev/null || true
    grep -qF '${want}' \"\${OUTPUT_DIR}/summary.md\""
done

echo
if [ "${problems}" -gt 0 ]; then
  echo "${problems} of ${checks} checks failed"
  exit 1
fi
echo "all ${checks} checks passed"
