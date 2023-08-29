package model

import "gorm.io/gorm"

type ArticleTag struct {
	gorm.Model
	TagID     uint32 `json:"tag_id" gorm:"type:int;not null;default:0;comment:标签 ID"`
	ArticleID uint32 `json:"article_id" gorm:"type:int;not null;default:0;comment:文章 ID"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
