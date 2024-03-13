package utils

import "html/template"

// Unescaped 解析html
func Unescaped(x string) interface{} {
	return template.HTML(x)
}
