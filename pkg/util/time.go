// 获取Time相关数据
package util

import (
	"math/rand"
	"strconv"
	"time"
)

var (
	weekdays = map[string]int{
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6,
		"Sunday":    7,
	}
)

// 获取当前时间戳
func GetTime(t string) string {
	now_time := time.Now().UnixNano()/1e6 + 120*1000
	if t == "13" {
		return strconv.FormatInt(now_time, 10)
	}
	now_time = time.Now().Unix()
	return strconv.FormatInt(now_time, 10)
}

// 获取当前时间
func GetNowTimeStr() string {
	return strconv.Itoa(int(time.Now().Unix()))
}
func ItoTime(i int) time.Time {
	return time.Unix(int64(i), 0)
}

// 获取当前时间戳
func GetTimeStr(t int, formate string) string {
	now_time := time.Now()
	formateStr := "20060102150405"
	if formate == "Y-m-d" {
		formateStr = "2006-01-02"
	}
	if formate == "Ymd" {
		formateStr = "20060102"
	}
	if formate == "Y-m-d H:i" {
		formateStr = "2006-01-02 15:04"
	}
	if formate == "Y-m-d H:i:s" {
		formateStr = "2006-01-02 15:04:05"
	}
	if formate == "Y" {
		formateStr = "2006"
	}
	if formate == "m月d日" {
		formateStr = "1月2日"
	}
	if t > 0 {
		now_time = time.Unix(int64(t), 0)
	}
	return now_time.Format(formateStr)
}

// 获取今日 Int
func GetTodayTime() int {
	the_time, _ := time.ParseInLocation("2006-01-02", GetTimeStr(0, "Ymd"), time.Local)
	return int(the_time.Unix())
}

// 计算前后时间0点
func GetDuTime(_day int) time.Time {

	beforeTime := time.Now().AddDate(0, 0, _day)
	//the_time, _ :=time.ParseInLocation("2006-01-02", GetTimeStr(beforeTime, "Ymd"), time.Local)
	return beforeTime
}

// 获取今日周几
func GetTodayWeekday() int {
	return weekdays[time.Now().Weekday().String()]
}

// 获取凌晨时间戳
func Lingchen() time.Time {
	formatLayout := "2006-01-02"
	today := time.Now().Format(formatLayout)
	// fmt.Print(today)
	t, _ := time.ParseInLocation(formatLayout, today, time.Local)
	// fmt.Print(t.Unix())
	return t
}

// 获取凌晨时间戳
func DateTime() int {
	formatLayout := "2006-01-02"
	today := time.Now().Format(formatLayout)
	// fmt.Print(today)
	t, _ := time.ParseInLocation(formatLayout, today, time.Local)
	// fmt.Print(t.Unix())
	return int(t.Unix())
}

// 获取当前时间 int
func GetNowInt() int {
	return int(time.Now().Unix())
}

// time 转 Int
func Time2Int(t time.Time) int {
	return int(t.Unix())
}

// time 转 string
func Time2Str(t time.Time, formate string) string {
	i := Time2Int(t)
	return GetTimeStr(i, formate)
}

// string 转 time
func Str2Time(t string) (time.Time, error) {
	var timeLayoutStr = "20060102150405"
	return time.Parse(timeLayoutStr, t)
}

// 获取今日int  日
//func GetTodayDayInt() int {
//return StoI(GetTimeStr(0, "Ymd"))
//}

// int 秒 转时间量
func GetValdTime(seconds int) map[string]interface{} {
	dv := 86400
	dh := 3600
	dm := 60
	day := seconds / dv
	hour := (seconds % dv) / dh
	min := (seconds % dh) / dm
	sec := seconds % dm
	return map[string]interface{}{
		"day":    day,
		"hour":   hour,
		"minute": min,
		"second": sec,
	}

}

// 获取指定日期0点0分
func TimeZeroPoint(t time.Time) time.Time {
	formatLayout := "2006-01-02"
	t1 := t.Format(formatLayout)
	t2, _ := time.ParseInLocation(formatLayout, t1, time.Local)
	return t2
}

// 计算日期相差多少天
func TimeSubDays(t1, t2 time.Time) int {
	t1 = TimeZeroPoint(t1)
	t2 = TimeZeroPoint(t2)
	hours := t1.Sub(t2).Hours()
	// print(hours)
	return int(hours / 24)
}

// 获取随机数
func GetIndex(max int) int {
	return rand.Intn(max)
}
