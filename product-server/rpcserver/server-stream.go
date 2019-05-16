package rpcserver

import (
	"errors"
	"fmt"
	"github.com/chenzhihao/grpc-showcase/product-server/pb"
	"log"
	"time"
)

func (s *Server) GetProductStream(in *pb.ProductListRequest, stream pb.Product_GetProductStreamServer) error {
	log.Printf("Received: %v \n", in.Id)
	for _, Id := range in.Id {
		if product, ok := products[Id]; ok == true {
			err := stream.Send(&pb.ProductReply{Name: product.Name, Price: product.Price})
			time.Sleep(1 * time.Second)
			if err != nil {
				return err
			}
		} else {
			return errors.New(fmt.Sprintf("no product found for ID %s", Id))
		}
	}

	return nil
}
