package model

// Tag 标签
type Tag struct {
	Count int     `json:"count"`
	Name  string  `json:"name"`
	Posts []*Post `json:"posts"`
	Url   string  `json:"url"`
}
