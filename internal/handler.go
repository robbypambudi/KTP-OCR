package internal

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/otiai10/gosseract/v2"
)

// handler dependencies
type Handler struct {
	tesseractClient *gosseract.Client
	ktpRegex        KtpRegex
}

func NewHandler(
	tesseractClient *gosseract.Client,
	ktpRegex KtpRegex,
) *Handler {
	return &Handler{
		tesseractClient,
		ktpRegex,
	}
}

func (h *Handler) KtpReader(c *gin.Context) {
	file, err := c.FormFile("ktp")
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			BuildResponseFailed("Gagal membaca KTP", err.Error(), nil),
		)
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			BuildResponseFailed("Gagal membaca KTP", err.Error(), nil),
		)
		return
	}
	defer func(c *gin.Context) {
		err := openedFile.Close()
		if err != nil {
			slog.Error(
				"Error closing file",
				"error", err.Error(),
			)
		}
	}(c)

	fileBytes, err := io.ReadAll(openedFile)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			BuildResponseFailed("Gagal membaca KTP", err.Error(), nil),
		)
		return
	}

	res := ReadImageBytes(fileBytes, h.tesseractClient, h.ktpRegex)
	c.JSON(http.StatusOK, BuildResponseSuccess("Berhasil membaca KTP", res))
}
