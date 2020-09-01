# gol

`gol` is `log`, but backwards.

* master: [![Build Status](https://travis-ci.org/xrash/gol.svg?branch=master)](http://travis-ci.org/xrash/gol)

# How it works

`gol` has three elements you must know:

 - Formatters;
 - Handlers;
 - The logger;

Here is how to create everything:

```go
logger := gol.NewLogger()

basicFormatter := formatters.NewBasicFormatter()
stdoutHandler := handlers.NewStdoutHandler()

logger.AddHandler(stdoutHandler, basicFormatter, gol.LEVEL_DEBUG)
```

The formatter implements the interface `gol.Formatter`, the handler implements `gol.Handler` and the logger is the struct `gol.Logger`.

# Useful formatters and handlers

There are some useful formatters in the `github.com/xrash/gol/v2/formatters` package, and useful handlers in the `github.com/xrash/gol/v2/handlers` package. Check them out:

Useful formatters:

 - github.com/xrash/gol/v2/formatters.BasicFormatter

Useful handlers:

 - github.com/xrash/gol/v2/handlers.WriterHandler
 - github.com/xrash/gol/v2/handlers.StdoutHandler
 - github.com/xrash/gol/v2/handlers.FileHandler
 - github.com/xrash/gol/v2/handlers.MemoryQueueHandler
