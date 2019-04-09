grpc: 
	 protoc -I pb service.proto --go_out=plugins=grpc:product-service/pb
	 protoc -I pb service.proto --go_out=plugins=grpc:product-client/pb
