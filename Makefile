.PHONY: examples
examples: custom-handler-example extension-example otel-example simple-example webserver-example

.PHONY: custom-handler-example
custom-handler-example:
	@go run examples/custom-handler/main.go || true

.PHONY: extension-example
extension-example:
	@go run examples/extension/main.go

.PHONY: otel-example
otel-example:
	@go run examples/open-telemetry/main.go

.PHONY: simple-example
simple-example:
	@go run examples/simple/main.go || true

.PHONY: webserver-example
webserver-example:
	@echo "Open http://localhost:8080 in your browser to see the example.\nPress Ctrl+C to stop the server."
	@go run examples/webserver/main.go
