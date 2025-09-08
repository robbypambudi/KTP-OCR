package main

import (
	"bytes"
	"image/png"
	"log/slog"

	"github.com/otiai10/gosseract/v2"
	"gocv.io/x/gocv"
)

func ReadImageBytes(imageBytes []byte, tesseractInstance *gosseract.Client) ktpData {
	imageMat, err := gocv.IMDecode(imageBytes, gocv.IMReadColor)
	defer func() {
		err := imageMat.Close()
		if err != nil {
			slog.Error("Could not close image file", "error", err)
		}
	}()
	if err != nil {
		slog.Error(
			"Could not read image bytes",
			"error", err,
		)
		return ktpData{}
	}

	processedImageBytes, err := processImage(imageMat)
	if err != nil {
		slog.Error(
			"Error processing image",
			"error", err.Error(),
		)
		return ktpData{}
	}

	return ExtractText(processedImageBytes, tesseractInstance)
}

func processImage(srcMat gocv.Mat) ([]byte, error) {
	processedImage := gocv.NewMat()

	err := gocv.CvtColor(srcMat, &processedImage, gocv.ColorBGRToGray)
	if err != nil {
		return nil, err
	}

	_ = gocv.Threshold(
		processedImage, &processedImage, 130, 255, gocv.ThresholdTrunc,
	)

	img, _ := processedImage.ToImage()
	buff := new(bytes.Buffer)
	_ = png.Encode(buff, img)

	// return processedImage.ToBytes(), nil
	return buff.Bytes(), nil
}
