package btime

import (
	"time"
)

// 时间戳转换为给定的格式，
// 如果给定模板为空，则转换为 "2006-01-02 15:04:05"
// 如果给定时间戳为0，则默认使用当前时间
func FormatTimestamp(timestamp int64, timeTemplate string) (timeString string) {
	tt := "2006-01-02 15:04:05"
	if timeTemplate != "" {
		tt = timeTemplate
	}
	ts := time.Now().Unix()
	if timestamp != 0 {
		ts = timestamp
	}
	timeString = time.Unix(ts, 0).Format(tt)
	return timeString
}

func String2Time(timeString string) (time.Time, error) {
	tt := "2006-01-02 15:04:05"
	st, err := time.Parse(tt, timeString)
	return st, err
}

// 获取当前时刻的n月前第一天的当前时刻
func FirstDayOfMonthAfterAddDate(t time.Time, years int, months int) time.Time {
	year, month, _ := t.Date()
	hour, min, sec := t.Clock()
	// AddDate后月份的第一天
	firstDayOfMonthAfterAddDate := time.Date(year+years, month+time.Month(months), 1,
		hour, min, sec, t.Nanosecond(), t.Location())
	return firstDayOfMonthAfterAddDate
}

// 获取当前时刻的n月前最后一天
func LastDayOfMonthAfterAddDate(t time.Time, years int, months int) time.Time {
	year, month, _ := t.Date()
	// AddDate后月份的第一天
	firstDayOfMonthAfterAddDate := time.Date(year+years, month+time.Month(months), 1,
		0, 0, 0, 0, t.Location())
	// AddDate后月份的最后一个时刻
	lastDayOfMonthAfterAddDate := firstDayOfMonthAfterAddDate.AddDate(0, 1, 0).Add(-time.Nanosecond)
	return lastDayOfMonthAfterAddDate
}
