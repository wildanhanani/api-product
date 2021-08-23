package controller

import (
	"api-product/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (Db *Db) CreateUser(c *gin.Context) {

	form := &struct {
		Username string `form:"username" validate:"required"`
		Password string `form:"password" validate:"required"`
	}{}

	err := c.Bind(form)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user := &models.User{}

	user.Username = form.Username
	user.Password = form.Password

	Db.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"msg":    "succes create user",
		"status": 200,
		"data": map[string]interface{}{
			"username": user.Username,
			"password": user.Password,
		},
	})
}
