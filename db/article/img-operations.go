package article

import (
	"github.com/gin-gonic/gin"
	"fmt"
  "time"
  "strconv"
)

func SaveImage (c *gin.Context) (string, error) {
	imageFile, err := c.FormFile("ImageFile")
	if err != nil {
		fmt.Println(err)
		return "", nil
	}

  timeStr := strconv.FormatInt(time.Now().Unix(), 10)

	imagePath := "public/images/" + imageFile.Filename + "-" + timeStr

	if saveErr := c.SaveUploadedFile(imageFile, imagePath); saveErr != nil {
    fmt.Println(saveErr.Error())
		return "", saveErr
	}

  imagePath = "/" + imagePath // to store in db right adress

	return imagePath, nil
}
