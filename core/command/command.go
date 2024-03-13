package command

import (
	"fmt"
	"lacerate/core/config"
	"lacerate/core/log"
	"net/http"
	"os"
	"strconv"
)

const (
	// HELP 帮助信息
	HELP = `
Usage:

lacerate command [args...]

	Initialize the blog folder
    	lacerate init

	Create a new markdown file
    	lacerate new filename

	Compile the blog
    	lacerate compile/c

    Open the file listener
    	lacerate watch/w

	Open the file server
    	lacerate http/web [port]

    Run all Lacerate services
    	lacerate run [port]
	`
)

// PrintHelp 打印帮助信息
func PrintHelp() {
	fmt.Println(HELP)
}

// Initialize 初始化操作
func Initialize() {
	config.CreateConf()
	CreateDir()
	log.Log.Debug("the initialization is successful!")
}

// ListenHttpServer 启动http服务
func ListenHttpServer(port int) {
	log.Log.Info("open the built-in web server...")
	p := strconv.Itoa(port)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(config.GlobalConf.Html+"/assets/"))))
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(config.GlobalConf.Html))))
	log.Log.Debugf("the built-in web server is successfully turned on and the port is monitored: %d...", port)
	err := http.ListenAndServe(":"+p, nil)
	if err != nil {
		log.Log.Errorf("listen http serve error: %s", err)
	}
}

// CreateDir 创建博客目录
func CreateDir() {
	_, err := os.Stat(config.GlobalConf.Html)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(config.GlobalConf.Html, os.ModePerm); err != nil {
			panic(err)
		}
	}
	_, err = os.Stat(config.GlobalConf.Markdown)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(config.GlobalConf.Markdown, os.ModePerm); err != nil {
			panic(err)
		}
	}
	_, err = os.Stat(config.GlobalConf.Storage)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(config.GlobalConf.Storage, os.ModePerm); err != nil {
			panic(err)
		}
	}
	_, err = os.Stat(config.GlobalConf.Theme)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(config.GlobalConf.Theme, os.ModePerm); err != nil {
			panic(err)
		}
	}
}
