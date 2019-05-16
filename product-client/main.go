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
	c := pb.NewProductClient(conn)

	ch := make(chan bool, 3)
	go getProduct(c, ch)
	go getProductStream(c, ch)
	go chat(c, ch)
	<-ch
	<-ch
	<-ch
}

func chat(c pb.ProductClient, done chan<- bool) {
	clientDeadline := time.Now().Add(time.Duration(5000) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	stream, err := c.Chat(ctx)
	if err != nil {
		log.Fatal(err)
	}

	messages := []pb.Message{{Message: "zhihao"}, {Message: "test"}}

	go func() {
		defer cancel()
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				done <- true
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Got message %s \n", in.Message)
		}
	}()

	for _, msg := range messages {
		time.Sleep(1 * time.Second)
		err := stream.Send(&msg)
		if err != nil {
			log.Fatalf("Failed to send a message: %v", err)
		}
	}

	err = stream.CloseSend()
	if err != nil {
		log.Fatal(err)
	}
}

func getProduct(c pb.ProductClient, done chan<- bool) {
	clientDeadline := time.Now().Add(time.Duration(5000) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	r, err := c.GetProduct(ctx, &pb.ProductRequest{Id: "1"})
	if err != nil {
		log.Fatalf("can't get product: %v", err)
	}
	log.Printf("[RECEIVED RESPONSE]: %v\n", r)
	done <- true
}

func getProductStream(c pb.ProductClient, done chan<- bool) {
	clientDeadline := time.Now().Add(time.Duration(50000) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

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
		log.Printf("[RECEIVED STREAM RESPONSE]: %v\n", resp)
	}

	done <- true
}
