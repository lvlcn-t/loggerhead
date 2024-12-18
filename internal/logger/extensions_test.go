package logger

import (
	"context"
	"log/slog"
	"testing"

	"github.com/lvlcn-t/loggerhead/internal/logger/test"
)

type logFunc func(l Provider, msg string, args ...any)

func TestLogger_LevelExtensions(t *testing.T) {
	tests := []struct {
		name    string
		logFunc logFunc
		handler test.MockHandler
	}{
		{
			name: "debug level disabled",
			logFunc: func(l Provider, msg string, args ...any) {
				l.Debugf(msg, args...)
			},
			handler: test.MockHandler{
				EnabledFunc: func(ctx context.Context, level slog.Level) bool {
					return false
				},
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					t.Error("Handle should not be called")
					return nil
				},
			},
		},
		{
			name: "debug level enabled",
			logFunc: func(l Provider, msg string, args ...any) {
				l.Debugf(msg, args...)
			},
			handler: test.MockHandler{
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					level := LevelDebug
					if r.Level != slog.Level(level) {
						t.Errorf("Expected level to be [%s], got [%s]", level, r.Level)
					}
					if r.NumAttrs() != 0 {
						t.Errorf("Expected 0 attributes, got %d", r.NumAttrs())
					}
					return nil
				},
			},
		},
		{
			name: "info level disabled",
			logFunc: func(l Provider, msg string, args ...any) {
				l.Infof(msg, args...)
			},
			handler: test.MockHandler{
				EnabledFunc: func(ctx context.Context, level slog.Level) bool {
					return false
				},
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					t.Error("Handle should not be called")
					return nil
				},
			},
		},
		{
			name: "info level enabled",
			logFunc: func(l Provider, msg string, args ...any) {
				l.Infof(msg, args...)
			},
			handler: test.MockHandler{
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					level := LevelInfo
					if r.Level != slog.Level(level) {
						t.Errorf("Expected level to be [%s], got [%s]", level, r.Level)
					}
					if r.NumAttrs() != 0 {
						t.Errorf("Expected 0 attributes, got %d", r.NumAttrs())
					}
					return nil
				},
			},
		},
		{
			name: "warn level disabled",
			logFunc: func(l Provider, msg string, args ...any) {
				l.Warnf(msg, args...)
			},
			handler: test.MockHandler{
				EnabledFunc: func(ctx context.Context, level slog.Level) bool {
					return false
				},
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					t.Error("Handle should not be called")
					return nil
				},
			},
		},
		{
			name: "warn level enabled",
			logFunc: func(l Provider, msg string, args ...any) {
				l.Warnf(msg, args...)
			},
			handler: test.MockHandler{
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					level := LevelWarn
					if r.Level != slog.Level(level) {
						t.Errorf("Expected level to be [%s], got [%s]", level, r.Level)
					}
					if r.NumAttrs() != 0 {
						t.Errorf("Expected 0 attributes, got %d", r.NumAttrs())
					}
					return nil
				},
			},
		},
		{
			name: "error level disabled",
			logFunc: func(l Provider, msg string, args ...any) {
				l.Errorf(msg, args...)
			},
			handler: test.MockHandler{
				EnabledFunc: func(ctx context.Context, level slog.Level) bool {
					return false
				},
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					t.Error("Handle should not be called")
					return nil
				},
			},
		},
		{
			name: "error level enabled",
			logFunc: func(l Provider, msg string, args ...any) {
				l.Errorf(msg, args...)
			},
			handler: test.MockHandler{
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					level := LevelError
					if r.Level != slog.Level(level) {
						t.Errorf("Expected level to be [%s], got [%s]", level, r.Level)
					}
					if r.NumAttrs() != 0 {
						t.Errorf("Expected 0 attributes, got %d", r.NumAttrs())
					}
					return nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLogger(Options{Handler: tt.handler})
			tt.logFunc(l, "test")
		})
	}
}

