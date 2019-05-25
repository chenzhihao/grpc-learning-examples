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
	c := pb.NewWarehouseClient(conn)

	ch := make(chan bool, 4)
	go getProduct(c, ch)
	go listProducts(c, ch)
	go chat(c, ch)
	go createProducts(c, ch)
	<-ch
	<-ch
	<-ch
	<-ch
}

func chat(c pb.WarehouseClient, done chan<- bool) {
	clientDeadline := time.Now().Add(time.Duration(5000) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	chatClient, err := c.Chat(ctx)
	if err != nil {
		log.Fatal(err)
	}

	messages := []pb.Message{{Message: "zhihao"}, {Message: "test"}}

	go func() {
		defer cancel()
		for {
			in, err := chatClient.Recv()
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
		err := chatClient.Send(&msg)
		if err != nil {
			log.Fatalf("Failed to send a message: %v", err)
		}
	}

	err = chatClient.CloseSend()
	if err != nil {
		log.Fatal(err)
	}
}

func getProduct(c pb.WarehouseClient, done chan<- bool) {
	clientDeadline := time.Now().Add(time.Duration(5000) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	r, err := c.GetProduct(ctx, &pb.GetProductRequest{Id: "1"})
	if err != nil {
		log.Fatalf("can't get product: %v", err)
	}
	log.Printf("[RECEIVED RESPONSE]: %v\n", r)
	done <- true
}

func listProducts(c pb.WarehouseClient, done chan<- bool) {
	clientDeadline := time.Now().Add(time.Duration(50000) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	stream, err := c.ListProducts(ctx, &pb.ListProductsRequest{Requests: []*pb.GetProductRequest{{Id: "1"}, {Id: "2"}}})
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

func createProducts(c pb.WarehouseClient, done chan<- bool) {
	clientDeadline := time.Now().Add(time.Duration(5000) * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
	defer cancel()

	products := []pb.Product{{Name: "new p1", Price: 100}, {Name: "new p2", Price: 200}}
	createProductClient, err := c.CreateProducts(ctx)
	if err != nil {
		log.Fatalf("can't create product: %v", err)
	}

	for _, product := range products {
		err := createProductClient.Send(&pb.CreateProductRequest{Product: &product})
		if err != nil {
			log.Fatalf("receive error: %v", err)
		}
	}
	_, err = createProductClient.CloseAndRecv()
	if err == io.EOF {
		done <- true
		return
	}

	if err != nil {
		log.Fatal(err)
	}
}
