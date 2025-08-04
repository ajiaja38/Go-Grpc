package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=go_grpc port=5432 Timezone=Asia/Jakarta"), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err.Error())
	}

	return db
}
