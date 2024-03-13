package service

import (
	"lacerate/core/model"
	"lacerate/core/utils"
	"sort"
	"time"
)

// GetArchive 获取归档信息
func GetArchive() []*model.PublishedYear {
	archiveYear := make(model.PublishedYears, 0)
	_archiveYear := make(map[string]*model.PublishedYear)
	for _, post := range postList {
		yearStr := utils.Year(post.CreatedAt)
		monthStr := utils.Month(post.CreatedAt)
		_month := time.Unix(post.CreatedAt, 0).Month()
		year := _archiveYear[yearStr]
		if year == nil {
			year = &model.PublishedYear{YearStr: yearStr, Months: make([]*model.PublishedMonth, 0), MonthDict: make(map[string]*model.PublishedMonth)}
			_archiveYear[yearStr] = year
		}
		month := year.MonthDict[monthStr]
		if month == nil {
			month = &model.PublishedMonth{MonthStr: monthStr, Posts: []*model.Post{}, Month: _month}
			year.MonthDict[monthStr] = month
		}
		month.Posts = append(month.Posts, post)
	}
	for _, year := range _archiveYear {
		monthArray := make(model.PublishedMonths, 0)
		for _, month := range year.MonthDict {
			monthArray = append(monthArray, month)
		}
		sort.Sort(monthArray)
		year.MonthDict = nil
		year.Months = monthArray
		archiveYear = append(archiveYear, year)
	}
	sort.Sort(archiveYear)
	return archiveYear
}
