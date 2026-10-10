-include .env
include migrations/Makefile

all: # @HELP builds the docker image for our application
all:
	docker build \
	--progress=plain \
	--build-arg GO_VERSION=$(GO_VERSION) --build-arg OS=$(OS) --build-arg ARCH=$(ARCH) \
	-t $(REGISTRY)/$(DOCKER_IMAGE) \
	-f Dockerfile .

dev: # @HELP run the application locally with go run
dev:
	mkdir -p .db
	go run ./cmd/api

run: # @HELP run the application image locally with docker compose
run: all
	DOCKER_IMAGE=$(DOCKER_IMAGE) \
	REGISTRY=$(REGISTRY) \
	docker compose up

test: # @HELP run any tests
test:
	go test ./...

lint: # @HELP lints project files and scripts
lint:
	sh ./build/scripts/lint.sh

clean: # @HELP cleans temporary files and build artifacts created during development
clean:
	DOCKER_IMAGE=$(DOCKER_IMAGE) \
	REGISTRY=$(REGISTRY) \
	docker compose down

help-vars:
	@echo "VARIABLES:"
	@grep -hE '^# *[A-Z0-9_]+: ' $(MAKEFILE_LIST) \
		| sed -E 's/^# *([A-Z0-9_]+): (.*)/  \1: \2/'

help-targets:
	@printf "\nTARGETS\n"
	@grep -hE '^.*: *# *@HELP' $(MAKEFILE_LIST)			\
			| awk '																		\
					BEGIN {FS = ": *# *@HELP"};						\
					{ printf "	%-30s %s\n", $$1, $$2 };	\
			'

help: # @HELP prints help messages
help: help-vars help-targets

# REGISTRY: Where should Docker images be pushed once they're built?
REGISTRY ?= localhost:5000
# DOCKER_IMAGE: Name of the Docker image to build
DOCKER_IMAGE ?= api
# GO_VERSION: Go version of the image used to build the application
GO_VERSION ?= 1.26
# OS: Used internally. Users should pass GOOS and/or GOARCH.
OS := $(if $(GOOS),$(GOOS),$(shell GOTOOLCHAIN=local go env GOOS))
# ARCH: Used internally. Users should pass GOOS and/or GOARCH.
ARCH := $(if $(GOARCH),$(GOARCH),$(shell GOTOOLCHAIN=local go env GOARCH))

SHELL := /usr/bin/env bash -o errexit -o pipefail -o nounset
.DEFAULT_GOAL = all
.PHONY: all clean dev help help-targets help-vars lint run test
