// Copyright 2014 The sutil Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stime

import (
	"fmt"
	"strings"
	"time"
	//"fmt"
)

func GetHM(seconds int64) string {

	hours := seconds / 3600
	minutes := (seconds % 3600) / 60

	return fmt.Sprintf("%dH%dM", hours, minutes)
}

func GetDHMS(seconds int) string {
	days := seconds / (24 * 3600)
	hours := (seconds % (24 * 3600)) / 3600
	minutes := (seconds % 3600) / 60
	seconds = seconds % 60

	var parts []string
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%dD", days))
	}
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%dH", hours))
	}
	if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%dM", minutes))
	}
	if seconds > 0 {
		parts = append(parts, fmt.Sprintf("%dS", seconds))
	}

	if len(parts) == 0 {
		return "0S"
	}

	return strings.Join(parts, "")

}

func GetDHMSMilli(milliSeconds int) string {
	seconds := milliSeconds / 1000
	resMilli := milliSeconds % 1000
	days := seconds / (24 * 3600)
	hours := (seconds % (24 * 3600)) / 3600
	minutes := (seconds % 3600) / 60
	seconds = seconds % 60

	var parts []string
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%dD", days))
	}
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%dH", hours))
	}
	if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%dM", minutes))
	}
	if seconds > 0 || resMilli > 0 {
		parts = append(parts, fmt.Sprintf("%d.%dS", seconds, resMilli))
	}

	if len(parts) == 0 {
		return "0S"
	}

	return strings.Join(parts, "")

}

func FmtTimeMDHMS(secondStamp int64) string {

	// 将毫秒转换为秒
	//seconds := secondStamp
	//msTimestamp := secondStamp * 1000
	//nanoseconds := (msTimestamp % 1000) * 1000000
	timestamp := time.Unix(secondStamp, 0)

	// 格式化时间
	// 例如："2006-01-02 15:04:05" 是Go中的标准时间格式字符串
	//formattedTime := timestamp.Format("2006-01-02 15:04:05")
	formattedTime := timestamp.Format("01-02 15:04:05")

	return formattedTime

}

func FmtTimeYMDHMSEast8(secondStamp int64) string {
	return fmtTimeEast8("06-01-02 15:04:05", secondStamp)

}

func FmtTimeMDHMSEast8(secondStamp int64) string {
	return fmtTimeEast8("01-02 15:04:05", secondStamp)
}

func fmtTimeEast8(format string, secondStamp int64) string {

	// 将毫秒转换为秒
	//seconds := secondStamp
	//msTimestamp := secondStamp * 1000
	//nanoseconds := (msTimestamp % 1000) * 1000000
	timestamp := time.Unix(secondStamp, 0)

	// 定义东八区时区
	location, err := time.LoadLocation("Asia/Shanghai") // 东八区，例如北京时间
	if err != nil {
		return ""
	}

	// 将UTC时间转换为东八区时间
	cstTime := timestamp.In(location)

	// 格式化时间
	// 例如："2006-01-02 15:04:05" 是Go中的标准时间格式字符串
	//formattedTime := timestamp.Format("2006-01-02 15:04:05")
	formattedTime := cstTime.Format(format)

	return formattedTime

}

func FmtTimeYMDEast8(secondStamp int64) string {

	// 将毫秒转换为秒
	//seconds := secondStamp
	//msTimestamp := secondStamp * 1000
	//nanoseconds := (msTimestamp % 1000) * 1000000
	timestamp := time.Unix(secondStamp, 0)

	// 定义东八区时区
	location, err := time.LoadLocation("Asia/Shanghai") // 东八区，例如北京时间
	if err != nil {
		return ""
	}

	// 将UTC时间转换为东八区时间
	cstTime := timestamp.In(location)

	// 格式化时间
	// 例如："2006-01-02 15:04:05" 是Go中的标准时间格式字符串
	//formattedTime := timestamp.Format("2006-01-02 15:04:05")
	formattedTime := cstTime.Format("2006-01-02")

	return formattedTime

}

