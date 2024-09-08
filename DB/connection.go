package DB

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"os"
	"log"
	"fmt"
)

func DBConnect() *gorm.DB  {
	if err := godotenv.Load(); err != nil {
    	log.Fatal("Error loading .env file")
  	}

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("user=%s dbname=%s sslmode=%s", dbUser, dbName, dbSSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to Database %v", err)
		return nil
	}

	return db
}