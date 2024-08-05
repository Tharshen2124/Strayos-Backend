package DB

import (
	"database/sql"
	"github.com/joho/godotenv"
	"os"
	"log"
	"fmt"
	_ "github.com/lib/pq"
)

func DBConnect() *sql.DB  {
	if err := godotenv.Load(); err != nil {
    	log.Fatal("Error loading .env file")
  	}

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s", dbUser, dbName, dbSSLMode)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Failed to connect to Database %v", err)
		return nil
	}

	return db
}