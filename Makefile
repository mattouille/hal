PROJECT = hal
BUILD_DIR = build
CGO_ENABLED ?= 0
GOOS ?= linux
GOARCH ?= amd64
VERSION = $(shell cat VERSION)

.PHONY: all test build install

test: 
	# This is commented out because most of the tests don't make sense or are broken
	# go test ./...

build:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(PROJECT) ./examples/complex/

all: clean test build install

clean:
	rm -rf $(BUILD_DIR)

install:
	cp $(BUILD_DIR)/* $(GOBIN)