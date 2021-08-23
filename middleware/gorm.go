package middleware

import "gorm.io/gorm"

type Db struct {
	DB *gorm.DB
}
