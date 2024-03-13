package utils

import (
	bf "github.com/russross/blackfriday"
	"log"
	"regexp"
	"strings"
)

// TocTitle 目录标题
var TocTitle = "<h4>目录:</h4>"

// nav正则
var navRegex = regexp.MustCompile(`(?ismU)<nav>(.*)</nav>`)

// MarkdownToHtml markdown转html
func MarkdownToHtml(content string) (str string) {
	defer func() {
		e := recover()
		if e != nil {
			str = content
			log.Println("Render markdown err:", e)
		}
	}()
	htmlFlags := 0
	if strings.Contains(strings.ToLower(content), "[toc]") {
		htmlFlags |= bf.HTML_TOC
	}
	htmlFlags |= bf.HTML_USE_XHTML
	htmlFlags |= bf.HTML_USE_SMARTYPANTS
	htmlFlags |= bf.HTML_SMARTYPANTS_FRACTIONS
	htmlFlags |= bf.HTML_SMARTYPANTS_LATEX_DASHES
	htmlFlags |= bf.HTML_FOOTNOTE_RETURN_LINKS
	renderer := bf.HtmlRenderer(htmlFlags, "", "")
	extensions := 0
	extensions |= bf.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= bf.EXTENSION_TABLES
	extensions |= bf.EXTENSION_FENCED_CODE
	extensions |= bf.EXTENSION_AUTOLINK
	extensions |= bf.EXTENSION_STRIKETHROUGH
	extensions |= bf.EXTENSION_SPACE_HEADERS
	extensions |= bf.EXTENSION_HARD_LINE_BREAK
	extensions |= bf.EXTENSION_FOOTNOTES
	str = string(bf.Markdown([]byte(content), renderer, extensions))
	if htmlFlags&bf.HTML_TOC != 0 {
		found := navRegex.FindIndex([]byte(str))
		if len(found) > 0 {
			toc := str[found[0]:found[1]]
			toc = TocTitle + toc
			str = str[found[1]:]
			reg := regexp.MustCompile(`\[toc\]|\[TOC\]`)
			str = reg.ReplaceAllString(str, toc)
		}
	}
	return str
}
