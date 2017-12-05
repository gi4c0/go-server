package article

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const textLength = 10
const titleLength = 10

func ValidateNewArticle (c *gin.Context) bool {
	text := c.PostForm("Text")
	title := c.PostForm("Title")

	if text == "" || title == "" {
		c.JSON(400, gin.H{"message": "Not enough data provided"})
		return false
	}

	if len(text) < textLength {
		c.JSON(400, gin.H{"message": "Too short 'Text' field. Should be at least " + strconv.Itoa(textLength) + " characters"})
		return false
	}

	if len(title) < titleLength {
		c.JSON(400, gin.H{"message": "Too short 'Title' field. Should be at least " + strconv.Itoa(titleLength) + " characters"})
		return false
	}

	return true
}