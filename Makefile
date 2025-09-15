# Makefile for serein-cli

# Use git describe to get a version string.
# Example: v1.0.0-3-g1234567
# Fallback to 'dev' if not in a git repository.
VERSION ?= $(shell git describe --tags --always --dirty --first-parent 2>/dev/null || echo "dev")

# Go parameters
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_RUN=$(GO_CMD) run
GO_FMT=$(GO_CMD) fmt
GO_CLEAN=$(GO_CMD) clean
GO_INSTALL=$(GO_CMD) install

# Binary name
BINARY_NAME=serein
NIX_BUILD=result

# Build flags
LDFLAGS = -ldflags="-s -w -X main.version=$(VERSION)"

.PHONY: all build run fmt clean install

all: build

build:
	@echo "Building $(BINARY_NAME) version $(VERSION)..."
	$(GO_BUILD) $(LDFLAGS) -o $(BINARY_NAME) .

run:
	$(GO_RUN) . --

fmt:
	@echo "Formatting code..."
	$(GO_FMT) ./...

clean:
	@echo "Cleaning..."
	$(GO_CLEAN)
	rm -f $(BINARY_NAME)
	rm -rf $(NIX_BUILD)

install:
	@echo "Installing $(BINARY_NAME) to $(shell $(GO_CMD) env GOPATH)/bin..."
	$(GO_INSTALL) .
