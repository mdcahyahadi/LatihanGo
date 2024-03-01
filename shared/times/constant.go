package times

// constant declare for time needed
const (
	TimeFormat string = "2006-01-02 15:04:05" // TimeFormat is a const for default time format
	TimeGmt    int    = 7                     // TimeGmt is an alias for GMT(+7)
)

// constant declare for formatting date time
const (
	TimeDefaultFormat            string = "2006-01-02 15:04:05"
	DateDefaultFormat            string = "2006-01-02"
	TimeDefaultFormatMonthString string = "02 Jan 2006 15:04"
	DateWithoutHyphenFormat      string = "20060102"
	DateWithSlashFormat          string = "01/Jan/2006"
	TimeRFC3339FormatDateTime    string = "2006-01-02T15:04:05"
)
