package models

import "time"

// Transactions struct contains Transaction object
type Transactions struct {
	ID													int				`json:"id" form:"id"`
	KodeInvoice									string		`gorm:"size:25" json:"kode_invoice" form:"kode_invoice"`
	TanggalInvoice							time.Time	`json:"tanggal_invoice" form:"tanggal_invoice"`
	TanggalTenggatPembayaran		time.Time	`json:"tanggal_tenggat_pembayaran" form:"tanggal_tenggat_pembayaran"`
	TanggalSelesaiPembayaran		time.Time	`json:"tanggal_selesai_pembayaran" form:"tanggal_selesai_pembayaran"`
	TanggalKonfirmasiPembayaran	time.Time	`json:"tanggal_konfirmasi_pembayaran" form:"tanggal_konfirmasi_pembayaran"`
	TanggalTerimaPesanan				time.Time	`json:"tanggal_terima_pesanan" form:"tanggal_terima_pesanan"`
	TanggalSelesaiTransaksi			time.Time	`json:"tanggal_selesai_transaksi" form:"tanggal_selesai_transaksi"`
	BuktiTransfer								string		`gorm:"size:255" json:"bukti_transfer" form:"bukti_transfer"`
	Status											int				`json:"status" form:"status"`
}
