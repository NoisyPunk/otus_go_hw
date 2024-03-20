package internalhttp

import (
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

//nolint:unused
func loggingMiddleware(next http.Handler, l *zap.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		l.Info("request stats", zap.String("latency", strconv.Itoa(int(time.Since(start).Nanoseconds()))))
	})
}
