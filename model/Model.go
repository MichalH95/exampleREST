package model

import (
	"gorm.io/gorm"
	"time"
)

const ClientTypeCompany = "Company"
const ClientTypePerson = "Person"

type Company struct {
	ID               uint `gorm:"primarykey"`
	Name             string
	ICO              string
	ContactFirstName string
	ContactLastName  string
	ClientId         uint
}

type Person struct {
	ID        uint `gorm:"primarykey"`
	FirstName string
	LastName  string
	BirthDate time.Time
	ClientId  uint
}

type Client struct {
	gorm.Model
	ClientType string
	Company    Company `gorm:"constraint:OnDelete:CASCADE;"`
	Person     Person  `gorm:"constraint:OnDelete:CASCADE;"`
}
