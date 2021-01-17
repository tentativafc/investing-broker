package config

import (
	"os"
)

func GetDbConfig() string {
	var dbConection = os.Getenv("DB_CONNECTION")
	if len(dbConection) == 0 {
		dbConection = "host=localhost user=postgres password=123456 dbname=postgres port=5432"
	}
	return dbConection
}

func GetDbConfigSts() string {
	var dbConection = os.Getenv("DB_CONNECTION_STS")
	if len(dbConection) == 0 {
		dbConection = "host=localhost user=postgres password=123456 dbname=postgres port=5433"
	}
	return dbConection
}

func GetGrpcStsServer() string {

	var gRpcStsConnection = os.Getenv("GRPC_STS")
	if len(gRpcStsConnection) == 0 {
		gRpcStsConnection = "localhost:50051"
	}
	return gRpcStsConnection
}
