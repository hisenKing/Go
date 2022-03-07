package timeutil

import (
	"fmt"
	"strconv"
	"time"
)

// TimestampToDateTime 10位时间戳转换为DateTime 例：1609527845 -> 2021-01-02 03:04:05
func TimestampToDateTime(timestamp int64) string {
	return TimeToDateTime(time.Unix(timestamp, 0))
}

// TimeToDateTime time.Time转换为DateTime 例：time.Time -> 2021-01-02 03:04:05
func TimeToDateTime(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}

// TimestampToDate 10位时间戳转换为Date(日期) 例：1609527845 -> 2021-01-02
func TimestampToDate(timestamp int64) string {
	return TimeToDate(time.Unix(timestamp, 0))
}

// TimeToDate time.Time转换为Date(日期) 例：time.Time -> 2021-01-02
func TimeToDate(time time.Time) string {
	return time.Format("2006-01-02")
}

// TimestampToYearMonth 10位时间戳转换为Date(日期) 例：1609527845 -> 2021-01
func TimestampToYearMonth(timestamp int64) string {
	return TimeToYearMonth(time.Unix(timestamp, 0))
}

// TimeToYearMonth time.Time转换为Date(日期) 例：time.Time -> 2021-01
func TimeToYearMonth(time time.Time) string {
	return time.Format("2006-01")
}

// TimestampToDateWithDelimiter 10位时间戳转换为Date(日期) 例：1609527845 分隔符(-) -> 2021-01-02 例2: 1609527845 分隔符(.) -> 2021.01.02
func TimestampToDateWithDelimiter(timestamp int64, delimiter string) string {
	var format = fmt.Sprintf("2006%s01%s02", delimiter, delimiter)
	return TimeToDateWithFormat(time.Unix(timestamp, 0), format)
}

// DateTimeToTimestamp 年月日转时间戳 （注：转换出错结果为0） 2021-01-02 03:04:05 -> 1609527845
func DateTimeToTimestamp(dateTime string) int64 {
	ts, _ := DateTimeToTimestampWithError(dateTime)
	return ts
}

// DateTimeToTimestampWithError 年月日转时间戳 2021-01-02 03:04:05 -> 1609527845
func DateTimeToTimestampWithError(dateTime string) (int64, error) {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", dateTime, time.Local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// 时间转时间戳
func DateToTimestamp(dateTime, format string) int64 {
	t, _ := time.ParseInLocation(format, dateTime, time.Local)
	return t.Unix()
}

// YmdToDate 例：20211001->2021年10月1日
func YmdToDate(day string) string {
	t, _ := time.ParseInLocation("20060102", day, time.Local)
	t.Format("2006年01月02日")
	return TimeToDateWithFormat(t, "2006年01月02日")
}

// YmToDate 例：202110->2021年10月
func YmToDate(month string) string {
	t, _ := time.ParseInLocation("200601", month, time.Local)
	return TimeToDateWithFormat(t, "2006年01月")
}

func TimeToDateWithFormat(time time.Time, format string) string {
	return time.Format(format)
}

// TimeToYm time.Time转换为Date(日期) 例：time.Time -> 202101
func TimeToYm(time time.Time) string {
	return time.Format("200601")
}

// TimestampToYmdHis 10位时间戳转换 例：1609527845 -> 2021-01-02 03:04:05
// 注：如果为0则返回空字符串
func TimestampToYmdHis(timestamp int64) string {
	if timestamp == 0 {
		return ""
	}
	return TimeToDateTime(time.Unix(timestamp, 0))
}

// TimestampToYmd 10位时间戳转换为日期 例：1609527845 -> 2021-01-02
// 注：如果为0则返回空字符串
func TimestampToYmd(timestamp int64) string {
	if timestamp == 0 {
		return ""
	}
	return TimeToDate(time.Unix(timestamp, 0))
}

// TimestampToYm 10位时间戳转换为日期 例：1609527845 -> 202101
// 注：如果为0则返回空字符串
func TimestampToYm(timestamp int64) string {
	if timestamp == 0 {
		return ""
	}
	return TimeToYm(time.Unix(timestamp, 0))
}

// GetMonthFirstDay 获取月份的第一天 例：202110 -> 20211001
func GetMonthFirstDay(month string) string {
	t, _ := time.ParseInLocation("200601", month, time.Local)
	return t.Format("20060102")
}

// GetMonthEndDay 获取月份的最后一天 例：202110 -> 20211031
func GetMonthEndDay(month string) string {
	t, _ := time.ParseInLocation("200601", month, time.Local)
	y, m, _ := t.Date()
	stamp := time.Date(y, m, 1, 0, 0, 0, 0, t.Location()).AddDate(0, 1, -1)
	return stamp.Format("20060102")
}

// GetTimeNextDay 获取天的下一天 例：20211001 -> 20211002
func GetTimeNextDay(day string) string {
	t, _ := time.ParseInLocation("20060102", day, time.Local)
	y, m, d := t.Date()
	stamp := time.Date(y, m, d, 0, 0, 0, 0, t.Location()).AddDate(0, 0, 1)
	return stamp.Format("20060102")
}

// GetDayAndNextStartTime 获取天的开始时间戳和下一天的开始时间戳  20211001 ->1633017600,1633104000
func GetDayAndNextStartTime(day string) (int64, int64) {
	t, _ := time.ParseInLocation("20060102", day, time.Local)
	y, m, d := t.Date()
	startTime := time.Date(y, m, d, 0, 0, 0, 0, t.Location()).Unix()
	nextStartTime := time.Date(y, m, d, 0, 0, 0, 0, t.Location()).AddDate(0, 0, 1).Unix()
	return startTime, nextStartTime
}

// GetMonthAndNextStartTime 获取天所在月的开始时间戳和下一个月的开始时间戳  20211005 ->1633017600,1635696000
func GetMonthAndNextStartTime(day string) (int64, int64) {
	t, _ := time.ParseInLocation("20060102", day, time.Local)
	y, m, _ := t.Date()
	startTime := time.Date(y, m, 1, 0, 0, 0, 0, t.Location()).Unix()
	nextStartTime := time.Date(y, m, 1, 0, 0, 0, 0, t.Location()).AddDate(0, 1, 0).Unix()
	return startTime, nextStartTime
}

// GetTimestampDateFirstTimestamp 时间戳所在当天第一秒
func GetTimestampDateFirstTimestamp(timestamp int64) int64 {
	var t = time.Unix(timestamp, 0)
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location()).Unix()
}

