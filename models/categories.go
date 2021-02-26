package models

// Categories struct contains Category object
type Categories struct {
	ID   int    `json:"id" form:"id"`
	Nama string `gorm:"size:24" json:"kategori" form:"kategori"`
	Kode string `gorm:"size:10" json:"kode" form:"kode"`
}
