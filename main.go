package main

import (
	"github.com/labstack/echo/v4"
	"test/controller"
)

func main() {
	e := echo.New()
	h := controller.NewItem()
	e.GET("/", contro)
}
