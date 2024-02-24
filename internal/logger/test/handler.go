package test

import (
	"context"
	"log/slog"
)

type MockHandler struct {
	EnabledFunc   func(ctx context.Context, level slog.Level) bool
	HandleFunc    func(ctx context.Context, r slog.Record) error
	WithAttrsFunc func(attrs []slog.Attr) slog.Handler
	WithGroupFunc func(name string) slog.Handler
}

func (m MockHandler) Enabled(ctx context.Context, level slog.Level) bool {
	if m.EnabledFunc != nil {
		return m.EnabledFunc(ctx, level)
	}
	return true
}

func (m MockHandler) Handle(ctx context.Context, r slog.Record) error {
	if m.HandleFunc != nil {
		return m.HandleFunc(ctx, r)
	}
	return nil
}

func (m MockHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if m.WithAttrsFunc != nil {
		return m.WithAttrsFunc(attrs)
	}
	return m
}

func (m MockHandler) WithGroup(name string) slog.Handler {
	if m.WithGroupFunc != nil {
		return m.WithGroupFunc(name)
	}
	return m
}
