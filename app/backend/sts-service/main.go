package main

import (
	"github.com/tentativafc/investing-broker/app/backend/sts-service/repo"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/route"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/rpc"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/service"
)

func main() {
	r := repo.NewClientCredentialsRepository()
	ss := service.NewStsService(r)
	go rpc.StartServer(ss)
	route.CreateRoutes(ss)
}
