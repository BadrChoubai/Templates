include .env

all: # @HELP builds the docker image for our application
all:
	docker build \
	--progress=plain \
	-t $(REGISTRY)/$(DOCKER_IMAGE) \
	-f Dockerfile .

run: # @HELP run the application image locally with docker compose
run:
	DOCKER_IMAGE=$(DOCKER_IMAGE) REGISTRY=$(REGISTRY) docker compose up

test: # @HELP run any tests
test:
	@echo "No tests to run"

lint: # @HELP lints project files and scripts
lint:
	sh ./scripts/lint.sh

format: # @HELP formats project files and scripts
format:
	sh ./scripts/format.sh

sync: # @HELP install dependencies
sync:
	uv sync
	@echo "To activate the environment, run: source .venv/bin/activate"

clean: # @HELP cleans temporary files and build artifacts
clean:
	DOCKER_IMAGE=$(DOCKER_IMAGE) REGISTRY=$(REGISTRY) docker compose down
	docker image rm $(REGISTRY)/$(DOCKER_IMAGE)
	rm -rf __pycache__
	rm -rf $(VENV)

help: # @HELP prints this help message
help:
	@echo "TARGETS"
	@grep -hE '^.*: *# *@HELP' $(MAKEFILE_LIST)     \
	    | awk '                                   \
	        BEGIN {FS = ": *# *@HELP"};           \
	        { printf "  %-20s %-30s\n", $$1, $$2 };  \
	    '

# REGISTRY: Where should Docker images be pushed once they're built?
REGISTRY ?= localhost:5000

SHELL := /usr/bin/env bash -o errexit -o pipefail -o nounset
.DEFAULT_GOAL = all
.PHONY: activate all clean format help lint run sync test
