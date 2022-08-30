proto gateway:
	protoc pkg/**/=pb/*.proto --go_out=. --go-grpc_out=.

proto auth:
	protoc pkg/**/pb/*.proto --go_out=. --go-grpc_out=.

proto journal:
	protoc pkg/**/pb/*.proto --go_out=. --go-grpc_out=.

gateway server:
	go run ./src/st-gateway/cmd/main.go

auth server:
	go run ./src/st-auth-svc/cmd/main.go

journal server:
	go run ./src/st-journal-svc/cmd/main.go