package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (sr *statusRecorder) WriteHeader(status int) {
	sr.status = status
	sr.ResponseWriter.WriteHeader(status)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		sr := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(sr, r)

		duration := time.Since(startTime)

		slog.InfoContext(
			r.Context(),
			"http request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", sr.status,
			"duration_ms", duration.Milliseconds(),
			"request_id", RequestIDFromContext(r.Context()),
		)
	})
}
