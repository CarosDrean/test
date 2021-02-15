package main

import (
	"github.com/labstack/echo/v4"
	"test/controller"
)

func main() {
	e := echo.New()
	cc := controller.Item{}
	e.GET("/", cc.GetAll)
	e.GET("/:id", cc.GetByID)
	e.POST("/", cc.Create)
	e.PUT("/:id", cc.Update)
	e.DELETE("/:id", cc.Delete)
	e.Logger.Fatal(e.Start(":8000"))
}
