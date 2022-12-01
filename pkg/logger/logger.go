package logger

import (
	"encoding/json"
	"io"
	"strconv"
	"sync"
	"time"

	errorHandling "github.com/go-park-mail-ru/2022_2_VDonate/pkg/errors"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	Logrus *logrus.Logger
}

// New settings of logger.
func New() *Logger {
	newLogger := Logger{Logrus: logrus.New()}
	newLogger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   time.RFC3339,
		DisableHTMLEscape: true,
		PrettyPrint:       true,
	})
	return &newLogger
}

var (
	lock         sync.Mutex
	SingleLogger *Logger
)

func GetInstance() *Logger {
	lock.Lock()
	defer lock.Unlock()

	if SingleLogger == nil {
		SingleLogger = New()
	}
	return SingleLogger
}

func ToLevel(level string) log.Lvl {
	switch level {
	case "debug":
		return log.DEBUG
	case "info":
		return log.INFO
	case "warn":
		return log.WARN
	case "error":
		return log.ERROR
	default:
		return log.INFO
	}
}

func toLogrusLevel(level log.Lvl) logrus.Level {
	switch level {
	case log.DEBUG:
		return logrus.DebugLevel
	case log.INFO:
		return logrus.InfoLevel
	case log.WARN:
		return logrus.WarnLevel
	case log.ERROR:
		return logrus.ErrorLevel
	}

	return logrus.InfoLevel
}

func toEchoLevel(level logrus.Level) log.Lvl {
	switch level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.InfoLevel:
		return log.INFO
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	}

	return log.OFF
}

func (l *Logger) Println(v ...interface{}) {
	l.Logrus.Println(v)
}

func (l *Logger) Output() io.Writer {
	return l.Logrus.Out
}

func (l *Logger) SetOutput(w io.Writer) {
	l.Logrus.SetOutput(w)
}

func (l *Logger) Level() log.Lvl {
	return toEchoLevel(l.Logrus.Level)
}

func (l *Logger) SetLevel(v log.Lvl) {
	l.Logrus.Level = toLogrusLevel(v)
}

func (l *Logger) SetHeader(_ string) {}

func (l *Logger) Formatter() logrus.Formatter {
	return l.Logrus.Formatter
}

func (l *Logger) SetFormatter(formatter logrus.Formatter) {
	l.Logrus.Formatter = formatter
}

func (l *Logger) Prefix() string {
	return ""
}

func (l *Logger) SetPrefix(_ string) {}

func (l *Logger) Print(i ...interface{}) {
	l.Logrus.Print(i...)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.Logrus.Printf(format, args...)
}

func (l *Logger) Printj(j log.JSON) {
	b, _ := json.Marshal(j)
	l.Logrus.Println(string(b))
}

func (l *Logger) Debug(i ...interface{}) {
	l.Logrus.Debug(i...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Logrus.Debugf(format, args...)
}

func (l *Logger) Debugj(j log.JSON) {
	b, _ := json.Marshal(j)
	l.Logrus.Debugln(string(b))
}

func (l *Logger) Info(i ...interface{}) {
	l.Logrus.Info(i...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.Logrus.Infof(format, args...)
}

func (l *Logger) Infoj(j log.JSON) {
	b, _ := json.Marshal(j)
	l.Logrus.Infoln(string(b))
}

func (l *Logger) Warn(i ...interface{}) {
	l.Logrus.Warn(i...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Logrus.Warnf(format, args...)
}

func (l *Logger) Warnj(j log.JSON) {
	b, _ := json.Marshal(j)
	l.Logrus.Warnln(string(b))
}

func (l *Logger) Error(i ...interface{}) {
	l.Logrus.Error(i...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Logrus.Errorf(format, args...)
}

func (l *Logger) Errorj(j log.JSON) {
	b, _ := json.Marshal(j)
	l.Logrus.Errorln(string(b))
}

func (l *Logger) Fatal(i ...interface{}) {
	l.Logrus.Fatal(i...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Logrus.Fatalf(format, args...)
}

func (l *Logger) Fatalj(j log.JSON) {
	b, _ := json.Marshal(j)
	l.Logrus.Fatalln(string(b))
}

func (l *Logger) Panic(i ...interface{}) {
	l.Logrus.Panic(i...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.Logrus.Panicf(format, args...)
}

func (l *Logger) Panicj(j log.JSON) {
	b, _ := json.Marshal(j)
	l.Logrus.Panicln(string(b))
}

func Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			start := time.Now()
			err := next(c)
			if err != nil {
				c.Error(err)
			}
			stop := time.Now()

			p := req.URL.Path

			bytesIn := req.Header.Get(echo.HeaderContentLength)

			message := errorHandling.CutCode(err)

			ctxLog := GetInstance().Logrus.WithFields(map[string]interface{}{
				"remote_ip":     c.RealIP(),
				"host":          req.Host,
				"uri":           req.RequestURI,
				"method":        req.Method,
				"path":          p,
				"error":         message,
				"referer":       req.Referer(),
				"user_agent":    req.UserAgent(),
				"status":        res.Status,
				"latency":       strconv.FormatInt(stop.Sub(start).Nanoseconds()/1000, 10),
				"latency_human": stop.Sub(start).String(),
				"bytes_in":      bytesIn,
				"bytes_out":     strconv.FormatInt(res.Size, 10),
			})
			if err != nil {
				ctxLog.Error("ERROR REQUEST")

				return err
			}
			ctxLog.Debug("REQUEST")

			return nil
		}
	}
}
