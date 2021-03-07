package controller

import (
	"fmt"
	"net/http"
	"ocrserver/db"
	"ocrserver/ocr"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Upload file function
func Upload(context *gin.Context) {

	file, err := context.FormFile("file")
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}

	// Check file extension
	fileExt := strings.ToLower(path.Ext(file.Filename))
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
		context.String(http.StatusBadRequest, "Unsupported file format")
		return
	}

	// Save file to local path
	filePath := fmt.Sprintf("%s-%d%s", strings.Replace(file.Filename, fileExt, "", 1), time.Now().Unix(), fileExt)
	filePath = fmt.Sprintf("%s/%s", "files", filePath)

	err = context.SaveUploadedFile(file, filePath)
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}

	// ocr inspection
	ocrClient := ocr.Default()
	err = ocrClient.SetImage(filePath)
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}

	text, err := ocrClient.Text()
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
		return
	}

	// Save result to postgresql
	ocrResult := new(db.OcrResult)
	ocrResult.UserName = context.Param("user")
	ocrResult.ImagePath = filePath
	ocrResult.ImageResult = text

	operator := db.OcrResultOperator{}
	err = operator.Insert(ocrResult)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(200, gin.H{
		"code":   200,
		"msg":    "Upload successed!",
		"result": text,
	})
}
