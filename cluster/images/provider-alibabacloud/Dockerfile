FROM alpine:3.22.1
RUN apk --no-cache add ca-certificates bash

ENV USER_ID=65532

# NOTE: This provider uses the no-fork architecture: it links the Terraform
# provider's Go SDK directly and calls its CRUD functions in-process. Neither
# the Terraform CLI nor the native provider plugin binary is needed at runtime,
# so nothing Terraform-related is installed here. Terraform is still used at
# build time to generate config/schema.json; see the top-level Makefile.
#
# The provider binary itself is layered on top of this family base image by
# `up xpkg batch`, which is why it is not ADDed here.

USER ${USER_ID}
EXPOSE 8080

ENTRYPOINT ["provider"]
