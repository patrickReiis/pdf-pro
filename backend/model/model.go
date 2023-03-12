package model

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbGorm, err = connectDb()

func connectDb() (*gorm.DB, error) {

	// I don't need to check for the env variables since they are checked in the `main` function
	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("DATABASE_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return db, errors.New(fmt.Sprintf("failed to connect database: %s", err))
	}

	return db, nil
}

// Returns the user by its Api Key
// If the user does not exists returns an empty string
func GetUserByApiKey(apiKey string) (user string) {
	return getUserByApiKeyImpl(apiKey)
}

func getUserByApiKeyImpl(apiKey string) string {
	return "implement"
}
