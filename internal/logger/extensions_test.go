package logger

import (
	"log/slog"
	"testing"
)

func TestGetLevel(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect Level
	}{
		{"Empty string", "", slog.LevelInfo},
		{"Debug level", "DEBUG", slog.LevelDebug},
		{"Info level", "INFO", slog.LevelInfo},
		{"Warn level", "WARN", slog.LevelWarn},
		{"Warning level", "WARNING", slog.LevelWarn},
		{"Error level", "ERROR", slog.LevelError},
		{"Invalid level", "UNKNOWN", slog.LevelInfo},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getLevel(tt.input)
			if got != tt.expect {
				t.Errorf("getLevel(%s) = %v, want %v", tt.input, got, tt.expect)
			}
		})
	}
}
