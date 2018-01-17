package article

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func SaveImage(text *string) (string, error) {
	timeStr := strconv.FormatInt(time.Now().UnixNano(), 10)
	imagePath := "public/images/" + timeStr + "-image"

	var finishedArticle string
	re := regexp.MustCompile(`src\s*=\s*"data:image/(\w{3});base64,(.+?)"`)

	srcAttr := re.FindStringSubmatch(*text)
	if len(srcAttr) == 0 {
		return *text, nil
	}

	fileType := srcAttr[1]
	base64Str := srcAttr[2]
	unbased, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", errors.New("Cannot decode base64 to img")
	}

	r := bytes.NewReader(unbased)
	var img image.Image
	var imgErr error
	// decode regarding image type
	if fileType == "png" {
		imagePath += ".png"
		img, imgErr = png.Decode(r)
		if imgErr != nil {
			return "", imgErr
		}
	} else if fileType == "jpg" {
		imagePath += ".jpg"
		img, imgErr = jpeg.Decode(r)
		if imgErr != nil {
			return "", imgErr
		}
	} else {
		return "", errors.New("Wrong image type. Use .jpg or .png")
	}

	f, fileErr := os.Create(imagePath)
	if fileErr != nil {
		return "", fileErr
	}

	png.Encode(f, img)
	savedImagePath := "src=\"/" + imagePath + "\""
	finishedArticle = strings.Replace(*text, srcAttr[0], savedImagePath, 1)

	return SaveImage(&finishedArticle)
}
