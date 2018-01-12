package article

import (
	"github.com/gin-gonic/gin"
  "fmt"
	"strconv"
  "errors"
)

const textLength = 10
const titleLength = 10

type article struct {
  Text string
  Title string
}

func ValidateNewArticle (c *gin.Context) error {
  var a article
  c.ShouldBindJSON(&a)

  fmt.Println(a.Text)
  
	if a.Text == "" || a.Title == "" {
		return errors.New("Not enough data provided")
	}

	if len(a.Text) < textLength {
		return errors.New("Too short 'Text' field. Should be at least " + strconv.Itoa(textLength) + " characters")
  }

	if len(a.Title) < titleLength {
		return errors.New("Too short 'Title' field. Should be at least " + strconv.Itoa(titleLength) + " characters")
	}

	return nil
}
