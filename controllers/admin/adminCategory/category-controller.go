package adminCategory

import (
	"github.com/gin-gonic/gin"
	"go-server/db/category"
)

func ChangeCategoryName (c *gin.Context) {
  newCategoryName := c.Param("new")
  oldCategoryName := c.Param("old")
  if newCategoryName == oldCategoryName {
    c.Status(200)
    return
  }

  err := category.Change(oldCategoryName, newCategoryName)
  if err != nil {
    c.JSON(500, gin.H{"message": err.Error()})
    return
  }

  c.Status(200)
}
