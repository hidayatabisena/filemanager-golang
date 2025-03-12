.PHONY: build install test clean

BINARY_NAME=filemanager
GOFLAGS=-ldflags="-s -w"

build:
	@echo "Building $(BINARY_NAME)..."
	@go build $(GOFLAGS) -o $(BINARY_NAME) ./cmd/filemanager

install: build
	@echo "Installing $(BINARY_NAME)..."
	@go install ./cmd/filemanager

test:
	@echo "Running tests..."
	@go test -v ./...

clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
