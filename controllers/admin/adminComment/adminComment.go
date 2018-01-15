package adminComment

import (
	"go-server/db/comment"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	articleId, intErr := strconv.Atoi(c.Param("articleId"))
	if intErr != nil {
		c.JSON(400, gin.H{"message": "articleId must be a number"})
		return
	}
	userId := c.MustGet("userId").(int)

	var text string
	jsonErr := c.ShouldBindJSON(&text)
	if jsonErr != nil {
		c.JSON(400, gin.H{"message": "Wrong req body format"})
		return
	}

	newComment := comment.NewComment{
		ArticleId: articleId,
		UserId:    userId,
		Text:      text,
	}

	err := newComment.AddAdminComment()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.Status(200)
}

func GetComments(c *gin.Context) {
	articleId, intErr := strconv.Atoi(c.Param("articleId"))
	if intErr != nil {
		c.JSON(400, gin.H{"message": "articleId must be a number"})
		return
	}

	comments, err := comment.GetCommentsByArticleId(articleId, "admin")
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"comments": comments})
}
