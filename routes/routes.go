package routes

import (
	"project/controllers"
	"project/constants"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	//categories
	e.GET("/categories", controllers.GetCategoriesControllers)
	e.GET("/categories/:id", controllers.GetCategoryControllers)

	//token
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.POST("/categories", controllers.CreateCategoriesController)

	return e
}
