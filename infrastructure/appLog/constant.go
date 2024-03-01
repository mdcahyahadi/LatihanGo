package appLog

const (
	ServiceName string = "mygram-svc"

	LogFolder         string = "logs/"
	LogTimeFormat     string = "2024-01-01T12:12:12-0700"
	LogFilenameFormat string = "log-2024-01-1.log"

	LogTypeRequest  string = "REQUEST"
	LogTypeResponse string = "RESPONSE"
	LogTypeWarning  string = "WARNING"
)
