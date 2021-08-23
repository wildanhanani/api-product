package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string `json:"nama"`
	Image       string `json:"image"`
	Description string `json:"deskripsi"`
}
