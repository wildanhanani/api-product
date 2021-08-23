package controller

import (
	"api-product/models"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (Db *Db) CreateProduct(c *gin.Context) {

	form := &struct {
		Image       *multipart.FileHeader `form:"image" validate:"required"`
		ProductName string                `form:"product_name" validate:"required"`
		Description string                `form:"description" validate:"required"`
	}{}

	err := c.Bind(form)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fileName := form.Image.Filename
	fmt.Println(fileName)
	filePath := fmt.Sprintf("public/file/%s", fileName)

	if err := c.SaveUploadedFile(form.Image, filePath); err != nil {
		fmt.Println(err.Error())
		return
	}
	product := &models.Product{}

	product.Image = fileName
	product.ProductName = form.ProductName
	product.Description = form.Description

	Db.DB.Create(product)
	c.JSON(http.StatusOK, gin.H{
		"msg":    "succes create product",
		"status": 200,
		"data": map[string]interface{}{
			"product_name": product.ProductName,
			"image":        product.Image,
			"description":  product.Description,
		},
	})
}

func (Db *Db) GetProduct(c *gin.Context) {
	var (
		product []models.Product
	)

	Db.DB.Find(&product)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success find Product",
		"data":    product,
	})
}

func (Db *Db) DeleteProduct(c *gin.Context) {

	id := c.Query("id")

	var (
		product models.Product
	)
	err := Db.DB.First(&product, id).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "data not found"})
		return
	}

	if err := Db.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success delete product"})

}

func (Db *Db) UpdateProduct(c *gin.Context) {

	form := &struct {
		ID          int                   `form:"id" validate:"required"`
		Image       *multipart.FileHeader `form:"image" validate:"required"`
		ProductName string                `form:"product_name" validate:"required"`
		Description string                `form:"description" validate:"required"`
	}{}

	err := c.Bind(form)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fileName := form.Image.Filename
	filePath := fmt.Sprintf("public/file/%s", fileName)

	if err := c.SaveUploadedFile(form.Image, filePath); err != nil {
		fmt.Println(err.Error())
		return
	}

	var (
		product models.Product
	)

	err = Db.DB.First(&product, form.ID).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "data not found"})
		return
	}

	product.ProductName = form.ProductName
	product.Image = fileName
	product.Description = form.Description

	err = Db.DB.Save(product).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "update succes",
		"data": map[string]interface{}{
			"product_name": product.ProductName,
			"image":        product.Image,
			"description":  product.Description,
			"UpdateAt":     product.UpdatedAt,
		},
	})

}
