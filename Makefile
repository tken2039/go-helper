# primaries
APPLICATION_NAME:=go-helper
PKG_ROOT:=github.com/tken2039/go-helper
VERSION:=$(shell git describe --always)
REVISION:=$(shell git rev-parse --short HEAD)

# go build dir name
BIN_DIR:=bin

# go options
GOVERSION:=$(shell go version)
# GOOS:=$(word 1,$(subst /, ,$(lastword $(GOVERSION))))
# GOARCH:=$(word 2,$(subst /, ,$(lastword $(GOVERSION))))
GOOS_LINUX:=linux
GOARCH_AMD:=amd64

# go commands
CMD_GO_BUILD:=go build
CMD_GO_INSTALL:= go install
CMD_GO_FMT:=gofmt
CMD_GO_VET:=go vet

# stack name
GO_HELPER_NAME:=go_helper

# go binaries
GO_BINS:= $(BIN_DIR)/$(VERSION).$(GOOS_LINUX)-$(GOARCH_AMD)/go-helper

.PHONY: build gobin gobuild install check check-no-staticcheck clean 

# build go programs
build: gobin

# execute building go programs with setting go environment
gobin:
	@$(MAKE) gobuild GOOS=linux GOARCH=amd64 

# execute building go programs
gobuild: $(GO_BINS)

# execute by `gobuild` target
$(BIN_DIR)/$(VERSION).$(GOOS_LINUX)-$(GOARCH_AMD)/%: $(GO_FILES) .git/HEAD
	$(CMD_GO_BUILD) -o $@ -tags netgo -installsuffix netgo \
		-ldflags '-extldflags "-static" -s -w -X main.version=$(VERSION) -X main.revision=$(REVISION)' \
		$(PKG_ROOT)/cmd/gohelper

install:
	$(CMD_GO_INSTALL) -tags netgo -installsuffix netgo \
		-ldflags '-extldflags "-static" -s -w -X main.version=$(VERSION) -X main.revision=$(REVISION)' \
		$(PKG_ROOT)/cmd/gohelper

# check go sources
check:
	$(CMD_GO_FMT) -l $(shell find . -name "*.go") | xargs -I{} sh -c 'tess -z {} || echo "{}"; exit 1'
	$(CMD_GO_VET) ./...
	staticcheck ./... 

# check go sources without staticcheck
# this target is intended for use with GitHub Actions
check-no-staticcheck:
	$(CMD_GO_FMT) -l $(shell find . -name "*.go") | xargs -I{} sh -c 'tess -z {} || echo "{}"; exit 1'
	$(CMD_GO_VET) ./...

# execute rm bin and images
clean:
	-@rm -r $(BIN_DIR)/*
