MAKE_DIR = $(PWD)
ROOT_DIR = $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/../)
PROTO_DIR = $(ROOT_DIR)/proto

# one for each sub dir that has its own makefile
MONO_HTTP_DIR := $(MAKE_DIR)/apps/mono-http

# packages - test only
# PKG_DATASTORE_DIR := $(MAKE_DIR)/pkg/datastore

.PHONY: test
test:
	@$(MAKE) -C $(MONO_HTTP_DIR) pytest


# regenerate all of the proto stuff
.PHONY: pb
pb:
	bash $(PROTO_DIR)/protoc-py.sh $(ROOT_DIR)
