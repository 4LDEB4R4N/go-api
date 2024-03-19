package main

import (
	"api/db"
	"api/models"
	"fmt"
)

func init() {
	db.Connect()
}

func main() {
	db.Db.SetupJoinTable(&models.Customer{}, "Vehicles", &models.CustomerVehicle{})
	err := db.Db.AutoMigrate(&models.Customer{}, &models.Vehicle{})
	if err != nil {
		fmt.Println("Error migrating the database")
	} else {
		fmt.Println("Database migrated")
	}
}