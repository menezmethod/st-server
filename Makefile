# Detect executables
PROTOC_GEN_GO := $(shell which protoc-gen-go)
PROTOC_GEN_GO_GRPC := $(shell which protoc-gen-go-grpc)
GOPATH := $(shell go env GOPATH)

# Define phony targets
.PHONY: all install_deps check_protoc_plugins proto generate_protos build run test help

# Default target
all: install_deps proto build

# Install dependencies
install_deps:
	go mod download

# Check if protoc plugins are installed
check_protoc_plugins:
ifndef PROTOC_GEN_GO
	$(error "protoc-gen-go is not installed")
endif
ifndef PROTOC_GEN_GO_GRPC
	$(error "protoc-gen-go-grpc is not installed")
endif

# Generate protobuf files for different services
proto: check_protoc_plugins proto_gateway proto_auth proto_journal generate_protos

proto_gateway:
	PATH="$(PATH):$(GOPATH)/bin" protoc src/st-gateway/versions/v1/auth/auth.proto \
	src/st-gateway/versions/v1/helper/helper.proto \
	src/st-gateway/versions/v1/journal/journal.proto \
	src/st-gateway/versions/v1/record/record.proto \
	--proto_path=src/st-gateway/versions/v1 \
	--proto_path=. \
	--go_out=paths=source_relative:src/st-gateway/pkg/pb \
	--go-grpc_out=paths=source_relative:src/st-gateway/pkg/pb \
	--grpc-gateway_out=paths=source_relative:src/st-gateway/pkg/pb

proto_auth:
	PATH="$(PATH):$(GOPATH)/bin" protoc st-auth-svc/pkg/pb/*.proto --go_out=. --go-grpc_out=.

proto_journal:
	PATH="$(PATH):$(GOPATH)/bin" protoc st-journal-svc/pkg/pb/*.proto --go_out=. --go-grpc_out=.

# Build executables for each service
build: build_gateway build_auth build_journal

build_gateway:
	go build -o ./bin/gateway ./src/st-gateway/cmd/main.go

build_auth:
	go build -o ./bin/auth ./src/st-auth-svc/cmd/main.go

build_journal:
	go build -o ./bin/journal ./src/st-journal-svc/cmd/main.go

# Run services
run: run_auth run_journal run_gateway

run_auth:
	./bin/auth

run_journal:
	./bin/journal

run_gateway:
	./bin/gateway

# Run tests across all go files
test:
	go test ./...

# Display help
help:
	@echo "Available commands:"
	@echo "  install_deps    - Install all dependencies"
	@echo "  proto           - Generate gRPC code for all services"
	@echo "  build           - Build all services"
	@echo "  run             - Run all services"
	@echo "  test            - Run tests"
	@echo "  help            - Show this help message"
