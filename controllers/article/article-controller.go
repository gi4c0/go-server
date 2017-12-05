package article

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"
	"go-server/db/article"
	"go-server/db/user"
)

func CreateArticle(c *gin.Context) {
	validInput := article.ValidateNewArticle(c)
	if !validInput { return }

	text := c.PostForm("Text")
	title := c.PostForm("Title")

	// save image path if provided
	imageFile, err := c.FormFile("Image")
	var imagePath = ""
	if err == nil {
		imagePath = "public/images/" + imageFile.Filename
	}

	userId := user.GetUserId(c.GetHeader("authorization"))
	if userId == -1 {
		c.JSON(401, gin.H{"message": "Your token has expired"})
		return
	}

	newArticle := article.NewArticle{
		Text: text,
		Title: title,
		UserId: userId,
		Image: imagePath,
	}

	articleCreated, articleError := article.Create(&newArticle)
	if !articleCreated {
		c.JSON(400, gin.H{"message": articleError})
		return
	}

	// save image if provided
	if imagePath != "" {
		if err := c.SaveUploadedFile(imageFile, imagePath); err != nil {
			c.String(400, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}
	}

	c.Status(200)
}

func GetArticles(c *gin.Context) {
	page, pageErr := strconv.Atoi(c.Param("page"))
	count, countErr := strconv.Atoi(c.Param("count"))
	if pageErr != nil || countErr != nil {
		c.JSON(400, gin.H{"message": "page and count must be numbers"})
		return
	}

	skip := (page - 1) * count

	articles, err := article.GetAll(skip, count)
	if err != "" {
		c.JSON(500, gin.H{"message": err})
		return
	}

	c.JSON(200, gin.H{"articles": articles})


}
