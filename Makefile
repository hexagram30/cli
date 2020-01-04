VERSION = $(shell cat VERSION)
GO ?= go
GOFMT ?= $(GO)fmt

DEFAULT_GOPATH=$(shell echo $$GOPATH|tr ':' '\n'|awk '!x[$$0]++'|sed '/^$$/d'|head -1)
ifeq ($(DEFAULT_GOPATH),)
DEFAULT_GOPATH := ~/go
endif
DEFAULT_GOBIN=$(DEFAULT_GOPATH)/bin
export PATH:=$(PATH):$(DEFAULT_GOBIN)

GOLANGCI_LINT=$(DEFAULT_GOBIN)/golangci-lint
RICH_GO = $(DEFAULT_GOBIN)/richgo

GODOC=godoc -index -links=true -notes="BUG|TODO|XXX|ISSUE"

DVCS_HOST = github.com
DVCS_ORG = hexagram30
PROJ = cli
BIN=hxgm30
CLIENT=rasterc

FQ_PROJ = $(DVCS_HOST)/$(DVCS_ORG)/$(PROJ)

DOCKER_ORG = hexagram30
DOCKER_DIR = build/docker

COMMIT_ID = $(shell git rev-parse --short HEAD)

LD_VERSION = -X $(FQ_PROJ)/common.vsn=$(VERSION)
LD_BUILDDATE = -X $(FQ_PROJ)/common.buildDate=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LD_GITCOMMIT = -X $(FQ_PROJ)/common.gitCommit=$(COMMIT_ID)
LD_GITBRANCH = -X $(FQ_PROJ)/common.gitBranch=$(shell git rev-parse --abbrev-ref HEAD)
LD_GITSUMMARY = -X $(FQ_PROJ)/common.gitSummary=$(shell git describe --tags --dirty --always)

LDFLAGS = -w -s $(LD_VERSION) $(LD_BUILDDATE) $(LD_GITBRANCH) $(LD_GITSUMMARY) $(LD_GITCOMMIT)

default: all

#############################################################################
###   Custom Installs   #####################################################
#############################################################################

GOLANGCI_LINT = $(DEFAULT_GOBIN)/golangci-lint
RICH_GO = $(DEFAULT_GOBIN)/richgo
GODA = $(DEFAULT_GOBIN)/goda

$(GOLANGCI_LINT):
	@echo ">> Couldn't find $(GOLANGCI_LINT); installing ..."
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
	sh -s -- -b $(DEFAULT_GOBIN) v1.21.0

$(RICH_GO):
	@echo ">> Couldn't find $(RICH_GO); installing ..."
	@GOPATH=$(DEFAULT_GOPATH) \
	GOBIN=$(DEFAULT_GOBIN) \
	GO111MODULE=on \
	$(GO) get -u github.com/kyoh86/richgo

$(GODA):
	@echo ">> Couldn't find $(GODA); installing ..."
	@GOPATH=$(DEFAULT_GOPATH) \
	GOBIN=$(DEFAULT_GOBIN) \
	GO111MODULE=on \
	$(GO) get -u github.com/loov/goda

#############################################################################
###   Build   ###############################################################
#############################################################################

all: clean lint test build

build: build-cli

# Set up build env:
# export GOPATH=~/go:~/lab/hexagram30/go
# export GOBIN=~/go/bin:~/lab/hexagram30/go/bin
# export PATH=$PATH:$GOBIN

bin:
	@mkdir ./bin

bin/$(BIN): bin
	@echo '>> Building CLI binary'
	@GO111MODULE=on $(GO) build -ldflags "$(LDFLAGS)" -o bin/$(BIN) ./cmd/$(BIN)

build-cli: | bin/$(BIN)

clean:
	@echo ">> Removing compiled binary files ..."
	@rm -f bin/*

clean-all: clean

install:
	@echo ">> Installing to $(DEFAULT_GOBIN) ..."
	@cp bin/$(BIN) $(DEFAULT_GOBIN)/

#############################################################################
###   Docker   ##############################################################
#############################################################################


#############################################################################
###   Linting & Testing   ###################################################
#############################################################################

show-linter:
	@echo $(GOLANGCI_LINT)

lint-silent: $(GOLANGCI_LINT)
	@$(GOLANGCI_LINT) \
	--enable=typecheck \
	--enable=golint \
	--enable=gocritic \
	--enable=misspell \
	--enable=nakedret \
	--enable=unparam \
	--enable=lll \
	--enable=goconst \
	--out-format=colored-line-number \
	run ./...

lint:
	@echo '>> Linting source code'
	@GO111MODULE=on $(MAKE) lint-silent

test: $(RICH_GO)
	@echo '>> Running all tests'
	@GO111MODULE=on $(RICH_GO) test ./... -v

test-nocolor:
	@echo '>> Running all tests'
	@GO111MODULE=on $(GO) test ./... -v

#############################################################################
###   Release Process   #####################################################
#############################################################################

tag:
	@echo "Tags:"
	@git tag|tail -5
	@git tag "v$(VERSION)"
	@echo "New tag list:"
	@git tag|tail -6

tag-and-push: tag
	@git push --tags

tag-delete: VERSION ?= 0.0.0
tag-delete:
	@git tag --delete v$(VERSION)
	@git push --delete origin v$(VERSION)

#############################################################################
###   Misc   ################################################################
#############################################################################

version:
	@echo $(VERSION)

commit-id:
	@echo $(COMMIT_ID)

show-ldflags:
	@echo $(LDFLAGS)

clean-cache:
	@echo '>> Purging Go mod cahce ...'
	@$(GO) clean -cache
	@$(GO) clean -modcache

rebuild-all: clean-all clean-cache all

check-modules:
	@echo '>> Checking modules ...'
	@GO111MODULE=on $(GO) mod tidy
	@GO111MODULE=on $(GO) mod verify

update-modules:
	@GO111MODULE=on go get -u ./...

list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | \
	awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | \
	sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

.PHONY: default build