func TestLogger_CustomLevels(t *testing.T) {
	tests := []struct {
		name      string
		attrs     []any
		logFunc   logFunc
		handler   test.MockHandler
		wantPanic bool
	}{
		{
			name:  "trace level",
			attrs: []any{"key", "value"},
			logFunc: func(l Provider, msg string, args ...any) {
				l.Trace(msg, args...)
			},
			handler: test.MockHandler{
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					return assertRecordLevel(t, &r, LevelTrace, true)
				},
			},
		},
		{
			name: "tracef level",
			logFunc: func(l Provider, msg string, args ...any) {
				l.Tracef(msg, args...)
			},
			handler: test.MockHandler{
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					return assertRecordLevel(t, &r, LevelTrace, false)
				},
			},
		},
		{
			name:  "trace context level",
			attrs: []any{"key", "value"},
			logFunc: func(l Provider, msg string, args ...any) {
				l.TraceContext(context.Background(), msg, args...)
			},
			handler: test.MockHandler{
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					return assertRecordLevel(t, &r, LevelTrace, true)
				},
			},
		},
		{
			name:  "notice level",
			attrs: []any{"key", "value"},
			logFunc: func(l Provider, msg string, args ...any) {
				l.Notice(msg, args...)
			},
			handler: test.MockHandler{
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					return assertRecordLevel(t, &r, LevelNotice, true)
				},
			},
		},
		{
			name: "noticef level",
			logFunc: func(l Provider, msg string, args ...any) {
				l.Noticef(msg, args...)
			},
			handler: test.MockHandler{
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					return assertRecordLevel(t, &r, LevelNotice, false)
				},
			},
		},
		{
			name:  "notice context level",
			attrs: []any{"key", "value"},
			logFunc: func(l Provider, msg string, args ...any) {
				l.NoticeContext(context.Background(), msg, args...)
			},
			handler: test.MockHandler{
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					return assertRecordLevel(t, &r, LevelNotice, true)
				},
			},
		},
		{
			name:  "panic level",
			attrs: []any{"key", "value"},
			logFunc: func(l Provider, msg string, args ...any) {
				l.Panic(msg, args...)
			},
			handler: test.MockHandler{
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					return assertRecordLevel(t, &r, LevelPanic, true)
				},
			},
			wantPanic: true,
		},
		{
			name: "panicf level",
			logFunc: func(l Provider, msg string, args ...any) {
				l.Panicf(msg, args...)
			},
			handler: test.MockHandler{
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					return assertRecordLevel(t, &r, LevelPanic, false)
				},
			},
			wantPanic: true,
		},
		{
			name:  "panic context level",
			attrs: []any{"key", "value"},
			logFunc: func(l Provider, msg string, args ...any) {
				l.PanicContext(context.Background(), msg, args...)
			},
			handler: test.MockHandler{
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					return assertRecordLevel(t, &r, LevelPanic, true)
				},
			},
			wantPanic: true,
		},
		{
			name:  "fatal level",
			attrs: []any{"key", "value"},
			logFunc: func(l Provider, msg string, args ...any) {
				l.Fatal(msg, args...)
			},
			handler: test.MockHandler{
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					return assertRecordLevel(t, &r, LevelFatal, true)
				},
			},
			wantPanic: true,
		},
		{
			name: "fatalf level",
			logFunc: func(l Provider, msg string, args ...any) {
				l.Fatalf(msg, args...)
			},
			handler: test.MockHandler{
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					return assertRecordLevel(t, &r, LevelFatal, false)
				},
			},
			wantPanic: true,
		},
		{
			name:  "fatal context level",
			attrs: []any{"key", "value"},
			logFunc: func(l Provider, msg string, args ...any) {
				l.FatalContext(context.Background(), msg, args...)
			},
			handler: test.MockHandler{
				HandleFunc: func(ctx context.Context, r slog.Record) error {
					return assertRecordLevel(t, &r, LevelFatal, true)
				},
			},
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exit = func(code int) {
				if code != 1 {
					t.Errorf("Expected exit code 1, got %d", code)
				}
				panic("os.Exit(1)")
			}

			l := NewLogger(Options{Handler: tt.handler})
			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Error("Expected panic")
					}
				}()
			}
			tt.logFunc(l, "test", tt.attrs...)
		})
	}
}

func assertRecordLevel(t *testing.T, r *slog.Record, level Level, wantAttrs bool) error {
	t.Helper()
	if r.Level != slog.Level(level) {
		t.Errorf("Expected level to be [%s], got [%s]", level, r.Level)
	}
	if !wantAttrs && r.NumAttrs() != 0 {
		t.Errorf("Expected 0 attributes, got %d", r.NumAttrs())
	}
	if wantAttrs && r.NumAttrs() == 0 {
		t.Errorf("Expected attributes, got 0")
	}
	return nil
}
