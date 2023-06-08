package syslog

import (
	"log"
	"log/syslog"
)

// SYSLOG_DEFAULT defines syslog default priority
const SYSLOG_DEFAULT syslog.Priority = syslog.LOG_ERR

// logger Defines file log data structure
type logger struct {
	logType syslog.Priority
	logger  *log.Logger
}

// Create return pointer to logger
func Create(logType syslog.Priority) *logger {
	if logType == 0 {
		logType = SYSLOG_DEFAULT
	}
	return &logger{
		logType: logType,
	}
}

// Config Set Type SysLog
func (f *logger) Config() {
	var err error
	f.logger, err = syslog.NewLogger(f.logType, log.Lmicroseconds|log.Ltime)
	if err != nil {
		panic(err)
	}
}

// Log Print Log To SysLog
func (f *logger) Log(massage string) bool {
	f.logger.Println(massage)
	return true
}
