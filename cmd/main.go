package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"ktp-reader-ocr/internal"
)

// function to test ktp reader using cmdline
func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Printf("usage: %s <nama_file>", args[0])
		return
	}

	filePath := args[1]

	tesseractClient := internal.NewTesseractConfig()
	ktpRegex := internal.NewKtpRegex()

	imageData, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error(
			"Error reading file",
			"error", err.Error(),
		)
		return
	}

	ktpData := internal.ReadImageBytes(imageData, tesseractClient, ktpRegex)

	b, err := json.MarshalIndent(ktpData, "", "  ")
	if err != nil {
		slog.Error(
			"Error converting to JSON",
			"error", err.Error(),
		)
		return
	}

	fmt.Println(string(b))
}
