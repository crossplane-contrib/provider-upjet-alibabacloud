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

# RAM SecurityPreference field migration

Use `mfaOperationForLogin` instead of `enforceMfaForLogin` in
`spec.forProvider` and `spec.initProvider`. The Terraform provider no longer
reads or writes `enforce_mfa_for_login`, so the deprecated field has also been
removed from the Crossplane parameters and observations.

Before upgrading, migrate manifests that use the old field and explicitly
choose the intended login MFA policy with `mfaOperationForLogin`. The provider
does not automatically translate the old boolean into a policy. Existing
configurations using `mfaOperationForLogin` are unaffected. Deploy the updated
provider package and SecurityPreference CRD together, then verify the effective
MFA policy in RAM.
