#!/usr/bin/env bash
# Deletes the DBNode before the Instance it belongs to.
#
# uptest issues every delete with --wait=false and only waits for completion in
# a later step, so listing DBNode ahead of Instance in the manifest does not
# make it finish first -- the two race, and the Instance usually wins. Once the
# Instance is gone, DescribeDBInstanceAttribute can no longer resolve the node,
# upstream Read returns an error rather than clearing the id, and upjet has no
# hook to treat a provider error as "resource gone". The DBNode then keeps its
# finalizer and teardown hangs until someone removes it by hand.
#
# Running as a pre-delete hook on the Instance keeps the node on the path that
# does work: deleted while its parent is still around.
set -euo pipefail

KUBECTL="${KUBECTL:-kubectl}"
DBNODE="dbnode.rds.alibabacloud.crossplane.io/default"
# Well inside the 40m chainsaw exec budget for the whole delete step, which
# this hook shares with every other resource in the example.
TIMEOUT="${UPTEST_DBNODE_DELETE_TIMEOUT:-20m}"

# Captured in a plain assignment rather than tested inline: errexit is
# suppressed inside an `if` condition, so a get that fails for a real reason --
# a transient API error, missing RBAC -- would produce empty output and read as
# "no DBNode here", skipping the wait and silently restoring the race this hook
# exists to prevent. As an assignment, a non-zero exit aborts the hook instead.
# --ignore-not-found keeps a genuine absence at exit 0 with empty output.
EXISTING=$("${KUBECTL}" get "${DBNODE}" --ignore-not-found -o name)
if [ -z "${EXISTING}" ]; then
  echo "pre-delete-hook: ${DBNODE} is not present, nothing to do"
  exit 0
fi

echo "pre-delete-hook: deleting ${DBNODE} while its parent Instance still exists"
# Redundant in the current delete order, since uptest has already issued its own
# --wait=false delete for the DBNode by the time this runs, and the wait below is
# what actually does the work. Kept so the hook stays correct on its own terms if
# the DBNode is ever reordered after the Instance, or the hook reused elsewhere.
"${KUBECTL}" delete "${DBNODE}" --wait=false --ignore-not-found
"${KUBECTL}" wait "${DBNODE}" --for=delete --timeout="${TIMEOUT}"
echo "pre-delete-hook: ${DBNODE} is gone, the Instance can now be deleted"
