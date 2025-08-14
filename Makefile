# Binary name
BINARY_NAME=github_user_activity
BINARY_DIR=bin

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build directory
$(BINARY_DIR):
	mkdir -p $(BINARY_DIR)

# Build the project
.PHONY: build
build: $(BINARY_DIR)
	$(GOBUILD) -o $(BINARY_DIR)/$(BINARY_NAME) -v

# Clean build files
.PHONY: clean
clean:
	$(GOCLEAN)
	rm -rf $(BINARY_DIR)

# Run tests
.PHONY: test
test:
	$(GOTEST) -v ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	$(GOTEST) -v ./... -coverprofile=coverage.out
	$(GOCMD) tool cover -html=coverage.out

# Update dependencies
.PHONY: deps
deps:
	$(GOGET) -v -t -d ./...
	$(GOMOD) tidy

# Install the binary
.PHONY: install
install: build
	mv $(BINARY_DIR)/$(BINARY_NAME) $(GOPATH)/bin/

# Run the application (for development)
.PHONY: run
run:
	$(GOCMD) run main.go

# Help command
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build          - Build the binary"
	@echo "  make clean         - Clean build files"
	@echo "  make test          - Run tests"
	@echo "  make test-coverage - Run tests with coverage"
	@echo "  make deps          - Update dependencies"
	@echo "  make install       - Install the binary"
	@echo "  make run           - Run the application"
	@echo "  make help          - Show this help message"
