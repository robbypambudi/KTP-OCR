package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otiai10/gosseract/v2"
)

type Configuration struct {
	GinConfig       *gin.Engine
	TesseractConfig *gosseract.Client
}

func NewConfiguration() Configuration {
	return Configuration{
		GinConfig:       NewGinConfig(),
		TesseractConfig: NewTesseractConfig(),
	}
}

func NewGinConfig() *gin.Engine {
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// multipart forms memory limit (2 MB)
	// for image uploads
	r.MaxMultipartMemory = 2 << 20

	return r
}

func NewTesseractConfig() *gosseract.Client {
	t := gosseract.NewClient()
	_ = t.SetLanguage("ind") // set to Bahasa Indonesia

	return t
}
