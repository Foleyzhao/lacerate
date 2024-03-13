package model

import "time"

// PublishedYears 年份归档列表
type PublishedYears []*PublishedYear

// PublishedMonths 月份归档列表
type PublishedMonths []*PublishedMonth

// PublishedYear 年份归档
type PublishedYear struct {
	YearStr   string                     `json:"year"`
	Months    []*PublishedMonth          `json:"months"`
	MonthDict map[string]*PublishedMonth `json:"-"`
}

// PublishedMonth 月份归档
type PublishedMonth struct {
	MonthStr string     `json:"month"`
	Posts    []*Post    `json:"posts"`
	Month    time.Month `json:"-"`
}

func (y PublishedYears) Len() int {
	return len(y)
}

func (y PublishedYears) Swap(i, j int) {
	y[i], y[j] = y[j], y[i]
}

func (y PublishedYears) Less(i, j int) bool {
	return y[i].YearStr > y[j].YearStr
}

func (m PublishedMonths) Len() int {
	return len(m)
}

func (m PublishedMonths) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m PublishedMonths) Less(i, j int) bool {
	return m[i].Month > m[j].Month
}
