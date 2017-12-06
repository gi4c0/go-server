package comment

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strconv"
)

const commentTextLength = 10

func ParseAndValidateNewComment (c *gin.Context) (NewComment, bool) {
	var comment NewComment

	err := c.BindJSON(&comment)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "Wrong data format"})
		return comment, false
	}

	if comment.Text == "" || len(comment.Text) < commentTextLength {
		c.JSON(400, gin.H{"message": "Comment Text should be at least " + strconv.Itoa(commentTextLength) + " chars"})
		return comment, false
	}

	comment.UserId = c.MustGet("userId").(int)

	articleId, idErr := strconv.Atoi(c.Param("articleId"))
	if idErr != nil {
		c.JSON(400, gin.H{"message": "Article id should be a number"})
		return comment, false
	}

	comment.ArticleId = articleId

	return comment, true
}
