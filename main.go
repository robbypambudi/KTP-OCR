package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func main() {
	ginInstance := GinConfig()
	ktpRegex := NewKtpRegex()
	tesseractInstance := TesseractConfig()
	defer func() {
		err := tesseractInstance.Close()
		if err != nil {
			slog.Error("Error closing tesseract client",
				"error", err.Error(),
			)
		}
	}()

	ginInstance.POST("/api/reader", func(ctx *gin.Context) {
		file, err := ctx.FormFile("ktp")
		if err != nil {
			fmt.Println(err)
		}

		openedFile, err := file.Open()
		if err != nil {
			fmt.Println(err)
		}
		defer func(ctx *gin.Context) {
			err := openedFile.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(ctx)

		fileBytes, err := io.ReadAll(openedFile)
		if err != nil {
			fmt.Println(err)
		}

		res := ReadImageBytes(fileBytes, tesseractInstance, ktpRegex)
		ctx.JSON(200, res)
	})

	if err := ginInstance.Run("localhost:8090"); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
