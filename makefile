.PHONY: grpc-client
grpc-client: ## Compile all the protocol buffer files to client folders
	 rm -rf ./product-client/pb;\
	 mkdir -p ./product-client/pb;\
	 protoc -I pb service.proto --go_out=plugins=grpc:product-client/pb;

.PHONY: grpc-server
grpc-server: ## Compile all the protocol buffer files to server folders
	 rm -rf ./product-server/pb;\
	 mkdir -p ./product-server/pb;\
	 protoc -I pb service.proto --go_out=plugins=grpc:product-server/pb;

.PHONY: grpc
grpc: grpc-client grpc-server

.PHONY: run-server
run-server:
	 cd product-server && go run main.go && echo "Running gRPC server"
.PHONY: server
server: grpc-server run-server ## Run gRPC server

.PHONY: run-client
run-client:
	 cd product-client && go run main.go && echo "Running gRPC client"
.PHONY: client
client: grpc-client run-client ## Run gRPC client

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

