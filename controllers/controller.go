package controllers

import (
	"api/db"
	"api/models"

	"github.com/gin-gonic/gin"
)

func CreateCustomer(ctx *gin.Context) {

	var body struct {
		Name string
		CNPJ string
	}

	ctx.Bind(&body)

	customer := models.Customer{Name: body.Name, CNPJ: body.CNPJ}
	result := db.Db.Create(&customer)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		ctx.JSON(200, gin.H{"customer": map[string]interface{}{
			"ID":   customer.ID,
			"Name": customer.Name,
			"CNPJ": customer.CNPJ,
		}})
	}

}

func GetAllCustomers(ctx *gin.Context) {

	var customers []models.Customer

	err := db.Db.Model(&models.Customer{}).Preload("Vehicles").Find(&customers).Error

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, gin.H{"customers": customers})
	}

}

func GetCustomer(ctx *gin.Context) {

	var customer models.Customer

	id := ctx.Param("id")

	err := db.Db.Model(&models.Customer{}).Preload("Vehicles").First(&customer, id)

	if err.Error != nil {
		ctx.JSON(500, gin.H{"error": err.Error.Error()})
	} else {
		ctx.JSON(200, gin.H{"customer": customer})
	}

}

func UpdateCustomer(ctx *gin.Context) {

	var body struct {
		Name string
		CNPJ string
	}

	ctx.Bind(&body)

	id := ctx.Param("id")

	var customer models.Customer

	db.Db.First(&customer, id)

	customer.Name = body.Name
	customer.CNPJ = body.CNPJ

	err := db.Db.Save(&customer)

	if err.Error != nil {
		ctx.JSON(500, gin.H{"error": err.Error.Error()})
	} else {
		ctx.JSON(200, gin.H{"customer": customer})
	}

}

func DeleteCustomer(ctx *gin.Context) {

	var customer models.Customer
	id := ctx.Param("id")

	err := db.Db.Delete(&customer, id)

	if err.Error != nil {
		ctx.JSON(500, gin.H{"error": err.Error.Error()})
	} else {
		ctx.JSON(200, gin.H{"customer": customer})
	}

}

func CreateVehicle(ctx *gin.Context) {

	var body struct {
		LicensePlate string
		Color        string
	}

	ctx.Bind(&body)

	vehicle := models.Vehicle{LicensePlate: body.LicensePlate, Color: body.Color}
	result := db.Db.Create(&vehicle)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		ctx.JSON(200, gin.H{"vehicle": vehicle})
	}

}

func DeleteVehicle(ctx *gin.Context) {

	id := ctx.Param("id")

	var vehicle models.Vehicle

	err := db.Db.Delete(&vehicle, id)

	if err.Error != nil {
		ctx.JSON(500, gin.H{"error": err.Error.Error()})
	} else {
		ctx.JSON(200, gin.H{"message": vehicle})
	}
}

func UpdateVehicle(ctx *gin.Context) {

	var body struct {
		LicensePlate string
		Color        string
	}

	ctx.Bind(&body)

	id := ctx.Param("id")

	var vehicle models.Vehicle

	db.Db.First(&vehicle, id)

	vehicle.LicensePlate = body.LicensePlate
	vehicle.Color = body.Color

	err := db.Db.Save(&vehicle)

	if err.Error != nil {
		ctx.JSON(500, gin.H{"error": err.Error.Error()})
	} else {
		ctx.JSON(200, gin.H{"vehicle": vehicle})
	}

}

func GetAllVehicles(ctx *gin.Context) {

	var vehicles []models.Vehicle

	err := db.Db.Find(&vehicles)

	if err.Error != nil {
		ctx.JSON(500, gin.H{"error": err.Error.Error()})
	} else {
		ctx.JSON(200, gin.H{"vehicles": vehicles})
	}
}

func GetVehicle(ctx *gin.Context) {

	var vehicle models.Vehicle

	id := ctx.Param("id")

	err := db.Db.First(&vehicle, id)

	if err.Error != nil {
		ctx.JSON(500, gin.H{"error": err.Error.Error()})
	} else {
		ctx.JSON(200, gin.H{"vehicle": vehicle})
	}
}

func AddVehicleToCustomer(ctx *gin.Context) {

	id := ctx.Param("id")
	vehicleID := ctx.Param("vehicleID")
	link := ctx.Param("link")

	if !models.IsValidLink(string(link)) {
		ctx.JSON(400, gin.H{"error": "Invalid link value"})
		return
	}

	var customer models.Customer
	var vehicle models.Vehicle

	db.Db.First(&customer, id)
	db.Db.First(&vehicle, vehicleID)

	customerVehicle := models.CustomerVehicle{
		CustomerID: customer.ID,
		VehicleID:  vehicle.ID,
		Link:       models.Link(link),
	}

	err := db.Db.Create(&customerVehicle).Error

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})

	} else {
		ctx.JSON(200, gin.H{"customer": customer, "vehicle": vehicle})
	}

}

func ListVehiclesByCustomer(ctx *gin.Context) {
	var customer models.Customer

	id := ctx.Param("id")

	err := db.Db.First(&customer, id).Error
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var vehicles []models.Vehicle

	err = db.Db.Model(&customer).Association("Vehicles").Find(&vehicles)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, gin.H{"vehicles": vehicles})
	}

}

func ListVehiclesByLink(ctx *gin.Context) {
	var customerVehicles []models.CustomerVehicle
	var vehicles []models.Vehicle

	link := ctx.Param("link")

	if !models.IsValidLink(link) {
		ctx.JSON(400, gin.H{"error": "Invalid link value"})
		return
	}

	err := db.Db.Where("link = ?", link).Find(&customerVehicles).Error
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for _, customerVehicle := range customerVehicles {
		var vehicle models.Vehicle
		err = db.Db.First(&vehicle, customerVehicle.VehicleID).Error
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		vehicles = append(vehicles, vehicle)
	}

	ctx.JSON(200, gin.H{"vehicles": vehicles})
}

func ListVehiclesByLinkAndCustomer(ctx *gin.Context) {
	var customerVehicles []models.CustomerVehicle
	var vehicles []models.Vehicle

	link := ctx.Param("link")
	userID := ctx.Param("id")

	if !models.IsValidLink(link) {
		ctx.JSON(400, gin.H{"error": "Invalid link value"})
		return
	}

	err := db.Db.Where("link = ? AND user_id = ?", link, userID).Find(&customerVehicles).Error
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for _, userVehicle := range customerVehicles {
		var vehicle models.Vehicle
		err = db.Db.First(&vehicle, userVehicle.VehicleID).Error
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		vehicles = append(vehicles, vehicle)
	}

	ctx.JSON(200, gin.H{"vehicles": vehicles})
}

func RemoveVehicleFromCustomer(ctx *gin.Context) {

	id := ctx.Param("id")
	vehicleID := ctx.Param("vehicleID")

	var customer models.Customer
	var vehicle models.Vehicle

	db.Db.First(&customer, id)
	db.Db.First(&vehicle, vehicleID)

	err := db.Db.Model(&customer).Association("Vehicles").Delete(&vehicle)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"user": customer})

}

func RemoveAllVehiclesFromCustomer(ctx *gin.Context) {

	var customer models.Customer

	id := ctx.Param("id")

	db.Db.First(&customer, id)

	err := db.Db.Model(&customer).Association("Vehicles").Clear()

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"costumer": customer})

}
