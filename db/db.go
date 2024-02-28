package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Database instance
var db *sql.DB

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func Connect() error {
	var err error
	godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
		return err
	}
	host := goDotEnvVariable("db_host")
	port := goDotEnvVariable("db_port")
	user := goDotEnvVariable("db_user")
	password := goDotEnvVariable("db_pass")
	dbname := goDotEnvVariable("db_name")

	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}
