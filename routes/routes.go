package routes

import (
	"project/controllers"
	"project/constants"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// users
	e.POST("/users/login", controllers.LoginController)
	e.POST("/users/register", controllers.RegisterUserController)

	//categories
	e.GET("/categories", controllers.GetCategoriesControllers)
	e.GET("/categories/:id", controllers.GetCategoryControllers)

	//products
	e.GET("/products", controllers.GetProductsControllers)
	e.GET("/product", controllers.GetProductsWithCategoryControllers)
	e.GET("/products/:id", controllers.GetProductControllers)

	//token
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/users", controllers.GetUsersController)
	r.GET("/users/:id", controllers.GetUserController)
	r.POST("/categories", controllers.CreateCategoriesController)
	r.POST("/products", controllers.CreateProductsController)

	return e
}
