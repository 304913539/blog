package model

import (
	"blog-service/pkg/app"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name  string `json:"name" gorm:"type:varchar(100);not null;default:'';comment:标签名称"`
	State uint8  `json:"state" gorm:"type:int;not null;default:1;comment:状态 0 为禁用、1 为启用"`
}

func (t *Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t *Tag) Count(db *gorm.DB) (int64, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t *Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (t *Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t *Tag) Update(db *gorm.DB) error {
	return db.Model(&Tag{}).Where("id = ?", t.ID).Update(t.Name, t).Error
}

func (t *Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ?", t.ID).Delete(&t).Error
}
