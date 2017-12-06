package middleware

import (
	"github.com/gin-gonic/gin"
	"go-server/db/user"
	"fmt"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("authorization")
		fmt.Println(token)
		if token == "" {
			c.JSON(401, gin.H{"message": "Please login first"})
			c.Abort()
			return
		}

		validToken, _ := user.VerifyToken(token)

		if !validToken {
			c.JSON(401, gin.H{"message": "Your token has expired"})
			c.Abort()
			return
		}

		userId := user.GetUserId(token)
		c.Set("userId", userId)
	}
}
