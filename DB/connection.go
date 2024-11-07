package DB

import (
	"example/main/utils"
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnect() *gorm.DB  {
	dbHost := utils.GetEnv("DB_HOST")
	dbPort := utils.GetEnv("DB_PORT")
	dbUser := utils.GetEnv("DB_USER")
	dbPassword := utils.GetEnv("DB_PASSWORD")
	dbName := utils.GetEnv("DB_NAME")
	// dbSSLMode := utils.GetEnv("DB_SSLMODE")

	// PostgreSQL connection string format:
	// "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to Database %v", err)
		return nil
	}

	return db
}