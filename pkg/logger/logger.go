package logger

import (
	"bytes"
	"fmt"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
)

// Formatter implements logrus.Formatter interface without time
type formatter struct {
	prefix string
}

var logger = logrus.New()

var (
	Debug  = logger.Debug
	Debugf = logger.Debugf
	Info   = logger.Info
	Infof  = logger.Infof
	Warn   = logger.Warn
	Warnf  = logger.Warnf
	Error  = logger.Error
	Errorf = logger.Errorf
	Fatal  = logger.Fatal
	Fatalf = logger.Fatalf
)

func init() {
	logger.Level = logrus.InfoLevel
	logger.Formatter = &formatter{}
	logger.SetReportCaller(true)
}

func GetLogger() *logrus.Logger {
	return logger
}

func SetLogLevel(level logrus.Level) {
	logger.Level = level
}

// Format builds log message WITHOUT time
func (f *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	var sb bytes.Buffer
	sb.WriteString(strings.ToUpper(entry.Level.String()))
	sb.WriteString(" ")
	sb.WriteString(f.prefix)
	sb.WriteString(fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line))
	sb.WriteString(" ")
	sb.WriteString(entry.Message)
	sb.WriteString("\n")
	return sb.Bytes(), nil
}
