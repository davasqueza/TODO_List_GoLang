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
		defer func() {
			var totalTime = time.Now().Sub(startTime) / time.Millisecond
			loggerMiddleware.logger.Printf("Server response time: %dms\n", totalTime)
		}()
		next(w, r)
	}
}
