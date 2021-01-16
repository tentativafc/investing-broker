package config

import (
	"os"
)

func GetDbConfig() string {
	var dbConection = os.Getenv("db_connection")
	if len(dbConection) == 0 {
		dbConection = "host=localhost user=postgres password=123456 dbname=postgres port=5432"
	}
	return dbConection
}

func GetDbConfigSts() string {
	var dbConection = os.Getenv("db_connection_sts")
	if len(dbConection) == 0 {
		dbConection = "host=localhost user=postgres password=123456 dbname=postgres port=5433"
	}
	return dbConection
}
