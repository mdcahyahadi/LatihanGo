package appLog

import (
	"fmt"
	"io"
	"latihan/infrastructure/environment"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type AppLog struct {
	Logger *logrus.Logger
}

func New(env environment.AppEnvironment, options ...loggerOption) *AppLog {
	loggerOptions := createDefaultOptions()
	logger := logrus.New()
	if options != nil {
		loggerOptions.apply(options)
	}
	logger.SetLevel(loggerOptions.getLevel(env))

	return &AppLog{Logger: logger}
}

func (log *AppLog) setLogFile() string {
	currentTime := time.Now()
	timestamp := currentTime.Format(LogTimeFormat)
	filename := LogFolder + currentTime.Format(LogFilenameFormat)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Logger.SetOutput(io.Discard)
	} else {
		log.Logger.SetOutput(file)
	}
	return timestamp
}

func (log *AppLog) setJSON() {
	formatter := new(logrus.JSONFormatter)
	formatter.DisableTimestamp = true
	log.Logger.SetLevel(log.Logger.Level)
	log.Logger.SetFormatter(formatter)
}

func (log *AppLog) setText() {
	formatter := new(logrus.TextFormatter)
	formatter.FullTimestamp = true
	formatter.DisableColors = false
	formatter.DisableQuote = true
	log.Logger.SetLevel(log.Logger.Level)
	log.Logger.SetFormatter(formatter)
}

func (log *AppLog) LogDebug(args ...interface{}) {
	log.setText()
	timestamp := log.setLogFile()
	log.Logger.Debug(fmt.Sprintf("%s %s", timestamp, args))
}

func (log *AppLog) LogWarn(userID, message string, err error) {
	log.setText()
	timestamp := log.setLogFile()
	log.Logger.WithFields(logrus.Fields{
		"service":   ServiceName,
		"userID":    userID,
		"error":     err,
		"message":   message,
		"timestamp": timestamp,
	}).Warn(LogTypeWarning)
}

func (log *AppLog) LogRequest(r *http.Request, req interface{}) {
	log.setJSON()
	timestamp := log.setLogFile()
	log.Logger.WithFields(logrus.Fields{
		"service":       ServiceName,
		"type":          LogTypeRequest,
		"url":           fmt.Sprintf("%s %s%s", r.Method, r.Host, r.URL),
		"requestHeader": r.Header,
		"requestBody":   req,
		"timestamp":     timestamp,
	}).Info(LogTypeRequest)
}

func (log *AppLog) LogResponse(statusCode int, r *http.Request, req, resp interface{}) {
	log.setJSON()
	timestamp := log.setLogFile()
	log.Logger.WithFields(logrus.Fields{
		"service":    ServiceName,
		"type":       LogTypeResponse,
		"url":        fmt.Sprintf("%s %s%s", r.Method, r.Host, r.URL),
		"statusCode": statusCode,
		"request":    req,
		"response":   resp,
		"timestamp":  timestamp,
	}).Info(LogTypeResponse)
}
