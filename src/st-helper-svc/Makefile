.PHONY: all proto server clean

# Variables
PROTO_FILE := ./protos/helper.proto
OUTPUT_DIR := .
INCLUDE_PATH := /Volumes/MenezSSD/Projects/st-server

# Targets
all: proto server

proto:
	python -m grpc_tools.protoc -I./protos -I$(INCLUDE_PATH) --python_out=$(OUTPUT_DIR) --grpc_python_out=$(OUTPUT_DIR) $(PROTO_FILE)

server:
	python server.py

clean:
	rm -f $(OUTPUT_DIR)/*_pb2.py $(OUTPUT_DIR)/*_pb2_grpc.py
