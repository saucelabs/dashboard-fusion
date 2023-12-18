# Copyright 2023 Sauce Labs Inc., all rights reserved.

export GOBIN ?= $(CURDIR)/bin
export PATH  := $(GOBIN):$(PATH)

include .version

ifneq ($(shell expr $(MAKE_VERSION) \>= 4), 1)
$(error This Makefile requires GNU Make version 4 or higher, got $(MAKE_VERSION))
endif

ifneq ($(GO_VERSION),$(shell go version | grep -o -E '1\.[0-9\.]+'))
$(error Go version mismatch, got $(shell go version | grep -o -E '1\.[0-9\.]+'), expected $(GO_VERSION))
endif

.PHONY: install-dependencies
install-dependencies:
	@rm -Rf bin && mkdir -p $(GOBIN)
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)
