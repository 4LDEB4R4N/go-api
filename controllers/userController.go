package controllers

// import (
// 	"api/db"
// 	"api/models"
// 	"github.com/gin-gonic/gin"
// )


// 

// func RemoveAllVehiclesFromUser(ctx *gin.Context) {

// 	var user models.User

// 	id := ctx.Param("id")

// 	db.Db.First(&user, id)

// 	err := db.Db.Model(&user).Association("Vehicles").Clear()

// 	if err != nil {
// 		ctx.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(200, gin.H{"user": user})

// }

// 
