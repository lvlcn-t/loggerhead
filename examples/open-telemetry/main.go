package main

import (
	"context"
	"os"

	"github.com/lvlcn-t/loggerhead/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	os.Setenv("LOG_FORMAT", "text")
	log := logger.NewNamedLogger("otel-example", logger.Opts{
		OpenTelemetry: true,
	})

	m1, err := baggage.NewMember("user_id", "123")
	if err != nil {
		log.Error("Failed to create baggage member", "error", err)
	}
	m2, err := baggage.NewMember("user_name", "jane.doe")
	if err != nil {
		log.Error("Failed to create baggage member", "error", err)
	}
	bag, err := baggage.New(m1, m2)
	if err != nil {
		log.Error("Failed to create baggage", "error", err)
	}
	ctx := baggage.ContextWithBaggage(context.Background(), bag)

	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Error("Failed to create stdout exporter", "error", err)
	}
	tp := trace.NewTracerProvider(trace.WithBatcher(exporter))
	otel.SetTracerProvider(tp)

	ctx, span := otel.Tracer("example").Start(ctx, "operation")
	defer span.End()

	log.InfoContext(ctx, "This is a log message with baggage and trace context.")
}
