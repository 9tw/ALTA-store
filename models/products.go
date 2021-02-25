package models

// Products struct contains Product object
type Products struct {
	ID						int 	 `json:"id" form:"id"`
	Nama					string `gorm:"size:90" gorm:"size:24" json:"produk" form:"produk"`
	Stok					int 	 `json:"stok" form:"stok"`
	Deskripsi			string `json:"deskripsi" form:"deskripsi"`
	Harga					int 	 `json:"harga" form:"harga"`
	CategoriesID	int 	 `json:"kode" form:"kode"`
	Gambar				string `gorm:"size:255" json:"gambar" form:"gambar"`
}

// RequiredNotNull returns true if Products data fulfil the not null requirements
func (p Products) RequiredNotNull() bool {
	if p.Nama == "" ||
		 p.Stok < 0 ||
		 p.Harga < 0 ||
		 p.CategoriesID < 0 {
			 return false
		 }
	return true
}