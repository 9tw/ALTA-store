package database

import (
	"fmt"
	"project/config"
	"project/models"
	"time"
)

// GetTransactions return all transactions and query by user if any
func GetTransactions(transactionIDs []int) (interface{}, error) {
	var transactions []models.Transactions

	if err := config.DB.Find(&transactions, transactionIDs).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetTransaction return transaction by given transaction id or kode invoice if any
func GetTransaction(id int) (interface{}, error) {
	var transactions models.Transactions

	if err := config.DB.First(&transactions, []int{id}).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// CreateTransaction to store new transaction data
func CreateTransaction(transaction *models.Transactions, userID int) (interface{}, error) {
	
	transaction.TanggalInvoice = time.Now()
	transaction.KodeInvoice = fmt.Sprintf("%d%d", userID, transaction.TanggalInvoice.Unix())
	transaction.TanggalTenggatPembayaran = transaction.TanggalInvoice.Add(time.Hour * 24 * 3)
	transaction.Status = 1

	if err := config.DB.Save(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

// UpdateTransactionStatus to update transaction status and insert date on related attribute
func UpdateTransactionStatus(transaction *models.Transactions) (interface{}, error) {
	switch transaction.Status {
	case 1:
		transaction.TanggalSelesaiPembayaran = time.Now() 
		transaction.Status = 2
	case 2:
		transaction.TanggalKonfirmasiPembayaran = time.Now()
		transaction.Status = 3
	case 3:
		transaction.TanggalTerimaPesanan = time.Now()
		transaction.Status = 4
	case 4:
		transaction.TanggalSelesaiTransaksi = time.Now()
		transaction.Status = 5
	}

	if err := config.DB.Save(transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}
