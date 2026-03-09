# provider-upjet-alibabacloud

## Build

Run the following to build the provider locally.

```
make submodules
make generate
```

## Authentication

The provider supports two authentication methods:

### 1. Traditional Authentication (Access Key)

Set up authentication using Alibaba Cloud Access Keys by creating a ProviderConfig with a secret reference:

```yaml
apiVersion: alibabacloud.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: example-creds
      namespace: crossplane-system
      key: credentials
```

The secret should contain a JSON with access_key, secret_key, and region:

```json
{
  "access_key": "YOUR_ACCESS_KEY",
  "secret_key": "YOUR_SECRET_KEY",
  "region": "cn-hangzhou"
}
```

### 2. OIDC Authentication

For enhanced security, the provider supports OpenID Connect (OIDC) authentication using Projected Service Account Tokens (PSAT).

To use OIDC authentication:

1. **Set up OIDC Provider in Alibaba Cloud RAM**

   First, create an OIDC provider in Alibaba Cloud RAM console that matches your identity provider's issuer URL. For example:
   - Provider name: `example-oidc-provider`
   - Issuer URL: `https://your-identity-provider.example.com`
   - Audience: `acs:ram::1234567890123456:oidc-provider/example-oidc-provider`

   > **Important**: The issuer URL configured in Alibaba Cloud must exactly match the `iss` claim in the OIDC tokens that will be presented to Alibaba Cloud STS API.

2. **Create a Service Account** (if using Kubernetes service account tokens)

   ```yaml
   apiVersion: v1
   kind: ServiceAccount
   metadata:
     name: provider-alibabacloud
     namespace: crossplane-system
   ```

3. **Configure the Provider with DeploymentRuntimeConfig**

   Create a DeploymentRuntimeConfig to mount the PSAT token:

   ```yaml
   apiVersion: pkg.crossplane.io/v1beta1
   kind: DeploymentRuntimeConfig
   metadata:
     name: alibaba-provider-oidc
   spec:
     deploymentTemplate:
       spec:
         selector: {}
         template:
           spec:
             serviceAccountName: provider-alibabacloud
             volumes:
               - name: oidc-token
                 projected:
                   sources:
                   - serviceAccountToken:
                       path: token
                       expirationSeconds: 3600
                       audience: acs:ram::1234567890123456:oidc-provider/example-oidc-provider
             containers:
               - name: package-runtime
                 volumeMounts:
                   - name: oidc-token
                     mountPath: /var/run/secrets/upbound.io/provider
                     readOnly: true
                 env:
                   - name: PSAT_TOKEN_PATH
                     value: /var/run/secrets/upbound.io/provider/token
   ```

   Then reference it in your Provider:

   ```yaml
   apiVersion: pkg.crossplane.io/v1
   kind: Provider
   metadata:
     name: provider-alibabacloud
   spec:
     package: xpkg.crossplane.io/crossplane-contrib/provider-alibabacloud:v1.2.0
     runtimeConfigRef:
       name: alibaba-provider-oidc
   ```

4. **Create a ProviderConfig with OIDC parameters**

   ```yaml
   apiVersion: alibabacloud.crossplane.io/v1beta1
   kind: ProviderConfig
   metadata:
     name: oidc-example
   spec:
     credentials:
       source: WebIdentity
       oidc:
         roleArn: acs:ram::1234567890123456:role/example-oidc-role
         providerArn: acs:ram::1234567890123456:oidc-provider/example-oidc-provider
   ```

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
