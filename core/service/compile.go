package service

import (
	"html/template"
	"lacerate/core/config"
	"lacerate/core/log"
	"lacerate/core/utils"
	"os"
	"path"
	"strings"
)

// 全局数据
var data = map[string]interface{}{
	"title":       config.GlobalConf.Title,
	"subtitle":    config.GlobalConf.SubTitle,
	"description": config.GlobalConf.Description,
	"keywords":    config.GlobalConf.Keywords,
	"author":      config.GlobalConf.Author,
	"avatar":      config.GlobalConf.Avatar,
	"github":      config.GlobalConf.Github,
	"email":       config.GlobalConf.Email,
}

// html模板函数字典
var funcMaps = template.FuncMap{
	"unescaped": utils.Unescaped,
	"cmonth":    utils.CMonth,
	"format":    utils.Format,
	"count":     utils.Count,
	"lt":        utils.Lt,
	"gt":        utils.Gt,
	"eq":        utils.Eq,
	"md5":       utils.Xmd5,
}

// Compile 编译博客
func Compile() {
	defer func() {
		if r := recover(); r != nil {
			log.Log.Errorf("panic recovered from: %v", r)
		}
	}()
	log.Log.Info("start compiling your blog...")
	checkThemeFile()
	copyAssetsFile()
	LoadPostList()
	// 创建页面
	CompileHome()
	CompilePost()
	CompileArchive()
	CompileTagPage()
	CompileTag()
	CompileCategoryPage()
	CompileCategory()
	CompileAbout()
	storageBlogMap()
	log.Log.Debug("compilation complete...")
}

// 存储文章
func storageBlogMap() {
	storage, err := utils.NewStorage(config.GlobalConf.Storage, "storage.json")
	if err != nil {
		panic(err)
	}
	err = storage.Store(GetPostList())
	if err != nil {
		panic(err)
	}
}

// CompileHome 编译主页
func CompileHome() {
	title := config.GlobalConf.HomeTitle
	if len(strings.TrimSpace(title)) == 0 {
		data["title"] = "home page"
	} else {
		data["title"] = title
	}
	data["postList"] = GetHomePostList()
	data["categoryList"] = GetCategoryList()
	data["tagList"] = GetTagList()
	err := utils.MkDir(config.GlobalConf.Html)
	if err != nil {
		panic(err)
	}
	homePath := path.Join(config.GlobalConf.Html, "index.html")
	htmlFile, err := os.Create(homePath)
	if err != nil {
		panic(err)
	}
	t, err := template.New("main.tpl").Funcs(funcMaps).ParseFiles(config.GlobalConf.Theme+"/layout/main.tpl", config.GlobalConf.Theme+"/layout/home.tpl")
	if err != nil {
		panic(err)
	}
	err = t.Execute(htmlFile, data)
	if err != nil {
		panic(err)
	}
}

// 拷贝静态文件
func copyAssetsFile() {
	err := utils.CopyDir(path.Join(config.GlobalConf.Theme, "assets"), path.Join(config.GlobalConf.Html, "assets"))
	if err != nil {
		panic(err)
	}
}

// 校验模板文件
func checkThemeFile() {
	if _, err := os.Stat(config.GlobalConf.Theme); os.IsNotExist(err) {
		panic("you need to initialize and add the template file first.")
	}
}

// CompileCategoryPage 编译分类导航页
func CompileCategoryPage() {
	subTitle := config.GlobalConf.CategoryTitle
	if len(strings.TrimSpace(subTitle)) == 0 {
		data["subtitle"] = "article category"
	} else {
		data["subtitle"] = subTitle
	}
	data["categoryList"] = GetCategoryList()
	data["tagList"] = GetTagList()
	filepath := path.Join(config.GlobalConf.Html, "category")
	err := utils.MkDir(filepath)
	if err != nil {
		panic(err)
	}
	filename := path.Join(filepath, "index.html")
	htmlFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	t, err := template.New("main.tpl").Funcs(funcMaps).ParseFiles(config.GlobalConf.Theme+"/layout/category.tpl", config.GlobalConf.Theme+"/layout/main.tpl")
	if err != nil {
		panic(err)
	}
	err = t.Execute(htmlFile, data)
	if err != nil {
		panic(err)
	}
}

// CompileCategory 编译分类页
func CompileCategory() {
	cateList := GetCategoryList()
	data["categoryList"] = cateList
	data["tagList"] = GetTagList()
	for _, cate := range cateList {
		data["subtitle"] = cate.Name
		data["pageTitle"] = cate.Name
		data["content"] = cate.Posts
		data["count"] = cate.Count
		filepath := path.Join(config.GlobalConf.Html, "category", cate.Name)
		err := utils.MkDir(filepath)
		if err != nil {
			panic(err)
		}
		filename := path.Join(filepath, "index.html")
		htmlFile, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		t, err := template.New("main.tpl").Funcs(funcMaps).ParseFiles(config.GlobalConf.Theme+"/layout/page.tpl", config.GlobalConf.Theme+"/layout/main.tpl")
		if err != nil {
			panic(err)
		}
		err = t.Execute(htmlFile, data)
		if err != nil {
			panic(err)
		}
	}
}

