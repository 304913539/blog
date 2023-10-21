package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UserId           uint32 `json:"user_id" gorm:"type:int;not null;default:0"`
	Conversation     string `json:"conversation" gorm:"type:string;not null;default:''"`
	UserContent      string `json:"user_content" gorm:"type:text;not null"`
	AssistantContent string `json:"assistant_content" gorm:"type:text;not null"`
}

func (m Message) TableName() string {
	return "blog_message"
}
