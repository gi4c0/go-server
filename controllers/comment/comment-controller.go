package comment

import (
	"github.com/gin-gonic/gin"
	"go-server/db/comment"
)

func AddComment(c *gin.Context) {
	newComment, success := comment.ParseAndValidateNewComment(c)
	if !success { return }

	newComment.Save()
}
