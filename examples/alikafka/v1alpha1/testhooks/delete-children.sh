#!/usr/bin/env bash
# Deletes the Kafka instance's child resources before the Instance itself.
#
# uptest issues every delete with --wait=false and only waits for completion in
# a later step, so the children and the Instance race, and the Instance usually
# wins. Once it is released, deleting a consumer group (and likely the other
# children) fails with BIZ.INSTANCE.STATUS.ERROR "instance is in RELEASED":
# upstream's describe calls treat that code as "not found", but its delete calls
# do not, so the child keeps its finalizer and teardown hangs. The RDS DBNode
# hook (examples/rds/v1alpha1/testhooks/delete-dbnode.sh) solves the same race.
#
# Running as a pre-delete hook on the Instance keeps the children on the path
# that does work: deleted while their instance is still around.
set -euo pipefail

KUBECTL="${KUBECTL:-kubectl}"
GROUP="alikafka.alibabacloud.crossplane.io"
KINDS="consumergroups topics saslusers saslacls instanceallowedipattachments scheduledscalingrules"
# Well inside the chainsaw exec budget for the whole delete step, which this hook
# shares with every other resource in the example.
TIMEOUT="${UPTEST_ALIKAFKA_CHILD_DELETE_TIMEOUT:-20m}"

# Captured in a plain assignment rather than tested inline, so a get that fails
# for a real reason aborts the hook instead of reading as "nothing to delete";
# see the DBNode hook for the full reasoning.
CHILDREN=""
for kind in ${KINDS}; do
  FOUND=$("${KUBECTL}" get "${kind}.${GROUP}" --ignore-not-found -o name)
  CHILDREN="${CHILDREN} ${FOUND}"
done

if [ -z "${CHILDREN// /}" ]; then
  echo "pre-delete-hook: no Kafka child resources, nothing to do"
  exit 0
fi

echo "pre-delete-hook: deleting Kafka child resources while their Instance still exists:${CHILDREN}"
# shellcheck disable=SC2086 # CHILDREN is a space-separated list of names.
"${KUBECTL}" delete ${CHILDREN} --wait=false --ignore-not-found
# shellcheck disable=SC2086
"${KUBECTL}" wait ${CHILDREN} --for=delete --timeout="${TIMEOUT}"
echo "pre-delete-hook: Kafka child resources are gone, the Instance can now be deleted"
