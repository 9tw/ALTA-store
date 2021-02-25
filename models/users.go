package models

// User struct contains User object
type Users struct {
  ID          int     `json:"id" form:"id"`
  Nama        string  `gorm:"size:255" json:"nama" form:"nama"`
  Email       string  `gorm:"size:320" json:"email" form:"email"`
  Password    string  `gorm:"size:40" json:"password" form:"password"`
  Token       string  `gorm:"size:160" json:"token" form:"token"`
  NoHp        string  `gorm:"size:13" json:"no_hp" form:"no_hp"`
  Foto        string  `gorm:"size:255" json:"foto" form:"foto"`
  Alamat      string  `gorm:"size:255" json:"alamat" form:"alamat"`
  Status      string  `gorm:"size:10" json:"status" form:"status"`
  IsVerified  bool    `json:"is_verified" form:"is_verified"`
}

