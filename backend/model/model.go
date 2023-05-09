package model

import (
	"errors"
	"fmt"
	"log"
	"os"
	model "pdfPro/model/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// My GORM database instance
var dbGorm *gorm.DB

func connectDb() error {

	// I don't need to check for the env variables since they are checked in the `main` function
	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("DATABASE_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return errors.New(fmt.Sprintf("failed to connect database: %s", err))
	}

	// migrate schemas
	db.AutoMigrate(&model.UserAccount{})

	dbGorm = db

	return nil
}

func InitDatabase() error {
	if dbGorm == nil {
		err := connectDb()
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

// Returns the secret used to sign JWT web tokens
func GetJwtSecret() (string, error) {
	secret, isSet := os.LookupEnv("JWT_SECRET")
	if isSet == false || secret == "" {
		return "", errors.New("The env variable JWT_SECRET is not set or is set to an empty value")
	}

	return secret, nil
}
