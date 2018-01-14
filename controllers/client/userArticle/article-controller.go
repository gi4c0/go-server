package userArticle

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"go-server/db/article"
  "fmt"
)

type reqArticle struct {
  Title string
  Text string
}

func CreateArticle(c *gin.Context) {
  userId := c.MustGet("userId").(int)
  permission := c.MustGet("permission").(string)

  var reqData reqArticle
  jsonErr := c.ShouldBindJSON(&reqData)
  if jsonErr != nil {
    fmt.Println(jsonErr)
    c.JSON(400, gin.H{"message": jsonErr.Error()})
    return
  }

  newArticle := article.NewArticle{
    Text: reqData.Text,
    Title: reqData.Title,
    UserId: userId,
  }

  articleCreated, articleError := article.Create(&newArticle, permission)
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
  if err != nil {
    c.JSON(500, gin.H{"message": err.Error()})
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
  if validationErr != nil {
    c.JSON(400, gin.H{"message": validationErr.Error()})
    return
  }

  userId := c.MustGet("userId").(int)

  articleId, idErr := strconv.Atoi(c.Param("id"))
  if idErr != nil {
    c.JSON(400, gin.H{"message": "Article id must be a number"})
    return
  }

  var reqData reqArticle
	jsonErr := c.ShouldBindJSON(&reqData)
	if jsonErr != nil {
		c.JSON(400, gin.H{"message": jsonErr.Error()})
		return
	}

  newArticle := article.NewArticle{
    Text: reqData.Text,
    Title: reqData.Title,
    UserId: userId,
    ArticleId: articleId,
  }

  articleUpdated, articleError := article.Update(&newArticle)
  if !articleUpdated {
    c.JSON(400, gin.H{"message": articleError})
    return
  }

  c.Status(200)
}

// func DeleteImage(c *gin.Context) {
//   articleId, idErr := strconv.Atoi(c.Param("id"))
//   if idErr != nil {
//     c.JSON(400, gin.H{"message": "Article id must be a number"})
//     return
//   }
//
//   userId := c.MustGet("userId").(int)
//
//   res, success := article.DeleteImage(articleId, userId)
//   if !success {
//     c.JSON(400, gin.H{"message": res})
//     return
//   }
// }

func GetUserArticlesPreview (c *gin.Context) {
  userId := c.MustGet("userId").(int)
  articles, err := article.UserArticles(userId)
  if err != nil {
    c.JSON(500, gin.H{"message": err.Error()})
    return
  }

  c.JSON(200, gin.H{"articles": articles})
  return
}
