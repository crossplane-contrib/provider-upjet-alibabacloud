#!/bin/bash

# setup-family-providers.sh - Dynamically create provider-specific image directories and Dockerfiles
# This script creates the necessary image directories for each provider family

set -e

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
BASE_IMAGE_DIR="$ROOT_DIR/cluster/images/provider-alibabacloud"
IMAGES_DIR="$ROOT_DIR/cluster/images"

# Get the provider families from SUBPACKAGES environment variable, fallback to all if not set
FAMILY_PROVIDERS="${SUBPACKAGES:-ack ackone alb alidns cdn cloudmonitorservice ecs kms messageservice oss polardb privatelink quotas ram tair vpc config}"

echo "Setting up family provider image directories for: $FAMILY_PROVIDERS"

# Create image directories for each family provider
for provider in $FAMILY_PROVIDERS; do
    
    provider_image_dir="$IMAGES_DIR/provider-alibabacloud-$provider"
    echo "Creating image directory for $provider: $provider_image_dir"
    
    # Create the directory
    mkdir -p "$provider_image_dir"
    
    # Copy the base Makefile
    cp "$BASE_IMAGE_DIR/Makefile" "$provider_image_dir/"
    
    # Create provider-specific Dockerfile
    cat > "$provider_image_dir/Dockerfile" << EOF
FROM alpine:3.20.3
RUN apk --no-cache add ca-certificates bash

ARG TARGETOS
ARG TARGETARCH

ADD "bin/\${TARGETOS}_\${TARGETARCH}/$provider" /usr/local/bin/provider

ENV USER_ID=65532

# NOTE: This provider uses the no-fork architecture: it links the Terraform
# provider's Go SDK directly and calls its CRUD functions in-process. Neither
# the Terraform CLI nor the native provider plugin binary is needed at runtime,
# so nothing Terraform-related is installed or configured here.

USER \${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
EOF
    
    echo "Created Dockerfile for $provider"
done

echo "Family provider image directories setup complete!"
echo "Created image directories for: $FAMILY_PROVIDERS"