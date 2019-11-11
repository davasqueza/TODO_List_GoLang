package middlewares

import (
	"log"
	"net/http"
	"time"
)

type LoggerMiddleware struct {
	logger *log.Logger
}

func NewLogger(logger *log.Logger) *LoggerMiddleware {
	return &LoggerMiddleware{
		logger: logger,
	}
}

func (loggerMiddleware *LoggerMiddleware) ServerTime(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer loggerMiddleware.logger.Printf("Server response time: %.2fs\n", time.Now().Sub(startTime).Seconds())
		next(w, r)
	}
}
