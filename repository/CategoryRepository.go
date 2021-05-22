package repository

import (
	"github.com/jinzhu/gorm"
	"lzh.practice/ginessential/common"
	"lzh.practice/ginessential/model"
)

type CategoryRepositpry struct {
	DB *gorm.DB
}

func NewCategoryRepository() CategoryRepositpry {
	return CategoryRepositpry{DB: common.GetDB()}
}

func (c CategoryRepositpry) Create(name string) (*model.Category, error) {
	category := model.Category{
		Name: name,
	}
	if err := c.DB.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepositpry) Update(category model.Category, name string) (*model.Category, error) {
	if err := c.DB.Model(&category).Update("name", name).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepositpry) SelectById(id int) (*model.Category, error) {
	var category model.Category
	if err := c.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepositpry) DeletById(id int) error {
	var category model.Category
	if err := c.DB.Delete(&category, id).Error; err != nil {
		return err
	}
	return nil
}
