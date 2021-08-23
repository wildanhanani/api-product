package controller

import "gorm.io/gorm"

type Db struct {
	DB *gorm.DB
}
