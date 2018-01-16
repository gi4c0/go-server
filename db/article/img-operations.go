package article

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image/png"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// TODO: make with go routines
func SaveImage(text *string) (string, error) {
	timeStr := strconv.FormatInt(time.Now().UnixNano(), 10)
	imagePath := "public/images/" + timeStr + "-image.png"
	fmt.Println(imagePath)

	var finishedArticle string
	re := regexp.MustCompile(`src\s*=\s*"data:image/(\w{3});base64,(.+?)"`)

	image := re.FindStringSubmatch(*text)
	if len(image) == 0 {
		return *text, nil
	}

	base64Str := image[2]
	unbased, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", errors.New("Cannot decode base64 to img")
	}

	r := bytes.NewReader(unbased)
	img, imgErr := png.Decode(r)
	if imgErr != nil {
		return "", imgErr
	}

	f, fileErr := os.Create(imagePath)
	if fileErr != nil {
		return "", fileErr
	}

	png.Encode(f, img)
	savedImagePath := "src=\"/" + imagePath + "\""
	finishedArticle = strings.Replace(*text, image[0], savedImagePath, 1)

	return SaveImage(&finishedArticle)
}
