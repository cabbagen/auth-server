package provider

import (
	"os"
	logger "github.com/sirupsen/logrus"
)

func init() {
	logger.SetFormatter(&logger.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logger.InfoLevel)
}

func MakeupCommonLogger(fields map[string]string) *logger.Entry {
	loggerFields := logger.Fields {
		"type": "common request",
	}

	for key, value := range fields {
		loggerFields[key] = value
	}
	return logger.WithFields(loggerFields)
}

func MakeupApplicationLogger(fields map[string]string) *logger.Entry {
	loggerFields := logger.Fields {
		"type": "application request",
	}

	for key, value := range fields {
		loggerFields[key] = value
	}
	return logger.WithFields(loggerFields)
}
