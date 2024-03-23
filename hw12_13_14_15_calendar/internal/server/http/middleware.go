package internalhttp

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler, l *zap.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		dateTime := time.DateTime
		rw := NewLoggingResponseWriter(w)
		next.ServeHTTP(rw, r)

		l.Info(
			"request stats",
			zap.Duration("latency", time.Since(start)),
			zap.String("client_addr", r.RemoteAddr),
			zap.String("date_time", dateTime),
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.Int("response_code", rw.statusCode),
			zap.String("http_version", r.Proto),
			zap.String("user_agent", r.UserAgent()),
		)
	})
}
