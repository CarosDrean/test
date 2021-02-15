package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"test/model"
)

type Item struct {}

var Items []model.Item

func (cc Item) Create(c echo.Context) error {
	item := model.Item{}
	err := c.Bind(&item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "error")
	}
	id := 0
	for _, _ = range Items {
		id++
	}
	Items = append(Items, item)
	return c.JSON(http.StatusCreated, "Creado con exito")
}

func (cc Item) Update(c echo.Context) error {
	id := c.Param("id")
	item := model.Item{}
	err := c.Bind(&item)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "error")
	}
	idInt, _ := strconv.Atoi(id)
	for i, e := range Items {
		if e.Id == idInt {
			Items[i] = e
		}
	}
	Items = append(Items, item)
	return c.JSON(http.StatusOK, "Actualizado con exito")
}

func (cc Item) Delete(c echo.Context) error {
	id := c.Param("id")
	var newItems []model.Item
	idInt, _ := strconv.Atoi(id)
	validator := false
	for _, e := range Items {
		if e.Id == idInt {
			validator = true
		} else {
			newItems = append(newItems, e)
		}
	}
	Items = newItems
	if validator {
		return c.JSON(http.StatusNotFound, "No se encontro el elemento")
	}
	return c.JSON(http.StatusOK, "Eliminado con exito")
}

func (cc Item) GetByID(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	for _, e := range Items {
		if e.Id == idInt {
			return c.JSON(http.StatusOK, e)
		}
	}
	return c.JSON(http.StatusNotFound, "No se encontro el elemento")
}

func (cc Item) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, Items)
}
