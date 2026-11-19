#!/usr/bin/env python3
# SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
#
# SPDX-License-Identifier: Apache-2.0

"""Generate examples/namespaced/ from the curated examples in examples/cluster/.

Every namespaced example is the cluster-scoped one with three changes:

- Managed resources move to the namespaced API group, e.g.
  vpc.alibabacloud.crossplane.io becomes vpc.alibabacloud.m.crossplane.io.
- Every object is placed in the "default" namespace. An object that already
  sets a namespace, such as a Secret in crossplane-system, is moved, because a
  namespaced managed resource reads its secret references from its own
  namespace.
- The namespace is dropped from secret references, which are local on
  namespaced managed resources (only name and key).

Test hook scripts get the same API group change. examples/namespaced/providerconfig/
is not generated: the namespaced ProviderConfig and ClusterProviderConfig kinds
differ from the cluster-scoped ProviderConfig, so those examples are written by
hand.

The transformation works line by line, so comments and formatting are kept.

Usage: python3 scripts/generate-namespaced-examples.py  (from the repo root)
"""

import os
import re
import shutil
import sys

SRC = os.path.join("examples", "cluster")
DST = os.path.join("examples", "namespaced")
KEEP = {"providerconfig"}
NAMESPACE = "default"

CLUSTER_GROUP = re.compile(r"\b([a-z0-9]+)\.alibabacloud\.crossplane\.io\b")
MR_API_VERSION = re.compile(r"^apiVersion:\s*[a-z0-9]+\.alibabacloud\.crossplane\.io/")
# A mapping key, optionally the first key of a list item ("- key:"), which sits
# two columns further in than the dash.
KEY_LINE = re.compile(r"^(\s*)(- )?([A-Za-z0-9_.-]+):")


def namespaced_group(text):
    return CLUSTER_GROUP.sub(r"\1.alibabacloud.m.crossplane.io", text)


def is_secret_ref(key):
    return key.endswith("SecretRef") or key == "writeConnectionSecretToRef"


def transform_document(lines):
    """Transforms one YAML document, given as a list of lines."""
    is_mr = any(MR_API_VERSION.match(line) for line in lines)
    out = []
    parents = []  # (indent, key) of the mapping keys enclosing the line
    for line in lines:
        m = KEY_LINE.match(line)
        if m:
            indent = len(m.group(1)) + (2 if m.group(2) else 0)
            key = m.group(3)
            while parents and parents[-1][0] >= indent:
                parents.pop()
            parent = parents[-1][1] if parents else None
            # The namespace is written right after "metadata:" below, so an
            # existing one is dropped rather than duplicated.
            if key == "namespace" and indent == 2 and parent == "metadata":
                continue
            if key == "namespace" and is_mr and parent and is_secret_ref(parent):
                continue
            parents.append((indent, key))
            if indent == 0 and key == "metadata":
                out.append(line)
                out.append(f"  namespace: {NAMESPACE}\n")
                continue
        if is_mr and line.startswith("apiVersion:"):
            line = namespaced_group(line)
        out.append(line)
    return out


def transform_yaml(text):
    docs, current = [], []
    for line in text.splitlines(keepends=True):
        if line.rstrip() == "---":
            docs.append((current, line))
            current = []
        else:
            current.append(line)
    docs.append((current, ""))
    return "".join("".join(transform_document(d)) + sep for d, sep in docs)


def main():
    if not os.path.isdir(SRC):
        sys.exit(f"{SRC} not found; run this from the repository root")
    for entry in os.listdir(DST) if os.path.isdir(DST) else []:
        if entry in KEEP:
            continue
        path = os.path.join(DST, entry)
        if os.path.isdir(path):
            shutil.rmtree(path)
        else:
            os.remove(path)

    count = 0
    for root, _, files in os.walk(SRC):
        rel = os.path.relpath(root, SRC)
        if rel.split(os.sep)[0] in KEEP:
            continue
        os.makedirs(os.path.join(DST, rel), exist_ok=True)
        for name in sorted(files):
            src, dst = os.path.join(root, name), os.path.join(DST, rel, name)
            with open(src) as f:
                text = f.read()
            if name.endswith(".yaml"):
                text = transform_yaml(text)
            elif name.endswith(".sh"):
                text = namespaced_group(text)
            with open(dst, "w") as f:
                f.write(text)
            shutil.copymode(src, dst)
            count += 1
    print(f"generate-namespaced-examples: wrote {count} files to {DST}")


if __name__ == "__main__":
    main()
