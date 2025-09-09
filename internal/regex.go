package internal

import (
	"log/slog"
	"regexp"
)

// KtpRegex struct represents regex used
// when scanning KTP. It uses regex compile
// so it compiles once only
type KtpRegex struct {
	TempatTanggalLahirRegex *regexp.Regexp
	JenisKelaminRegex       *regexp.Regexp
	GolDarRegex             *regexp.Regexp
	AgamaRegex              *regexp.Regexp
	KawinRegex              *regexp.Regexp
	BerlakuHinggaRegex      *regexp.Regexp
}

func NewKtpRegex() KtpRegex {
	// tempat tanggal lahir regex
	ttlRegex, err := regexp.Compile(
		`[A-Za-z\s\-.']+,\s*\b(0[1-9]|[12][0-9]|3[01])-(0[1-9]|1[0-2])-((19|20)\d{2})\b`,
	)
	if err != nil {
		slog.Error("Error compiling ttl regex",
			"error", err,
		)
		panic(err)
	}

	// jenis kelamin regex
	jenisKelaminRegex, err := regexp.Compile(
		"(LAKI-LAKI|LAKI|LELAKI|PEREMPUAN)",
	)
	if err != nil {
		slog.Error("Error compiling jenis kelamin regex",
			"error", err,
		)
		panic(err)
	}

	// golongan darah regex
	golDarRegex, err := regexp.Compile(
		`^(A|B|AB|O)?$`,
	)
	if err != nil {
		slog.Error("Error compiling golongan darah regex",
			"error", err,
		)
		panic(err)
	}

	// agama regex
	agamaRegex, err := regexp.Compile(
		"(ISLAM|KRISTEN|KATHOLIK|HINDU|BUDDHA|KONGHUCU)",
	)
	if err != nil {
		slog.Error("Error compiling agama regex",
			"error", err,
		)
		panic(err)
	}

	// perkawinan regex
	kawin, err := regexp.Compile(
		"(KAWIN|BELUM KAWIN|CERAI HIDUP|CERAI MATI)",
	)
	if err != nil {
		slog.Error("Error compiling agama regex",
			"error", err,
		)
		panic(err)
	}

	// berlaku hingga regex
	berlakuHingga, err := regexp.Compile(
		`(0[1-9]|[12][0-9]|3[01])-(0[1-9]|1[0-2])-(19|20)\d{2}`,
	)
	if err != nil {
		slog.Error("Error compiling berlaku hingga regex",
			"error", err,
		)
		panic(err)
	}

	return KtpRegex{
		TempatTanggalLahirRegex: ttlRegex,
		JenisKelaminRegex:       jenisKelaminRegex,
		GolDarRegex:             golDarRegex,
		AgamaRegex:              agamaRegex,
		KawinRegex:              kawin,
		BerlakuHinggaRegex:      berlakuHingga,
	}
}
