package article

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const textLength = 10
const titleLength = 10

func ValidateNewArticle (c *gin.Context) string {
	text := c.PostForm("Text")
	title := c.PostForm("Title")

	if text == "" || title == "" {
		return "Not enough data provided"
	}

	if len(text) < textLength {
		return "Too short 'Text' field. Should be at least " + strconv.Itoa(textLength) + " characters"
	}

	if len(title) < titleLength {
		return "Too short 'Title' field. Should be at least " + strconv.Itoa(titleLength) + " characters"
	}

	return ""
}