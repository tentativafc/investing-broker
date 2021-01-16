package rpc

import (
	"context"
	"log"
	"net"

	"github.com/tentativafc/investing-broker/app/backend/sts-service/dto"
	errSts "github.com/tentativafc/investing-broker/app/backend/sts-service/error"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/service"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/stspb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	s service.StsService
}

func (s *Server) HandleError(err error) error {

	switch err.(type) {
	case *errSts.AuthError:
		error := err.(*errSts.AuthError)
		return status.Errorf(codes.PermissionDenied, error.Error())
	case *errSts.NotFoundError:
		error := err.(*errSts.NotFoundError)
		return status.Errorf(codes.NotFound, error.Error())
	case *errSts.BadRequestError:
		error := err.(*errSts.BadRequestError)
		return status.Errorf(codes.InvalidArgument, error.Error())
	case *errSts.GenericError:
		error := err.(*errSts.GenericError)
		return status.Errorf(codes.Internal, error.Error())
	}
	return status.Errorf(codes.Internal, err.(error).Error())

}

func (s *Server) GenerateToken(ctx context.Context, tr *stspb.TokenRequest) (*stspb.TokenResponse, error) {

	dtr := dto.TokenRequest{ClientId: tr.ClientId, ClientSecret: tr.ClientSecret}

	token, err := s.s.CreateToken(dtr)

	if err != nil {
		return nil, s.HandleError(err)
	}

	return &stspb.TokenResponse{Token: token}, nil
}

func (s *Server) ValidateToken(ctx context.Context, tr *stspb.ValidateTokenRequest) (*stspb.ValidateTokenResponse, error) {

	dreq := dto.ValidateTokenRequest{Token: tr.Token}

	dresp, err := s.s.ValidateToken(dreq)

	if err != nil {
		return nil, s.HandleError(err)
	}

	return &stspb.ValidateTokenResponse{Token: dresp.Token, ClientId: dresp.ClientId, ClientName: dresp.ClientName}, nil
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