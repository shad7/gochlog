MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash
.SHELLFLAGS := -o pipefail -euc
.DEFAULT_GOAL := build

.PHONY: clean test cover lint format xcompile

IMPORT_PATH := github.com/shad7/gochlog
VERSION?=$(shell awk -F\" '/^const Version/ { print $$2; exit }' core/info.go)

# Get the git commit
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
GIT_DESCRIBE=$(shell git describe --tags --always)
LDFLAGS := -X ${IMPORT_PATH}/core.GitCommit='${GIT_COMMIT}${GIT_DIRTY}' -X ${IMPORT_PATH}/core.GitDescribe='${GIT_DESCRIBE}'

ROOT := /gochlog
DEV_IMAGE := gochlog_dev

DOCKERRUN := docker run --rm \
	-v ${ROOT}/vendor:/go/src \
	-v ${ROOT}:${ROOT}/src/${IMPORT_PATH} \
	-w ${ROOT}/src/${IMPORT_PATH} \
	${DEV_IMAGE}

DOCKERBUILD := docker run --rm \
	-e VERSION="${VERSION}" \
	-e LDFLAGS="${LDFLAGS}" \
	-v ${ROOT}/vendor:/go/src \
	-v ${ROOT}:${ROOT}/src/${IMPORT_PATH} \
	-w ${ROOT}/src/${IMPORT_PATH} \
	${DEV_IMAGE}

DOCKERNOVENDOR := docker run --rm -i \
	-v ${ROOT}:${ROOT}/src/${IMPORT_PATH} \
	-w ${ROOT}/src/${IMPORT_PATH} \
	${DEV_IMAGE}


clean:
	@rm -rf bin build cover release vendor

# ----------------------------------------------
# docker build

# default top-level target
build: build/dev

build/dev:  generate check */*.go
	@mkdir -p bin/
	${DOCKERBUILD} go build -o bin/${ROOT} -ldflags "$(LDFLAGS)"

# builds the builder container
build/image_build:
	docker build -t ${DEV_IMAGE} -f Dockerfile.dev .

# top-level target for vendoring our packages: glide install requires
# being in the package directory so we have to run this for each package
vendor: build/image_build
	${DOCKERBUILD} glide install

# fetch a dependency via go get, vendor it, and then save into the parent
# package's glide.yml
# usage DEP=github.com/owner/package make add-dep
add-dep: build/image_build
	${DOCKERNOVENDOR} bash -c "DEP=$(DEP) ./scripts/add_dep.sh"

# ----------------------------------------------
# develop and test

generate: vendor
	${DOCKERNOVENDOR} bash ./scripts/gen.sh

format:
	${DOCKERNOVENDOR} bash ./scripts/fmt.sh

lint: format
	${DOCKERBUILD} bash ./scripts/lint.sh

check: lint
	${DOCKERBUILD} bash ./scripts/check.sh


# run unit tests and write out test coverage
test: check
	${DOCKERRUN} bash ./scripts/test.sh

cover: check
	@rm -rf cover/
	@mkdir -p cover
	${DOCKERRUN} bash ./scripts/cover.sh

# compile into binary for the current version across multiple different
# platforms (architectures and OSes)
xcompile: check
	@rm -rf build/
	@mkdir -p build
	${DOCKERBUILD} bash ./scripts/xcompile.sh

# zips each xcompiled binary and uploads to Artifactory
dist: xcompile
	@rm -rf release/
	@mkdir -p release/
	${DOCKERBUILD} bash ./scripts/dist.sh

docs: check
	@rm -rf docs/
	@mkdir -p docs/
	${DOCKERNOVENDOR} bash ./scripts/docs.sh

# ------ Docker Machine Helpers
mount:
	docker-machine ssh default 'sudo mkdir -p /gochlog ; sudo mount -t vboxsf gochlog /gochlog'

# ------ Docker Helpers
drma:
	docker rm $(shell docker ps -a -q)

drmia:
	docker rmi $(shell docker images -q --filter "dangling=true")

drmvu:
	docker volume rm $(shell docker volume ls -qf dangling=true)
