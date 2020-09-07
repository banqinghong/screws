package btime

import (
	"time"
)

// 时间戳转换为给定的格式，如果给定模板为空，则转换为 "2006-01-02 15:04:05"
func FormatTimestamp(ts int64, timeTemplate string)(timeString string){
	tt := "2006-01-02 15:04:05"
	if timeTemplate != "" {
		tt = timeTemplate
	}
	timeString = time.Unix(ts, 0).Format(tt)
	return timeString
}
