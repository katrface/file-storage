include .env.example
export

LOCAL_BIN:=$(CURDIR)/bin
PATH:=$(LOCAL_BIN):$(PATH)


PROJECT_NAME ?= file-storage
REGISTRY_IMAGE ?= registry/file-storage
NETWORK_NAME ?= development_default
API_PORT ?= 8000


# Версия составляется по приоритету git tag > branch > commit
VERSION = $(shell git describe --tags --exact-match 2> /dev/null || git symbolic-ref -q --short HEAD || git rev-parse --short HEAD || echo "noversion")


# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

DOCKER_SETTINGS = --env PG_URL=$(PG_URL) \
	--env LOG_LEVEL=$(LOG_LEVEL)


build: ### Build docker image
	docker build --pull -t $(REGISTRY_IMAGE):$(VERSION) .
.PHONY: build

up: ### Up docker container
	make down
	@docker network create $(NETWORK_NAME) || true
	@docker run --name $(PROJECT_NAME) \
		$(DOCKER_SETTINGS) \
		--network=$(NETWORK_NAME) \
		-p 127.0.0.1:${API_PORT}:3000 \
		--detach \
		$(REGISTRY_IMAGE):$(VERSION)
.PHONY: up

down: ### Down docker container
	docker stop $(PROJECT_NAME) && docker container rm $(PROJECT_NAME) || true
.PHONY: down

run: ### Local start
	go run cmd/file_storage_server/main.go
.PHONY: run
