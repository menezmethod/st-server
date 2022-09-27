proto auth-svc:
	protoc pkg/pb/*.proto -I=$GOPATH/src/st-auth-svc/vendor --gofast_out=plugins=grpc:. --proto_path=.

proto journal-svc:
	protoc pkg/pb/*.proto -I=$GOPATH/src/st-auth-svc/vendor --gofast_out=plugins=grpc:. --proto_path=.

proto gateway-auth:
	protoc pkg/auth/pb/*.proto -I=$GOPATH/src/st-auth-svc/vendor --gofast_out=plugins=grpc:. --proto_path=.

proto gateway-journal:
	protoc pkg/journal/pb/*.proto -I=$GOPATH/src/st-auth-svc/vendor --gofast_out=plugins=grpc:. --proto_path=.

gateway server:
	go build ./src/st-gateway/cmd/main.go

auth server:
	go build ./src/st-auth-svc/cmd/main.go

journal server:
	go build ./src/st-journal-svc/cmd/main.go

all servers:
	go build ./src/st-gateway/cmd/main.go ./src/st-auth-svc/cmd/main.go ./src/st-journal-svc/cmd/main.go