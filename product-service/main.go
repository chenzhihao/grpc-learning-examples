package main

import (
	"context"
	"log"
	"net"

	"github.com/chenzhihao/grpc-showcase/product-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

const (
	port = ":50051"
)

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
	return &pb.ProductReply{Name: "Fancy product", Price: 442.50}, nil
}
