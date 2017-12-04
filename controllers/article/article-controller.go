package article

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"go-server/db/article"
	"go-server/db/user"
)

func CreateArticle(c *gin.Context) {
	text := c.PostForm("Text")
	title := c.PostForm("Title")

	if text == "" || title == "" {
		c.JSON(400, gin.H{"message": "Not enough data provided"})
		return
	}

	file, err := c.FormFile("Image")
	if err != nil {
		c.String(400, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	imagePath := "./public/images/" + file.Filename

	userId := user.GetUserId(c.GetHeader("authorization"))
	if userId == -1 {
		c.JSON(401, gin.H{"message": "Your token has expired"})
		return
	}

	newArticle := article.NewArticle{
		Text: text,
		Title: title,
		Image: imagePath,
		UserId: userId,
	}

	articleCreated, articleError := article.Create(&newArticle)
	if !articleCreated {
		c.JSON(400, gin.H{"message": articleError})
		return
	}

	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		c.String(400, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	c.Status(200)
}
