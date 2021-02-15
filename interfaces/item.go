package interfaces

import "test/model"

type Item interface {
	Create(item model.Item)
	Update(Id string, item model.Item)
	Delete(Id string)
	GetByID(Id string) model.Item
	GetAll() []model.Item
}
