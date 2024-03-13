package service

import (
	"lacerate/core/model"
)

// 分类列表
var categoryList map[string]*model.Category

// 初始化
func init() {
	categoryList = make(map[string]*model.Category)
}

// GetCategoryList 获取菜单列表
func GetCategoryList() map[string]*model.Category {
	return categoryList
}
