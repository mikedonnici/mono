# MRV Monorepo


## Setting up

See <https://grpc.io/docs/languages/go/quickstart/>

- Install `protoc` (protocol buffer compiler)

Mac OS X

```shell
brew install protobuf 
```

Linux

```shell

```

- Install Go plugins

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

- Ensure go path 

```shell
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Style Guides

- [Protocol Buffers Files](https://developers.google.com/protocol-buffers/docs/style)


## Testing with `grpcurl`

