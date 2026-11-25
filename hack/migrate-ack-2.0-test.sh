#!/usr/bin/env bash
# Tests hack/migrate-ack-2.0.sh against a stub `kubectl` on PATH, so nothing
# reaches a cluster. Run it from anywhere: hack/migrate-ack-2.0-test.sh
set -euo pipefail

SCRIPT="$(cd "$(dirname "$0")" && pwd)/migrate-ack-2.0.sh"
WORK="$(mktemp -d)"
trap 'rm -rf "${WORK}"' EXIT
mkdir -p "${WORK}/bin"

# The stub serves `get crd` for the kinds with a file in $STUB_DIR/items, lists
# their items from that file, and logs every patch and annotate to
# $STUB_DIR/writes.
cat >"${WORK}/bin/kubectl" <<'STUB'
#!/usr/bin/env bash
set -euo pipefail
case "$1 $2" in
  "get crd")
    [ -f "${STUB_DIR}/items/${3%%.*}" ] || exit 1
    ;;
  get\ *)
    jq -s '{items: .}' "${STUB_DIR}/items/${2%%.*}"
    ;;
  patch\ * | annotate\ *)
    echo "$*" >>"${STUB_DIR}/writes"
    ;;
  *)
    echo "stub kubectl: unexpected call: $*" >&2
    exit 1
    ;;
esac
STUB
chmod +x "${WORK}/bin/kubectl"

ANNOTATION=migrate.alibabacloud.crossplane.io/ack-runtime
failures=0
check() {
  local description="$1"
  shift
  if "$@"; then
    echo "ok: ${description}"
  else
    echo "FAIL: ${description}"
    sed 's/^/    /' "${STUB_DIR}/out"
    failures=$((failures + 1))
  fi
}
run() {
  set +e
  # KUBECTL is pinned to the stub: the script honours it, and one exported in
  # the caller's environment would reach a real cluster.
  STUB_DIR="${STUB_DIR}" PATH="${WORK}/bin:${PATH}" KUBECTL=kubectl bash "${SCRIPT}" "$@" >"${STUB_DIR}/out" 2>&1
  echo $? >"${STUB_DIR}/status"
  set -e
}
status_is() { [ "$(cat "${STUB_DIR}/status")" = "$1" ]; }
output_has() { grep -qF -- "$1" "${STUB_DIR}/out"; }
wrote() { grep -qF -- "$1" "${STUB_DIR}/writes" 2>/dev/null; }
did_not_write() { ! grep -qF -- "$1" "${STUB_DIR}/writes" 2>/dev/null; }
no_writes() { [ ! -s "${STUB_DIR}/writes" ]; }

start() {
  export STUB_DIR="${WORK}/$1"
  mkdir -p "${STUB_DIR}/items"
  echo "--- $1"
}

start save
cat >"${STUB_DIR}/items/kubernetes" <<'EOF'
{"metadata": {"name": "k1"}, "spec": {"forProvider": {"runtime": {"name": "containerd", "version": "1.6.28"}}, "initProvider": {"runtime": {"name": "containerd"}}}}
{"metadata": {"name": "k2"}, "spec": {"forProvider": {}}}
EOF
cat >"${STUB_DIR}/items/managedkubernetes" <<'EOF'
{"metadata": {"name": "mk"}, "status": {"atProvider": {"connections": {"apiServerInternet": "x"}}}}
EOF
run save
check "save records the runtime in an annotation" \
  wrote "annotate --overwrite kubernetes.ack.alibabacloud.crossplane.io k1 ${ANNOTATION}={\"forProvider\":{\"name\":\"containerd\",\"version\":\"1.6.28\"},\"initProvider\":{\"name\":\"containerd\"}}"
check "save skips a resource with no runtime" did_not_write " k2 "
check "save leaves ManagedKubernetes alone" did_not_write "managedkubernetes"
check "save patches nothing" did_not_write "patch "
check "save says what to do next" output_has "Saved the runtime of 1 resource(s). Upgrade, then run"

start after-upgrade
cat >"${STUB_DIR}/items/kubernetes" <<EOF
{"metadata": {"name": "saved", "annotations": {"${ANNOTATION}": "{\"forProvider\":{\"name\":\"containerd\",\"version\":\"1.6.28\"}}"}}, "spec": {"forProvider": {"runtime": {}}}, "status": {"atProvider": {"connections": {}, "runtime": {}}}}
{"metadata": {"name": "unsaved"}, "spec": {"forProvider": {"runtime": {}}}, "status": {"atProvider": {"connections": {}}}}
{"metadata": {"name": "kept"}, "spec": {"forProvider": {"runtime": {"name": "containerd"}}}}
{"metadata": {"name": "done"}, "spec": {"forProvider": {"runtime": [{"name": "containerd"}]}}, "status": {"atProvider": {"connections": [{"serviceDomain": "x"}]}}}
EOF
cat >"${STUB_DIR}/items/managedkubernetes" <<'EOF'
{"metadata": {"name": "mk"}, "status": {"atProvider": {"certificateAuthority": {"client_cert": "c"}}}}
EOF
run
check "a dry run writes nothing" no_writes
check "a dry run lists what it would convert" output_has "would convert kubernetes.ack.alibabacloud.crossplane.io/saved: /spec/forProvider/runtime,"
run --apply
check "a pruned runtime is rebuilt from the saved values, and the annotation removed" \
  wrote 'patch kubernetes.ack.alibabacloud.crossplane.io saved --type=json -p [{"op":"replace","path":"/spec/forProvider/runtime","value":[{"name":"containerd","version":"1.6.28"}]},{"op":"remove","path":"/metadata/annotations/migrate.alibabacloud.crossplane.io~1ack-runtime"}]'
check "object-form status fields are removed through the status subresource" \
  wrote 'patch kubernetes.ack.alibabacloud.crossplane.io saved --subresource=status --type=json -p [{"op":"remove","path":"/status/atProvider/connections"},{"op":"remove","path":"/status/atProvider/runtime"}]'
check "a runtime with no saved values is not overwritten" did_not_write "unsaved --type=json"
check "it is reported" output_has "cannot convert kubernetes.ack.alibabacloud.crossplane.io/unsaved spec.forProvider.runtime: its values were not saved before the upgrade"
check "its status is still cleaned" wrote "unsaved --subresource=status"
check "and the run fails, so it is not missed" status_is 1
check "an unpruned runtime keeps its own values" \
  wrote 'patch kubernetes.ack.alibabacloud.crossplane.io kept --type=json -p [{"op":"replace","path":"/spec/forProvider/runtime","value":[{"name":"containerd"}]}]'
check "a 1.4 map in status is removed too" \
  wrote 'patch managedkubernetes.ack.alibabacloud.crossplane.io mk --subresource=status --type=json -p [{"op":"remove","path":"/status/atProvider/certificateAuthority"}]'
check "a resource already in 2.0 form is left alone" did_not_write " done "

start nothing-to-do
cat >"${STUB_DIR}/items/managedkubernetes" <<'EOF'
{"metadata": {"name": "mk"}, "status": {"atProvider": {"connections": [{"apiServerInternet": "x"}]}}}
EOF
run --apply
check "nothing to convert is reported" output_has "No ACK resources hold 1.x field shapes."
check "and nothing is written" no_writes
check "and the run succeeds" status_is 0

start bad-argument
run --force
check "an unknown argument shows the usage" output_has "usage:"
check "and fails" status_is 2

echo
if [ "${failures}" -gt 0 ]; then
  echo "${failures} check(s) failed"
  exit 1
fi
echo "all checks passed"
