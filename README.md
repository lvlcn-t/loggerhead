# Loggerhead - Logging Library<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<p align="center">
    <a href="/../../commits/" title="Last Commit"><img src="https://img.shields.io/github/last-commit/lvlcn-t/Loggerhead?style=flat"></a>
    <a href="/../../issues" title="Open Issues"><img src="https://img.shields.io/github/issues/lvlcn-t/Loggerhead?style=flat"></a>
    <a href="/../../pulls" title="Open Pull Requests"><img src="https://img.shields.io/github/issues-pr/lvlcn-t/Loggerhead?style=flat"></a>
</p>

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
    - [Usage Scenarios](#usage-scenarios)
    - [Extending Loggerhead](#extending-loggerhead)
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
go get github.com/lvlcn-t/loggerhead@v0.2.0
```

### Basic Usage

Here's a simple example of how to use it in your Go application:

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

This function initializes a new logger with default settings. It's the starting point for using Loggerhead in your application. 
```go
log := logger.NewLogger()
```

#### NewNamedLogger

Creates a new logger with a specified name. This is useful for identifying logs originating from different parts of your application.
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

### Usage Scenarios

- **Application Logging**: Use Loggerhead for general application logging, capturing information ranging from debug messages to critical errors.
- **Request Tracking**: In web applications, Loggerhead can be used to log and track HTTP requests, providing valuable insights for debugging and monitoring.
- **Contextual Information**: With context-aware logging, Loggerhead is ideal for applications that require detailed logs with contextual information, especially useful in microservices and distributed systems.

### Extending Loggerhead

Loggerhead is designed to be extendable. Developers can write their own log handling functions, customize log formats, and integrate with different logging backends or systems.

For further examples and detailed usage, please refer to the [_tbd_](./examples) directory in our repository.

## Contributing

Contributions are welcome! Please refer to the [CONTRIBUTING](#CONTRIBUTING.md) file for guidelines on how to contribute to this project.

## License
This library is licensed under the [MIT License](LICENSE), see the LICENSE file for details.
