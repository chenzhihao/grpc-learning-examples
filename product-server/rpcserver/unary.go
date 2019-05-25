package rpcserver

import (
	"context"
	"github.com/chenzhihao/grpc-showcase/product-server/pb"

	"errors"
	"fmt"
	"log"
)

type Product struct {
	Name  string
	Price float64
}

var products = map[string]Product{
	"1": {
		Name:  "MacBook",
		Price: 2000,
	},
	"2": {
		Name:  "iPhone",
		Price: 800,
	},
	"3": {
		Name:  "airPods",
		Price: 200,
	},
}

type Server struct{}

func (s *Server) GetProduct(ctx context.Context, in *pb.GetProductRequest) (*pb.Product, error) {
	log.Printf("Received: %v \n", in.Id)
	if product, ok := products[in.Id]; ok == true {
		return &pb.Product{Name: product.Name, Price: product.Price}, nil
	}
	return nil, errors.New(fmt.Sprintf("no product found for ID %s", in.Id))
}
