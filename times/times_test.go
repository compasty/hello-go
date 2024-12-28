package times

import (
	"fmt"
	"testing"
	"time"
)

func TestDemoUse(t *testing.T) {
	now := time.Now()
	// 返回当前秒时间戳
	fmt.Println(now.Unix())
	// 返回当前纳秒时间戳
	fmt.Println(now.UnixNano())
	// 时间戳小数部分 单位：纳秒
	fmt.Println(now.Nanosecond())
	// 返回日期
	year, month, day := now.Date()
	fmt.Printf("year:%d, month:%d, day:%d\n", year, month, day)
	// 年
	fmt.Println(now.Year())
	// 月
	fmt.Println(now.Month())
	// 日
	fmt.Println(now.Day())

	// 时分秒
	hour, minute, second := now.Clock()
	fmt.Printf("hour:%d, minute:%d, second:%d\n", hour, minute, second)
	// 时
	fmt.Println(now.Hour())
	// 分
	fmt.Println(now.Minute())
	// 秒
	fmt.Println(now.Second())

	// 返回星期
	fmt.Println(now.Weekday())
	// 返回一年中对应的第几天
	fmt.Println(now.YearDay())
	// 返回时区
	fmt.Println(now.Location())

	// 返回一年中第几天
	fmt.Println(now.YearDay())
}

func TestFormatTime(t *testing.T) {
	now := time.Now()

	fmt.Println(now.Format("2006-01-02 15:03:04"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:03:04"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
}

func TestCreateTime(t *testing.T) {
	now := time.Now()
	layout := "2006-01-02 15:04:05"
	// 根据秒数和纳秒数创建时间
	t1 := time.Unix(now.Unix(), 0)
	fmt.Println(t1.Format(layout))
	// 分别指定年，月，日，时，分，秒，纳秒，时区
	t2 := time.Date(2011, time.Month(3), 12, 15, 30, 20, 0, now.Location())
	fmt.Println(t2.Format(layout))
}

func TestParseTime(t *testing.T) {

}
