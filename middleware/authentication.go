package middleware

import (
	"github.com/gin-gonic/gin"
	"go-server/db/user"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("authorization")
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

		userId, permission := user.GetUserId(token)
		c.Set("userId", userId)
    c.Set("permission", permission)
	}
}

func RequireModerator() gin.HandlerFunc {
	return func (c *gin.Context) {
		token := c.GetHeader("authorization")
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

		userId, permission := user.VerifyPermission(token, "moderator")
		if userId < 1 {
			c.JSON(403, gin.H{"message": "You don't have permissions enough"})
			c.Abort()
			return
		}

		c.Set("userId", userId)
    c.Set("permission", permission)
	}
}

func RequireAdmin() gin.HandlerFunc {
	return func (c *gin.Context) {
		token := c.GetHeader("authorization")
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

		userId, _ := user.VerifyPermission(token, "admin")
		if userId < 1 {
			c.JSON(403, gin.H{"message": "You don't have permissions enough"})
			c.Abort()
			return
		}

		c.Set("userId", userId)
    c.Set("permission", "admin")
	}
}
