package main

import (
	"context"
	"log"
	"net"

	"github.com/tentativafc/investing-broker/sts-service/stspb"
	"github.com/tentativafc/investing-broker/sts-service/util"
	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) GenerateToken(ctx context.Context, tr *stspb.TokenRequest) (*stspb.TokenResponse, error) {

	clientId := tr.ClientId

	token, err := util.CreateToken(clientId)

	if err != nil {
		return nil, err
	}

	return &stspb.TokenResponse{Token: token}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}

	s := grpc.NewServer()
	stspb.RegisterStsServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}
}
