package middleware

import (
	"context"
	"log/slog"
)

type ContextHandler struct {
	slog.Handler
}

func (h *ContextHandler) Handle(ctx context.Context, record slog.Record) error {
	requestId := RequestIDFromContext(ctx)
	if requestId != "" {
		record.AddAttrs(slog.String("request_id", requestId))
	}
	return h.Handler.Handle(ctx, record)
}
