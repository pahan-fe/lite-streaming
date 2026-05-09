package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const requestIDKey contextKey = "request_id"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		requestID := uuid.New().String()
		w.Header().Set("X-Request-ID", requestID)

		ctx := context.WithValue(req.Context(), requestIDKey, requestID)

		next.ServeHTTP(w, req.WithContext(ctx))
	})
}

func RequestIDFromContext(ctx context.Context) string {
	if requestID, ok := ctx.Value(requestIDKey).(string); ok {
		return requestID
	}

	return ""
}
