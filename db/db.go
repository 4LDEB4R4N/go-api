package db

import (
	"fmt"
	"os"
	"strconv"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

var Db *gorm.DB

func Connect(){

	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlSetup), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database")
		panic(err)
	} else {
		Db = db
		fmt.Println("Database connected")
	}


}