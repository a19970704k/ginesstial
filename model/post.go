package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//文章的增删改查
type Post struct {
	//文章要以uuid作为主键
	ID uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	//作者id
	UserId uint `json:"user_id" gorm:"not null"`
	//文章分类id
	CategoryId uint `json:"category_id" gorm:"not null"`
	//文章分类实体
	Category  *Category
	Title     string `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg   string `json:"head_img"`
	Content   string `json:"content" gorm:"type:text;not null"`
	CreatedAt Time   `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt Time   `json:"updated_at" gorm:"type:timestamp"`
}

func (post *Post) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}
