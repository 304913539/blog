package model

import (
	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	AppKey    string `json:"app_key" gorm:"type:varchar(20);not null;default:''"`
	AppSecret string `json:"app_secret" gorm:"type:varchar(50);not null;default:''"`
}

func (a *Auth) TableName() string {
	return "blog_auth"
}

func (a *Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("app_key = ? AND app_secret = ?", a.AppKey, a.AppSecret)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}
	return auth, nil
}
