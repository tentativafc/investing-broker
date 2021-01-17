package main

import (
	"log"

	"github.com/tentativafc/investing-broker/app/backend/user-service/config"
	"github.com/tentativafc/investing-broker/app/backend/user-service/repo"
	"github.com/tentativafc/investing-broker/app/backend/user-service/route"
	"github.com/tentativafc/investing-broker/app/backend/user-service/service"
	"github.com/tentativafc/investing-broker/app/backend/user-service/stspb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial(config.GetGrpcStsServer(), grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect to sts server: %v", err)
	}

	sc := stspb.NewStsClient(cc)
	ur := repo.NewUserRepository()
	us := service.NewUserService(ur, sc)

	route.CreateRoutes(us)
}
