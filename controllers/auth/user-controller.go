package auth

import (
	"github.com/gin-gonic/gin"
	"test/db/user"
)

func Register (c *gin.Context) {
	var newUSer user.User
	err := c.BindJSON(&newUSer)

	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	if newUSer.Username == "" || newUSer.Password == "" {
		c.JSON(401, gin.H{"message": "Not enough data"})
		return
	}

	newUserErr := user.CreateUser(newUSer)
	if newUserErr != nil {
		c.JSON(newUserErr.GetCode(), gin.H{"message": newUserErr.GetError()})
		return
	}

	c.JSON(200, gin.H{"username": newUSer.Username})
}

func Login (c *gin.Context) {
	var existUser user.User
	err := c.BindJSON(&existUser)

	if err != nil {
		c.JSON(400, gin.H{"error": "Bad format"})
		return
	}
	if existUser.Username == "" || existUser.Password == "" {
		c.JSON(401, gin.H{"message": "Not enough data"})
		return
	}

	existUserErr := user.VerifyUser(existUser)
	if existUserErr != nil {
		c.JSON(existUserErr.GetCode(), gin.H{"message": existUserErr.GetError()})
		return
	}

	token := user.GenerateToken(existUser.Username)

	c.JSON(200, gin.H{"token": token})
}

func Auth (c *gin.Context) {
	token := c.GetHeader("authorization")

	validToken := user.VerifyToken(token)

	if !validToken {
		c.JSON(401, gin.H{"message": "Your token has expired"})
		return
	}

	c.Status(200)
}
