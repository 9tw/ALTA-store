package controllers

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"net/http"
	"project/lib/database"
	"project/models"
	"strconv"

	"github.com/labstack/echo"
)

// GetUsersController get all users
func GetUsersController(c echo.Context) error {
  users, err := database.GetUsers()
	
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
    "status": "success",
    "data": users,
  })
}

// GetUserController get a user by given user ID
func GetUserController(c echo.Context) error {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }

	users, getErr := database.GetUser(id)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

  return c.JSON(http.StatusOK, map[string]interface{}{
    "status": "success",
    "data": users,
  })
}

// LoginController check if given email and password correct and return fake token
func LoginController(c echo.Context) error {
  inputUser := models.Users{}
  c.Bind(&inputUser)

  if inputUser.Email == "" || inputUser.Password == "" {
    return echo.NewHTTPError(
			http.StatusBadRequest, 
			errors.New("Email atau password tidak boleh kosong"),
		)
  }

	// hash input password
	sha := sha1.New()
	sha.Write([]byte(inputUser.Password))
	inputUser.Password = fmt.Sprintf("%x", sha.Sum(nil))

	users, getErr := database.Login(&inputUser)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}

  return c.JSON(http.StatusOK, map[string]interface{}{
    "status": "success",
    "data": users,
  })
}

// RegisterUserController create a new user
func RegisterUserController(c echo.Context) error {
  inputUser := models.Users{}
  c.Bind(&inputUser)

  if inputUser.Email == "" || inputUser.Password == "" {
    return echo.NewHTTPError(
			http.StatusBadRequest, 
			errors.New("Email atau password tidak boleh kosong"),
		)
  }

	// hash input password
	sha := sha1.New()
	sha.Write([]byte(inputUser.Password))
	inputUser.Password = fmt.Sprintf("%x", sha.Sum(nil))

	users, getErr := database.Register(&inputUser)
	if getErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, getErr.Error())
	}

  return c.JSON(http.StatusOK, map[string]interface{}{
    "status": "success",
    "data": users,
  })
}
