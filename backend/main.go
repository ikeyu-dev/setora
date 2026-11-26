package main

import (
	"log"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS("http://127.0.0.1:8181"))
	e.GET("/health", health)

	if err := e.Start("127.0.0.1:8180"); err != nil {
		log.Fatal(err)
	}
}
