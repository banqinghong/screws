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
