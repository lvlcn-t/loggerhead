package logger

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	clog "github.com/charmbracelet/log"
	otel "github.com/remychantenay/slog-otel"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name     string
		opts     []Options
		wantErr  bool
		levelEnv string
	}{
		{
			name:     "No handler with default log level",
			opts:     nil,
			wantErr:  false,
			levelEnv: "",
		},
		{
			name:     "No handler with DEBUG log level",
			opts:     nil,
			wantErr:  false,
			levelEnv: "DEBUG",
		},
		{
			name:     "No handler with NOTICE log level",
			opts:     nil,
			wantErr:  false,
			levelEnv: "NOTICE",
		},
		{
			name:     "No handler with ERROR log level",
			opts:     nil,
			wantErr:  false,
			levelEnv: "ERROR",
		},
		{
			name: "Custom handler provided",
			opts: []Options{
				{Handler: slog.NewJSONHandler(os.Stdout, nil)},
			},
			wantErr:  false,
			levelEnv: "",
		},
		{
			name: "Otel enabled",
			opts: []Options{
				{OpenTelemetry: true},
			},
			wantErr:  false,
			levelEnv: "",
		},
		{
			name: "Otel enabled with WARN log level",
			opts: []Options{
				{OpenTelemetry: true},
			},
			wantErr:  false,
			levelEnv: "WARN",
		},
		{
			name: "No handler with env ERROR log level and options WARN log level",
			opts: []Options{
				{Level: "WARN"},
			},
			wantErr:  false,
			levelEnv: "ERROR",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("LOG_LEVEL", tt.levelEnv)

			log := NewLogger(tt.opts...)

			if (log == nil) != tt.wantErr {
				t.Errorf("NewLogger() error = %v, expectedErr %v", log == nil, tt.wantErr)
			}

			if tt.levelEnv != "" {
				want := getLevel(tt.levelEnv)
				got := log.Enabled(context.Background(), slog.Level(want))
				if !got {
					t.Errorf("Expected log level: %v", want)
				}
			}

			if len(tt.opts) > 0 {
				if tt.opts[0].OpenTelemetry {
					if _, ok := log.Handler().(*otel.OtelHandler); !ok {
						t.Errorf("Want %T, got %T", &otel.OtelHandler{}, log.Handler())
					}
					return
				}
				if tt.opts[0].Handler != nil {
					if !reflect.DeepEqual(log.Handler(), tt.opts[0].Handler) {
						t.Errorf("Want %T, got %T", tt.opts[0].Handler, log.Handler())
					}
				}
			}
		})
	}
}

func TestNewContextWithLogger(t *testing.T) {
	tests := []struct {
		name      string
		parentCtx context.Context
	}{
		{
			name:      "With Background context",
			parentCtx: context.Background(),
		},
		{
			name:      "With already set logger in context",
			parentCtx: context.WithValue(context.Background(), logger{}, NewLogger()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := NewContextWithLogger(tt.parentCtx)
			defer cancel()

			log := ctx.Value(ctxKey{})
			if _, ok := log.(Logger); !ok {
				t.Errorf("Context does not contain Logger, got %T", log)
			}
			if ctx == tt.parentCtx {
				t.Errorf("NewContextWithLogger returned the same context as the parent")
			}
		})
	}
}

func TestFromContext(t *testing.T) {
	tests := []struct {
		name string
		ctx  context.Context
		want Logger
	}{
		{
			name: "Context with logger",
			ctx:  IntoContext(context.Background(), NewLogger(Options{Handler: slog.NewJSONHandler(os.Stdout, nil)})),
			want: NewLogger(Options{Handler: slog.NewJSONHandler(os.Stdout, nil)}),
		},
		{
			name: "Context without logger",
			ctx:  context.Background(),
			want: NewLogger(),
		},
		{
			name: "Nil context",
			ctx:  nil,
			want: NewLogger(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromContext(tt.ctx)
			_, ok := got.(*logger)
			if !ok {
				t.Errorf("FromContext() = %T, want %T", got, tt.want)
			}
		})
	}
}

func TestFromSlog(t *testing.T) {
	tests := []struct {
		name string
		l    *slog.Logger
		want Logger
	}{
		{
			name: "Slog logger",
			l:    slog.New(slog.NewJSONHandler(os.Stdout, nil)),
			want: NewLogger(Options{Handler: slog.NewJSONHandler(os.Stdout, nil)}),
		},
		{
			name: "Nil slog logger",
			l:    nil,
			want: NewLogger(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromSlog(tt.l)
			if _, ok := got.(*logger); !ok {
				t.Errorf("FromSlog() = %T, want %T", got, tt.want)
			}

			if reflect.TypeOf(got.Handler()) != reflect.TypeOf(tt.want.Handler()) {
				t.Errorf("FromSlog().Handler() = %v, want %v", got.Handler(), tt.want.Handler())
			}
		})
	}
}

func TestLogger_ToSlog(t *testing.T) {
	tests := []struct {
		name string
		l    Logger
	}{
		{
			name: "Logger",
			l:    NewLogger(Options{Handler: slog.NewJSONHandler(os.Stdout, nil)}),
		},
		{
			name: "Nil logger",
			l:    &logger{nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.l != nil {
				got := tt.l.ToSlog()
				if got == nil {
					t.Errorf("ToSlog() = %v, want %v", got, tt.l)
				}
			}
		})
	}
}

func TestMiddleware(t *testing.T) {
	tests := []struct {
		name        string
		parentCtx   context.Context
		expectInCtx bool
	}{
		{
			name:        "With logger in parent context",
			parentCtx:   IntoContext(context.Background(), NewLogger()),
			expectInCtx: true,
		},
		{
			name:        "Without logger in parent context",
			parentCtx:   context.Background(),
			expectInCtx: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			middleware := Middleware(tt.parentCtx)
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, ok := r.Context().Value(ctxKey{}).(Logger)
				if tt.expectInCtx != ok {
					t.Errorf("Middleware() did not inject logger correctly, got %v, want %v", ok, tt.expectInCtx)
				}
			})

			req := httptest.NewRequest("GET", "/", http.NoBody)
			w := httptest.NewRecorder()

			middleware(handler).ServeHTTP(w, req)
		})
	}
}

func TestNewBaseHandler(t *testing.T) {
	tests := []struct {
		name      string
		format    string
		level     string
		wantLevel int
	}{
		{
			name:      "Default handler",
			format:    "",
			level:     "",
			wantLevel: int(slog.LevelInfo),
		},
		{
			name:      "Text handler with custom log level",
			format:    "TEXT",
			level:     "DEBUG",
			wantLevel: int(clog.DebugLevel),
		},
		{
			name:      "JSON handler with custom log level",
			format:    "JSON",
			level:     "WARN",
			wantLevel: int(slog.LevelWarn),
		},
		{
			name:      "Invalid log level",
			format:    "TEXT",
			level:     "UNKNOWN",
			wantLevel: int(clog.InfoLevel),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("LOG_FORMAT", tt.format)
			t.Setenv("LOG_LEVEL", tt.level)
			opts := newDefaultOptions()
			handler := newBaseHandler(opts)

			if tt.format == "TEXT" {
				if _, ok := handler.(*clog.Logger); !ok {
					t.Errorf("Expected handler to be of type *log.Logger")
				}
			} else {
				if _, ok := handler.(*slog.JSONHandler); !ok {
					t.Errorf("Expected handler to be of type *slog.JSONHandler")
				}
			}

			ok := handler.Enabled(context.Background(), slog.Level(tt.wantLevel))
			if !ok {
				t.Errorf("Expected log level: %v", tt.wantLevel)
			}
		})
	}
}
