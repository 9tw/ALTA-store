package controllers

import (
	"net/http"
	"project/lib/database"
	"project/middlewares"
	"project/models"
	"strconv"

	"github.com/labstack/echo"
)

// GetCartsControllers get all transaction items from a cart of given User ID
func GetCartsControllers(c echo.Context) error {
	userID := middlewares.ExtractTokenUserID(c)
	_, getUserErr := database.GetUser(userID)
	if getUserErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getUserErr.Error())
	}

	transactionItems, getCartsErr := database.GetCarts(userID)
	if getCartsErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getCartsErr.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"items": transactionItems,
	})
}

// CreateCartsControllers create new transaction items into a cart
func CreateCartsControllers(c echo.Context) error {
	transactionItem := models.TransactionItems{}
	if err := c.Bind(&transactionItem); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// check if user ID valid
	userID := middlewares.ExtractTokenUserID(c)
	_, getUserErr := database.GetUser(userID)
	if getUserErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getUserErr.Error())
	}

	// save to database
	transactionItem.UsersID = userID
	transactionItemResult, err := database.AddCarts(&transactionItem)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"items": transactionItemResult,
	})
}

// UpdateCartsController update existing transaction items on a cart
func UpdateCartsController(c echo.Context) error {
	inputTransactionItem := models.TransactionItems{}
	c.Bind(&inputTransactionItem)

	transactionItemID, strconvErr := strconv.Atoi(c.Param("id"))
	if strconvErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, strconvErr.Error())
	}

	updatedTransactionItem, err := database.UpdateCarts(&inputTransactionItem, transactionItemID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"items": updatedTransactionItem,
	})
}

// DeleteCartsController delete existing transaction items from cart
func DeleteCartsController(c echo.Context) error {
	transactionItemID, strconvErr := strconv.Atoi(c.Param("id"))
	if strconvErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, strconvErr.Error())
	}

	if _, err := database.DeleteCarts(transactionItemID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}
