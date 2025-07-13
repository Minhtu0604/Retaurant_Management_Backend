package controller

import (
	"github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Logic to get all users
		c.JSON(200, gin.H{"message": "Get all users"})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func HashPassword(password string) (string, error) {

}

func VerifyPassword(userPassword string, providePassword string) (bool, string) {

}
