package routes

import (
	"project/controllers"
	"project/constants"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// New returns echo object that contains all routes.
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

	// transactions
	jwtGroup := e.Group("")
	jwtGroup.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	jwtGroup.GET("/transactions", controllers.GetTransactionsControllers)
	jwtGroup.GET("/transactions/:id", controllers.GetTransactionControllers)
	jwtGroup.POST("/transactions", controllers.CreateTransactionControllers)
	jwtGroup.PUT("/transactions", controllers.UpdateTransactionStatusControllers)

	// carts
	jwtGroup.GET("/carts", controllers.GetCartsControllers)
	jwtGroup.POST("/carts", controllers.CreateCartsControllers)
	jwtGroup.PUT("/carts/:id", controllers.UpdateCartsController)
	jwtGroup.DELETE("/carts/:id", controllers.DeleteCartsController)

	//token
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.POST("/categories", controllers.CreateCategoriesController)
	r.POST("/products", controllers.CreateProductsController)

	return e
}
