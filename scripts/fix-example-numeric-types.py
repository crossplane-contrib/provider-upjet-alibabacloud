#!/usr/bin/env python3
"""Unquotes numeric-looking string values in generated example manifests
that the resource's CRD schema types as number/integer.

upjet's example generator (pkg/examples) carries scraped Terraform HCL
example values through without coercing them to the field's schema type,
so a field documented upstream as e.g. pending_window_in_days = "7" is
written to examples-generated/**/*.yaml as the YAML string "7" even
though the CRD types it as a number. Applying such a manifest fails
Create validation outright.

Run as part of `make generate` (see apis/generate.go) so regeneration
is self-correcting and `make check-diff` stays stable. Edits files via
targeted regex substitution rather than re-serializing them through
PyYAML, so nothing else in a file's formatting is disturbed.
"""
import glob
import re
import sys
from collections import defaultdict

try:
    import yaml
except ImportError:
    print(
        "fix-example-numeric-types: this script needs PyYAML, which isn't "
        "installed for this python3. Run `pip install pyyaml` (or "
        "`pip install --user pyyaml`) and re-run `make generate`.",
        file=sys.stderr,
    )
    sys.exit(1)


def extract_number_paths(node, prefix=""):
    """Return {path: schema type} for every number/integer leaf under an
    OpenAPI v3 schema node, recursing through object properties and array
    items."""
    result = {}
    if not isinstance(node, dict):
        return result
    for name, schema in node.get("properties", {}).items():
        path = f"{prefix}.{name}" if prefix else name
        t = schema.get("type")
        if t == "object" and "properties" in schema:
            result.update(extract_number_paths(schema, path))
        elif t == "array" and isinstance(schema.get("items"), dict):
            result.update(extract_number_paths(schema["items"], path + "[]"))
        elif t in ("number", "integer"):
            result[path] = t
    return result


def build_schema_index():
    """Map (group, kind) -> {field path: schema type} across all CRDs."""
    index = {}
    for f in glob.glob("package/crds/*.yaml"):
        try:
            docs = list(yaml.safe_load_all(open(f)))
        except yaml.YAMLError:
            continue
        for doc in docs:
            if not doc or doc.get("kind") != "CustomResourceDefinition":
                continue
            group = doc["spec"]["group"]
            kind = doc["spec"]["names"]["kind"]
            type_map = {}
            for v in doc["spec"].get("versions", []):
                spec_props = (
                    v.get("schema", {})
                    .get("openAPIV3Schema", {})
                    .get("properties", {})
                    .get("spec", {})
                    .get("properties", {})
                )
                for top in ("forProvider", "initProvider"):
                    if top in spec_props:
                        type_map.update(extract_number_paths(spec_props[top]))
            index[(group, kind)] = type_map
    return index


def find_violations(node, type_map, prefix=""):
    """Yield (leaf field name, quoted numeric value) for every scalar under
    node whose path is typed number/integer but whose value is a quoted
    numeric string."""
    if isinstance(node, dict):
        for k, v in node.items():
            path = f"{prefix}.{k}" if prefix else k
            if isinstance(v, dict):
                yield from find_violations(v, type_map, path)
            elif isinstance(v, list):
                for item in v:
                    if isinstance(item, dict):
                        yield from find_violations(item, type_map, path + "[]")
            elif isinstance(v, str) and path in type_map:
                if v.startswith("${") or v.startswith("$("):
                    continue
                if re.fullmatch(r"-?\d+(\.\d+)?", v):
                    yield (k, v)


def main():
    schema_index = build_schema_index()
    if not schema_index:
        print("fix-example-numeric-types: no CRDs found under package/crds/, skipping", file=sys.stderr)
        return 0

    by_file = defaultdict(set)
    for ef in glob.glob("examples-generated/**/*.yaml", recursive=True):
        try:
            docs = list(yaml.safe_load_all(open(ef)))
        except yaml.YAMLError as e:
            print(f"fix-example-numeric-types: cannot parse {ef}: {e}", file=sys.stderr)
            continue
        for doc in docs:
            if not doc or "apiVersion" not in doc or "kind" not in doc:
                continue
            group = doc["apiVersion"].split("/")[0]
            type_map = schema_index.get((group, doc["kind"]))
            if not type_map:
                continue
            spec = doc.get("spec", {})
            for top in ("forProvider", "initProvider"):
                if top in spec:
                    for leaf, val in find_violations(spec[top], type_map):
                        by_file[ef].add((leaf, val))

    total = 0
    for ef, pairs in by_file.items():
        content = open(ef).read()
        original = content
        for leaf, val in pairs:
            pattern = re.compile(
                r'(^[ \t]*(?:- )?' + re.escape(leaf) + r':\s*)"' + re.escape(val) + r'"\s*$',
                re.MULTILINE,
            )
            content, n = pattern.subn(lambda m: m.group(1) + val, content)
            total += n
        if content != original:
            open(ef, "w").write(content)
            print(f"fix-example-numeric-types: fixed {ef}")

    print(f"fix-example-numeric-types: {total} field(s) unquoted", file=sys.stderr)
    return 0


if __name__ == "__main__":
    sys.exit(main())
