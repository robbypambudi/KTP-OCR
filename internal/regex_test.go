package internal

import (
	"testing"
)

func TestTempatTanggalLahirRegex(t *testing.T) {
	ktpRegex := NewKtpRegex()

	// test "<place>, DD-MM-YYYY format"
	firstFormat := "Jogjakarta, 01-02-2003"
	wantFirstFormat := ktpRegex.TempatTanggalLahirRegex.Find([]byte(firstFormat))
	if string(wantFirstFormat) != firstFormat {
		t.Errorf(`TempatTanggalLahirRegex("01-02-2003") == %s, wants: %s`, string(wantFirstFormat), firstFormat)
	}

	// test empty string
	emptyString := ""
	wantEmpty := ktpRegex.TempatTanggalLahirRegex.Find([]byte(emptyString))
	if string(wantEmpty) != emptyString {
		t.Errorf(`TempatTanggalLahirRegex("") == %s, wants: %s`, string(wantEmpty), emptyString)
	}
}
