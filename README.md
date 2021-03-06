[![Build Status](https://travis-ci.org/xrash/gol.svg?branch=master)](http://travis-ci.org/xrash/gol)

# gol

`gol` is `log`, but backwards.

# How to import

You should import the v2:

```go
import "github.com/xrash/gol/v2"
```

# Documentation

Go to [https://pkg.go.dev/github.com/xrash/gol/v2](https://pkg.go.dev/github.com/xrash/gol/v2) for complete documentation.

# Example

```go
package main

import (
	"github.com/xrash/gol/v2"
	"github.com/xrash/gol/v2/formatters"
	"github.com/xrash/gol/v2/handlers"
)

func main() {
	logger := gol.NewLogger()

	basicFormatter := formatters.NewBasicFormatter()
	stdoutHandler := handlers.NewStdoutHandler()

	logger.AddHandler(stdoutHandler, basicFormatter, gol.LEVEL_DEBUG)

	// ...

	logger.Debug("Example %d, level %s", 1, "debug")
	logger.Notice("Example %d, level %s", 2, "notice")
	logger.Info("Example %d, level %s", 3, "info")
	logger.Warn("Example %d, level %s", 4, "warn")
	logger.Error("Example %d, level %s", 5, "error")
	logger.Crit("Example %d, level %s", 6, "crit")
	logger.Alert("Example %d, level %s", 7, "alert")
	logger.Emerg("Example %d, level %s", 8, "emerg")
}
```
