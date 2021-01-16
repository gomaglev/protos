SHELL := /bin/bash

PROTO_FILES ?= $(shell find $(PWD) -type f -path '*.proto' | grep -v "vendor")
PROTO_PB_FILES ?= $(shell find $(PWD) -type f -path '*.pb.go' | grep -v "vendor")

PROTOC := ${GOPATH}/bin/protoc
PROTOC_INJECT_TAG := ${GOPATH}/bin/protoc-go-inject-tag
PROTOS_DEST = $(PWD)
PROTOC_FLAGS ?= \
    -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
    -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
	--proto_path=${PROTOS_DEST} \
	--grpc-gateway_out=paths=source_relative,logtostderr=true:${PROTOS_DEST} \
	--go_out=paths=source_relative:$(PROTOS_DEST) \
	--go-grpc_out=paths=source_relative,require_unimplemented_servers=false:$(PROTOS_DEST) \
	--validate_out=paths=source_relative,lang=go:$(PROTOS_DEST) \
	--openapiv2_out $(PROTOS_DEST) --openapiv2_opt logtostderr=true \

protos: # generate protobuf files
	$(PROTOC) $(PROTOC_FLAGS) ${PROTO_FILES}

faker: # add faker tag to pb files
	$(PROTOC_INJECT_TAG) -input="${PROTO_PB_FILES}" -verbose=false
