package logger

import (
	"testing"
)

func TestGetLevel(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Level
	}{
		{"Empty string", "", LevelInfo},
		{"Trace level", "TRACE", LevelTrace},
		{"Debug level", "DEBUG", LevelDebug},
		{"Info level", "INFO", LevelInfo},
		{"Notice level", "NOTICE", LevelNotice},
		{"Warn level", "WARN", LevelWarn},
		{"Warning level", "WARNING", LevelWarn},
		{"Error level", "ERROR", LevelError},
		{"Invalid level", "UNKNOWN", LevelInfo},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newLevel(tt.input)
			if got != tt.want {
				t.Errorf("getLevel(%s) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
