lint:
	pre-commit run --hook-stage pre-commit -a
	pre-commit run --hook-stage pre-push -a

examples: custom-handler-example extension-example otel-example simple-example webserver-example

custom-handler-example:
	@go run examples/custom-handler/main.go || true

extension-example:
	@go run examples/extension/main.go

otel-example:
	@go run examples/open-telemetry/main.go

simple-example:
	@go run examples/simple/main.go || true

webserver-example:
	@echo "Open http://localhost:8080 in your browser to see the example.\nPress Ctrl+C to stop the server."
	@go run examples/webserver/main.go

.PHONY: lint examples custom-handler-example extension-example otel-example simple-example webserver-example