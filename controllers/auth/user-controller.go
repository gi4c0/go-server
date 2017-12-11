package auth

import (
	"github.com/gin-gonic/gin"
	"go-server/db/user"
	"go-server/db"
	//"fmt"
	"fmt"
)

func Register (c *gin.Context) {
	var newUSer user.User
	err := c.ShouldBindJSON(&newUSer)

	if err != nil {
		fmt.Println(err)
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

	existUserErr := user.VerifyUser(&existUser)
	if existUserErr != nil {
		c.JSON(existUserErr.GetCode(), gin.H{"message": existUserErr.GetError()})
		return
	}

	token := user.GenerateToken(existUser.Username, existUser.Permission)
	db.Con.Exec("UPDATE test.Users SET Token = ? WHERE Username = ?", token, existUser.Username)

	c.JSON(200, gin.H{"token": token})
}

func Auth (c *gin.Context) {
	token := c.GetHeader("authorization")
	var dbToken string

	validToken, username := user.VerifyToken(token)
	db.Con.QueryRow("SELECT Token from test.Users WHERE Username = ?", username).Scan(&dbToken)

	if !validToken || token != dbToken {
		c.JSON(401, gin.H{"message": "Your token has expired"})
		return
	}

	c.Status(200)
}

func Logout(c *gin.Context) {
	token := c.GetHeader("authorization")
	validToken, username := user.VerifyToken(token)

	if !validToken {
		c.JSON(401, gin.H{"message": "Your token has expired"})
		return
	}

	db.Con.Exec("UPDATE test.Users SET Token = \"\" WHERE Username = ?", username)

	c.Status(200)
}