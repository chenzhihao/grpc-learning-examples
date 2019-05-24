package rpcserver

import (
	"github.com/chenzhihao/grpc-showcase/product-server/pb"
	"io"
	"log"
	"time"
)

func (s *Server) Chat(stream pb.Product_ChatServer) error {
	for {
		time.Sleep(1 * time.Second)
		message, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			log.Printf("Chat finished at %v \n", endTime)
			err = stream.Send(&pb.Message{Message: "Bye"})
			return err
		}
		if err != nil {
			return err
		}

		log.Printf("Receive msg: %s \n", message.Message)
		err = stream.Send(&pb.Message{Message: "Server ACK"})
		if err != nil {
			return err
		}
	}
}
