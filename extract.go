package main

import (
	"log/slog"
	"strings"

	"github.com/otiai10/gosseract/v2"
)

func ExtractText(imageBytes []byte, tesseractClient *gosseract.Client) ktpData {
	err := tesseractClient.SetImageFromBytes(imageBytes)
	if err != nil {
		slog.Error("Error setting image",
			"error", err.Error(),
		)
		return ktpData{}
	}

	text, err := tesseractClient.Text()
	if err != nil {
		slog.Error("Error getting text from image",
			"error", err,
		)
		return ktpData{}
	}

	ktpData := NewKtpData()

	// handling data here
	splittedText := strings.SplitSeq(text, "\n")
	for currentLineText := range splittedText {
		// Tempat/Tgi Lahir : GROBOGAN. 02-09-1979
		// Jenis Kelamin : LAKI-LAKI Gol Darah :O
		// Alamat PRM PURI DOMAS D-3. SEMPU
		// RTARW 001024
		// Kel/Desa : WEDOMARTANI
		// Kecamatan : NGEMPLAK
		// Agama . ISLAM
		// Status Perkawinan: KAWIN
		// Pekerjaan : PEDAGANG
		// Kewarganegaraan: WNI
		// Berlaku Hingga  :02-09-2017

		// handle NIK
		if strings.Contains(currentLineText, "NIK") {
			nik := strings.Split(currentLineText, ":")[1]
			nik = strings.TrimSpace(nik)

			// NIK contains only number
			// handle wrong interpretation by Tesseract
			wordReplacement := map[string]string{
				"b": "6",
				"e": "2",
				"L": "1",
			}

			var res string
			for _, char := range nik {
				if replacement, ok := wordReplacement[string(char)]; ok {
					res += replacement
				} else {
					res += string(char)
				}
			}

			ktpData.NIK = res
		}

		// handle Nama
		if strings.Contains(currentLineText, "Nama") {
			nama := strings.Split(currentLineText, ":")[1]
			nama = strings.TrimSpace(nama)
			ktpData.Nama = nama
		}

		// handle tempat tanggal lahir
		if strings.Contains(currentLineText, "Tempat") {
		}
	}

	return ktpData
}
