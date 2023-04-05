package model

import (
	"errors"
	"fmt"
	"os"
	model "pdfPro/model/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// My GORM database instance
var dbGorm, _ = connectDb() // Ignoring possible returned error since it is checked in the `model_test.go` file

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
		return nil, errors.New(fmt.Sprintf("failed to connect database: %s", err))
	}

	// migrate schemas
	db.AutoMigrate(&model.UserAccount{})

	return db, nil
}
