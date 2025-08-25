FROM alpine:3.22.1
RUN apk --no-cache add ca-certificates bash

ARG TARGETOS
ARG TARGETARCH

ADD "bin/${TARGETOS}_${TARGETARCH}/monolith" /usr/local/bin/provider

ENV USER_ID=65532

# Setup Terraform environment
ARG TERRAFORM_VERSION=1.5.7
ARG TERRAFORM_PROVIDER_SOURCE=aliyun/alicloud
ARG TERRAFORM_PROVIDER_VERSION=1.247.0
ARG TERRAFORM_PROVIDER_DOWNLOAD_NAME=terraform-provider-alicloud
ARG TERRAFORM_NATIVE_PROVIDER_BINARY=terraform-provider-alicloud_v1.247.0
ARG TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX=https://releases.hashicorp.com/terraform-provider-alicloud/1.247.0

ENV PLUGIN_DIR=/terraform/provider-mirror/registry.terraform.io/${TERRAFORM_PROVIDER_SOURCE}/${TERRAFORM_PROVIDER_VERSION}/${TARGETOS}_${TARGETARCH}
ENV TF_CLI_CONFIG_FILE=/terraform/.terraformrc
ENV TF_FORK=0

RUN mkdir -p ${PLUGIN_DIR}

ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_${TARGETOS}_${TARGETARCH}.zip /tmp
ADD ${TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX}/terraform-provider-alicloud_${TERRAFORM_PROVIDER_VERSION}_${TARGETOS}_${TARGETARCH}.zip /tmp
RUN echo -e 'provider_installation {\n  filesystem_mirror {\n    path    = "/terraform/provider-mirror"\n    include = ["*/*"]\n  }\n  direct {\n    exclude = ["*/*"]\n  }\n}' > ${TF_CLI_CONFIG_FILE}

RUN unzip /tmp/terraform_${TERRAFORM_VERSION}_${TARGETOS}_${TARGETARCH}.zip -d /usr/local/bin \
  && chmod +x /usr/local/bin/terraform \
  && rm /tmp/terraform_${TERRAFORM_VERSION}_${TARGETOS}_${TARGETARCH}.zip \
  && unzip /tmp/terraform-provider-alicloud_${TERRAFORM_PROVIDER_VERSION}_${TARGETOS}_${TARGETARCH}.zip -d ${PLUGIN_DIR} \
  && chmod +x ${PLUGIN_DIR}/* \
  && rm /tmp/terraform-provider-alicloud_${TERRAFORM_PROVIDER_VERSION}_${TARGETOS}_${TARGETARCH}.zip \
  && chown -R ${USER_ID}:${USER_ID} /terraform

# Provider controller needs these environment variables at runtime
ENV TERRAFORM_VERSION=${TERRAFORM_VERSION}
ENV TERRAFORM_PROVIDER_SOURCE=${TERRAFORM_PROVIDER_SOURCE}
ENV TERRAFORM_PROVIDER_VERSION=${TERRAFORM_PROVIDER_VERSION}
ENV TERRAFORM_NATIVE_PROVIDER_PATH=${PLUGIN_DIR}/${TERRAFORM_NATIVE_PROVIDER_BINARY}

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
