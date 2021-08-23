package controller

import (
	"api-product/models"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type jwtCustomClaim struct {
	JwtStandart jwt.StandardClaims
}

func (Db *Db) Login(c *gin.Context) {
	var (
		user models.User
	)
	userForm := c.PostForm("username")
	passForm := c.PostForm("password")

	if userForm == "" || passForm == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "Authenticated Failed",
		})
	} else {
		if err := Db.DB.Where("username = ? and password = ?", userForm, passForm).First(&user); err.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"msg": "incorrect username or password",
			})
			return
		}

	}

	sign := jwt.New(jwt.GetSigningMethod("HS256"))

	claims := sign.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	token, err := sign.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":   "login succes",
		"token": token,
	})

}
