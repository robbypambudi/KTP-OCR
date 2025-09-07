package main

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/otiai10/gosseract/v2"
)

func ExtractText(imageBytes []byte) {
	// ktpData := NewKtpData()

	tesseractClient := gosseract.NewClient()
	defer func() {
		err := tesseractClient.Close()
		if err != nil {
			slog.Error("Error closing tesseract client",
				"error", err.Error(),
			)
			return
		}
	}()

	// use Bahasa Indonesia for KTP
	_ = tesseractClient.SetLanguage("ind")

	err := tesseractClient.SetImageFromBytes(imageBytes)
	if err != nil {
		slog.Error("Error setting image",
			"error", err.Error(),
		)
	}

	text, err := tesseractClient.Text()
	if err != nil {
		slog.Error("Error getting text from image",
			"error", err.Error(),
		)
	}

	splittedText := strings.SplitSeq(text, "\n")
	for currentLineText := range splittedText {
		fmt.Println(currentLineText)
	}
}
