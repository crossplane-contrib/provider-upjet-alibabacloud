#!/usr/bin/env bash
# Tests an upgrade of the provider between two published versions against the
# live test account: it installs the family packages at one version, creates
# the given examples, upgrades the packages to the other version, and checks
# that the managed resources survived. Driven by
# .github/workflows/upgrade-test.yaml; it expects a Kubernetes cluster with
# Crossplane installed, and jq.
#
#   upgrade-test.sh plan <from> <to> <example>...
#       Checks the arguments and prints the packages it would install. Needs no
#       cluster.
#
#   upgrade-test.sh run <from> <to> <example>...
#       Runs the test. Needs UPTEST_CLOUD_CREDENTIALS, as uptest does.
#
# Versions are package tags, such as v1.4.0 or v1.5.0-rc.0. Examples are
# curated files under examples/, and must not use uptest's ${...}
# substitutions, which only uptest fills in.
#
# The test fails when, after the upgrade, a managed resource is gone, has a
# different external name, is not Ready and Synced, or got a create or delete
# event; or when the new version cannot delete them. Updates to the cloud
# resource and changes to spec.forProvider are reported, but do not fail it.
#
# Optional environment:
#   KUBECTL          kubectl binary (default: kubectl)
#   PACKAGE_ORG      registry and org of the packages (default:
#                    xpkg.crossplane.io/crossplane-contrib, the one the
#                    service packages name for their family dependency)
#   READY_TIMEOUT    seconds to wait for Ready and Synced, and for deletion
#                    (default: 2400, or the longest uptest.upbound.io/timeout
#                    in the examples)
#   SETTLE_SECONDS   seconds to let the new version reconcile before checking
#                    (default: 300)
#   ACTIVATE         comma-separated ManagedResourceActivationPolicy patterns.
#                    When set, the default policy is replaced by one that
#                    activates only these, before the upgrade (Crossplane v2)
#   OUTPUT_DIR       where snapshots and the summary go (default:
#                    _output/upgrade-test)
#   DEBUG            "true" runs the providers with --debug, so their logs in
#                    the cluster dump show what they diff and change
set -euo pipefail

KUBECTL="${KUBECTL:-kubectl}"
PACKAGE_ORG="${PACKAGE_ORG:-xpkg.crossplane.io/crossplane-contrib}"
SETTLE_SECONDS="${SETTLE_SECONDS:-300}"
DEBUG="${DEBUG:-false}"
# The DeploymentRuntimeConfig the providers use when DEBUG is true.
DEBUG_RUNTIME_CONFIG=upgrade-test-debug
OUTPUT_DIR="${OUTPUT_DIR:-_output/upgrade-test}"
FAMILY=provider-family-alibabacloud

failures=0
SUMMARY=""

die() {
  echo "::error::$*" >&2
  exit 1
}

fail() {
  echo "::error::$*"
  failures=$((failures + 1))
  SUMMARY+="- ❌ $*"$'\n'
}

note() {
  echo "$*"
  SUMMARY+="- $*"$'\n'
}

check_version() {
  [[ "$1" =~ ^v[0-9]+\.[0-9]+\.[0-9]+(-[0-9A-Za-z.]+)?$ ]] ||
    die "\"$1\" is not a package version such as v1.4.0 or v1.5.0-rc.0"
}

