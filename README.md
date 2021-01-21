# protos

`protos` defines shared proto pkg for all protomicro golang repos.

## protoc installation

Download protoc from https://github.com/protocolbuffers/protobuf/releases
Copy BOTH `protoc` file and `include` folder to `${GOPATH}/bin/protoc`.

Or you can use apt on Linux, brew on Mac.

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

## Protobuf generation

protos - Generate protobuf, grpc-gateway files:

```shell
make protos
```

faker - Generate faker tags in pb files for testing:

```shell
make faker
```

## Usage

To use protos in another repo, execute
`go get github.com/gomaglev/protos` to update go.mod.

To use specific branch execute
`go get github.com/gomaglev/protos@{commit hash}`

## Troubleshooting

`google/protobuf/timestamp.proto: File not found.`
make sure the `include` folder in protoc release package was copied to same folder as `protoc`.
