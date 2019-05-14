package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/chenzhihao/grpc-showcase/product-client/pb"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductServiceClient(conn)

	clientDeadline := time.Now().Add(time.Duration(5000) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)

	defer cancel()
	r, err := c.GetProduct(ctx, &pb.ProductRequest{Id: "1"})
	if err != nil {
		log.Fatalf("can't get product: %v", err)
	}
	log.Printf("[RECEIVED RESPONSE]: %v\n", r)

	stream, err := c.GetProductStream(ctx, &pb.ProductListRequest{Id: []string{"1", "2", "3"}})
	if err != nil {
		log.Fatalf("can't get product list: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("receive error: %v", err)
		}
		log.Printf("[RECEIVED STREAM RESPONSE]: %v\n", resp) // 输出响应
	}
}