// CompileTagPage 编译标签导航页
func CompileTagPage() {
	subTitle := config.GlobalConf.TagTitle
	if len(strings.TrimSpace(subTitle)) == 0 {
		data["subtitle"] = "article tags"
	} else {
		data["subtitle"] = subTitle
	}
	data["categoryList"] = GetCategoryList()
	data["tagList"] = GetTagList()
	filePath := path.Join(config.GlobalConf.Html, "tag")
	err := utils.MkDir(filePath)
	if err != nil {
		panic(err)
	}
	fileName := path.Join(filePath, "index.html")
	htmlFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	t, err := template.New("main.tpl").Funcs(funcMaps).ParseFiles(config.GlobalConf.Theme+"/layout/tag.tpl", config.GlobalConf.Theme+"/layout/main.tpl")
	if err != nil {
		panic(err)
	}
	err = t.Execute(htmlFile, data)
	if err != nil {
		panic(err)
	}
}

// CompileTag 编译标签页
func CompileTag() {
	tags := GetTagList()
	data["categoryList"] = GetCategoryList()
	data["tagList"] = GetTagList()
	for _, tag := range tags {
		data["subtitle"] = tag.Name
		data["pageTitle"] = tag.Name
		data["content"] = tag.Posts
		data["count"] = tag.Count
		data["tpl"] = config.GlobalConf.Theme + "/layout/page.html"
		filepath := path.Join(config.GlobalConf.Html, "tag", tag.Name)
		err := utils.MkDir(filepath)
		if err != nil {
			panic(err)
		}
		filename := path.Join(filepath, "index.html")
		htmlFile, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		t, err := template.New("main.tpl").Funcs(funcMaps).ParseFiles(config.GlobalConf.Theme+"/layout/page.tpl", config.GlobalConf.Theme+"/layout/main.tpl")
		if err != nil {
			panic(err)
		}
		err = t.Execute(htmlFile, data)
		if err != nil {
			panic(err)
		}
	}
}

// CompileAbout 编译关于我页
func CompileAbout() {
	about, err := GetAbout()
	if err != nil {
		panic(err)
	}
	subTitle := config.GlobalConf.AboutTitle
	if len(strings.TrimSpace(subTitle)) == 0 {
		data["subtitle"] = "about me"
	} else {
		data["subtitle"] = subTitle
	}
	data["post"] = about
	data["categoryList"] = GetCategoryList()
	data["tagList"] = GetTagList()
	filePath := path.Join(config.GlobalConf.Html, "about.html")
	htmlFile, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	t, err := template.New("main.tpl").Funcs(funcMaps).ParseFiles(config.GlobalConf.Theme+"/layout/post.tpl", config.GlobalConf.Theme+"/layout/main.tpl")
	if err != nil {
		panic(err)
	}
	err = t.Execute(htmlFile, data)
	if err != nil {
		panic(err)
	}
}

// CompilePost 编译文章页
func CompilePost() {
	data["categoryList"] = GetCategoryList()
	data["tagList"] = GetTagList()
	for _, post := range postList {
		data["subtitle"] = post.Title
		data["description"] = strings.TrimSpace(post.Summary)
		data["keywords"] = strings.Join(post.Tags, ",")
		data["post"] = post
		url := CreatePostLink(post)
		filePath := path.Join(config.GlobalConf.Html, url)
		err := utils.MkDir(filePath)
		if err != nil {
			panic(err)
		}
		fileName := path.Join(filePath, "index.html")
		htmlFile, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		t, err := template.New("main.tpl").Funcs(funcMaps).ParseFiles(config.GlobalConf.Theme+"/layout/post.tpl", config.GlobalConf.Theme+"/layout/main.tpl")
		if err != nil {
			panic(err)
		}
		err = t.Execute(htmlFile, data)
		if err != nil {
			panic(err)
		}
	}
}

// CompileArchive 编译归档页
func CompileArchive() {
	subTitle := config.GlobalConf.ArchiveTitle
	if len(strings.TrimSpace(subTitle)) == 0 {
		data["subtitle"] = "article archiving"
	} else {
		data["subtitle"] = subTitle
	}
	data["archive"] = GetArchive()
	data["categoryList"] = GetCategoryList()
	data["tagList"] = GetTagList()
	filePath := path.Join(config.GlobalConf.Html, "archive")
	err := utils.MkDir(filePath)
	if err != nil {
		panic(err)
	}
	fileName := path.Join(filePath, "index.html")
	htmlFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	t, err := template.New("main.tpl").Funcs(funcMaps).ParseFiles(config.GlobalConf.Theme+"/layout/archive.tpl", config.GlobalConf.Theme+"/layout/main.tpl")
	if err != nil {
		panic(err)
	}
	err = t.Execute(htmlFile, data)
	if err != nil {
		panic(err)
	}
}
