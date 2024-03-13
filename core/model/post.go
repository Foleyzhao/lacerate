package model

// PostList 文章列表
type PostList []*Post

// Post 文章
type Post struct {
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Summary     string   `json:"summary"`
	Content     string   `json:"content"`
	Tags        []string `json:"tags"`
	Category    []string `json:"category"`
	CreatedAt   int64    `json:"created_at"`
	UpdatedAt   int64    `json:"updated_at"`
	Url         string   `json:"url"`
}

func (p PostList) Len() int {
	return len(p)
}

func (p PostList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PostList) Less(i, j int) bool {
	return p[i].CreatedAt > p[j].CreatedAt
}
