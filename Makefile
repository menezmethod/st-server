PROTOC_GEN_GO := $(shell which protoc-gen-go)
PROTOC_GEN_GO_GRPC := $(shell which protoc-gen-go-grpc)
PROTOC_GEN_GO_GRPC_MOCK := $(shell which protoc-gen-go-grpc-mock)

.PHONY: all proto_gateway proto_auth proto_journal build_gateway build_auth build_journal

all: proto_gateway proto_auth proto_journal build_gateway build_auth build_journal

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

proto_gateway: check_protoc_plugins
	protoc pkg/gateway/pb/*.proto --go_out=. --go-grpc_out=. --go-grpc-mock_out=.

proto_auth: check_protoc_plugins
	protoc pkg/auth/pb/*.proto --go_out=. --go-grpc_out=. --go-grpc-mock_out=.

proto_journal: check_protoc_plugins
	protoc pkg/journal/pb/*.proto --go_out=. --go-grpc_out=. --go-grpc-mock_out=.

build_gateway:
	go build ./src/st-gateway/cmd/main.go

build_auth:
	go build ./src/st-auth-svc/cmd/main.go

build_journal:
	go build ./src/st-journal-svc/cmd/main.go