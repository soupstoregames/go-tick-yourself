package logging

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var (
	standardFields logrus.Fields
	logger *logrus.Logger
)

func init() {
	logger = logrus.New()
	logger.Formatter = customFormatter{&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	}}
	logger.Out = os.Stdout
}

func SetStandardFields(service, version string) {
	hostname, _ := os.Hostname()
	standardFields = logrus.Fields{
		"service": service,
		"version": version,
		"hostname": hostname,
	}
}

func Info(msg interface{}) {
	logger.Info(msg)
}

func Warn(msg interface{}) {
	logger.Warn(msg)
}

func Error(msg interface{}) {
	logger.Error(msg)
}

func Fatal(msg interface{}) {
	logger.Fatal(msg)
}

func WithErr(value error) *logrus.Entry {
	return logger.WithField("error", value)
}