package main

import (
	"context"
	"log"
	"net"

	"github.com/tentativafc/investing-broker/sts-service/sts"
	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) GenerateToken(ctx context.Context, tr *sts.TokenRequest) (*sts.TokenResponse, error) {
	return &sts.TokenResponse{Token: "ABC"}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}

	s := grpc.NewServer()
	sts.RegisterStsServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}
}
