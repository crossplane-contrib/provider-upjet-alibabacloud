#!/usr/bin/env bash
# Installs the pinned Alibaba Cloud CLI for linux/amd64 into /usr/local/bin.
# Used by .github/workflows/cloud-cleanup.yaml, and by cloud-cleanup-check.yaml
# so that a version bump is tested before the cleanup workflow needs it.
#
# The archive is checked against the release's SHASUMS256.txt. That catches a
# corrupt or truncated download, not a replaced release: release assets are
# mutable, and the checksums file comes from the same release. The archives'
# .minisig files use the CLI's own signature format (aliyun-cli#1473), which
# stock minisign cannot verify.
set -euo pipefail

# Renovate updates this version; see .github/renovate.json5.
ALIYUN_CLI_VERSION="3.5.2"

base="https://github.com/aliyun/aliyun-cli/releases/download/v${ALIYUN_CLI_VERSION}"
archive="aliyun-cli-linux-${ALIYUN_CLI_VERSION}-amd64.tgz"
work="$(mktemp -d)"
trap 'rm -rf "${work}"' EXIT

cd "${work}"
curl -fsSLO "${base}/${archive}"
curl -fsSL "${base}/SHASUMS256.txt" | grep " ${archive}\$" | sha256sum -c -
tar -xzf "${archive}" aliyun
sudo install -m 0755 aliyun /usr/local/bin/aliyun
aliyun version
