package service

import (
	"lacerate/core/config"
	"lacerate/core/model"
	"lacerate/core/utils"
	"os"
	"path"
	"time"
)

// GetAbout 获取about内容
func GetAbout() (post *model.Post, err error) {
	post = &model.Post{}
	about := path.Join(config.GlobalConf.Markdown, "/about.md")
	if _, err := os.Stat(about); os.IsNotExist(err) {
		return post, nil
	}
	content, err := os.ReadFile(about)
	if err != nil {
		return nil, err
	}
	post.Title = ""
	post.Content = utils.MarkdownToHtml(string(content))
	post.CreatedAt = time.Now().Unix()
	return post, nil
}
