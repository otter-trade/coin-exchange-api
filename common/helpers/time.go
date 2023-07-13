package helpers

import "time"

// Time time()
func Time() int64 {
	return time.Now().Unix()
}

func ParseInLocationDate(str string) (dateTime time.Time) {
	var LOC, _ = time.LoadLocation("Asia/Shanghai")
	dateTime, _ = time.ParseInLocation("2006-01-02 15:04:05", str, LOC)
	return dateTime
}
