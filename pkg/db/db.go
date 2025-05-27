package db

import (
	"app/test/configs"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(config *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(config.Db.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("[DB] - [NewDb] - [ERROR] %s", err)
	}

	log.Printf("[DB] - [NewDb] - [INFO] connected to database")

	return &Db{DB: db}
}
