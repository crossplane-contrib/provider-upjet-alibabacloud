# ====================================================================================
# Setup Project

include ../../../build/makelib/common.mk

# ====================================================================================
#  Options

include ../../../build/makelib/imagelight.mk

# ====================================================================================
# Setup k8s tools

-include ../../../build/makelib/k8s_tools.mk

# ====================================================================================
# Setup xpkg

-include ../../../build/makelib/xpkg.mk

# ====================================================================================
# Targets

img.build:
	@$(INFO) Family base image to build: $(IMAGE)
	@$(INFO) Building image $${IMAGE}; \
	$(MAKE) BUILD_ARGS="--load ${BUILD_ARGS}" IMAGE=$${IMAGE} XPKG_REG_ORGS=$(XPKG_REG_ORGS) img.build.shared; \
	if [[ "$${LOAD_PACKAGES}" == "true" ]]; then \
		$(MAKE) batch-process SUBPACKAGES="$(SUBPACKAGES)" XPKG_DIR=$(XPKG_DIR) XPKG_IGNORE=$(XPKG_IGNORE) BATCH_PLATFORMS=$(PLATFORM) BUILD_ONLY=true STORE_PACKAGES="$$(tr ' ' ',' <<< "$(SUBPACKAGES)")" && \
		for s in $(SUBPACKAGES); do \
			export t=$$(docker load -qi "$(XPKG_OUTPUT_DIR)/$(PLATFORM)/$(PROJECT_NAME)-$${s}-$(VERSION).xpkg") && \
			docker tag $${t##*:} $(BUILD_REGISTRY)/$(PROJECT_NAME)-$${s}-$(ARCH); \
		done; \
	fi || $(FAIL)
	@$(OK) docker build $${IMAGE}

img.build.shared:
	@cp Dockerfile $(IMAGE_TEMP_DIR) || $(FAIL)
	@cp -r $(OUTPUT_DIR)/bin/ $(IMAGE_TEMP_DIR)/bin || $(FAIL)
	@docker buildx build $(BUILD_ARGS) \
		--platform $(IMAGE_PLATFORMS) \
		-t $(IMAGE) \
		$(IMAGE_TEMP_DIR) || $(FAIL)

img.publish: batch-process

img.promote:
	@$(INFO) Skipping image promotion from $(FROM_IMAGE) to $(TO_IMAGE)
	@echo Promote is deferred to xpkg machinery
	@$(OK) Image promotion skipped for $(FROM_IMAGE) to $(TO_IMAGE)

CONCURRENCY ?= 30
# The service packages depend on the family config package. When it is
# missing, Crossplane installs the highest tag that satisfies this constraint,
# and when it is installed, keeps it as long as it satisfies it. So limit it to
# this release's major version: an open ">= 0.0.0" would install a newer
# major's family on a fresh 1.x install, and would keep a 1.x family under 2.x
# services. Pre-release builds also accept pre-release families. VERSION_MAJOR
# comes from build/makelib/common.mk.
ifeq ($(shell echo '$(VERSION_MAJOR)' | grep -xE '[0-9]+'),)
    $(error cannot read a major version from VERSION "$(VERSION)")
endif
VERSION_NEXT_MAJOR := $(shell echo $$(($(VERSION_MAJOR) + 1)))
DEP_CONSTRAINT := >= $(VERSION_MAJOR).0.0, < $(VERSION_NEXT_MAJOR).0.0
ifeq (-,$(findstring -,$(VERSION)))
    DEP_CONSTRAINT = >= $(VERSION_MAJOR).0.0-0, < $(VERSION_NEXT_MAJOR).0.0-0
endif
BUILD_ONLY ?= false
STORE_PACKAGES ?= ""
batch-process: $(CROSSPLANE_CLI)
	@$(INFO) Batch processing smaller provider packages for: "$(SUBPACKAGES)"
	@mkdir -p "$(XPKG_OUTPUT_DIR)/$(PLATFORM)" && \
	$(CROSSPLANE_CLI) xpkg batch --smaller-providers "$$(tr ' ' ',' <<< "$(SUBPACKAGES)")" \
		--family-base-image $(BUILD_REGISTRY)/$(PROJECT_NAME) \
		--platform $(BATCH_PLATFORMS) \
		--provider-name $(PROJECT_NAME) \
		--family-package-url-format $(XPKG_REG_ORGS)/%s:$(VERSION) \
		--package-repo-override monolith=$(PROJECT_NAME) --package-repo-override config=provider-family-$(PROVIDER_NAME) \
		--provider-bin-root $(OUTPUT_DIR)/bin \
		--output-dir $(XPKG_OUTPUT_DIR) \
		--store-packages "$(STORE_PACKAGES)" \
		--build-only=$(BUILD_ONLY) \
		--examples-root $(ROOT_DIR)/examples \
		--examples-group-override monolith=* --examples-group-override config=providerconfig \
		--crd-root $(ROOT_DIR)/package/crds \
		--crd-group-override monolith=* --crd-group-override config=$(CONFIG_CRD_GROUP) \
		--package-metadata-template $(ROOT_DIR)/package/crossplane.yaml.tmpl \
		--template-var XpkgRegOrg=$(CONFIG_DEPENDENCY_REG_ORG) --template-var DepConstraint="$(DEP_CONSTRAINT)" --template-var ProviderName=$(PROVIDER_NAME) --template-var ProviderAuthGroup=$(PROVIDER_AUTH_GROUP) \
		--concurrency $(CONCURRENCY) \
		--push-retry 10 || $(FAIL)
	@$(OK) Done processing smaller provider packages for: "$(SUBPACKAGES)"
