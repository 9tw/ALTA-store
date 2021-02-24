package models

type Categories struct {
	Id						string `json: "id" form: "id"`
	NamaKategori	string `json: "kategori" form: "kategori"`
	KodeKategori	string `json: "kode" form: "kode"`
}
