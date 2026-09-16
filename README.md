# provider-upjet-alibabacloud

## Build

Run the following to build the provider locally.

```
make submodules
make generate
```

## API groups and scopes

The provider serves every resource twice, in two API groups:

| Group | Scope | Use |
|---|---|---|
| `<service>.alibabacloud.crossplane.io` | Cluster | Legacy cluster-scoped managed resources. Works on Crossplane v1 and v2. |
| `<service>.alibabacloud.m.crossplane.io` | Namespaced | Crossplane v2 namespaced managed resources. Composable in v2. |

The two are functionally equivalent; existing cluster-scoped resources keep
working unchanged. Namespaced resources differ in three ways:

- `spec.providerConfigRef` is a typed reference, with `kind` as well as `name`.
  Omitted, it defaults to `kind: ClusterProviderConfig, name: default`.
- `spec.writeConnectionSecretToRef` and any `...SecretRef` under
  `spec.forProvider` are local references: they resolve in the resource's own
  namespace and take no `namespace` field.
- `spec.publishConnectionDetailsTo` does not exist. External Secret Stores were
  an alpha feature that Crossplane v2 dropped, so it is gone from **all**
  resources, cluster-scoped ones included.

```yaml
apiVersion: vpc.alibabacloud.m.crossplane.io/v1alpha1
kind: VPC
metadata:
  name: example
  namespace: default
spec:
  providerConfigRef:
    kind: ClusterProviderConfig
    name: default
  forProvider:
    cidrBlock: 10.0.0.0/8
```

Examples for both scopes live under [examples/cluster](examples/cluster) and
[examples/namespaced](examples/namespaced).

## Authentication

Provider credentials are configured with a provider config. Static credentials
and STS session credentials are read from a Kubernetes `Secret`; AssumeRole and
AssumeRoleWithOIDC options are configured as structured fields on its spec.

There are three kinds, and a managed resource may only reference one from its
own API group:

| Kind | Group | Scope | Referenced by |
|---|---|---|---|
| `ProviderConfig` | `alibabacloud.crossplane.io` | Cluster | cluster-scoped MRs |
| `ProviderConfig` | `alibabacloud.m.crossplane.io` | Namespaced | namespaced MRs, same namespace |
| `ClusterProviderConfig` | `alibabacloud.m.crossplane.io` | Cluster | namespaced MRs, any namespace |

Use a namespaced `ProviderConfig` when each namespace should carry its own
credential, and a `ClusterProviderConfig` to share one across namespaces. The
spec is identical in all three; the examples below apply to each.

Do not put `assume_role` or `assume_role_with_oidc` blocks in the credentials
secret. The credentials secret should contain only credential material such as
`access_key`, `secret_key`, `security_token`, and `region`.

### Static Credentials

The default static AK/SK example uses `credentials.source: Secret`:

```bash
kubectl apply -f examples/providerconfig/v1beta1/secret.yaml.tmpl
kubectl apply -f examples/providerconfig/v1beta1/providerconfig.yaml
```

The credentials secret uses JSON in the `credentials` key:

```json
{
  "access_key": "...",
  "secret_key": "...",
  "region": "cn-hangzhou"
}
```

`region_id` is still accepted as a compatibility fallback, but new examples
should use `region`.

### STS Session Credentials

STS session credentials are supported by adding `security_token` to the secret:

```bash
kubectl apply -f examples/providerconfig/v1beta1/secret-sts-token.yaml.tmpl
kubectl apply -f examples/providerconfig/v1beta1/providerconfig-sts-token.yaml
```

### AssumeRole

For AssumeRole, use a base credential secret for the caller identity and put
the role parameters in `spec.assumeRole`:

```bash
kubectl apply -f examples/providerconfig/v1beta1/secret.yaml.tmpl
kubectl apply -f examples/providerconfig/v1beta1/providerconfig-assume-role.yaml
```

```yaml
apiVersion: alibabacloud.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: assume-role
spec:
  credentials:
    source: Secret
    secretRef:
      name: example-creds
      namespace: crossplane-system
      key: credentials
  assumeRole:
    roleARN: acs:ram::<account-id>:role/<role-name>
    sessionName: crossplane-assume-role
    sessionExpiration: 3600
```

### AssumeRoleWithOIDC

For AssumeRoleWithOIDC, configure the role, OIDC provider, and token source in
`spec.assumeRoleWithOIDC`:

```bash
kubectl apply -f examples/providerconfig/v1beta1/providerconfig-assume-role-with-oidc.yaml
```

```yaml
apiVersion: alibabacloud.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: assume-role-with-oidc
spec:
  credentials:
    source: None
  assumeRoleWithOIDC:
    roleARN: acs:ram::<account-id>:role/<role-name>
    oidcProviderARN: acs:ram::<account-id>:oidc-provider/<provider-name>
    oidcTokenFile: /var/run/secrets/ack.alibabacloud.com/rrsa-tokens/token
    roleSessionName: crossplane-oidc
    sessionExpiration: 3600
```

