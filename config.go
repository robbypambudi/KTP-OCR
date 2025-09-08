package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otiai10/gosseract/v2"
)

func GinConfig() *gin.Engine {
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

	// multipart forms memory limit
	r.MaxMultipartMemory = 2 << 20

	return r
}

func TesseractConfig() *gosseract.Client {
	t := gosseract.NewClient()
	_ = t.SetLanguage("ind")

	return t
}
