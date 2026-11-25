#!/usr/bin/env bash
# Converts the ACK resources that provider 1.x created to the field shapes of
# 2.0.
#
#   migrate-ack-2.0.sh save      before upgrading: saves runtime values
#   migrate-ack-2.0.sh           after upgrading: shows what it would change
#   migrate-ack-2.0.sh --apply   after upgrading: changes it
#
# Terraform provider 2.0 changed certificateAuthority, connections and runtime
# from maps to lists of one object. Provider 1.x stored them as objects, and
# 2.0 reports a resource that still holds one Synced=False until it is
# converted.
#
# Run `save` on 1.x, before upgrading the ack package. Once the 2.0 CRDs are
# installed, the API server reads an object-form spec.forProvider.runtime as
# {}, dropping its name and version, and the next write of the resource stores
# that. So `save` records the runtime of each Kubernetes and EdgeKubernetes in
# the annotation below, which the upgrade leaves alone.
#
# After upgrading, for each cluster-scoped EdgeKubernetes, Kubernetes and
# ManagedKubernetes, --apply:
#
#   - wraps spec.forProvider.runtime and spec.initProvider.runtime in a list,
#     with the saved values: {name, version} becomes [{name, version}], and
#     removes the annotation. A runtime with no saved values is not changed;
#     set it from your manifest instead, in the list form;
#   - removes status.atProvider.certificateAuthority, connections and runtime,
#     which the provider fills in again on its next observation.
#
# Fields that are already lists are left alone, so it is safe to run again.
# Change manifests and Compositions that set runtime to the list form too, or
# the next apply of the object form is rejected.
#
# Needs kubectl and jq. Optional environment:
#   KUBECTL   kubectl binary (default: kubectl)
set -euo pipefail

KUBECTL="${KUBECTL:-kubectl}"
GROUP=ack.alibabacloud.crossplane.io
ANNOTATION=migrate.alibabacloud.crossplane.io/ack-runtime

mode=show
case "${1:-}" in
  "") ;;
  --apply) mode=apply ;;
  save) mode=save ;;
  *)
    echo "usage: $0 save | $0 [--apply]" >&2
    exit 2
    ;;
esac

# The runtime values to save: {"forProvider": {...}, "initProvider": {...}},
# with only the object-form, non-empty ones.
SAVED='{forProvider: .spec.forProvider.runtime, initProvider: .spec.initProvider.runtime}
  | with_entries(select(.value | type == "object" and length > 0))'

# For one resource after the upgrade: the JSON patches for its spec and
# metadata and for its status, and the runtimes that cannot be recovered.
# shellcheck disable=SC2016 # $a, $p, $f and $saved are jq's, not the shell's.
PLAN='(.metadata.annotations[$a] // "{}" | fromjson) as $saved
  | [["forProvider", "initProvider"][] as $p
     | select(.spec[$p].runtime | type == "object")
     | {p: $p, value: ($saved[$p] // (.spec[$p].runtime | if length > 0 then . else null end))}] as $runtimes
  | {
      spec: ([$runtimes[] | select(.value != null)
              | {op: "replace", path: "/spec/\(.p)/runtime", value: [.value]}]
             + if (.metadata.annotations[$a] != null) and ([$runtimes[] | select(.value == null)] == [])
               then [{op: "remove", path: "/metadata/annotations/\($a | gsub("~"; "~0") | gsub("/"; "~1"))"}]
               else [] end),
      status: [["certificateAuthority", "connections", "runtime"][] as $f
               | select(.status.atProvider[$f] | type == "object")
               | {op: "remove", path: "/status/atProvider/\($f)"}],
      lost: [$runtimes[] | select(.value == null) | "spec.\(.p).runtime"]
    }'

kinds=(edgekubernetes kubernetes managedkubernetes)
[ "${mode}" = save ] && kinds=(edgekubernetes kubernetes)

changed=0
lost=0
for kind in "${kinds[@]}"; do
  resource="${kind}.${GROUP}"
  if ! "${KUBECTL}" get crd "${resource}" >/dev/null 2>&1; then
    continue
  fi
  while IFS= read -r item; do
    name=$(jq -r '.metadata.name' <<<"${item}")
    if [ "${mode}" = save ]; then
      saved=$(jq -c "${SAVED}" <<<"${item}")
      [ "${saved}" = "{}" ] && continue
      "${KUBECTL}" annotate --overwrite "${resource}" "${name}" "${ANNOTATION}=${saved}" >/dev/null
      echo "saved the runtime of ${resource}/${name}"
      changed=$((changed + 1))
      continue
    fi
    plan=$(jq -c --arg a "${ANNOTATION}" "${PLAN}" <<<"${item}")
    spec=$(jq -c '.spec' <<<"${plan}")
    status=$(jq -c '.status' <<<"${plan}")
    for p in $(jq -r '.lost[]' <<<"${plan}"); do
      echo "cannot convert ${resource}/${name} ${p}: its values were not saved before the upgrade; set it from your manifest, as a list"
      lost=$((lost + 1))
    done
    [ "${spec}" = "[]" ] && [ "${status}" = "[]" ] && continue
    changed=$((changed + 1))
    fields=$(jq -r '[(.spec + .status)[] | .path] | join(", ")' <<<"${plan}")
    if [ "${mode}" = show ]; then
      echo "would convert ${resource}/${name}: ${fields}"
      continue
    fi
    if [ "${spec}" != "[]" ]; then
      "${KUBECTL}" patch "${resource}" "${name}" --type=json -p "${spec}" >/dev/null
    fi
    if [ "${status}" != "[]" ]; then
      "${KUBECTL}" patch "${resource}" "${name}" --subresource=status --type=json -p "${status}" >/dev/null
    fi
    echo "converted ${resource}/${name}: ${fields}"
  done < <("${KUBECTL}" get "${resource}" -o json | jq -c '.items[]')
done

case "${mode}" in
  save)
    echo "Saved the runtime of ${changed} resource(s). Upgrade, then run $0 --apply."
    ;;
  show)
    if [ "${changed}" -eq 0 ] && [ "${lost}" -eq 0 ]; then
      echo "No ACK resources hold 1.x field shapes."
    elif [ "${changed}" -gt 0 ]; then
      echo "Run with --apply to convert these ${changed} resource(s)."
    fi
    ;;
  apply)
    [ "${changed}" -eq 0 ] && [ "${lost}" -eq 0 ] && echo "No ACK resources hold 1.x field shapes."
    ;;
esac
[ "${lost}" -eq 0 ]