When running in ACK with RRSA enabled, use the token file path injected by RRSA.
For local testing, create a projected ServiceAccount token and point
`oidcTokenFile` at that local file.

After applying a `ProviderConfig`, managed resources can reference it with
`spec.providerConfigRef.name`. If omitted, resources use the `default`
`ProviderConfig`.

## Test

Add an environment variable to set the credentials for the target Alibaba
account for the tests as follows and then run `make e2e`.

```
export UPTEST_CLOUD_CREDENTIALS='{
    "access_key": "...",
    "secret_key": "...",
    "region": "us-west-1"
}'
```

## Hiding a field from a generated CRD

This provider uses upjet's no-fork architecture: it links the upstream Terraform
provider's Go SDK and calls its CRUD functions in-process, rather than shelling
out to `terraform`.

That changes what `delete(r.TerraformResource.Schema, "field")` does, and the
difference is easy to miss. Under CLI reconciliation `TerraformResource` was a
throwaway map rebuilt from `config/schema.json`, so deleting from it only shaped
the generated CRD. Under no-fork, upjet assigns `TerraformResource` straight from
the live `*schema.Provider` **without copying it**, so the same delete strips the
field from the map the upstream CRUD functions execute against. Depending on how
upstream reads that field, the result is anything from a silently unset value to
a panic on every create:

```go
// alicloud_oss_bucket Create, upstream:
oss.ACL(oss.ACLType(d.Get("acl").(string)))   // d.Get returns nil once "acl" is
                                              // gone from the schema
```

So per-resource configurators in `config/<group>/config.go` must **not** call
`delete(r.TerraformResource.Schema, ...)`. Add the field to
`generationOnlySchemaOmissions` in [config/schema_omissions.go](config/schema_omissions.go)
instead. Those omissions are applied only when `GetProvider` is building the
code-generation provider, which shapes the CRD exactly as before while leaving
the runtime schema intact.

`TestRuntimeSchemaIsPristine` enforces this. If it fails after you add a resource
— or after merging a branch that adds one — the fix is to **move** the delete
into the table, not to revert the test:

```
resource "alicloud_alikafka_instance" is missing field "topic_quota" from its
runtime Terraform schema; schema fields may only be omitted from the
code-generation provider (see generationOnlySchemaOmissions)
```

## Submit PR

- `make reviewable` before submitting a new PR
- git commit -s -m "sign every commit"

## Release New Provider Version

### Determine Version

Identify the version to be released by increasing the minor version by one. For example, if the provider's latest version is v1.1.0, the new version will be v1.2.0.

According to the semantic versioning specification, a version number is represented as MAJOR.MINOR.PATCH. For 1.2.0 : MAJOR=1, MINOR=2, PATCH=0 

### Create Release Branch

From the GitHub UI, create a new branch from the main branch with the name release-<major>-<minor><patch>.

To cut the release v1.2.0, we will name our branch release-1.2.0.

### Build Release Candidate

GitHub should automatically trigger a `CI` workflow run on the newly created branch and produce a package.  You can check it from the GitHub UI by clicking `Actions => CI`.

If it does not, you can manually run the GitHub workflow named CI on the release branch to produce a package.

### Cut The Release

Tag the release branch with the version by running the GitHub workflow named `Tag` on the release branch.

### Prerequisite for Publishing New Smaller Scoped Providers

Please ask an authorized person to create a smaller scoped provider repository in the upbound.io crossplane-contrib organization, so that the consecutive "Publish The Providers" step can succeed. The respository needs to be publicly accessible and be published to the Upbound marketplace itself.

### Publish The Providers

Build and push the family packages using the `Publish Provider Packages` Github Actions workflow. To do this, you need to provide the values of the following parameters:

- subpackages (to be built individually, e.g. config ram): config ack ackone alb alidns cdn cloudmonitorservice ecs fcv3 kms messageservice oss polardb privatelink quotas ram slb tair vpc
- size (Number of smaller provider packages to build and push with each build job): 30
- concurrency (Number of parallel package builds within each build job): 1
- version (Version string to use while publishing the packages,e.g. v1.2.0): v1.2.0
- go-version (Go version to use if building needs to be done): 1.24

Your release build will be published once the `Publish Provider Packages` job if
releasing a family of providers succeeds. Check their availability in the
Upbound marketplace [here](https://marketplace.upbound.io/providers/crossplane-contrib/provider-family-alibabacloud).

### Add Release Notes

Go [here](https://github.com/crossplane-contrib/provider-upjet-alibabacloud) and
click on releases on the left side. 

On the releases page, click on "Draft New Release".
- As target select your release branch that you created above
- Select the corresponding release tag
- Use your version as Release Title, e.g. v1.2.0
- Click "Generate release notes"
- Click "Publish release"
