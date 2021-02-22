package main

import (
  "net/http"
  "github.com/labstack/echo"
)

func Hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello ALTA Store Customer")
}

func main() {
  e := echo.New()
  e.Get("/hello", Hello)
  e.Logger.Fatal(e.Start(":8000"))
}
