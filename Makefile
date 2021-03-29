# Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
# See LICENSE.txt for license information.

## Docker Build Versions
DOCKER_BUILD_IMAGE = golang:1.15.8
DOCKER_BASE_IMAGE = alpine:3.13

# Variables
GO = go
APP := rotatorctl
APPNAME := rotatorctl
TAG     := test
CHECKSUM = $(shell cat * | md5 | cut -c1-8)
ROTATORCTL_NAME ?= mattermost/rotatorctl

################################################################################

export GO111MODULE=on

all: check-style fmt

.PHONY: check-style
check-style: govet
	@echo Checking for style guide compliance

.PHONY: vet
govet:
	@echo Running govet
	$(GO) vet ./...
	@echo Govet success

.PHONY: fmt
fmt: ## Run go fmt against code
	@echo Running go fmt
	go fmt ./...
	@echo Go fmt success

.PHONY: tests
tests: ## Run go test against code
	@echo Running go test
	$(GO) test ./... -v
	@echo Go test success

# Build for linux distribution
.PHONY: build
build:
	@echo Building Mattermost Rotatorctl for Linux
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GO) build -gcflags all=-trimpath=$(PWD) -asmflags all=-trimpath=$(PWD) -a -installsuffix cgo -o build/_output/bin/$(APP)  ./cmd/$(APP)

# Build for MacOS distribution
.PHONY: build-mac
build-mac:
	@echo Building Mattermost Rotatorctl for Mac
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 $(GO) build -gcflags all=-trimpath=$(PWD) -asmflags all=-trimpath=$(PWD) -a -installsuffix cgo -o build/_output/bin/$(APP)-darwin  ./cmd/$(APP)

# Builds the docker image
.PHONY: build-image
build-image:
	@echo Building Rotatorctl Docker Image
	docker build \
	--build-arg DOCKER_BUILD_IMAGE=$(DOCKER_BUILD_IMAGE) \
	--build-arg DOCKER_BASE_IMAGE=$(DOCKER_BASE_IMAGE) \
	. -f build/Dockerfile \
	-t $(ROTATORCTL_NAME):$(TAG)_$(CHECKSUM) -t $(ROTATORCTL_NAME):$(TAG) \
	--no-cache

# Build for all distros
.PHONY: distros
distros: build build-mac build-image

# Cut a release
.PHONY: release
release:
	@echo Cut a release
	sh ./scripts/release.sh