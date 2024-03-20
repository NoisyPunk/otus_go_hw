package internalhttp

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler, l *zap.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		dateTime := time.DateTime
		next.ServeHTTP(w, r)

		l.Info(
			"request stats",
			zap.Duration("latency", time.Since(start)),
			zap.String("client_addr", r.RemoteAddr),
			zap.String("date_time", dateTime),
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("http_version", r.Proto),
			zap.String("user_agent", r.UserAgent()),
		)
	})
}
