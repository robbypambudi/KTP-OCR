package main

import (
	"log"
	"log/slog"

	"ktp-reader-ocr/internal"
)

func main() {
	configuration := internal.NewConfiguration()
	ktpRegex := internal.NewKtpRegex()

	defer func() {
		err := configuration.TesseractConfig.Close()
		if err != nil {
			slog.Error("Error closing tesseract client",
				"error", err.Error(),
			)
		}
	}()

	internal.CreateRoutes(configuration.GinConfig, configuration.TesseractConfig, ktpRegex)

	if err := configuration.GinConfig.Run(":8090"); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
