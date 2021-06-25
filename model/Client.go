package model

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	FirstName string
	Surname   string
	Company   Company
}
