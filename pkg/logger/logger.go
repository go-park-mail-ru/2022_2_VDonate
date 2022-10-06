package logger

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	Logrus *logrus.Entry
}

func NewLogrus(host, port string) (*Logger, error) {
	logrusLogger := logrus.WithFields(logrus.Fields{
		"host": host,
		"port": port,
	})
	logrusLogger.Logger.SetFormatter(&logrus.TextFormatter{})
	logrusLogger.Logger.SetLevel(logrus.DebugLevel)

	return &Logger{logrusLogger}, nil
}
