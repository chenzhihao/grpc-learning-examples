package main

import (
	"context"
	"log"
	"time"

	"github.com/chenzhihao/grpc-showcase/product-client/pb"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetProduct(ctx, &pb.ProductRequest{Id: "1"})
	if err != nil {
		log.Fatalf("can't get product: %v", err)
	}
	log.Printf("product: %s, price: %f", r.Name, r.Price)
}
