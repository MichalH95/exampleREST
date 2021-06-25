package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name     string
	ICO      string
	ClientId uint
}
