proto gateway:
	protoc pkg/**/pb/*.proto --go_out=. --go-grpc_out=. --go-grpc-mock_out=.

proto auth:
	protoc pkg/pb/*.proto --go_out=. --go-grpc_out=. --go-grpc-mock_out=.

proto journal:
	protoc pkg/pb/*.proto --go_out=. --go-grpc_out=. --go-grpc-mock_out=.

gateway server:
	go build ./src/st-gateway/cmd/main.go

auth server:
	go build ./src/st-auth-svc/cmd/main.go

journal server:
	go build ./src/st-journal-svc/cmd/main.go