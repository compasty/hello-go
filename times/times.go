package times

import "time"

func DaysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func GetCurrentFormatStr(fmtStr string) string {
	if fmtStr == "" {
		fmtStr = "2006-01-02 15:04:05"
	}
	return time.Now().Format(fmtStr)
}

func TimeStr2Time(fmtStr, valueStr, locStr string) int64 {
	loc := time.Local
	if locStr != "" {
		loc, _ = time.LoadLocation(locStr) // 设置时区
	}
	if fmtStr == "" {
		fmtStr = "2006-01-02 15:04:05"
	}
	t, _ := time.ParseInLocation(fmtStr, valueStr, loc)
	return t.Unix()
}
