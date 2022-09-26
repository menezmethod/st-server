protoc pkg/**/pb/*.proto --go_out=. --go-grpc_out=.
	protoc pkg/**/pb/*.proto -I=$GOPATH/src/ --gofast_out=plugins=grpc:. --proto_path=.


proto gateway:
	protoc pkg/**/pb/*.proto -I=$GOPATH/src/ --gofast_out=plugins=grpc:. --proto_path=.

proto auth:
	protoc pkg/pb/*.proto -I=$GOPATH/src/ --gofast_out=plugins=grpc:. --proto_path=.

proto journal:
	protoc pkg/pb/*.proto -I=$GOPATH/src/ --gofast_out=plugins=grpc:. --proto_path=.

gateway server:
	go build ./src/st-gateway/cmd/main.go

auth server:
	go build ./src/st-auth-svc/cmd/main.go

journal server:
	go build ./src/st-journal-svc/cmd/main.go

all servers:
	go build ./src/st-gateway/cmd/main.go ./src/st-auth-svc/cmd/main.go ./src/st-journal-svc/cmd/main.go