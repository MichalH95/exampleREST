package database

import (
	"gorm.io/gorm"
)

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "pass"
	Dbname   = "examplerest_db"
)

var (
	DBConn *gorm.DB
)
