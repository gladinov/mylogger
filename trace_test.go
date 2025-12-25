package mylogger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWithTraceID(t *testing.T) {
	cases := []struct {
		name    string
		ctx     context.Context
		traceID string
	}{
		{
			name:    "success",
			ctx:     context.Background(),
			traceID: "alkjgljdsglk",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx2 := WithTraceID(tc.ctx, tc.traceID)
			got, _ := ctx2.Value(TraceIDKey).(string)

			require.Equal(t, tc.traceID, got)
		})
	}
}

func TestTraceIDFromContext(t *testing.T) {
	cases := []struct {
		name    string
		ctx     context.Context
		traceID string
		ok      bool
	}{
		{
			name:    "success",
			ctx:     context.Background(),
			traceID: "alkjgljdsglk",
			ok:      true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx2 := WithTraceID(tc.ctx, tc.traceID)
			got, gotOk := TraceIDFromContext(ctx2)

			require.Equal(t, tc.traceID, got)
			require.Equal(t, tc.ok, gotOk)
		})
	}
}
