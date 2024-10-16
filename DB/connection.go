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

	dsn := fmt.Sprintf("user=%s dbname=%s sslmode=%s", dbUser, dbName, dbSSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to Database %v", err)
		return nil
	}

	return db
}