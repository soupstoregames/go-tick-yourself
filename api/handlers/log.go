package handlers

import (
	"github.com/soupstoregames/go-tick-yourself/logging"
	"github.com/soupstoregames/go-tick-yourself/metrics"
	"net/http"
)

func LoggingAndMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logging.LogHTTPRequest(r)
		metrics.CountHTTPRequest(r)

		res := newResponseCatcher(w, r)
		next.ServeHTTP(res, r)

		metrics.CountHTTPResponse(&res.response)
		logging.LogHTTPResponse(&res.response)
	})
}