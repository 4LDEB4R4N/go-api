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
	// drop the table if it exists
	db.Db.Migrator().DropTable(&models.Customer{}, &models.Vehicle{}, &models.CustomerVehicle{})
	// migrate the database
	err := db.Db.AutoMigrate(&models.Customer{}, &models.Vehicle{}, &models.CustomerVehicle{})
	if err != nil {
		fmt.Println("Error migrating the database")
	} else {
		fmt.Println("Database migrated")
	}
}