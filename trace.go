package mylogger

import "context"

type ctxKey struct{}

var TraceIDKey = ctxKey{}

func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceIDKey, traceID)
}

func TraceIDFromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(TraceIDKey).(string)
	return id, ok
}
