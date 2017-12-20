package auth

import (
	"github.com/gin-gonic/gin"
	"go-server/db/user"
	"go-server/db"
	"fmt"
)

func Register (c *gin.Context) {
	var newUser user.User
	err := c.ShouldBindJSON(&newUser)

	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": err})
		return
	}
	if newUser.Username == "" || newUser.Password == "" {
		c.JSON(401, gin.H{"message": "Not enough data"})
		return
	}

	newUserErr := user.CreateUser(newUser)
	if newUserErr != nil {
		c.JSON(newUserErr.GetCode(), gin.H{"message": newUserErr.GetError()})
		return
	}

	token := user.GenerateToken(newUser.Username, "user")
	c.JSON(200, gin.H{"username": newUser.Username, "token": token})
}

func Login (c *gin.Context) {
	var existUser user.User
	err := c.ShouldBindJSON(&existUser)

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
	if token == "" {
		c.JSON(400, gin.H{"message": "Token not provided"})
		return
	}
	validToken, username := user.VerifyToken(token)

	if !validToken {
		c.JSON(401, gin.H{"message": "Your token has expired"})
		return
	}

	db.Con.Exec("UPDATE test.Users SET Token = \"\" WHERE Username = ?", username)

	c.Status(200)
}

func CheckUsername(c *gin.Context) {
	type Data struct {
		Username string `json:"username"`
	}
	var data Data

	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	userExist := user.CheckUsername(data.Username)
	if !userExist {
		c.JSON(400, gin.H{"message": "This username is already taken"})
		return
	}

	c.Status(200)
}