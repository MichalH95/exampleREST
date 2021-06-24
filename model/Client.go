package model

import "github.com/jinzhu/gorm"

type Client struct {
	gorm.Model
	FirstName string
	Surname   string
}
