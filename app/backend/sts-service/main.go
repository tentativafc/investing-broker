package main

import (
	"github.com/tentativafc/investing-broker/app/backend/sts-service/route"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/rpc"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/service"
)

func main() {

	s := service.NewStsService()

	go rpc.StartServer(s)
	route.CreateRoutes(s)

}
