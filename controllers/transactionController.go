package controllers

import (
	"net/http"
	"project/lib/database"
	"project/middlewares"
	"project/models"
	"strconv"

	"github.com/labstack/echo"
)

// GetTransactionsControllers get all transactions
func GetTransactionsControllers(c echo.Context) error {
	userID := middlewares.ExtractTokenUserID(c)
	_, getUserErr := database.GetUser(userID)
	if getUserErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getUserErr.Error())
	}

	// Get Unique Transactions ID from TransactionItems where UserID is equal
	uniqueTransactionIDs, uniqueErr := database.GetUniqueTransactionID(userID)
	if uniqueErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, uniqueErr.Error())
	}

	var uniqueIDs = []int{}
	for _, transaction := range uniqueTransactionIDs.([]models.TransactionItems) {
		uniqueIDs = append(uniqueIDs, transaction.TransactionsID)
	}

	transactions, err := database.GetTransactions(uniqueIDs)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"transactions": transactions,
	})
}

// GetTransactionControllers get specific transaction by given transaction ID
func GetTransactionControllers(c echo.Context) error {
	userID := middlewares.ExtractTokenUserID(c)
	_, getUserErr := database.GetUser(userID)
	if getUserErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getUserErr.Error())
	}

	
	id, strconvErr := strconv.Atoi(c.Param("id"))
	if strconvErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, strconvErr.Error())
	}
	

	transactions, err := database.GetTransaction(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	transactionsModel := transactions.(models.Transactions)
	items, getItemsErr := database.GetItemsByTransactionsID(userID, transactionsModel.ID)
	if getItemsErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getItemsErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"transactions": transactions,
		"items": items,
	})
}

// CreateTransactionControllers create a new transaction
func CreateTransactionControllers(c echo.Context) error {
	var transaction models.Transactions
	userID := middlewares.ExtractTokenUserID(c)

	transactions, err := database.CreateTransaction(&transaction, userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Mengubah Transaction Items pada Cart untuk menginformasikan 
	// bahwa Item sudah masuk ke Payment Process
	transactionsModel := transactions.(*models.Transactions)
	if _, checkoutCartsErr := database.CheckoutCarts(userID, transactionsModel.ID); checkoutCartsErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, checkoutCartsErr.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"transactions": transactions,
	})
}

// UpdateTransactionStatusControllers update existing transaction data
func UpdateTransactionStatusControllers(c echo.Context) error {
	transactionID, strconvErr := strconv.Atoi(c.FormValue("id"))
	if strconvErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, strconvErr.Error())
	}

	existingTransaction, getErr := database.GetTransaction(transactionID)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}

	existingTransactionModel := existingTransaction.(models.Transactions)
	buktiTransfer := c.FormValue("bukti_transfer")
	if buktiTransfer != "" && existingTransactionModel.Status == 1 {
		existingTransactionModel.BuktiTransfer = buktiTransfer
	}
	
	var err error
	if existingTransaction, err = database.UpdateTransactionStatus(&existingTransactionModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"transactions": existingTransaction,
	})
}
