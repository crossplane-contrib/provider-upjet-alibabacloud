#!/usr/bin/env bash
# Checks that this machine can reach the VPC endpoint of every region the given
# examples use, so a run on a runner whose route to Alibaba Cloud has dropped
# fails within minutes instead of timing out every create. Any HTTP response
# counts as reachable; the endpoint answers a bare request with 400. The same
# check runs inline in .github/workflows/e2e.yaml.
#
#   check-alibaba-connectivity.sh <example>...
set -uo pipefail

regions=$({ grep -hE '^[[:space:]]+region:' "$@" || true; } | awk '{print $2}' | tr -d "\"'" | sort -u)
if [ -z "${regions}" ]; then
  echo "No region checks required: the examples set no region"
  exit 0
fi
failed=0
for region in ${regions}; do
  [[ "${region}" =~ ^[a-z0-9-]+$ ]] || continue
  host="vpc.${region}.aliyuncs.com"
  reached=""
  for attempt in 1 2 3 4 5 6; do
    if curl -sS -o /dev/null --connect-timeout 10 --max-time 20 "https://${host}/"; then
      reached=1
      break
    fi
    echo "attempt ${attempt}: cannot reach ${host}"
    sleep 20
  done
  if [ -n "${reached}" ]; then
    echo "${host} is reachable"
  else
    echo "::error::cannot reach ${host} from this runner; rerun later"
    failed=1
  fi
done
exit "${failed}"
