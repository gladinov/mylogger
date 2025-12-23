package sl

import (
	"log/slog"
	"os"

	"github.com/gladinov/mylogger/slogtrace"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func NewLogger(env string) *slog.Logger {
	var (
		level   slog.Level
		handler slog.Handler
	)

	switch env {
	case envLocal:
		level = slog.LevelDebug
		base := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
		handler = slogtrace.New(base)

	case envDev:
		level = slog.LevelDebug
		base := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
		handler = slogtrace.New(base)

	case envProd:
		level = slog.LevelInfo
		base := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
		handler = slogtrace.New(base)

	default:
		level = slog.LevelInfo
		base := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
		handler = slogtrace.New(base)

	}
	return slog.New(handler)
}
