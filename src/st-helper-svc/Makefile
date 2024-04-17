.PHONY: all proto server clean

all: proto server

# Path to the proto file
PROTO_FILE=./protos/helper.proto
PYTHON_OUT_DIR=./helper

# Generate Python code from proto file
proto:
   python -m grpc_tools.protoc --proto_path=protos --python_out=$(PYTHON_OUT_DIR) --grpc_python_out=$(PYTHON_OUT_DIR) $(PROTO_FILE)

# Run the gRPC server
server:
   python -m st_helper.server

# Clean generated files
clean:
   rm -f $(PYTHON_OUT_DIR)/*_pb2.py $(PYTHON_OUT_DIR)/*_pb2_grpc.py