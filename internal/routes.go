package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/otiai10/gosseract/v2"
)

// put all the routes here
func CreateRoutes(server *gin.Engine, tesseract *gosseract.Client, ktpRegex KtpRegex) {
	handler := NewHandler(tesseract, ktpRegex)

	server.POST("/api/reader", handler.KtpReader)
}
