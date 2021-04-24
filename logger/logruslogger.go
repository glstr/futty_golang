package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusLogger(option *LogOption) (*LogrusLogger, error) {
	f, err := os.OpenFile(option.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	logger := logrus.New()
	//logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(f)
	logger.SetReportCaller(true)

	return &LogrusLogger{logger}, nil
}

func (l *LogrusLogger) Debug(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

func (l *LogrusLogger) Notice(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}

func (l *LogrusLogger) Warn(format string, v ...interface{}) {
	l.logger.Warnf(format, v...)
}
