package internal

import (
	"log/slog"
	"strings"

	"github.com/otiai10/gosseract/v2"
)

func ExtractText(
	imageBytes []byte,
	tesseractClient *gosseract.Client,
	ktpRegex KtpRegex,
) ktpData {
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
		key := strings.Split(currentLineText, ":")
		if len(key) == 1 {
			continue
		}

		// handle NIK
		if strings.Contains(key[0], "NIK") {
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

			continue
		}

		// handle Nama
		if strings.Contains(key[0], "Nama") {
			nama := strings.Split(currentLineText, ":")[1]
			nama = strings.TrimSpace(nama)
			ktpData.Nama = nama

			continue
		}

		// handle tempat tanggal lahir
		if strings.Contains(key[0], "Tempat") {
			ttl := strings.Split(currentLineText, ":")[1]
			ttl = strings.TrimSpace(ttl)

			wordReplacement := map[string]string{
				".": ",",
			}

			var res string
			for _, char := range ttl {
				if replacement, ok := wordReplacement[string(char)]; ok {
					res += replacement
				} else {
					res += string(char)
				}
			}

			fullText := ktpRegex.TempatTanggalLahirRegex.Find(
				[]byte(res),
			)
			if fullText != nil {
				ktpData.TempatTanggalLahir = string(fullText)
			}

			continue
		}

		// handle jenis kelamin and golongan darah
		if strings.Contains(key[0], "Kelamin") {
			jenisKelamin := ktpRegex.JenisKelaminRegex.Find(
				[]byte(currentLineText),
			)
			if jenisKelamin != nil {
				ktpData.JenisKelamin = string(jenisKelamin)
			}

			golDarah := ktpRegex.GolDarRegex.Find(
				[]byte(currentLineText),
			)
			if golDarah != nil {
				ktpData.GolonganDarah = string(golDarah)
			}

			continue
		}

		// handle alamat
		if strings.Contains(key[0], "Alamat") {
			alamat := strings.Split(currentLineText, ":")
			if len(alamat) == 1 {
				continue
			}

			actualAlamat := alamat[1]
			actualAlamat = strings.TrimSpace(actualAlamat)

			ktpData.Alamat = actualAlamat
		}

		// handle rt/rw
		if strings.Contains(key[0], "RT") ||
			strings.Contains(key[0], "RW") {
			fullText := strings.Split(currentLineText, ":")

			if len(fullText) == 1 {
				continue
			}

			rtAndRw := strings.TrimSpace(fullText[1])

			splittedRtRw := strings.Split(rtAndRw, "/")
			if len(splittedRtRw) == 1 {
				continue
			}

			ktpData.RT = splittedRtRw[0]
			ktpData.RW = splittedRtRw[1]

			continue
		}

		// handle kelurahan atau desa
		if strings.Contains(key[0], `Kel/Desa`) {
			kelDesa := strings.Split(currentLineText, ":")
			if len(kelDesa) == 1 {
				continue
			}

			actualKelDesa := strings.TrimSpace(kelDesa[1])

			ktpData.KelurahanAtauDesa = actualKelDesa

			continue
		}

		// handle kecamatan
		if strings.Contains(key[0], "Kecamatan") {
			kecamatan := strings.Split(currentLineText, ":")[1]
			kecamatan = strings.TrimSpace(kecamatan)

			ktpData.Kecamatan = kecamatan

			continue
		}

		// handle agama
		if strings.Contains(key[0], "Agama") {
			agama := ktpRegex.AgamaRegex.Find(
				[]byte(currentLineText),
			)
			if agama != nil {
				ktpData.Agama = string(agama)
			}

			continue
		}

		// handle status perkawinan
		if strings.Contains(key[0], "Perkawinan") {
			kawin := ktpRegex.KawinRegex.Find(
				[]byte(currentLineText),
			)
			if kawin != nil {
				ktpData.StatusPerkawinan = string(kawin)
			}

			continue
		}

		// handle pekerjaan
		if strings.Contains(key[0], "Pekerjaan") {
			pekerjaan := strings.Split(currentLineText, ":")[1]
			pekerjaan = strings.TrimSpace(pekerjaan)

			ktpData.Pekerjaan = pekerjaan

			continue
		}

		// handle kewarganegaraan
		if strings.Contains(key[0], "Kewarganegaraan") {
			kewarganegaraan := strings.Split(currentLineText, ":")[1]
			kewarganegaraan = strings.TrimSpace(kewarganegaraan)

			ktpData.Kewarganegaraan = kewarganegaraan

			continue
		}

		if strings.Contains(key[0], "Berlaku Hingga") {
			berlakuHingga := ktpRegex.BerlakuHinggaRegex.Find(
				[]byte(currentLineText),
			)
			if berlakuHingga != nil {
				ktpData.BerlakuHingga = string(berlakuHingga)
			}

			continue
		}
	}

	return ktpData
}
