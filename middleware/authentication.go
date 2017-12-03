package middleware

import (
	"github.com/gin-gonic/gin"
	"go-server/db/user"
)

func RequireAuth() gin.HandlerFunc{
	return func(c *gin.Context) {
		token := c.GetHeader("authorization")

		validToken := user.VerifyToken(token)

		if !validToken {
			c.JSON(401, gin.H{"message": "Your token has expired"})
			return
		}

		c.Next()
	}
}
