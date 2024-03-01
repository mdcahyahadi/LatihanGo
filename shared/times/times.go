package times

import (
	"strconv"
	"time"
)

// Now is a func to get current time string
func Now(hourOffset int, format string) string {
	return time.Now().UTC().Add(time.Hour * time.Duration(hourOffset)).Format(format)
}

// TimeStampString a timestamp string format that used by X-HEADER on Request ESB
func TimeStampString() string {
	dt := time.Now()
	timestamp := int64(time.Nanosecond) * dt.UnixNano() / int64(time.Millisecond)
	timestampStr := strconv.Itoa(int(timestamp))
	return timestampStr
}

// DateTimeFormatGeneralTimeStamp parses format 2019-12-02 00:00:00.0 or 2019-12-02 00:00:00.0
// To format 02 Jan 06 15:04 (RFC822) without location
// dateStr argument should be a string
func DateTimeFormatGeneralTimeStamp(dateStr string) string {
	t, _ := time.Parse(TimeDefaultFormat, dateStr)
	return t.Format(TimeDefaultFormatMonthString)
}

// DateFormatMutationPeriod parses format 20060102 to format 01/Jan/200
// dateStr argument should be a string
func DateFormatMutationPeriod(dateStr string) string {
	t, _ := time.Parse(DateWithoutHyphenFormat, dateStr)
	return t.Format(DateWithSlashFormat)
}

func TimeStringToDateNoHyphen(timeStr string) string {
	date, _ := time.Parse(TimeFormat, timeStr)
	return date.Format("20060102")
}

func GetLastDayOfMonth(date time.Time) int {
	currentYear, currentMonth, _ := date.Date()
	currentLocation := date.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	return lastOfMonth.Day()
}