func DayBeginStamp(now int64) int64 {

	_, offset := time.Now().Zone()
	//fmt.Println(zone, offset)
	return now - (now+int64(offset))%int64(3600*24)
	//return (now + int64(offset))/int64(3600 * 24) * int64(3600 * 24) - int64(offset)

}

func HourBeginStamp(now int64) int64 {

	_, offset := time.Now().Zone()
	//fmt.Println(zone, offset)
	return now - (now+int64(offset))%int64(3600)
	//return (now + int64(offset))/int64(3600 * 24) * int64(3600 * 24) - int64(offset)

}

// 获取指定天的时间范围
// 天格式 2006-01-02
// 为空时候返回当天的
func DayBeginStampFromStr(day string) (int64, error) {
	nowt := time.Now()
	now := nowt.Unix()

	var begin int64
	if len(day) > 0 {
		tm, err := time.ParseInLocation("2006-01-02", day, nowt.Location())
		if err != nil {
			return 0, err
		}

		begin = tm.Unix()

	} else {
		begin = DayBeginStamp(now)

	}

	return begin, nil

}

func DayBeginStampFromDayStrEast8(day string) (int64, error) {

	if len(day) == 0 {
		return 0, fmt.Errorf("day empty")
	}

	location, err := time.LoadLocation("Asia/Shanghai") // 东八区，例如北京时间
	if err != nil {
		return 0, err
	}

	var begin int64
	tm, err := time.ParseInLocation("2006-01-02", day, location)
	if err != nil {
		return 0, err
	}

	begin = tm.Unix()

	return begin, nil

}

func WeekScope(stamp int64) (int64, int64) {
	now := time.Unix(stamp, 0)
	weekday := time.Duration(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	year, month, day := now.Date()
	currentZeroDay := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	begin := currentZeroDay.Add(-1 * (weekday - 1) * 24 * time.Hour).Unix()
	return begin, begin + 24*3600*7 - 1
}

func MonthScope(stamp int64) (int64, int64) {
	now := time.Unix(stamp, 0)
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	return firstOfMonth.Unix(), lastOfMonth.Unix() + 3600*24 - 1
}

func MonthScopeEast8(secondStamp int64) (int64, int64) {
	//now := time.Unix(stamp, 0)
	now := time.Unix(secondStamp, 0)

	// 定义东八区时区
	location, err := time.LoadLocation("Asia/Shanghai") // 东八区，例如北京时间
	if err != nil {
		return 0, 0
	}

	// 将UTC时间转换为东八区时间
	now = now.In(location)

	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	return firstOfMonth.Unix(), lastOfMonth.Unix() + 3600*24 - 1
}

var (
	Since2014 int64 = time.Date(2014, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano() / 1000
)

func Timestamp2014() uint64 {
	return uint64(time.Now().UnixNano()/1000 - Since2014)

}

type runTimeStat struct {
	//logkey string
	since time.Time
}

//func (m *runTimeStat) StatLog() string {
//	return fmt.Sprintf("%s RUNTIME:%d", m.logkey, m.Duration())
//}

func (m *runTimeStat) Millisecond() int64 {
	return m.Microsecond() / 1000
}

func (m *runTimeStat) Microsecond() int64 {
	return m.Duration().Nanoseconds() / 1000

}

func (m *runTimeStat) Nanosecond() int64 {
	return m.Duration().Nanoseconds()
}

func (m *runTimeStat) Duration() time.Duration {
	return time.Since(m.since)
}

func (m *runTimeStat) Reset() {
	m.since = time.Now()
}

func (m *runTimeStat) BeginTime() time.Time {
	return m.since
}

// func NewTimeStat(key string) *runTimeStat {
func NewTimeStat() *runTimeStat {
	return &runTimeStat{
		//logkey: key,
		since: time.Now(),
	}
}
