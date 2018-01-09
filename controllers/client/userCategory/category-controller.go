package userCategory

import (
	"github.com/gin-gonic/gin"
	"go-server/db/category"
)

func GetAll(c *gin.Context) {
    categories, err := category.GetAll()
    if err != nil {
    	c.JSON(500, gin.H{"message": err.Error()})
    	return
	}

	c.JSON(200, gin.H{"categories": categories})
	return
}

