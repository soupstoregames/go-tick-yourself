package metrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
)

var (
	httpRequests = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests",
		Help: "A count of all HTTP requests",
	}, []string{"method", "endpoint"})

	httpResponses = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_responses",
		Help: "A count of all HTTP responses",
	}, []string{"method", "endpoint", "code"})
)

func CountHTTPRequest(req *http.Request) {
	httpRequests.With(prometheus.Labels{
		"method": req.Method,
		"endpoint": req.RequestURI,
	}).Inc()
}

func CountHTTPResponse(res *http.Response) {
	httpResponses.With(prometheus.Labels{
		"method": res.Request.Method,
		"endpoint": res.Request.RequestURI,
		"code": fmt.Sprintf("%d", res.StatusCode),
	}).Inc()
}