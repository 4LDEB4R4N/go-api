package models

import (
	"gorm.io/gorm"
)

type Link string

const (
	AUTONOMO Link = "AUTONOMO"
	AGREGADO Link = "AGREGADO"
	FROTA    Link = "FROTA"
)

func IsValidLink(link string) bool {
	switch Link(link) {
	case AUTONOMO, AGREGADO, FROTA:
		return true
	}
	return false
}

type Customer struct {
	gorm.Model
	Name     string
	CNPJ     string
	Vehicles []Vehicle `gorm:"many2many:customer_vehicles;"`
}

type Vehicle struct {
	gorm.Model
	LicensePlate string
	Color        string
	Customers        []Customer `gorm:"many2many:customer_vehicles;"`
}

type CustomerVehicle struct {
	gorm.Model
	CustomerID    uint
	VehicleID uint
	Link      Link
}

func (CustomerVehicle) BeforeCreate(db *gorm.DB) (err error) {
	return db.SetupJoinTable(&Customer{}, "Vehicles", &CustomerVehicle{})
}
