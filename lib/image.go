package lib

import (
	"crypto/sha1"
	"encoding/hex"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateImage (file *multipart.FileHeader, ctx *gin.Context) (string, error) {
	outputDir := "image/"
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.Mkdir(outputDir, os.ModePerm)
		if err != nil {
			return err.Error(), nil
		}
	}
	
	hash := sha1.New()
	hashInBytes := hash.Sum([]byte(file.Filename))
	hashString := hex.EncodeToString(hashInBytes)
	outputFileName := hashString + "_" + file.Filename
	path := outputDir + outputFileName
	ctx.SaveUploadedFile(file, path)

	return outputFileName, nil
}