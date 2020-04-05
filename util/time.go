package util

import (
"time"
)

// NowTime 获取当前时间: 2006-01-02 15:04:05
func NowTime() (now string) {
	now = time.Now().Format("2006-01-02 15:04:05") //这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
	return
}

// NowDate 获取当前时间: 2006-01-02
func NowDate() (now string) {
	now = time.Now().Format("2006-01-02") //这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
	return
}

func StrToTimestamp(date string) int64 {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, date, loc)
	return theTime.Unix()
}

func GetTodayBeginTime() int64 {
	now := time.Now().Format("2006-01-02 00:00:00")
	timeLayout := "2006-01-02 00:00:00"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, now, loc)
	return theTime.Unix()
}
