package utils

import "time"

// Format 日期格式化
func Format(unix int64) string {
	t := time.Unix(unix, 0)
	return t.Format("2006-01-02")
}

// Month 月份格式化
func Month(unix int64) string {
	t := time.Unix(unix, 0)
	return t.Format("1")
}

// Year 年份格式化
func Year(unix int64) string {
	t := time.Unix(unix, 0)
	return t.Format("2006")
}

// CMonth 日期格式化（不带年份）
func CMonth(unix int64) string {
	t := time.Unix(unix, 0)
	return t.Format("01-02")
}

// Str2Unix 字符串转时间
func Str2Unix(layout, timeStr string) int64 {
	tm, _ := time.Parse(layout, timeStr)
	return tm.Unix()
}
