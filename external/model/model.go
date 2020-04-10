package model

// Biodata : struct model for biodata
type Biodata struct {
	NIK          int32
	Nama         string `db:"nama"`
	JenisKelamin string `db:"jk"`
	Alamat       string `db:"alamat"`
}
