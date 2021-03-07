package ocr

import (
	"github.com/otiai10/gosseract/v2"
)

var ocrClient *gosseract.Client

// Default db
func Default() *gosseract.Client {
	if ocrClient == nil {
		ocrClient = gosseract.NewClient()
	}

	return ocrClient
}
