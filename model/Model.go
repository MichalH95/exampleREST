package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

const ClientTypeCompany int = 1
const ClientTypePerson int = 2

type Company struct {
	gorm.Model
	Name             string
	ICO              string
	ContactFirstName string
	ContactLastName  string
	Client           Client
}

type Person struct {
	gorm.Model
	FirstName string
	LastName  string
	BirthDate time.Time
	Client    Client
}

type Client struct {
	ID         uint `gorm:"primarykey"`
	CompanyId  sql.NullInt64
	PersonId   sql.NullInt64
	ClientType int
}
