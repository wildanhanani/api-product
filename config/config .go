package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	var userDB, passDB, hostDB, namaDB, port string

	err := godotenv.Load()

	if err != nil {
		panic("Failed to load env file")
	} else {

		userDB = os.Getenv("DATABASE_USER")
		passDB = os.Getenv("DATABASE_PASS")
		hostDB = os.Getenv("DATABASE_HOST")
		namaDB = os.Getenv("DATABASE_NAME")
		port = os.Getenv("DATABASE_PORT")

	}

	conn := userDB + ":" + passDB + "@tcp(" + hostDB + ":" + port + ")/" + namaDB + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		panic("Gagal masuk ke database")
	} else {
		fmt.Println("Koneksi ke database berhasil")
	}
	return db
}
