# Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
# See LICENSE.txt for license information.

## Docker Build Versions
DOCKER_BUILD_IMAGE = golang:1.16.3
DOCKER_BASE_IMAGE = alpine:3.13.2

# Binaries
TOOLS_BIN_DIR := $(abspath bin)
GO_INSTALL = ./scripts/go_install.sh

GOLINT_VER := master
GOLINT_BIN := golint
GOLINT_GEN := $(TOOLS_BIN_DIR)/$(GOLINT_BIN)

OUTDATED_VER := master
OUTDATED_BIN := go-mod-outdated
OUTDATED_GEN := $(TOOLS_BIN_DIR)/$(OUTDATED_BIN)

# Variables
GO = go
APP := rotatorctl
APPNAME := rotatorctl
BUILD_TIME := $(shell date -u +%Y%m%d.%H%M%S)
BUILD_HASH := $(shell git rev-parse HEAD)
ROTATORCTL_NAME ?= mattermost/rotatorctl

################################################################################

export GO111MODULE=on

all: check-style fmt

.PHONY: check-style
check-style: govet lint
	@echo Checking for style guide compliance

.PHONY: vet
govet:
	@echo Running govet
	$(GO) vet ./...
	@echo Govet success

.PHONY: lint
lint: $(GOLINT_GEN)
	@echo Running lint
	$(GOLINT_GEN) -set_exit_status ./...
	@echo Golint success

.PHONY: fmt
fmt: ## Run go fmt against code
	@echo Running go fmt
	go fmt ./...
	@echo Go fmt success

.PHONY: check-modules
check-modules: $(OUTDATED_GEN) ## Check outdated modules
	@echo Checking outdated modules
	$(GO) list -u -m -json all | $(OUTDATED_GEN) -update -direct

.PHONY: tests
tests: ## Run go test against code
	@echo Running go test
	$(GO) test ./... -v
	@echo Go test success

# Build for linux distribution
.PHONY: build
build:
	@echo Building Mattermost Rotatorctl for Linux
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GO) build -gcflags all=-trimpath=$(PWD) -asmflags all=-trimpath=$(PWD) -a -installsuffix cgo -mod=mod -o build/_output/bin/$(APP)  ./cmd/$(APP)

# Build for MacOS distribution
.PHONY: build-mac
build-mac:
	@echo Building Mattermost Rotatorctl for Mac
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 $(GO) build -gcflags all=-trimpath=$(PWD) -asmflags all=-trimpath=$(PWD) -a -installsuffix cgo -mod=mod -o build/_output/bin/$(APP)-darwin  ./cmd/$(APP)

# Builds the docker image
.PHONY: build-image
build-image:
	@echo Building Rotatorctl Docker Image
	docker build \
	--build-arg DOCKER_BUILD_IMAGE=$(DOCKER_BUILD_IMAGE) \
	--build-arg DOCKER_BASE_IMAGE=$(DOCKER_BASE_IMAGE) \
	. -f build/Dockerfile \
	-t $(ROTATORCTL_NAME):$(BUILD_HASH)_$(BUILD_TIME) -t $(ROTATORCTL_NAME):$(BUILD_HASH) \
	--no-cache

# Build for all distros
.PHONY: distros
distros: build build-mac build-image

# Cut a release
.PHONY: release
release:
	@echo Cut a release
	sh ./scripts/release.sh

.PHONY: tidy
tidy:
	@echo Go mod tidy
	$(GO) mod tidy
	@echo Go mod tidy success

## --------------------------------------
## Tooling Binaries
## --------------------------------------

$(OUTDATED_GEN): ## Build go-mod-outdated.
	GOBIN=$(TOOLS_BIN_DIR) $(GO_INSTALL) github.com/psampaz/go-mod-outdated $(OUTDATED_BIN) $(OUTDATED_VER)

$(GOLINT_GEN): ## Build golint.
	GOBIN=$(TOOLS_BIN_DIR) $(GO_INSTALL) golang.org/x/lint/golint $(GOLINT_BIN) $(GOLINT_VER)