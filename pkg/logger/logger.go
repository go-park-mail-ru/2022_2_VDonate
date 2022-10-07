package logger

import (
	"github.com/sirupsen/logrus"
)

type Log struct {
	Logrus *logrus.Logger
}

func NewLogrus() *Log {
	logrusLogger := Log{Logrus: logrus.New()}
	logrusLogger.Logrus.SetFormatter(&logrus.TextFormatter{})
	logrusLogger.Logrus.SetLevel(logrus.DebugLevel)

	return &logrusLogger
}
