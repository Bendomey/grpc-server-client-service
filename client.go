package main

import (
	"grpc-service/chat"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect")
	}

	defer conn.Close()

	//define a chat service client
	c := chat.NewChatServiceClient(conn)
	//interace with grpc server
	message := chat.Message{
		Body: "Hello from the client!",
	}
	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling say hello %v", err)
	}

	log.Printf("Response from Server: %s", response.Body)
}
