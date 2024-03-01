PROTOC_GEN_GO := $(shell which protoc-gen-go)
PROTOC_GEN_GO_GRPC := $(shell which protoc-gen-go-grpc)
PROTOC_GEN_GO_GRPC_MOCK := $(shell which protoc-gen-go-grpc-mock)

.PHONY: all proto_gateway proto_auth proto_journal build_gateway build_auth build_journal install_deps run_auth run_journal run_gateway test help

all: install_deps proto_gateway proto_auth proto_journal build_gateway build_auth build_journal

install_deps:
	go mod download

prepare_mock_dirs:
	@mkdir -p ./pkg/auth/pb/mock
	@mkdir -p ./pkg/journal/pb/mock
	@mkdir -p ./st-auth-svc/pkg/pb/mock
	@mkdir -p ./st-journal-svc/pkg/pb/mock

check_protoc_plugins:
ifndef PROTOC_GEN_GO
	$(error "protoc-gen-go is not installed")
endif
ifndef PROTOC_GEN_GO_GRPC
	$(error "protoc-gen-go-grpc is not installed")
endif
ifndef PROTOC_GEN_GO_GRPC_MOCK
	$(error "protoc-gen-go-grpc-mock is not installed")
endif

proto_gw_auth: check_protoc_plugins prepare_mock_dirs
	protoc pkg/auth/pb/*.proto --go_out=. --go-grpc_out=. --go-grpc-mock_out=pkg/auth/pb/mock

proto_gw_journal: check_protoc_plugins prepare_mock_dirs
	protoc pkg/journal/pb/*.proto --go_out=. --go-grpc_out=. --go-grpc-mock_out=pkg/journal/pb/mock

proto_auth: check_protoc_plugins prepare_mock_dirs
	protoc st-auth-svc/pkg/pb/*.proto --go_out=. --go-grpc_out=. --go-grpc-mock_out=st-auth-svc/pkg/pb/mock

proto_journal: check_protoc_plugins prepare_mock_dirs
	protoc st-journal-svc/pkg/pb/*.proto --go_out=. --go-grpc_out=. --go-grpc-mock_out=st-journal-svc/pkg/pb/mock

build_gateway:
	go build -o ./bin/gateway ./src/st-gateway/cmd/main.go

build_auth:
	go build -o ./bin/auth ./src/st-auth-svc/cmd/main.go

build_journal:
	go build -o ./bin/journal ./src/st-journal-svc/cmd/main.go

run_auth:
	./bin/auth

run_journal:
	./bin/journal

run_gateway:
	./bin/gateway

test:
	go test ./...

help:
	@echo "Available commands:"
	@echo "  install_deps  - Install all dependencies"
	@echo "  proto         - Generate gRPC code for all services"
	@echo "  build         - Build all services"
	@echo "  run           - Run all services"
	@echo "  test          - Run tests"
	@echo "  help          - Show this help message"

proto: proto_gw_auth proto_gw_journal proto_auth proto_journal

build: build_gateway build_auth build_journal

run: run_auth run_journal run_gateway
