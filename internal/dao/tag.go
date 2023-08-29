package dao

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
	"gorm.io/gorm"
)

func (d *Dao) CountTag(name string, state uint8) (int64, error) {
	tag := model.Tag{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetTagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateTag(name string, state uint8) error {
	tag := model.Tag{
		Name:  name,
		State: state,
	}

	return tag.Create(d.engine)
}

func (d *Dao) UpdateTag(id uint, name string, state uint8) error {
	tag := model.Tag{
		Name:  name,
		State: state,
		Model: gorm.Model{ID: id},
	}

	return tag.Update(d.engine)
}

func (d *Dao) DeleteTag(id uint) error {
	tag := model.Tag{Model: gorm.Model{ID: id}}
	return tag.Delete(d.engine)
}
