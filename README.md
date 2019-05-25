# gRPC Showcase

The purpose of this repo is providing some simple examples for people to learn gRPC.

Including:

- [x] [unary gRPC](product-server/rpcserver/unary.go)

- [x] [server streaming gRPC](product-server/rpcserver/server-streaming.go)

- [x] [client streaming gRPC](product-server/rpcserver/client-streaming.go)

- [x] [bidirectional streaming gRPC](product-server/rpcserver/bidirectional-streaming.go)

- [ ] gRPC interceptor

- [ ] Authentication

- [ ] [gRPC-web](https://github.com/grpc/grpc-web)

- [ ] gRPC [REST Gateway](https://github.com/grpc-ecosystem/grpc-gateway)



### Dependencies

[Golang](https://golang.org/) 1.12.*

[protobuf](https://github.com/protocolbuffers/protobuf) 

### How-to
```bash
## get the help
make help

### run gRPC server
make run-server

### run gRPC client
make run-client
```