# check_examples fails unless every example is a curated file that can be
# applied as it is.
check_examples() {
  local example
  [ "$#" -gt 0 ] || die "no examples given"
  for example in "$@"; do
    [[ "${example}" == examples/* && "${example}" == *.yaml ]] ||
      die "\"${example}\" is not a file under examples/"
    [[ "${example}" != *..* ]] || die "\"${example}\" must not contain \"..\""
    [ -f "${example}" ] || die "${example} does not exist"
    if grep -nE '^[^#]*\$\{' "${example}" >&2; then
      die "${example} uses uptest substitutions, which only uptest fills in"
    fi
  done
}

# services prints the service packages the examples need, one per line.
services() {
  grep -hE '^apiVersion:' "$@" | awk '{print $2}' | cut -d/ -f1 |
    sed -nE 's/^([a-z0-9]+)\.alibabacloud(\.m)?\.crossplane\.io$/\1/p' | sort -u
}

packages() {
  echo "${FAMILY}"
  services "$@" | sed 's/^/provider-alibabacloud-/'
}

ready_timeout() {
  local longest
  longest=$(grep -hoE 'uptest\.upbound\.io/timeout: *"?[0-9]+' "$@" | grep -oE '[0-9]+$' | sort -n | tail -1 || true)
  longest="${longest:-0}"
  echo "${READY_TIMEOUT:-$((longest > 2400 ? longest : 2400))}"
}

# Helm returns before Crossplane has installed its own CRDs, so wait for the
# deployment and the Provider CRD before applying any package.
wait_crossplane() {
  "${KUBECTL}" -n crossplane-system rollout status deployment/crossplane --timeout=5m
  "${KUBECTL}" wait --for=condition=Established crd/providers.pkg.crossplane.io --timeout=2m
}

# provider_name <package> prints the name Crossplane gives the package when it
# installs it as a dependency: its repository path without the registry, with
# "/" as "-", such as crossplane-contrib-provider-family-alibabacloud.
provider_name() {
  local path="${PACKAGE_ORG#*/}/$1"
  echo "${path//\//-}"
}

# runtime_config_ref prints the Provider spec lines that select the debug
# DeploymentRuntimeConfig, or nothing when DEBUG is not true.
runtime_config_ref() {
  [ "${DEBUG}" = true ] || return 0
  printf '\n  runtimeConfigRef:\n    apiVersion: pkg.crossplane.io/v1beta1\n    kind: DeploymentRuntimeConfig\n    name: %s' "${DEBUG_RUNTIME_CONFIG}"
}

# create_debug_runtime_config creates the DeploymentRuntimeConfig that runs a
# provider with --debug.
create_debug_runtime_config() {
  "${KUBECTL}" apply -f - <<EOF
apiVersion: pkg.crossplane.io/v1beta1
kind: DeploymentRuntimeConfig
metadata:
  name: ${DEBUG_RUNTIME_CONFIG}
spec:
  deploymentTemplate:
    spec:
      selector: {}
      template:
        spec:
          containers:
          - name: package-runtime
            args:
            - --debug
EOF
}

# apply_and_wait <version> <package>... applies the packages at that version and
# waits until each runs it.
apply_and_wait() {
  local version="$1" package name deadline current
  shift
  for package in "$@"; do
    "${KUBECTL}" apply -f - <<EOF
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: $(provider_name "${package}")
spec:
  package: ${PACKAGE_ORG}/${package}:${version}$(runtime_config_ref)
EOF
  done
  deadline=$((SECONDS + 900))
  for package in "$@"; do
    name=$(provider_name "${package}")
    while :; do
      current=$("${KUBECTL}" get provider.pkg "${name}" -o json | jq -r --arg want "${PACKAGE_ORG}/${package}:${version}" '
        def cond(t): [.status.conditions[]? | select(.type == t) | .status][0];
        if (.status.currentIdentifier == $want) and cond("Installed") == "True" and cond("Healthy") == "True"
        then "ready" else "\(.status.currentIdentifier // "none")" end')
      [ "${current}" = ready ] && break
      if ((SECONDS > deadline)); then
        "${KUBECTL}" get provider.pkg "${name}" -o yaml || true
        die "${name} did not become healthy at ${version} within 15 minutes (it runs ${current})"
      fi
      sleep 10
    done
    echo "${name} is healthy at ${version}"
  done
}

# install <version> <package>... installs or upgrades the packages, under the
# names Crossplane gives dependencies: when a service resolves its family
# dependency, it then finds ours, instead of installing a second family that
# breaks dependency resolution for all of them. The family goes first, on its
# own, so an upgrade moves it before the services that depend on it.
install() {
  local version="$1" name services=()
  shift
  for name in "$@"; do
    [ "${name}" = "${FAMILY}" ] || services+=("${name}")
  done
  apply_and_wait "${version}" "${FAMILY}"
  [ "${#services[@]}" -eq 0 ] || apply_and_wait "${version}" "${services[@]}"
}

# snapshot <file> saves what the checks compare for every managed resource.
snapshot() {
  "${KUBECTL}" get managed -A -o json | jq '[.items[] | {
    uid: .metadata.uid,
    ref: "\(.kind)/\(.metadata.name)\(if .metadata.namespace then " in \(.metadata.namespace)" else "" end)",
    externalName: .metadata.annotations["crossplane.io/external-name"],
    ready: ([.status.conditions[]? | select(.type == "Ready") | .status][0] // "Unknown"),
    synced: ([.status.conditions[]? | select(.type == "Synced") | .status][0] // "Unknown"),
    forProvider: .spec.forProvider
  }]' >"$1"
}

# all_ready <file> succeeds when every resource in the snapshot is Ready and
# Synced.
all_ready() {
  jq -e 'length > 0 and all(.ready == "True" and .synced == "True")' "$1" >/dev/null
}

wait_ready() {
  local file="$1" timeout="$2" deadline
  deadline=$((SECONDS + timeout))
  while :; do
    snapshot "${file}"
    all_ready "${file}" && return 0
    if ((SECONDS > deadline)); then
      return 1
    fi
    sleep 30
  done
}

# activate <patterns> replaces the default activation policy with one that
# activates only the given patterns.
activate() {
  local patterns
  "${KUBECTL}" get crd managedresourceactivationpolicies.apiextensions.crossplane.io >/dev/null 2>&1 ||
    die "ACTIVATE needs Crossplane v2, which has ManagedResourceActivationPolicy"
  patterns=$(tr ',' '\n' <<<"$1" | sed -E 's/^[[:space:]]+//; s/[[:space:]]+$//' | grep . |
    jq -R . | jq -s .)
  "${KUBECTL}" delete managedresourceactivationpolicy default --ignore-not-found
  jq -n --argjson patterns "${patterns}" '{
    apiVersion: "apiextensions.crossplane.io/v1alpha1",
    kind: "ManagedResourceActivationPolicy",
    metadata: {name: "upgrade-test"},
    spec: {activate: $patterns}
  }' | "${KUBECTL}" apply -f -
  note "Replaced the default activation policy with: $(jq -r 'join(", ")' <<<"${patterns}")"
}

# check_events <since> <uids-file> fails on create and delete events since the
# upgrade, and reports updates and warnings.
check_events() {
  local since="$1" uids="$2" events
  events=$("${KUBECTL}" get events -A -o json | jq -r --arg since "${since}" --slurpfile uids "${uids}" '
    .items[]
    | select(.involvedObject.uid as $u | $uids[0] | any(.[]; . == $u))
    | select((.lastTimestamp // .eventTime // .firstTimestamp // "") >= $since)
    | "\(.type)\t\(.reason)\t\(.involvedObject.kind)/\(.involvedObject.name)\t\(.message | gsub("\n"; " ") | .[0:200])"')
  while IFS=$'\t' read -r type reason ref message; do
    [ -n "${reason}" ] || continue
    case "${reason}" in
      CreatedExternalResource | DeletedExternalResource)
        fail "${ref}: ${reason} after the upgrade" ;;
      UpdatedExternalResource)
        note "⚠️ ${ref} updated the cloud resource after the upgrade" ;;
      *)
        if [ "${type}" = Warning ]; then
          note "⚠️ ${ref}: warning ${reason}: ${message}"
        fi
        ;;
    esac
  done <<<"${events}"
}

# compare <before> <after> checks every resource from before against after.
compare() {
  local before="$1" after="$2" rows
  rows=$(jq -r --slurpfile after "${after}" '
    .[] as $b
    | ($after[0] | map(select(.uid == $b.uid)) | first) as $a
    | [$b.ref,
       (if $a == null then "gone" elif $a.externalName != $b.externalName then "changed" else "same" end),
       ($a.ready // "-"), ($a.synced // "-"),
       (if $a == null then "" else
          [($b.forProvider // {}) + ($a.forProvider // {}) | keys[]
           | select(($b.forProvider // {})[.] != ($a.forProvider // {})[.])] | join(" ") end)]
    | @tsv' "${before}")
  SUMMARY+=$'\n'"| Resource | External name | Ready | Synced | Changed in spec.forProvider |"$'\n'
  SUMMARY+="| --- | --- | --- | --- | --- |"$'\n'
  while IFS=$'\t' read -r ref name ready synced changed; do
    SUMMARY+="| ${ref} | ${name} | ${ready} | ${synced} | ${changed} |"$'\n'
  done <<<"${rows}"
  SUMMARY+=$'\n'
  while IFS=$'\t' read -r ref name ready synced changed; do
    case "${name}" in
      gone) fail "${ref} is gone after the upgrade" ;;
      changed) fail "${ref} has a different external name after the upgrade" ;;
    esac
    [ "${name}" = gone ] && continue
    [ "${ready}" = True ] && [ "${synced}" = True ] ||
      fail "${ref} is not Ready and Synced after the upgrade (Ready ${ready}, Synced ${synced})"
    [ -z "${changed}" ] || note "⚠️ ${ref}: spec.forProvider changed: ${changed}"
  done <<<"${rows}"
  local added
  added=$(jq -r --slurpfile before "${before}" '
    [.[] | select(.uid as $u | $before[0] | any(.[]; .uid == $u) | not) | .ref] | join(", ")' "${after}")
  if [ -n "${added}" ]; then
    note "⚠️ New managed resources appeared after the upgrade: ${added}"
  fi
}

# activation_summary reports how many managed resource definitions are
# active, when Crossplane has them.
activation_summary() {
  "${KUBECTL}" get managedresourcedefinitions >/dev/null 2>&1 || return 0
  note "Managed resource definitions: $("${KUBECTL}" get managedresourcedefinitions -o json | jq -r '
    [.items[] | {scope: (if (.spec.group | test("\\.m\\.crossplane\\.io$")) then "namespaced" else "cluster" end), state: (.spec.state // "Unknown")}]
    | group_by(.scope + " " + .state) | map("\(length) \(.[0].scope) \(.[0].state)") | join(", ")')"
}

wait_deleted() {
  local timeout="$1" deadline left
  deadline=$((SECONDS + timeout))
  while :; do
    left=$("${KUBECTL}" get managed -A -o name | wc -l | tr -d ' ')
    [ "${left}" = 0 ] && return 0
    ((SECONDS > deadline)) && return 1
    sleep 30
  done
}

write_summary() {
  local status=$?
  {
    echo "## Upgrade test: ${FROM} → ${TO}"
    echo
    printf '%s' "${SUMMARY}"
    echo
    if [ "${failures}" -eq 0 ] && [ "${status}" -eq 0 ]; then
      echo "**Passed.**"
    elif [ "${failures}" -eq 0 ]; then
      echo "**Failed: the test stopped early (exit status ${status}); see the log.**"
    else
      echo "**Failed: ${failures} problem(s).**"
    fi
  } >"${OUTPUT_DIR}/summary.md"
  if [ -n "${GITHUB_STEP_SUMMARY:-}" ]; then
    cat "${OUTPUT_DIR}/summary.md" >>"${GITHUB_STEP_SUMMARY}"
  fi
  cat "${OUTPUT_DIR}/summary.md"
}

run() {
  local timeout since example
  timeout=$(ready_timeout "${EXAMPLES[@]}")
  mapfile -t PACKAGES < <(packages "${EXAMPLES[@]}")
  mkdir -p "${OUTPUT_DIR}"
  trap write_summary EXIT

  note "Packages: ${PACKAGES[*]}, from ${PACKAGE_ORG}"
  note "Examples: ${EXAMPLES[*]}"

  echo "--- Installing ${FROM}"
  wait_crossplane
  if [ "${DEBUG}" = true ]; then
    note "Providers run with --debug"
    create_debug_runtime_config
  fi
  install "${FROM}" "${PACKAGES[@]}"
  KUBECTL="${KUBECTL}" cluster/test/setup.sh

  echo "--- Creating the examples"
  for example in "${EXAMPLES[@]}"; do
    "${KUBECTL}" apply -f "${example}"
  done
  if ! wait_ready "${OUTPUT_DIR}/before.json" "${timeout}"; then
    "${KUBECTL}" get managed -A || true
    fail "The examples did not become Ready and Synced on ${FROM} within ${timeout}s, so the upgrade was not tested"
    return
  fi
  note "$(jq length "${OUTPUT_DIR}/before.json") managed resources Ready and Synced on ${FROM}"
  jq '[.[].uid]' "${OUTPUT_DIR}/before.json" >"${OUTPUT_DIR}/uids.json"

  if [ -n "${ACTIVATE:-}" ]; then
    activate "${ACTIVATE}"
  fi

  echo "--- Upgrading to ${TO}"
  since=$(date -u +%Y-%m-%dT%H:%M:%SZ)
  install "${TO}" "${PACKAGES[@]}"
  echo "Letting ${TO} reconcile for ${SETTLE_SECONDS}s"
  sleep "${SETTLE_SECONDS}"
  # A resource that is still settling gets the rest of the timeout; the
  # comparison then reports whatever state it is in.
  wait_ready "${OUTPUT_DIR}/after.json" "${timeout}" || true
  compare "${OUTPUT_DIR}/before.json" "${OUTPUT_DIR}/after.json"
  check_events "${since}" "${OUTPUT_DIR}/uids.json"
  activation_summary

  echo "--- Deleting with ${TO}"
  "${KUBECTL}" delete managed -A --all --wait=false
  if wait_deleted "${timeout}"; then
    note "All managed resources deleted with ${TO}"
  else
    "${KUBECTL}" get managed -A || true
    fail "${TO} did not delete every managed resource within ${timeout}s"
  fi
}

main() {
  [ "$#" -ge 4 ] || die "usage: $0 plan|run <from> <to> <example>..."
  MODE="$1" FROM="$2" TO="$3"
  shift 3
  EXAMPLES=("$@")
  check_version "${FROM}"
  check_version "${TO}"
  check_examples "${EXAMPLES[@]}"

  case "${MODE}" in
    plan)
      echo "Upgrade ${FROM} → ${TO}, from ${PACKAGE_ORG}"
      echo "Packages: $(packages "${EXAMPLES[@]}" | tr '\n' ' ')"
      echo "Ready and delete timeout: $(ready_timeout "${EXAMPLES[@]}")s"
      ;;
    run)
      [ -n "${UPTEST_CLOUD_CREDENTIALS:-}" ] || die "UPTEST_CLOUD_CREDENTIALS is not set"
      run
      [ "${failures}" -eq 0 ]
      ;;
    *)
      die "unknown mode \"${MODE}\"; use plan or run"
      ;;
  esac
}

# Run unless sourced, which the tests do to reach the functions.
if [[ "${BASH_SOURCE[0]}" == "$0" ]]; then
  main "$@"
fi
