package main

import (
	"log/slog"

	"gocv.io/x/gocv"
)

func ReadFile(filePath string) {
}

func processImage(srcMat gocv.Mat) {
	processedImage := gocv.NewMat()

	err := gocv.CvtColor(srcMat, &processedImage, gocv.ColorBGRToGray)
	if err != nil {
		slog.Error(
			"Error processing image to grayscale: ",
			"error", err.Error(),
		)
		return
	}

	_ = gocv.Threshold(
		processedImage, &processedImage, 127, 255, gocv.ThresholdTrunc,
	)

	if gocv.IMWrite("final_test_golang.png", processedImage) {
		slog.Info("Image processing success")
	} else {
		slog.Error("Could not save processed image")
	}
}
