package utils

import (
	"strings"
)

// Convert 文章标题转换文章链接
func Convert(str string) string {
	str = strings.ToLower(str)
	ss := strings.SplitN(str, " ", -1)

	return strings.Join(ss, "-")
}
