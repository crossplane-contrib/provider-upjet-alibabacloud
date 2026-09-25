# Unreleased

## Breaking changes: RAM deprecated fields

The following fields have been removed from the released `v1alpha1` APIs,
including `spec.forProvider`, `spec.initProvider`, and `status.atProvider`:

| Resource | Removed field | Replacement input field |
| --- | --- | --- |
| Role | `name` | `roleName` |
| Role | `document` | `assumeRolePolicyDocument` |
| SecurityPreference | `enforceMfaForLogin` | `mfaOperationForLogin` |

Migrate manifests, including Composition templates and init-provider inputs,
before upgrading. There is no automatic conversion path. Removed fields are
no longer supported and may be rejected or pruned depending on API validation.
For SecurityPreference, explicitly choose the intended login MFA policy; the
old boolean is not automatically translated. Manifests already using the
replacement fields require no changes. Resource names (`metadata.name`) and
external-name annotations are unchanged.

Upgrade the provider package and generated CRDs together. This fixes Role
refresh conflicts caused by late initialization of deprecated aliases and
removes a SecurityPreference input that the Terraform provider no longer sends
to RAM. These changes do not require recreating RAM roles.
