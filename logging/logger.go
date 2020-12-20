package logging

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	standardFields logrus.Fields
	logger         *logrus.Logger
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
		"service":  service,
		"version":  version,
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

func WithError(value error) *logrus.Entry {
	return logger.WithField("error", value)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(fields)
}
