package article

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"go-server/db/article"
)

func CreateArticle(c *gin.Context) {
	validInputError := article.ValidateNewArticle(c)
	if validInputError != "" {
		c.JSON(400, gin.H{"message": validInputError})
		return
	}


	imagePath, imgErr := article.SaveImage(c)
	if imgErr {
		c.JSON(400, gin.H{"message": imagePath})
		return
	}

	userId := c.MustGet("userId").(int)

	newArticle := article.NewArticle{
		Text: c.PostForm("Text"),
		Title: c.PostForm("Title"),
		UserId: userId,
		Image: imagePath,
	}

	articleCreated, articleError := article.Create(&newArticle)
	if !articleCreated {
		c.JSON(400, gin.H{"message": articleError})
		return
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

func GetSingleArticle(c *gin.Context) {
    id := c.Param("id")
    articleId, convErr := strconv.Atoi(id)
    if convErr != nil {
    	c.JSON(400, gin.H{"message": "Article id should be a function"})
    	return
	}

	dbArticle, err := article.GetSingleArticle(articleId)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"article": dbArticle})
	return
}

func UpdateArticle(c *gin.Context) {
	validationErr := article.ValidateNewArticle(c)
	if validationErr != "" {
		c.JSON(400, gin.H{"message": validationErr})
		return
	}

	userId := c.MustGet("userId").(int)

	articleId, idErr := strconv.Atoi(c.Param("id"))
	if idErr != nil {
		c.JSON(400, gin.H{"message": "Article id must be a number"})
		return
	}

	imagePath := ""
	if img := c.PostForm("ImagePath"); img != "" {
		imagePath = img
	} else {
		resp, imgErr := article.SaveImage(c)
		if imgErr {
			c.JSON(400, gin.H{"message": resp})
			return
		}
		imagePath = resp
	}

	newArticle := article.NewArticle{
		Text: c.PostForm("Text"),
		Title: c.PostForm("Title"),
		UserId: userId,
		ArticleId: articleId,
		Image: imagePath,
	}

	articleUpdated, articleError := article.Update(&newArticle)
	if !articleUpdated {
		c.JSON(400, gin.H{"message": articleError})
		return
	}

	c.Status(200)
}

func DeleteImage(c *gin.Context) {
	articleId, idErr := strconv.Atoi(c.Param("id"))
	if idErr != nil {
		c.JSON(400, gin.H{"message": "Article id must be a number"})
		return
	}

	userId := c.MustGet("userId").(int)

	res, success := article.DeleteImage(articleId, userId)
	if !success {
		c.JSON(400, gin.H{"message": res})
		return
	}
}

func ApproveArticle(c *gin.Context) {
    id := c.Param("id")
    articleId, err := strconv.Atoi(id)
    if err != nil {
    	c.JSON(400, gin.H{"message": "Article id must be a number"})
    	return
	}

	approveErr := article.Approve(articleId)
	if approveErr != nil {
		c.JSON(500, gin.H{"message": approveErr.Error()})
		return
	}

	c.Status(200)
}