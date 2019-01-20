package logging

import "github.com/sirupsen/logrus"

// console color palette
const (
	colorRed    int = 31
	colorGreen  int = 32
	colorYellow int = 33
	colorBlue   int = 36
	colorGray   int = 37
)

func getColorByLogLevel(level logrus.Level) int {
	switch level {
	case logrus.DebugLevel, logrus.InfoLevel:
		return colorBlue
	case logrus.WarnLevel:
		return colorYellow
	case logrus.PanicLevel, logrus.ErrorLevel, logrus.FatalLevel:
		return colorRed
	case logrus.TraceLevel:
		return colorGreen
	default:
		return colorGray
	}
}
