package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Handler is the interface for bucky logger
type Handler interface {
	Info(description ...interface{})
	Debug(description ...interface{})
	Warn(err error)
	Error(err error)
	WithField(string, interface{}) *Logger
	WithFields(map[string]interface{}) *Logger
}

// Logger is the implementation of Handler interface
type Logger struct {
	handler *logrus.Entry
}

// New instantiates bucky logger instance
func New(appname string, opts Options) (Handler, error) {
	f, err := os.OpenFile("./build/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	mw := io.MultiWriter(os.Stdout, f)
	log := logrus.New()

	//switch opts.Format {
	//case JSONLogFormat:
	//	log.SetFormatter(&logrus.JSONFormatter{
	//		TimestampFormat: time.RFC3339,
	//	})
	//case SyslogLogFormat:
	//	log.SetFormatter(&logrus.TextFormatter{
	//		TimestampFormat: time.RFC3339,
	//		FullTimestamp:   true,
	//	})
	//}

	// log.SetReportCaller(true)
	log.SetOutput(mw)
	entry := log.WithFields(logrus.Fields{
		"app": appname,
	})

	return &Logger{
		handler: entry,
	}, nil
}

// Error logs for log-level error
func (l *Logger) Error(err error) {
	l.handler.Log(logrus.ErrorLevel, err.Error())
}

// Warn logs for log-level warning
func (l *Logger) Warn(err error) {
	l.handler.Log(logrus.WarnLevel, err.Error())
}

// Info logs for log-level info
func (l *Logger) Info(description ...interface{}) {
	l.handler.Log(logrus.InfoLevel,
		description...,
	)
}

// Debug logs for log-level debug
func (l *Logger) Debug(description ...interface{}) {
	l.handler.Log(logrus.DebugLevel,
		description...,
	)
}

// WithField logs with added field
func (l *Logger) WithField(s string, val interface{}) *Logger {
	obj := l.handler.WithField(s, val)
	return &Logger{
		handler: obj,
	}
}

// WithFields logs with given added fields
func (l *Logger) WithFields(d map[string]interface{}) *Logger {
	obj := l.handler.WithFields(logrus.Fields(d))
	return &Logger{
		handler: obj,
	}
}
