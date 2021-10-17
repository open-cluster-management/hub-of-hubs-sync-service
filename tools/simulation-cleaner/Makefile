# Copyright IBM Corp All Rights Reserved.
# Copyright London Stock Exchange Group All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#
# -------------------------------------------------------------
# This makefile defines the following targets
#
#   - all (default) - formats the code, runs liners, downloads vendor libs, and builds executable
#   - fmt - formats the code
#   - vendor - download all third party libraries and puts them inside vendor directory
#   - clean-vendor - removes third party libraries from vendor directory
#   - build - builds the controller
#   - clean - cleans the build directories
#   - clean-all - superset of 'clean' that also removes vendor dir
#   - lint - runs code analysis tools

COMPONENT := $(shell basename $(shell pwd))

.PHONY: all				##formats the code, runs liners, downloads vendor libs, and builds executable
all: clean-vendor fmt lint vendor build

.PHONY: fmt				##formats the code
fmt:
	@gci -w .
	@go fmt ./...
	@gofumpt -w .

.PHONY: vendor			##download all third party libraries and puts them inside vendor directory
vendor:
	@go mod vendor

.PHONY: clean-vendor			##removes third party libraries from vendor directory
clean-vendor:
	-@rm -rf vendor

.PHONY: build			##builds the controller
build:
	@go build -o bin/${COMPONENT} main.go

.PHONY: clean			##cleans the build directories
clean:
	@rm -rf bin

.PHONY: clean-all			##superset of 'clean' that also removes vendor dir
clean-all: clean-vendor clean

.PHONY: lint				##runs code analysis tools
lint: clean-vendor
	go vet ./...
	golint ./...
	golangci-lint run ./...
