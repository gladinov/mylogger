package slogtrace

import (
	"context"
	"log/slog"

	"github.com/gladinov/mylogger/trace"
)

type TraceHandler struct {
	next slog.Handler
}

func New(next slog.Handler) *TraceHandler {
	return &TraceHandler{next: next}
}

func (h *TraceHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.next.Enabled(ctx, level)
}

func (h *TraceHandler) Handle(ctx context.Context, r slog.Record) error {
	if traceID, ok := trace.TraceIDFromContext(ctx); ok {
		r.AddAttrs(slog.String("trace_id", traceID))
	}
	return h.next.Handle(ctx, r)
}

func (h *TraceHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &TraceHandler{next: h.next.WithAttrs(attrs)}
}

func (h *TraceHandler) WithGroup(name string) slog.Handler {
	return &TraceHandler{next: h.next.WithGroup(name)}
}
