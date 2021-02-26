package controllers

import (
	"net/http"
	"project/lib/database"
	"project/models"
	"github.com/labstack/echo"
)

func GetCategoriesControllers(c echo.Context) error {
	cats, e := database.GetCategories()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"categories":  cats,
	})
}

func GetCategoryControllers(c echo.Context) error {
	id := c.Param("id")
	cats, e := database.GetCategory(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"categories":  cats,
	})
}

func CreateCategoriesController(c echo.Context) error {
	cat := models.Categories{}
	c.Bind(&cat)
	cats, e := database.CreateCategories(&cat)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"categories":  cats,
	})
}
