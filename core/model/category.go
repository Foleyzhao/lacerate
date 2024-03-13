package model

// Category 分类
type Category struct {
	Count int     `json:"count"`
	Name  string  `json:"name"`
	Posts []*Post `json:"posts"`
	Url   string  `json:"url"`
}
