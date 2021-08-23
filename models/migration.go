package models

import (
	"fmt"

	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) {

	if check := db.Migrator().HasTable(&User{}); !check {
		db.Migrator().CreateTable(&User{})
		fmt.Println("Table User berhasil di create")
	}
	if check := db.Migrator().HasTable(&Product{}); !check {
		db.Migrator().CreateTable(&Product{})
		fmt.Println("Table Product berhasil di create")
	}
}
