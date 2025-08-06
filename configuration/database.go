package configuration

import (
	"log"
	"os"

	"github.com/go-pg/pg/v9"
)

func ConnectDB(config Configuration) *pg.DB {
	opts := &pg.Options{
		User:     config.DbConfig.DbUser,
		Password: config.DbConfig.DbPassword,
		Addr:     config.DbConfig.DbHost + ":" + config.DbConfig.DbPort,
		Database: config.DbConfig.DbName,
	}

	db := pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect to PotsgreSQL")
		os.Exit(100)
	}

	log.Printf("Connected to PostgreSQL")

	return db

}
