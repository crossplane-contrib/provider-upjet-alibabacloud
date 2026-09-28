# v1.4.0

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

## New resources

AliKafka (ApsaraMQ for Kafka) is available as a new `alikafka` API group with
seven managed resources:

- `Instance`
- `Topic`
- `ConsumerGroup`
- `SaslUser`
- `SaslAcl`
- `ScheduledScalingRule`
- `InstanceAllowedIpAttachment`

The provider now ships 199 CRDs, up from 192 in v1.3.0.

## Updates

- Terraform provider `aliyun/alicloud` 1.267.0 to 1.281.0
- Go 1.24.1 to 1.26.8
- crossplane-runtime 1.20.0-rc to 1.20.11
- uptest v1.1.2 to v2.2.0
- Security updates for `google.golang.org/grpc` and `github.com/antchfx/xpath`
- Renovate is active, and GitHub Actions are pinned to digests
- Documentation covers the prerequisites for the smaller-scoped provider families

`k8s.io/apimachinery` and `k8s.io/client-go` are deliberately held at v0.32.3,
and `crossplane-tools` is pinned ahead of its crossplane-runtime v2 migration.
upjet remains at v1.9.0 and controller-runtime at v0.19.0.
