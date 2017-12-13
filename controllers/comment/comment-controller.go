package comment

import (
	"github.com/gin-gonic/gin"
	"go-server/db/comment"
	"strconv"
)

func AddComment(c *gin.Context) {
	newComment, err := comment.ParseAndValidateNewComment(c)
	if err { return }

	newComment.Save()
}

func DeleteComment(c *gin.Context) {
	userId := c.MustGet("userId").(int)

	commentId, parseErr := strconv.Atoi(c.Param("commentId"))
	if parseErr != nil {
		c.JSON(400, gin.H{"message": "Comment id must be a number"})
		return
	}

	deleteCommentErr := comment.DeleteComment(commentId, userId)
	if deleteCommentErr != "" {
		c.JSON(400, gin.H{"message": deleteCommentErr})
		return
	}

	c.Status(200)
}

func UpdateComment(c *gin.Context) {
    var newComment comment.NewComment

    jsonErr := c.BindJSON(&newComment)
    if jsonErr != nil {
    	c.JSON(400, gin.H{"message": "Wrong format"})
    	return
	}

	commentId, convertErr := strconv.Atoi(c.Param("commentId"))
	if convertErr != nil {
		c.JSON(400, gin.H{"message": "Comment id should be a number"})
		return
	}

	newComment.CommentId = commentId
	newComment.UserId = c.MustGet("userId").(int)

	updateError := newComment.Update()
	if updateError != "" {
		c.JSON(400, gin.H{"message": updateError})
		return
	}
}