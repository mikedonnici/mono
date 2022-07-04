MAKE_DIR = $(PWD)
ROOT_DIR = $(abspath $(dir $(lastword $(MAKEFILE_LIST)))/../)
PROTO_DIR = $(ROOT_DIR)/proto

# one for each sub dir that has its own Makefile
GRPC_SRVR_DIR := $(MAKE_DIR)/cmd/grpc-srvr

# packages - test only
PKG_DATASTORE_DIR := $(MAKE_DIR)/pkg/datastore


checkdir:
	echo $(MAKE_DIR)
	echo $(ROOT_DIR)
	echo $(PROTO_DIR)


.PHONY: test
test:
	@$(MAKE) -C $(GRPC_SRVR_DIR) test
	@$(MAKE) -C $(PKG_DATASTORE_DIR) test


.PHONY: build
build:
	@$(MAKE) -C $(GRPC_SRVR_DIR) build

.PHONY: run
run:
	@$(MAKE) -C $(GRPC_SRVR_DIR) run

# runs the default target for all the specified sub dir makefiles
all:
	@$(MAKE) -C $(GRPC_SRVR_DIR)

.PHONY: clean
clean:
	@$(MAKE) -C $(GRPC_SRVR_DIR) clean

## regenerate all of the proto stuff
#.PHONY: pb
#pb:
#	bash $(PROTO_DIR)/protoc-go.sh $(ROOT_DIR)
