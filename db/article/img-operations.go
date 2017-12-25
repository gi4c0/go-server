package article

import (
	"github.com/gin-gonic/gin"
	"fmt"
  "time"
  "strconv"
)

func SaveImage (c *gin.Context) (string, bool) {
	imageFile, err := c.FormFile("ImageFile")
	if err != nil {
		fmt.Println(err)
		return "", false
	}

  timeStr := strconv.FormatInt(time.Now().Unix(), 10)

	imagePath := "public/images/" + imageFile.Filename + "-" + timeStr

	if saveErr := c.SaveUploadedFile(imageFile, imagePath); saveErr != nil {
		//c.String(400, fmt.Sprintf("get form err: %s", err.Error()))
		return saveErr.Error(), true
	}

	return imagePath, false
}
