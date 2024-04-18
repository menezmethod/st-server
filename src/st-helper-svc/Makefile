.PHONY: all proto server clean

all: proto server

# Path to the proto file
PROTO_FILE=./protos/helper.proto

# Generate Python and gRPC code from proto file
proto:
	python -m grpc_tools.protoc \
		-I./protos \
		--python_out=. \
		--grpc_python_out=. \
		$(PROTO_FILE)

# Run the gRPC server
server:
	python -m st_helper.server

# Clean generated files
clean:
	rm -f *_pb2.py *_pb2_grpc.py