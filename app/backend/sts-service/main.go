package main

import (
	"fmt"

	"github.com/tentativafc/investing-broker/app/backend/sts-service/repo"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/route"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/rpc"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/service"
)

func main() {
	fmt.Println("Starting repository...")
	r := repo.NewClientCredentialsRepository()
	fmt.Println("Starting Service...")
	ss := service.NewStsService(r)
	fmt.Println("Starting GRPC Server...")
	go rpc.StartServer(ss)
	fmt.Println("Starting Rest Server...")
	route.CreateRoutes(ss)
}
