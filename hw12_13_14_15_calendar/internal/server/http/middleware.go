package internalhttp

import (
	"net/http"
)

//nolint:unused
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _, _ = next, w, r
		// TODO
	})
}
