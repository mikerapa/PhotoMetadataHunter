package logging

import (
	"bytes"
	"fmt"
	logrus "github.com/sirupsen/logrus"
	"os"
	"strings"
)

const dateTimeFormat = "2 Jan 2006 15:04:05"

type myFormatter struct {
	logrus.TextFormatter
}

func (f *myFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	formattedLogEntry := fmt.Sprintln(entry.Time.Format(dateTimeFormat), strings.ToUpper(entry.Level.String()), "-", entry.Message)
	byteBuffer := &bytes.Buffer{}
	byteBuffer.WriteString(fmt.Sprintf("\x1b[%dm", getColorByLogLevel(entry.Level)))
	byteBuffer.WriteString(formattedLogEntry)
	byteBuffer.WriteString("\x1b[0m")
	return byteBuffer.Bytes(), nil

}

func getFormatter() logrus.Formatter {
	formatter := new(myFormatter)
	formatter.ForceColors = true
	return formatter
}

func GetLogger(logLevel logrus.Level) logrus.Logger {
	logger := logrus.Logger{
		Out:       os.Stderr,
		Level:     logLevel,
		Formatter: getFormatter(),
	}

	return logger
}

// (error, warning, info, debug, trace)
func GetLogLevelFromString(logLevelString string) logrus.Level {
	switch strings.ToLower(logLevelString) {
	case "info":
		return logrus.InfoLevel
	case "trace":
		return logrus.TraceLevel
	default:
		return logrus.ErrorLevel
	}
}
