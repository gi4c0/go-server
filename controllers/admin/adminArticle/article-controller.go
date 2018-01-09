package adminArticle

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"go-server/db/article"
)

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

func GetUnapproved(c *gin.Context) {
  articles, err := article.GetUnapproved()
  if err != nil {
    c.JSON(500, gin.H{"message": err.Error()})
    return
  }

  c.JSON(200, gin.H{"articles": articles})
  return
}