// GetTimestampPrevDateFirstTimestamp 时间戳所在上一天第一秒
func GetTimestampPrevDateFirstTimestamp(timestamp int64) int64 {
	return GetTimestampDateFirstTimestamp(timestamp - 86400)
}

// GetTimestampNextDateFirstTimestamp 时间戳所在下一天第一秒
func GetTimestampNextDateFirstTimestamp(timestamp int64) int64 {
	return GetTimestampDateFirstTimestamp(timestamp + 86400)
}

// GetTimestampMonthFirstTimestamp 时间戳所在当月第一秒
func GetTimestampMonthFirstTimestamp(timestamp int64) int64 {
	var t = time.Unix(timestamp, 0)
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location()).Unix()
}

// GetTimestampPrevMonthFirstTimestamp 时间戳所在上月第一秒
func GetTimestampPrevMonthFirstTimestamp(timestamp int64) int64 {
	var t = time.Unix(timestamp, 0)
	y, m, _ := t.Date()

	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location()).AddDate(0, -1, 0).Unix()

	//var year = y
	//var month = m - 1
	//if m <= 1 {
	//	month = 12
	//	year = y - 1
	//}
	//return time.Date(year, month, 1, 0, 0, 0, 0, t.Location()).AddDate(0, -1, 0).Unix()
}

// GetTimestampNextMonthFirstTimestamp 时间戳所在下月第一秒
func GetTimestampNextMonthFirstTimestamp(timestamp int64) int64 {
	var t = time.Unix(timestamp, 0)
	y, m, _ := t.Date()

	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location()).AddDate(0, 1, 0).Unix()

	//var year = y
	//var month = m + 1
	//if m >= 12 {
	//	month = 1
	//	year = y + 1
	//}
	//return time.Date(year, month, 1, 0, 0, 0, 0, t.Location()).Unix()
}

func GetCountdownToMHS(countdown int64) string {
	if countdown <= 0 {
		return ""
	}
	var day = countdown / 60 / 60 / 24
	var hour = countdown / 60 / 60 % 24
	var minute = countdown / 60 % 60
	countdownTips := strconv.FormatInt(minute, 10) + "分"
	if hour > 0 {
		countdownTips = strconv.FormatInt(hour, 10) + "时" + countdownTips
	}
	if day > 0 {
		countdownTips = strconv.FormatInt(day, 10) + "天" + countdownTips
	}
	return countdownTips
}
