package article

import (
	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	file, _ := c.FormFile("Image")
	c.SaveUploadedFile(file, "go-server/public/images/")

	c.Status(200)
}
