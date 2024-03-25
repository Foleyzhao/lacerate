package service

import (
	"bufio"
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"lacerate/core/config"
	"lacerate/core/log"
	"lacerate/core/model"
	"lacerate/core/utils"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

// 文章列表
var postList []*model.Post

// Content markdown内容
type Content struct {
	Title       string
	Description string
	Date        string
	Categories  []string
	Tags        []string
	Content     string
}

// GetPostList 获取文章列表
func GetPostList() []*model.Post {
	return postList
}

// CreateMarkdown 创建markdown文件
func CreateMarkdown(filename string) string {
	file := path.Join(config.GlobalConf.Markdown, filename+".md")
	_, err := os.Stat(file)
	if !os.IsNotExist(err) {
		log.Log.Errorf("The file already exists.")
		os.Exit(1)
	}
	src, err := utils.CreateFile(config.GlobalConf.Markdown, filename+".md")
	if err != nil {
		panic(err)
	}
	date := time.Now().Format("2006-01-02")
	now := time.Now().Format("15:04:05")
	informationContent := `---
date: ` + date + `
time: ` + now + `
title: ` + filename + `
categories:
-
tagList:
-
-
---`
	err = utils.WriteFile(src, informationContent)
	if err != nil {
		panic(err)
	}

	return src
}

// MarkdownList 获取markdown文件夹下所有文件
func MarkdownList(markdownDir string) (markdownList []string) {
	_ = filepath.Walk(markdownDir, func(path string, f os.FileInfo, err error) error {
		if err != nil { //忽略错误
			return err
		}
		if f.IsDir() {
			return nil
		}
		//if strings.ToLower(f.Name()) == "readme.md" {
		//	return nil
		//}
		if f.Name() == "about.md" {
			return nil
		}
		if strings.HasSuffix(f.Name(), ".md") {
			markdownList = append(markdownList, path)
		}
		return nil
	})

	return markdownList
}

// LoadPostList 加载文章列表
func LoadPostList() {
	postList = make([]*model.Post, 0)
	markdownList := MarkdownList(config.GlobalConf.Markdown)
	for _, markdown := range markdownList {
		post, err := loadMarkdownContent(markdown)
		if err == nil {
			post.Url = CreatePostLink(post)
			postList = append(postList, post)
			for _, _cate := range post.Category {
				if len(_cate) <= 0 {
					continue
				}
				category := categoryList[_cate]
				if category == nil {
					category = &model.Category{Count: 0, Name: _cate, Posts: make([]*model.Post, 0), Url: "/category/" + _cate}
					categoryList[_cate] = category
				}

				category.Count += 1
				category.Posts = append(category.Posts, post)
			}
			for _, _tag := range post.Tags {
				if len(_tag) <= 0 {
					continue
				}
				tag := tagList[_tag]
				if tag == nil {
					tag = &model.Tag{Count: 0, Name: _tag, Posts: make([]*model.Post, 0), Url: "/tag/" + _tag}
					tagList[_tag] = tag
				}
				tag.Count += 1
				tag.Posts = append(tag.Posts, post)
			}
		} else {
			panic(err)
		}
	}
	sort.Sort(model.PostList(postList))
}

// 加载markdown内容生成文章
func loadMarkdownContent(file string) (post *model.Post, err error) {
	post = &model.Post{}
	content, err := ReadMarkdownContent(file)
	if err != nil {
		return nil, err
	}
	if post.Summary == "" {
		summaryLine := config.GlobalConf.SummaryLine
		post.Summary, err = generateSummary(content.Content, summaryLine)

		if err != nil {
			return nil, err
		}
	}

	post.Title = content.Title
	post.Description = content.Description
	post.Category = content.Categories
	post.Tags = content.Tags
	post.Content = utils.MarkdownToHtml(content.Content)
	post.CreatedAt = utils.Str2Unix("2006-01-02", content.Date)

	return post, nil
}

// 生成摘要
func generateSummary(content string, lines int) (string, error) {
	buff := bufio.NewReader(bytes.NewBufferString(content))
	dst := ""
	for lines > 0 {
		line, err := buff.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		if strings.Contains(strings.ToLower(line), "[toc]") {
			continue
		}
		reg := regexp.MustCompile(`!\[(.*)\]\((.*)\)`)
		if reg.MatchString(line) {
			continue
		}
		if strings.Trim(line, "\r\n\t ") == "```" {
			continue
		}
		dst += line
		lines--
	}

	return utils.MarkdownToHtml(dst), nil
}

// ReadMarkdownContent 读取markdown内容
func ReadMarkdownContent(path string) (content *Content, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	br := bufio.NewReader(f)
	line, err := br.ReadString('\n')
	if err != nil {
		return nil, err
	}
	if !strings.HasPrefix(line, "---") {
		err = fmt.Errorf("markdown file format error, the file header must start with '---': " + path)
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	for {
		line, err = br.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
		}
		if strings.HasPrefix(line, "---") {
			break
		}
		buf.WriteString(line)
	}
	err = yaml.Unmarshal(buf.Bytes(), &content)
	contentByte, err := io.ReadAll(br)
	if err != nil {
		return nil, err
	}
	fi, _ := f.Stat()
	if content.Title == "" {
		content.Title = strings.Replace(strings.TrimRight(fi.Name(), ".md"), config.GlobalConf.Markdown+"/", "", 1)
	}
	if content.Date == "" {
		content.Date = utils.Format(fi.ModTime().Unix())
	}
	content.Content = string(contentByte)
	return
}

// CreatePostLink 创建文章链接
func CreatePostLink(art *model.Post) string {
	t := time.Unix(art.CreatedAt, 0)
	year, month, day := t.Date()
	link := fmt.Sprintf("/%s/%d/%d/%d/%s/", "post", year, month, day, utils.Convert(art.Title))
	return link
}

// GetHomePostList 获取首页文章列表
func GetHomePostList() []*model.Post {
	num := config.GlobalConf.HomePostNum
	if num == 0 || len(postList) <= num {
		num = len(postList)
	}

	homePostList := make([]*model.Post, num)
	copy(homePostList, postList)

	return homePostList
}
