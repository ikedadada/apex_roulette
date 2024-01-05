package logger

import (
	"apex_roulette/application_service/service"
	"context"
	"io"
	"log/slog"

	"cloud.google.com/go/logging"
)

type logger struct {
	logger *slog.Logger
}

func NewLogger(writer io.Writer) service.Logger {
	replacer := newLogKeyReplacer()
	return &logger{
		logger: slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{ReplaceAttr: replacer})),
	}
}

func (l *logger) StructLog(level service.LogLevel, message string) {
	serverity := l.convertToServerity(level)

	ctx := context.Background()
	l.logger.Log(ctx, slog.Level(serverity), message)
}

func (l *logger) convertToServerity(level service.LogLevel) logging.Severity {
	switch level {
	case service.LogLevelInfo:
		return logging.Info
	case service.LogLevelError:
		return logging.Error
	case service.LogLevelCritical:
		return logging.Critical
	default:
		return logging.Warning
	}
}

func newLogKeyReplacer() func(groups []string, a slog.Attr) slog.Attr {
	return func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.LevelKey {
			a.Key = "severity"
			v := logging.Severity(a.Value.Any().(slog.Level)).String()
			a.Value = slog.StringValue(v)
		}
		if a.Key == slog.MessageKey {
			a.Key = "message"
		}
		return a
	}
}
