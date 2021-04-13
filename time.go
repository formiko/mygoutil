package mygoutil

import (
	"math/rand"
	"time"
)

// 时间戳转时间字符串
func StampToStr(stamp int64) string {
	timeLayout := "2006-01-02 15:04:05"
	return time.Unix(stamp, 0).Format(timeLayout)
}

// 随机时间戳
func RandStamp() int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63()
}

// 随机时间字符串
func RandTimeStr() string {
	return StampToStr(RandStamp())
}

// 随机前后days天的时间戳
func TheseDaysStamp(days int) int64 {
	return time.Now().Unix() + rand.Int63n(86400 * int64(days) * 2) - 86400 * int64(days)
}

// 随机前后days天的时间字符串
func TheseDaysStr(days int) string {
	return StampToStr(TheseDaysStamp(days))
}