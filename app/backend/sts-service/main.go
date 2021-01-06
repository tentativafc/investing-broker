package main

import (
	"context"
	"log"
	"net"

	"github.com/tentativafc/investing-broker/sts-service/sts"
	"github.com/tentativafc/investing-broker/sts-service/util"
	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) GenerateToken(ctx context.Context, tr *sts.TokenRequest) (*sts.TokenResponse, error) {

	clientId := tr.ClientId

	token, err := util.GenerateToken(clientId)

	if err != nil {
		return nil, err
	}

	return &sts.TokenResponse{Token: token}, nil
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
