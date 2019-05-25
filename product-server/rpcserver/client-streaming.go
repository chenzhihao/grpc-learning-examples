package rpcserver

import (
	"github.com/chenzhihao/grpc-showcase/product-server/pb"
	"io"
	"log"
	"time"
)

func (s *Server) CreateProducts(server pb.Warehouse_CreateProductsServer) error {
	for {
		request, err := server.Recv()
		if err == io.EOF {
			endTime := time.Now()
			log.Printf("Create products finished at %v \n", endTime)
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Create new product: %v", request.Product)
	}
}
