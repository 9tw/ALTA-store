package database

import (
	"errors"
	"project/config"
	"project/models"
)

// GetCarts return all transactionItems in cart based on given User ID
func GetCarts(userID int) (interface{}, error) {
	var transactionItems []models.TransactionItems

	if err := config.DB.Where(
		"users_id = ? AND status = ?", 
		userID, 
		0,
	).Find(&transactionItems).Error; err != nil {
		return nil, err
	}
	return transactionItems, nil
}

// GetItemsByTransactionsID return all transactionItems by given Transactions ID
func GetItemsByTransactionsID(userID int, transactionsID int) (interface{}, error) {
	var transactionItems []models.TransactionItems

	if err := config.DB.Where(
		"users_id = ? AND status = ? AND transactions_id = ?",
		userID,
		1,
		transactionsID,
	).Find(&transactionItems).Error; err != nil {
		return nil, err
	}
	return transactionItems, nil
}

// GetUniqueTransactionID return all unique transaction ID based on given User ID
func GetUniqueTransactionID(userID int) (interface{}, error) {
	var uniqueTransactionsID []models.TransactionItems

	if err := config.DB.Select("transactions_id").Distinct("transactions_id").Where(
		"users_id = ? AND status = ?",
		userID,
		1,
	).Find(&uniqueTransactionsID).Error; err != nil {
		return nil, err
	}
	return uniqueTransactionsID, nil
}

// AddCarts store new transactionItems to cart
func AddCarts(transactionItem *models.TransactionItems) (interface{}, error) {
	// Validate transactionItem
	if transactionItem.Validate() != true {
		return nil, errors.New("Format Users ID, Products ID, Harga dan Jumlah yang diberikan tidak sesuai")
	}

	transactionItem.UpdateJumlah(transactionItem.Jumlah)
	transactionItem.Status = 0
	if err := config.DB.Save(transactionItem).Error; err != nil {
		return nil, err
	}
	return transactionItem, nil
}

// UpdateCarts update transactionItems data by given ID and new value
func UpdateCarts(newTransactionItem *models.TransactionItems, transactionItemID int) (interface{}, error) {
	existingTransactionItem := models.TransactionItems{}
	if getErr := config.DB.First(
		&existingTransactionItem,
		"id = ? AND status = ?", 
		transactionItemID,
		0,
	).Error; getErr != nil {
		return nil, getErr
	} 

	existingTransactionItem.UpdateJumlah(newTransactionItem.Jumlah)
	// Validate transactionItem
	if existingTransactionItem.Validate() != true {
		return nil, errors.New("Format Harga dan Jumlah yang diberikan tidak sesuai")
	}

	if err := config.DB.Save(existingTransactionItem).Error; err != nil {
		return nil, err
	}
	return existingTransactionItem, nil
}

// CheckoutCarts update transactionItems status and transactions ID
func CheckoutCarts(userID int, transactionsID int) (interface{}, error) {
	if err := config.DB.Model(models.TransactionItems{}).Where(
		"users_id = ? AND status = ?", 
		userID, 
		0,
	).Updates(map[string]interface{}{
		"status": 1,
		"transactions_id": transactionsID,
	}).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

// DeleteCarts delete existing transactionItems data
func DeleteCarts(transactionItemID int) (interface{}, error) {
	// Check if given transactionID is available
	existingTransactionItem := models.TransactionItems{}
	if getErr := config.DB.Where(
		"id = ?",
		transactionItemID,
	).First(&existingTransactionItem).Error; getErr != nil {
		return nil, getErr
	}

	if deleteErr := config.DB.Delete(&existingTransactionItem).Error; deleteErr != nil {
		return nil, deleteErr
	}
	return nil, nil
}
