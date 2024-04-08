package main

import (
	"api/controllers"
	"api/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	db.Connect()
}

func main() {
	router := gin.Default()

	router.POST("/customers", controllers.CreateCustomer)
	router.GET("/customers", controllers.GetAllCustomers)
	router.GET("/customers/:id", controllers.GetCustomer)
	router.PUT("/customers/:id", controllers.UpdateCustomer)
	router.DELETE("/customers/:id", controllers.DeleteCustomer)

	router.POST("/vehicles", controllers.CreateVehicle)
	router.GET("/vehicles", controllers.GetAllVehicles)
	router.GET("/vehicles/:id", controllers.GetVehicle)
	router.PUT("/vehicles/:id", controllers.UpdateVehicle)
	router.DELETE("/vehicles/:id", controllers.DeleteVehicle)

	router.POST("/customers/:id/vehicles/:vehicleID/links/:link", controllers.AddVehicleToCustomer)
	router.GET("/customers/:id/vehicles", controllers.ListVehiclesByCustomer)
	router.GET("/vehicles/links/:link", controllers.ListVehiclesByLink)

	router.GET("/customers/:id/vehicles/links/:link", controllers.ListVehiclesByLinkAndCustomer)

	router.DELETE("/customers/:id/vehicles/:vehicleID", controllers.RemoveVehicleFromCustomer)
	router.DELETE("/customers/:id/vehicles", controllers.RemoveAllVehiclesFromCustomer)

	err := router.Run(":8080")

	if err != nil {
		panic(err)
	}

	fmt.Println("Server is running on port 8080")
}


