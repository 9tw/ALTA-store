package models

type Products struct {
	Id					string `json: "id" form: "id"`
	NamaProduk	string `json: "produk" form: "produk"`
	Stok				string `json: "stok" form: "stok"`
	Deskripsi		string `json: "deskripsi" form: "deskripsi"`
	Harga				string `json: "harga" form: "harga"`
	IdKategori	string `json: "kode" form: "kode"`
	Gambar			string `json: "gambar form: "gambar"`
}