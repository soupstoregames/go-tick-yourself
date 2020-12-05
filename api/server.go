package api

import (
	"github.com/soupstoregames/go-tick-yourself/api/handlers"
	"net/http"
	"time"
)

func NewHTTPServer(addr string, routes http.Handler) *http.Server {
	r := handlers.LoggingAndMetrics(routes)

	server := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return server
}