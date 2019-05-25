package rpcserver

import (
	"errors"
	"fmt"
	"github.com/chenzhihao/grpc-showcase/product-server/pb"
	"log"
	"time"
)

func (s *Server) ListProducts(in *pb.ListProductsRequest, stream pb.Warehouse_ListProductsServer) error {
	log.Printf("Received: %v \n", in.Requests)
	for _, Request := range in.Requests {
		if product, ok := products[Request.Id]; ok == true {
			err := stream.Send(&pb.Product{Name: product.Name, Price: product.Price})
			time.Sleep(1 * time.Second)
			if err != nil {
				return err
			}
		} else {
			return errors.New(fmt.Sprintf("no product found for ID %s", Request.Id))
		}
	}

	return nil
}
