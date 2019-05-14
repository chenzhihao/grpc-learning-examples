package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/chenzhihao/grpc-showcase/product-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

type Product struct {
	Name  string
	Price float64
}

const (
	port = ":50051"
)

var products = map[string]Product{
	"1": {
		Name:  "MacBook",
		Price: 2000,
	},
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GetProduct(ctx context.Context, in *pb.ProductRequest) (*pb.ProductReply, error) {
	log.Printf("Received: %v", in.Id)
	if product, ok := products[in.Id]; ok == true {
		return &pb.ProductReply{Name: product.Name, Price: product.Price}, nil
	}
	return nil, errors.New(fmt.Sprintf("no product found for ID %s", in.Id))
}
