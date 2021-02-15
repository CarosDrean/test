package storage

import (
	"errors"
	"test/model"
)

var Items []model.Item

func Create(item model.Item) {
	Items = append(Items, item)
}

func Update(Id string, item model.Item) error {
	itemM := model.Item{}
	index := 0
	for i, e := range Items {
		if e.Id == Id {
			itemM = e
			index = 0
		}
	}
	if itemM.Name == "" {
		return errors.New("no se encontro el elemento")
	}
	Items[index] = item
	return nil
}

func Delete(Id string) error {
	var newList []model.Item
	index := 0
	for i, e := range Items {
		if e.Id != Id {
			newList = append(newList, e)
		} el
	}
}

func GetByID(Id string) error {

}

func GetAll() error {

}
