# RAM Role field migration

Use `spec.forProvider.roleName` and `spec.forProvider.assumeRolePolicyDocument`
for Role resources. The deprecated `name` and `document` fields have been
removed from Role parameters, init parameters, and observations. Update any
manifests using those fields, including `spec.initProvider`, before upgrading.
The resource's `metadata.name` and `crossplane.io/external-name` annotation are
unchanged.

Terraform Read returns both the canonical fields and their deprecated aliases.
Previously, late initialization could copy both pairs into the desired spec,
causing the next refresh to fail with `conflicts with` errors. The updated
provider ignores the deprecated aliases, including aliases already persisted by
an earlier provider alongside the canonical fields. No RAM role recreation is
required.

Deploy the updated provider package and its generated Role CRD together. After
upgrading, change `description` and verify that reconciliation succeeds and the
description is updated in RAM. Local regression tests cover parameter generation
after adoption and repeated observations; they do not replace this cloud check.
