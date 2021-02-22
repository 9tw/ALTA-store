package main

import (
  "net/http"
  "github.com/labstack/echo"
)

// Hello menginformasikan bahwa API berjalan dengan baik
func Hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello ALTA Store Customer")
}

func main() {
  e := echo.New()
  e.GET("/hello", Hello)
  e.Logger.Fatal(e.Start(":8000"))
}
