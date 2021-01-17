package main

import (
	"fmt"
	"log"

	"github.com/tentativafc/investing-broker/app/backend/user-service/config"
	"github.com/tentativafc/investing-broker/app/backend/user-service/repo"
	"github.com/tentativafc/investing-broker/app/backend/user-service/route"
	"github.com/tentativafc/investing-broker/app/backend/user-service/service"
	"github.com/tentativafc/investing-broker/app/backend/user-service/stspb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Start GRPC client...")
	cc, err := grpc.Dial(config.GetGrpcStsServer(), grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect to sts server: %v", err)
	}
	sc := stspb.NewStsClient(cc)
	fmt.Println("Starting repository...")
	ur := repo.NewUserRepository()
	fmt.Println("Starting Service...")
	us := service.NewUserService(ur, sc)
	fmt.Println("Starting Rest Server...")
	route.CreateRoutes(us)
}
