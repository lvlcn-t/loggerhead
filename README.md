# halog Logging Library

## Overview
halog is a versatile and extendable logging library for Go, designed to enhance the capabilities of standard Go logging with additional features and utilities. Building on top of the `slog.Logger`, it provides a robust solution for application logging, including extended log levels, formatted logging methods, and context-aware logging capabilities.

## Features
- **Extended Log Levels**: Includes standard levels like Debug, Info, Warn, and Error, along with custom levels for finer control.
- **Formatted Logging**: Methods like `Debugf`, `Infof`, `Warnf`, and `Errorf` for formatted output.
- **Context-Aware Logging**: Enhanced logging functionalities that leverage Go's context for richer, more informative log entries.
- **Thread-Safe Operations**: Ensures safe use in concurrent environments.
- **Customizable Handlers**: Easily integrate with different logging backends and formats.

## Getting Started
To start using halog, install the package using:

```bash
go get github.com/lvlcn-t/halog@latest
```

### Basic Usage
Here's a simple example of how to use it in your Go application:

```go
package main

import (
    "github.com/lvlcn-t/halog/logger"
    "log/slog"
)

func main() {
    log := logger.NewLogger()
    log.Infof("This is an info message: %s", "Hello, logger!")
}
```

## Documentation
For detailed documentation, including all available methods and configuration options, refer to the [tbd](#documentation).

## Contributing
Contributions are welcome! Please refer to the [tbd](#CONTRIBUTING.md) file for guidelines on how to contribute to this project.

## License
This library is licensed under the [MIT License](LICENSE), see the LICENSE file for details.
