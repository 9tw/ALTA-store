package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// User struct contains User object
type User struct {
  ID          int     `json:"id" form:"id"`
  Nama        string  `json:"nama" form:"nama"`
  Email       string  `json:"email" form:"email"`
  Password    string  `json:"password" form:"password"`
  Token       string  `json:"token" form:"token"`
  NoHp        string  `json:"no_hp" form:"no_hp"`
  Foto        string  `json:"foto" form:"foto"`
  Alamat      string  `json:"alamat" form:"alamat"`
  Status      string  `json:"status" form:"status"`
  IsVerified  bool    `json:"is_verified" form:"is_verified"`
}

// Array sebagai pengganti database
var users = []User{}

// Hello menginformasikan bahwa API berjalan dengan baik
func Hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello ALTA Store Customer")
}

// Controllers Users

// GetUsersController get all users
func GetUsersController(c echo.Context) error {
  return c.JSON(http.StatusOK, map[string]interface{}{
    "message": "successful operation",
    "data": users,
  })
}

// GetUserController get a user by given user ID
func GetUserController(c echo.Context) error {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, map[string]interface{}{
      "message": "bad request",
    })
  }

  for _, user := range users {
    if user.ID == id {
      return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "successful operation",
        "data": user,
      })
    }
  }

  return c.JSON(http.StatusOK, map[string]interface{}{
    "message": "successful operation",
    "data": nil,
  })
}

// LoginController check if given email and password correct and return fake token
func LoginController(c echo.Context) error {
  inputUser := User{}
  c.Bind(&inputUser)

  if inputUser.Email == "" || inputUser.Password == "" {
    return c.JSON(http.StatusBadRequest, map[string]interface{}{
      "message": "bad request",
    })
  }

  for _, user := range users {
    if user.Email == inputUser.Email &&
        user.Password == inputUser.Password {
          return c.JSON(http.StatusOK, map[string]interface{}{
            "message": "successful operation",
            "data": user,
          })
        }
  }

  return c.JSON(http.StatusOK, map[string]interface{}{
    "message": "successful operation",
    "data": nil,
  })
}

// RegisterUserController create a new user
func RegisterUserController(c echo.Context) error {
  inputUser := User{}
  c.Bind(&inputUser)

  if inputUser.Email == "" || inputUser.Password == "" {
    return c.JSON(http.StatusBadRequest, map[string]interface{}{
      "message": "bad request",
    })
  }

  inputUser.ID = len(users) + 1
  users = append(users, inputUser)
  return c.JSON(http.StatusOK, map[string]interface{}{
    "message": "successful operation",
    "data": inputUser,
  })
}

func main() {
  e := echo.New()
  e.GET("/hello", Hello)

  // Routing Users
  e.GET("/users", GetUsersController)
  e.GET("/users/:id", GetUserController)
  e.POST("/users/login", LoginController)
  e.POST("/users/register", RegisterUserController)

  e.Logger.Fatal(e.Start(":8000"))
}
