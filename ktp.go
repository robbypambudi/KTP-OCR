package main

type ktpData struct {
	NIK                string `json:"nik"`
	Nama               string `json:"nama"`
	TempatTanggalLahir string `json:"tempat_tanggal_lahir"`
	JenisKelamin       string `json:"jenis_kelamin"`
	Alamat             string `json:"alamat"`
	RT                 string `json:"rt"`
	RW                 string `json:"rw"`
	KelurahanAtauDesa  string `json:"kelurahan_atau_desa"`
	Kecamatan          string `json:"kecamatan"`
	Agama              string `json:"agama"`
	StatusPerkawinan   string `json:"status_perkawinan"`
	Pekerjaan          string `json:"pekerjaan"`
	Kewarganegaraan    string `json:"kewarganegaraan"`
	BerlakuHingga      string `json:"berlaku_hingga"`
}

func NewKtpData() ktpData {
	return ktpData{
		BerlakuHingga: "SEUMUR HIDUP", // UU Nomor 24 Tahun 2013
	}
}
