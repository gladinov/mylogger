package mylogger

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"testing"
)

func TestTraceHandler_AddsTraceID(t *testing.T) {
	var buf bytes.Buffer

	base := slog.NewJSONHandler(&buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	handler := New(base)
	logger := slog.New(handler)

	ctx := WithTraceID(context.Background(), "trace-123")

	logger.InfoContext(ctx, "hello")

	var record map[string]any
	if err := json.Unmarshal(buf.Bytes(), &record); err != nil {
		t.Fatalf("failed to parse log json: %v", err)
	}

	if got := record["trace_id"]; got != "trace-123" {
		t.Fatalf("expected trace_id=trace-123, got=%v", got)
	}
}

func TestTraceHandler_NoTraceID(t *testing.T) {
	var buf bytes.Buffer

	base := slog.NewJSONHandler(&buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	handler := New(base)
	logger := slog.New(handler)

	logger.Info("hello")

	var record map[string]any
	if err := json.Unmarshal(buf.Bytes(), &record); err != nil {
		t.Fatalf("failed to parse log json: %v", err)
	}

	if _, ok := record["trace_id"]; ok {
		t.Fatal("trace_id should not be present")
	}
}

func TestTraceHandler_WithAttrs(t *testing.T) {
	var buf bytes.Buffer

	base := slog.NewJSONHandler(&buf, nil)
	handler := New(base)

	logger := slog.New(handler).With(
		slog.String("service", "test"),
	)

	ctx := WithTraceID(context.Background(), "trace-456")
	logger.InfoContext(ctx, "hello")

	var record map[string]any
	if err := json.Unmarshal(buf.Bytes(), &record); err != nil {
		t.Fatalf("failed to parse log json: %v", err)
	}

	if record["trace_id"] != "trace-456" {
		t.Fatalf("trace_id lost after WithAttrs")
	}
}
