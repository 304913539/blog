package model

import (
	"blog-service/pkg/app"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title         string `json:"title" gorm:"type:varchar(100);not null;default:''"`
	Desc          string `json:"desc" gorm:"type:varchar(200);not null;default:''"`
	Content       string `json:"content" gorm:"type:text;not null;"`
	CoverImageUrl string `json:"cover_image_url" gorm:"type:varchar(255);not null;default:'';comment:封面图片地址"`
	State         uint8  `json:"state" gorm:"type:tinyint;not null;default:1;comment:状态 0 为禁用、1 为启用"`
}

func (a Article) TableName() string {
	return "blog_article"
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}
