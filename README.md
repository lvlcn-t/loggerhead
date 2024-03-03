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

Here's a simple example of how to use Loggerhead with default settings:

```go
package main

import (
  "github.com/lvlcn-t/loggerhead/logger"
)

func main() {
  log := logger.NewLogger()
  log.Infof("This is an info message: %s", "Hello, logger!")
}
```

## Documentation

Loggerhead provides a comprehensive set of features for advanced logging in Go applications. Here's an overview of its primary functionalities and how to use them effectively:

### Core Logging Functions

#### NewLogger

This function initializes a new logger with default settings or [custom handlers](#custom-handlers) if provided. It's the starting point for using Loggerhead in your application.

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

Loggerhead supports configuration of the default logger behavior through environment variables. This allows for easy customization of log level and format without altering the codebase:

- `LOG_LEVEL`: Determines the log output level (e.g., DEBUG, INFO, WARN). This setting influences which logs are output based on their severity.
- `LOG_FORMAT`: Specifies the log output format. Set to "TEXT" for human-readable logs or any other value for JSON format, which is the default.

### Extending Loggerhead

Loggerhead is designed to be highly extendable, offering several ways for developers to customize and enhance its functionality:

#### Custom Handlers

Developers have the flexibility to use their own `slog.Handler` implementations with Loggerhead. This allows for complete control over how logs are processed, formatted, and outputted. Whether integrating with third-party logging services, applying unique log formatting, or implementing custom log filtering logic, you can pass your custom handler to `NewLogger` or `NewNamedLogger` functions to replace the default behavior.

```go
  // Example of using a custom handler
  log := logger.NewLogger(myCustomHandler)
```

This feature is especially useful for applications with specific logging requirements not covered by the default handlers. By providing your own implementation, you can tailor the logging behavior to fit the needs of your application precisely.

#### Custom Log Formats

While Loggerhead supports text and JSON formats out of the box, through custom `slog.Handler` implementations, developers can define entirely custom formats. This is ideal for adhering to organizational logging standards or enhancing log readability.

#### Integration with Logging Backends

Custom handlers also enable integration with various logging backends and services. Whether you're sending logs to a file, a console, a database, or a cloud-based logging platform, you can encapsulate this logic within your handler and use it seamlessly with Loggerhead.

### Usage Scenarios

- **Application Logging**: Use Loggerhead for general application logging, capturing information ranging from debug messages to critical errors.
- **Request Tracking**: In web applications, Loggerhead can be used to log and track HTTP requests, providing valuable insights for debugging and monitoring.
- **Contextual Information**: With context-aware logging, Loggerhead is ideal for applications that require detailed logs with contextual information, especially useful in microservices and distributed systems.

For further examples and detailed usage, including how to implement and integrate custom `slog.Handler` instances, please refer to the [examples](./examples) directory in our repository.

## Contributing

Contributions are welcome! Please refer to the [CONTRIBUTING](CONTRIBUTING.md) file for guidelines on how to contribute to this project.

## License

This library is licensed under the [MIT License](LICENSE), see the LICENSE file for details.
