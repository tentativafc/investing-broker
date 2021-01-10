package rpc

import (
	"context"
	"log"
	"net"

	"github.com/tentativafc/investing-broker/app/backend/sts-service/service"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/stspb"
	"google.golang.org/grpc"
)

type Server struct {
	s service.StsService
}

func (s *Server) GenerateToken(ctx context.Context, tr *stspb.TokenRequest) (*stspb.TokenResponse, error) {

	clientId := tr.ClientId
	clientSecret := tr.ClientSecret

	token, err := s.s.CreateToken(clientId, clientSecret)

	if err != nil {
		return nil, err
	}

	return &stspb.TokenResponse{Token: token}, nil
}

func StartServer(sts service.StsService) {
	server := Server{s: sts}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}

	s := grpc.NewServer()
	stspb.RegisterStsServer(s, &server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}
}
