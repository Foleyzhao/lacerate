package config

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

var GlobalConf = Config()

var confFileName = "config.yml"

var confContent = `
# 站点信息
title: xxx's Blog
subtitle: 主页
description: xxx's Blog
keywords: blog
# 作者信息
author: HappyNewYear
avatar: /assets/avatar.jpg
github: https://github.com/Foleyzhao
email: foleyzhao@163.com
# 配置信息
summary_line: 6
home_post_num: 10
# 文件存储
theme: theme/blog
markdown: markdown
html: /data/www/html
storage: storage
# 文件监听
paths:
  - markdown
suffix:
  - md
  - yml
# 自定义信息
home_title: xxx's Blog
archive_title: 归档
tag_title: 标签
category_title: 分类
about_title: 关于我
`

// SystemConfig 系统配置
type SystemConfig struct {
	Title         string   `yaml:"title"`
	SubTitle      string   `yaml:"subtitle"`
	Description   string   `yaml:"description"`
	Keywords      string   `yaml:"keywords"`
	Author        string   `yaml:"name"`
	Avatar        string   `yaml:"avatar"`
	Github        string   `yaml:"github"`
	Email         string   `yaml:"email"`
	SummaryLine   int      `yaml:"summary_line"`
	HomePostNum   int      `yaml:"home_post_num"`
	Theme         string   `yaml:"theme"`
	Markdown      string   `yaml:"markdown"`
	Html          string   `yaml:"html"`
	Storage       string   `yaml:"storage"`
	Paths         []string `yaml:"paths"`
	Suffix        []string `yaml:"suffix"`
	HomeTitle     string   `yaml:"home_title,omitempty"`
	ArchiveTitle  string   `yaml:"archive_title,omitempty"`
	TagTitle      string   `yaml:"tag_title,omitempty"`
	CategoryTitle string   `yaml:"category_title,omitempty"`
	AboutTitle    string   `yaml:"about_title,omitempty"`
}

// 加载系统配置
func loadConf() ([]byte, error) {
	_, err := os.Stat(confFileName)
	if os.IsNotExist(err) {
		CreateConf()
	}
	file, err := os.Open(confFileName)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	return io.ReadAll(file)
}

// CreateConf 创建系统配置文件
func CreateConf() {
	_, err := os.Stat(confFileName)
	if os.IsNotExist(err) {
		_, err := os.OpenFile(confFileName, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			panic(err)
		}

		var confWrite = []byte(confContent)
		err = os.WriteFile(confFileName, confWrite, 0666)
		if err != nil {
			panic(err)
		}
	}
}

// Config 系统配置
func Config() SystemConfig {
	confContent, err := loadConf()
	if err != nil {
		panic("failed to load configuration file: " + err.Error())
	}

	c := SystemConfig{}
	err = yaml.Unmarshal(confContent, &c)
	if err != nil {
		panic("failed to parse the configuration file: " + err.Error())
	}

	return c
}
