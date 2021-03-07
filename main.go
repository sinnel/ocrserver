package main

import (
	"net/http"
	"ocrserver/controller"
	"ocrserver/db"
	"ocrserver/ocr"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	_ "github.com/otiai10/gosseract/v2"
)

func main() {

	// Initialzie ocr core
	client := ocr.Default()
	defer client.Close()

	// Initialize data base
	db := db.Default()
	if db == nil {
		return
	}
	defer db.Close()

	// Initialize router
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello gin",
		})
	})

	r.POST("/upload/:user", controller.Upload)
	r.Run(":8080")
}
