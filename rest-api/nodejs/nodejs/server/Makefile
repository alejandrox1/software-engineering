GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD 2>/dev/null)
GIT_BRANCH_CLEAN := $(shell echo $(GIT_BRANCH) | sed -e "s/[^[:alnum:]]/-/g")
DOCKER_IMAGE := nodejs-api$(if $(GIT_BRANCH_CLEAN),:$(GIT_BRANCH_CLEAN))

DOCKER_BUILD_ARGS ?= --force-rm
DOCKERFILE ?= Dockerfile

DOCKER_MOUNT := -v $(CURDIR):/opt/node-easy-notes-app
DOCKER_FLAGS := docker run --rm -it --net host $(DOCKER_MOUNT)

DOCKER_RUN := $(DOCKER_FLAGS) "$(DOCKER_IMAGE)"


.PHONY: build shell


build:
	docker build $(DOCKER_BUILD_ARGS) -f "$(DOCKERFILE)" -t "$(DOCKER_IMAGE)" .


shell: build
	$(DOCKER_RUN) bash

