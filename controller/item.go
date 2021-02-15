package controller

import (
	"github.com/labstack/echo/v4"
	"test/interfaces"
)

type Item struct {
	Store interfaces.Item
}

func NewItem(storage interfaces.Item) Item {
	return Item{Store: storage}
}

func (cc Item) Create(c echo.Context) {

}

func (cc Item) Update(c echo.Context) {

}

func (cc Item) Delete(c echo.Context) {

}

func (cc Item) GetByID(c echo.Context) {

}

func (cc Item) GetAll(c echo.Context) {

}
