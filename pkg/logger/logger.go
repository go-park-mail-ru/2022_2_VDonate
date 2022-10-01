package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	Logrus *logrus.Entry
}

func NewLogrus(port string, dbType string) (*Logger, error) {
	host, _ := os.Hostname()
	logrusLogger := logrus.WithFields(logrus.Fields{
		"host":     host,
		"port":     port,
		"database": dbType,
	})
	logrusLogger.Logger.SetFormatter(&logrus.TextFormatter{})
	logrusLogger.Logger.SetLevel(logrus.DebugLevel)

	return &Logger{logrusLogger}, nil
}
