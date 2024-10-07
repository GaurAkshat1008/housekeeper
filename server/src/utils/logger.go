package utils

import (
	"github.com/sirupsen/logrus"
)

// Create a global logger
var Logger *logrus.Logger

// Initialize the logger
func InitLogger() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
		ForceColors:     true,
	})
	Logger.SetLevel(logrus.InfoLevel)
}
