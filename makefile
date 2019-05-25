.PHONY: protobuf-client
protobuf-client: ## Compile all the protocol buffer files to client folders
	 rm -rf ./product-client/pb;\
	 mkdir -p ./product-client/pb;\
	 protoc -I pb service.proto --go_out=plugins=grpc:product-client/pb;

.PHONY: protobuf-server
protobuf-server: ## Compile all the protocol buffer files to server folders
	 rm -rf ./product-server/pb;\
	 mkdir -p ./product-server/pb;\
	 protoc -I pb service.proto --go_out=plugins=grpc:product-server/pb;

.PHONY: protobuf
protobuf: protobuf-client protobuf-server

.PHONY: run-server
run-server:
	 @cd product-server  && echo "Running gRPC server" && go run main.go

.PHONY: server
server: protobuf-server run-server ## Run gRPC server

.PHONY: run-client
run-client:
	 @cd product-client && echo "Running gRPC client" && go run main.go

.PHONY: client
client: protobuf-client run-client

.PHONY: go-mod-client
go-mod-client:
	@cd product-client && go mod tidy

.PHONY: go-mod-server
go-mod-server:
	@cd product-server && go mod tidy

.PHONY: go-mod
go-mod: go-mod-client go-mod-server ## `go mod tidy` for both server and client

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

