package logging

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func LogHTTPRequest(req *http.Request) {
	fields := logrus.Fields{
		"forwarded_for":  req.Header.Get("x-forwarded-for"),
		"protocol":       req.Proto,
		"remote_address": req.RemoteAddr,
		"url":            req.RequestURI,
		"method":         req.Method,
		"content_length": req.ContentLength,
	}

	logger.WithFields(fields).Info("HTTP request")
}

func LogHTTPResponse(res *http.Response) {
	fields := logrus.Fields{
		"forwarded_for":  res.Request.Header.Get("x-forwarded-for"),
		"protocol":       res.Proto,
		"remote_address": res.Request.RemoteAddr,
		"url":            res.Request.RequestURI,
		"method":         res.Request.Method,
		"content_length": res.ContentLength,
		"status":         res.StatusCode,
	}

	logger.WithFields(fields).Info("HTTP response")
}