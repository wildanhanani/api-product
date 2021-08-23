package main

import (
	"api-product/config"
	"api-product/controller"
	"api-product/middleware"
	"api-product/models"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Movie service up andn running",
		})
	})
	dbMySql := config.Connection()
	Dbs := controller.Db{DB: dbMySql}

	models.Migrations(dbMySql)
	r.POST("/user", Dbs.CreateUser)
	r.POST("/login", Dbs.Login)
	r.POST("/product", middleware.AuthJWT, Dbs.CreateProduct)
	r.PUT("/product/update", middleware.AuthJWT, Dbs.UpdateProduct)
	r.DELETE("/product/delete", middleware.AuthJWT, Dbs.DeleteProduct)
	r.GET("/product/find", middleware.AuthJWT, Dbs.GetProduct)

	r.Run(":8080")
}
