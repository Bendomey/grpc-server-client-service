package main

import (
	"log"
	"net"

	"grpc-service/chat"

	"google.golang.org/grpc"
)

// protoc --go_out=plugins=grpc:chat chat.proto : for generating code in terminal

func main() {
	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Could not listen on on port 5000: %v", err)
	}
	s := chat.Server{}

	gprcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(gprcServer, &s)

	if err := gprcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve grpc %v", err)

	}
}
