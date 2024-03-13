package utils

// Count 统计分片长度
func Count(sl []string) (num int) {
	num = 0
	for _, s := range sl {
		if s != "" {
			num += 1
		}
	}
	return
}

// Lt 判断小于
func Lt(a, b int) bool { return a < b }

// Eq 判断等于
func Eq(a, b int) bool { return a == b }

// Gt 判断大于
func Gt(a, b int) bool { return a > b }
