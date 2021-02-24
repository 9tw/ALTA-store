package routes

import (
	"project/controllers"
	"project/constants"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	//products
	e.GET("/products", controllers.GetProductsControllers)
	e.GET("/products/:id", controllers.GetProductControllers)

	//need token
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.POST("/products", controllers.CreateProductsController)

	return e
}
