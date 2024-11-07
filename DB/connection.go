package DB

import (
	"example/main/utils"
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB  {
	dbUser := utils.GetEnv("DB_USER")
	dbName := utils.GetEnv("DB_NAME")
	dbSSLMode := utils.GetEnv("DB_SSLMODE")
	dbHost := utils.GetEnv("DB_HOST")
	dbPassword := utils.GetEnv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbUser, dbPassword, dbName, dbSSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to Database %v", err)
		return nil
	}

	return db
}