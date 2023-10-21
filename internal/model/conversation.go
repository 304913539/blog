package model

import "gorm.io/gorm"

type Conversation struct {
	gorm.Model
	Name   string `json:"name" gorm:"type:string;not null;default:''"`
	UserId uint32 `json:"user_id" gorm:"type:int;not null;default:0"`
}

func (c Conversation) TableName() string {
	return "blog_conversation"
}
