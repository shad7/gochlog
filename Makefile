MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash
.SHELLFLAGS := -o pipefail -euc
.DEFAULT_GOAL := build

.PHONY: clean test cover check format xcompile

export PROJECT = gochlog
IMPORT_PATH := github.com/shad7/${PROJECT}

INSTALL_VENDOR := $(shell [ ! -d vendor ] && echo 1)

FORCE_VENDOR_INSTALL ?=
ifneq ($(strip $(FORCE_VENDOR_INSTALL)),)
	INSTALL_VENDOR := 1
endif

VERSION?=$(shell awk -F\" '/^const Version/ { print $$2; exit }' core/info.go)

# Get the git commit
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
GIT_DESCRIBE=$(shell git describe --tags --always)
LDFLAGS := -X ${IMPORT_PATH}/core.GitCommit='${GIT_COMMIT}${GIT_DIRTY}' -X ${IMPORT_PATH}/core.GitDescribe='${GIT_DESCRIBE}'

# Windows environment?
CYG_CHECK := $(shell hash cygpath 2>/dev/null && echo 1)
ifeq ($(CYG_CHECK),1)
	VBOX_CHECK := $(shell hash VBoxManage 2>/dev/null && echo 1)

	# Docker Toolbox (pre-Windows 10)
	ifeq ($(VBOX_CHECK),1)
		ROOT := /${PROJECT}
	else
		# Docker Windows
		ROOT := $(shell cygpath -m -a "$(shell pwd)")
	endif
else
	# all non-windows environments
	ROOT := $(shell pwd)
endif

DEV_IMAGE := ${PROJECT}_dev

DOCKERRUN := docker run --rm \
	-v ${ROOT}/vendor:/go/src \
	-v ${ROOT}:/${PROJECT}/src/${IMPORT_PATH} \
	-w /${PROJECT}/src/${IMPORT_PATH} \
	${DEV_IMAGE}

DOCKERBUILD := docker run --rm \
	-e PROJECT="${PROJECT}" \
	-e VERSION="${VERSION}" \
	-e LDFLAGS="${LDFLAGS}" \
	-v ${ROOT}/vendor:/go/src \
	-v ${ROOT}:/${PROJECT}/src/${IMPORT_PATH} \
	-w /${PROJECT}/src/${IMPORT_PATH} \
	${DEV_IMAGE}

DOCKERNOVENDOR := docker run --rm -i \
	-e IMPORT_PATH="${IMPORT_PATH}" \
	-v ${ROOT}:/${PROJECT}/src/${IMPORT_PATH} \
	-w /${PROJECT}/src/${IMPORT_PATH} \
	${DEV_IMAGE}


clean:
	@rm -rf bin build cover release vendor

# ----------------------------------------------
# docker build

# default top-level target
build: build/dev

build/dev: generate check */*.go
	@mkdir -p bin/
	${DOCKERBUILD} go build -o bin/${PROJECT} -ldflags "$(LDFLAGS)"

# builds the builder container
build/image_build:
	@echo "Building dev container"
	@docker build --quiet -t ${DEV_IMAGE} -f Dockerfile.dev .

# top-level target for vendoring our packages: glide install requires
# being in the package directory so we have to run this for each package
vendor: build/image_build
ifeq ($(INSTALL_VENDOR),1)
	${DOCKERRUN} glide install --skip-test
endif

# fetch a dependency via go get, vendor it, and then save into the parent
# package's glide.yml
# usage DEP=github.com/owner/package make add-dep
add-dep: build/image_build
ifeq ($(strip $(DEP)),)
	$(error "No dependency provided. Expected: DEP=<go import path>")
endif
	${DOCKERNOVENDOR} glide get --skip-test ${DEP}

# ----------------------------------------------
# develop and test

generate: vendor
	${DOCKERNOVENDOR} bash ./scripts/gen.sh

format: vendor
	${DOCKERNOVENDOR} bash ./scripts/fmt.sh

check: format
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

# zips each xcompiled binary and uploads to Github
dist: xcompile
	@rm -rf release/
	@mkdir -p release/
	${DOCKERBUILD} bash ./scripts/dist.sh

docs: check
	@rm -rf docs/
	@mkdir -p docs/
	${DOCKERNOVENDOR} bash ./scripts/docs.sh

# ------ Docker Helpers
drma:
	docker rm $(shell docker ps -a -q)

drmia:
	docker rmi $(shell docker images -q --filter "dangling=true")

drmvu:
	docker volume rm $(shell docker volume ls -qf dangling=true)
