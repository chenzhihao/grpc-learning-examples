.PHONY: grpc

grpc: ## Compile all the protocol buffer files to service and client folders
	 protoc -I pb service.proto --go_out=plugins=grpc:product-service/pb
	 protoc -I pb service.proto --go_out=plugins=grpc:product-client/pb

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

