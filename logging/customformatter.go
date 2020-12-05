package logging

import "github.com/sirupsen/logrus"

type customFormatter struct {
	logrus.Formatter
}

func (u customFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.UTC()
	e.Data = e.WithFields(standardFields).Data
	return u.Formatter.Format(e)
}