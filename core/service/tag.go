package service

import (
	"lacerate/core/model"
)

// 标签列表
var tagList map[string]*model.Tag

// 初始化
func init() {
	tagList = make(map[string]*model.Tag)
}

// GetTagList 获取标签列表
func GetTagList() map[string]*model.Tag {
	return tagList
}
