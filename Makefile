# Project settings
BINARY_NAME := taskmgr
BUILD_DIR   := bin
SRC_DIR     := ./cmd/$(BINARY_NAME)

# By default, build for the local OS/ARCH; override with e.g.
#   make GOOS=linux GOARCH=amd64 build
GOOS   ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

.PHONY: all build run test fmt vet lint install clean help

all: build

## Build the CLI binary
build:
	@echo "→ Building $(BINARY_NAME) for $(GOOS)/$(GOARCH)…"
	@mkdir -p $(BUILD_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)

## Run the freshly built binary (pass ARGS="…")
run: build
	@echo "→ Running $(BINARY_NAME)…"
	@$(BUILD_DIR)/$(BINARY_NAME) $(ARGS)

## Run all unit tests with verbose output
test:
	@echo "→ Running tests…"
	@go test ./... -v

## Format Go code
fmt:
	@echo "→ Formatting code (go fmt)…"
	@go fmt ./...

## Vet Go code
vet:
	@echo "→ Vetting code (go vet)…"
	@go vet ./...

## Lint = fmt + vet
lint: fmt vet

## Install into your $GOPATH/bin (or module bin dir)
install:
	@echo "→ Installing $(BINARY_NAME)…"
	@go install $(SRC_DIR)

## Clean out binaries
clean:
	@echo "→ Cleaning…"
	@rm -rf $(BUILD_DIR)

## Show help
help:
	@echo "Makefile targets:"
	@echo "  make           (alias for 'make build')"
	@echo "  make build     Build the CLI binary"
	@echo "  make run       Run the binary (set ARGS=\"…\" for args)"
	@echo "  make test      Run all tests"
	@echo "  make fmt       go fmt ./..."
	@echo "  make vet       go vet ./..."
	@echo "  make lint      fmt + vet"
	@echo "  make install   go install ./cmd/$(BINARY_NAME)"
	@echo "  make clean     rm -rf $(BUILD_DIR)"
	@echo "  make help      Show this message"
