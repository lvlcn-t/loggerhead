# Loggerhead - Logging Library<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- markdownlint-disable MD033 -->
<p align="center">
    <a href="https://pkg.go.dev/github.com/lvlcn-t/loggerhead"><img src="https://pkg.go.dev/badge/github.com/lvlcn-t/loggerhead.svg" alt="Go Reference"></a>
    <a href="/../../commits/" title="Last Commit"><img src="https://img.shields.io/github/last-commit/lvlcn-t/loggerhead?style=flat" alt="Last Commit"></a>
    <a href="/../../issues" title="Open Issues"><img src="https://img.shields.io/github/issues/lvlcn-t/loggerhead?style=flat" alt="Open Issues"></a>
    <a href="/../../pulls" title="Open Pull Requests"><img src="https://img.shields.io/github/issues-pr/lvlcn-t/loggerhead?style=flat" alt="Open Pull Requests"></a>
</p>
<!-- markdownlint-enable MD033 -->

<!-- code_chunk_output -->

- [Loggerhead - Logging Library](#loggerhead---logging-library)
  - [Overview](#overview)
  - [Features](#features)
  - [Getting Started](#getting-started)
    - [Basic Usage](#basic-usage)
  - [Documentation](#documentation)
    - [Core Logging Functions](#core-logging-functions)
      - [NewLogger](#newlogger)
      - [NewNamedLogger](#newnamedlogger)
      - [Formatted Logging Methods](#formatted-logging-methods)
    - [Contextual Logging](#contextual-logging)
      - [NewContextWithLogger](#newcontextwithlogger)
      - [FromContext](#fromcontext)
      - [IntoContext](#intocontext)
    - [Middleware Integration](#middleware-integration)
    - [Configuration via Environment Variables](#configuration-via-environment-variables)
    - [Extending Loggerhead](#extending-loggerhead)
      - [Custom Handlers](#custom-handlers)
      - [Custom Log Formats](#custom-log-formats)
      - [Integration with Logging Backends](#integration-with-logging-backends)
      - [OpenTelemetry Integration](#opentelemetry-integration)
        - [Enabling OpenTelemetry](#enabling-opentelemetry)
        - [Usage with Context](#usage-with-context)
        - [Advanced Scenarios](#advanced-scenarios)
    - [Usage Scenarios](#usage-scenarios)
  - [Contributing](#contributing)
  - [License](#license)

<!-- /code_chunk_output -->

## Overview

Loggerhead is a versatile and extendable logging library for Go, designed to enhance the capabilities of standard Go logging with additional features and utilities. Building on top of the `slog.Logger`, it provides a robust solution for application logging, including extended log levels, formatted logging methods, and context-aware logging capabilities.

## Features

- **Extended Log Levels**: Includes standard levels like Debug, Info, Warn, and Error, along with custom levels for finer control.
- **Formatted Logging**: Methods like `Debugf`, `Infof`, `Warnf`, and `Errorf` for formatted output.
- **Context-Aware Logging**: Enhanced logging functionalities that leverage Go's context for richer, more informative log entries.
- **Thread-Safe Operations**: Ensures safe use in concurrent environments.
- **Customizable Handlers**: Easily integrate with different logging backends and formats.

## Getting Started

To start using Loggerhead, install the package using:

```bash
go get -u github.com/lvlcn-t/loggerhead
```

### Basic Usage

To start using Loggerhead, you can create a logger instance with default settings or customize it using the `logger.Opts` struct. Loggerhead provides functions for both general and named loggers, allowing for easy differentiation between log sources.

```go
package main

import (
  "github.com/lvlcn-t/loggerhead/logger"
)

func main() {
  // Creating a default logger
  log := logger.NewLogger()

  // Creating a named logger with custom options
  opts := logger.Opts{Level: "INFO", Format: "JSON"}
  log := logger.NewNamedLogger("myServiceLogger", opts)

  // Logging a message
  log.Info("Hello, world!")
}
```

## Documentation

Loggerhead provides a comprehensive set of features for advanced logging in Go applications. Here's an overview of its primary functionalities and how to use them effectively:

### Core Logging Functions

#### NewLogger

This function initializes a new logger with default settings or a [custom handler](#custom-handlers) if provided. It's the starting point for using Loggerhead in your application.

```go
log := logger.NewLogger()
```

#### NewNamedLogger

Creates a new logger with a specified name and a [custom handlers](#custom-handlers) if provided. This is useful for identifying logs originating from different parts of your application.

```go
log := logger.NewNamedLogger("MyLogger")
```

#### Formatted Logging Methods

These methods allow you to log messages with a specific format, similar to `Printf` functions. They are handy for inserting variable content into your logs.

```go
log.Infof("User %s has logged in", username)
```

### Contextual Logging

#### NewContextWithLogger

Generates a new context that carries a logger. This is particularly useful in applications where you need to pass the logger through context.

```go
ctx, cancel := logger.NewContextWithLogger(parentContext)
defer cancel()
```

#### FromContext

Retrieves the logger embedded in the provided context. This is essential for getting the logger in different parts of your application where the context is passed.

```go
log := logger.FromContext(ctx)
```

#### IntoContext

Embeds a logger into a given context. This can be used to ensure that logging remains consistent across different parts of your application.

```go
newCtx := logger.IntoContext(ctx, log)
```

### Middleware Integration

Loggerhead offers a middleware function that can be integrated into HTTP server frameworks. This middleware injects the logger into every HTTP request's context, making it easier to log request-specific information.

```go
var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {}
http.Handle("/", logger.Middleware(context.Background())(handler))
```

### Configuration via Environment Variables

You can configure the logging behavior of the loggerhead library using environment variables. This allows you to adjust configurations without modifying your code.

The following environment variables are supported:

- `LOG_LEVEL`: Adjusts the minimum log level. This allows you to control the verbosity of the logs.
  Available options are the standard log levels. For example: `DEBUG`, `INFO`, `WARN`, `ERROR`.
- `LOG_FORMAT`: Sets the log format. This allows you to customize the format of the log messages.
  Available options are `TEXT` and `JSON`.

### Extending Loggerhead

Loggerhead is designed to be highly extendable, offering several ways for developers to customize and enhance its functionality:

#### Custom Handlers

Developers have the flexibility to use their own `slog.Handler` implementations with Loggerhead. This allows for complete control over how logs are processed, formatted, and outputted. Whether integrating with third-party logging services, applying unique log formatting, or implementing custom log filtering logic, you can pass your custom handler to `NewLogger` or `NewNamedLogger` functions to replace the default behavior.

```go
  // Example of using a custom handler
  log := logger.NewLogger(opts.Handler{
    Handler: myCustomHandler,
  })
```

This feature is especially useful for applications with specific logging requirements not covered by the default handlers. By providing your own implementation, you can tailor the logging behavior to fit the needs of your application precisely.

#### Custom Log Formats

While Loggerhead supports text and JSON formats out of the box, through custom `slog.Handler` implementations, developers can define entirely custom formats. This is ideal for adhering to organizational logging standards or enhancing log readability.

#### Integration with Logging Backends

Custom handlers also enable integration with various logging backends and services. Whether you're sending logs to a file, a console, a database, or a cloud-based logging platform, you can encapsulate this logic within your handler and use it seamlessly with Loggerhead.

#### OpenTelemetry Integration

Loggerhead supports integration with [OpenTelemetry](https://opentelemetry.io/docs/), allowing for enriched logging with trace and metrics data. This feature is invaluable for applications that require observability and tracing capabilities. By enabling OpenTelemetry support, you can ensure that your logs include relevant tracing information, such as trace IDs and span IDs, linking log entries to specific operations in your application.

##### Enabling OpenTelemetry

To enable OpenTelemetry integration, set the `OpenTelemetry` flag to `true` in the `Opts` struct when creating a logger. This will automatically enrich your logs with OpenTelemetry data.

```go
log := logger.NewLogger(logger.Opts{
  OpenTelemetry: true,
})
```

##### Usage with Context

When logging within a context that includes OpenTelemetry trace data, Loggerhead automatically includes this information in the logs. This integration allows for seamless correlation between logs and traces, making it easier to debug and monitor distributed applications.

```go
// Assuming the context has been configured with OpenTelemetry
log.InfoContext(ctx, "This is a log message with trace context.")
```

##### Advanced Scenarios

Loggerhead's OpenTelemetry integration is designed to be flexible, supporting advanced use cases such as:

- **Custom Trace Attributes**: You can add custom attributes to your traces that are automatically included in your logs, providing more detailed and contextual information for each log entry.
- **Baggage**: Utilize OpenTelemetry's baggage feature to pass additional metadata through the context, which can then be automatically included in logs for comprehensive tracing and debugging.

```go
// Example of using baggage to add user information to logs
m1, _ := baggage.NewMember("user_id", "12345")
m2, _ := baggage.NewMember("role", "admin")
bag, _ := baggage.New(m1, m2)
ctx := baggage.ContextWithBaggage(context.Background(), bag)

log.InfoContext(ctx, "User action logged with baggage")
```

### Usage Scenarios

- **Application Logging**: Use Loggerhead for general application logging, capturing information ranging from debug messages to critical errors.
- **Request Tracking**: In web applications, Loggerhead can be used to log and track HTTP requests, providing valuable insights for debugging and monitoring.
- **Contextual Information**: With context-aware logging, Loggerhead is ideal for applications that require detailed logs with contextual information, especially useful in microservices and distributed systems.
- **Enhanced Observability with OpenTelemetry**: When using OpenTelemetry, Loggerhead enriches logs with trace and metrics data, providing comprehensive observability for distributed applications. This includes linking logs to specific operations and tracing information, making it easier to debug and monitor complex systems.

For further examples and detailed usage, including how to implement and integrate custom `slog.Handler` instances or use OpenTelemetry with Loggerhead, please refer to the [examples](./examples) directory in our repository.

## Contributing

Contributions are welcome! Please refer to the [CONTRIBUTING](CONTRIBUTING.md) file for guidelines on how to contribute to this project.

## License

This library is licensed under the [MIT License](LICENSE), see the LICENSE file for details.
