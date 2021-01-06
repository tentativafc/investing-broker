package main

import (
	"log"
	"net"

	"github.com/tentativafc/investing-broker/sts/sts"
	"google.golang.org/grpc"
)

type server struct{}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}

	s := grpc.NewServer()
	sts.RegisterStsServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}
}
