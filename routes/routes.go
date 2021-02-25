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
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users/login", controllers.LoginController)
	e.POST("/users/register", controllers.RegisterUserController)

	//categories
	e.GET("/categories", controllers.GetCategoriesControllers)
	e.GET("/categories/:id", controllers.GetCategoryControllers)

	//products
	e.GET("/products", controllers.GetProductsControllers)
	e.GET("/products/:id", controllers.GetProductControllers)

	//token
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.POST("/categories", controllers.CreateCategoriesController)
	r.POST("/products", controllers.CreateProductsController)

	return e
}
