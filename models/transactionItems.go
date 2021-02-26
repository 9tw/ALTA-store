package models

// TransactionItems struct contains Transaction Item object
type TransactionItems struct {
	ID							int			`json:"id" form:"id"`
	UsersID					int			`json:"users_id" form:"users_id"`
	ProductsID			int			`json:"products_id" form:"products_id"`
	Gambar					string	`gorm:"size:255" json:"gambar" form:"gambar"`
	Status					int			`json:"status" form:"status"`
	Harga						int			`json:"harga" form:"harga"`
	Jumlah					int			`json:"jumlah" form:"jumlah"`
	HargaTotal			int			`json:"harga_total" form:"harga_total"`
	TransactionsID	int			`json:"transactions_id" form:"transactions_id"`
}

// Validate to validate if TransactionItem contains correct data
func (transactionItem *TransactionItems) Validate() bool {
	if transactionItem.Harga <= 0 ||
		transactionItem.Jumlah <= 0 ||
		transactionItem.UsersID == 0 ||
		transactionItem.ProductsID == 0 {
			return false
		} 
	return true
}

// UpdateJumlah to update Harga Total attribute by multiplying Harga and Jumlah
func (transactionItem *TransactionItems) UpdateJumlah(jumlah int) {
	transactionItem.Jumlah = jumlah
	transactionItem.HargaTotal = transactionItem.Harga * transactionItem.Jumlah
}
