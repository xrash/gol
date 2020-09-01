/*

Package gol provides a simple but powerful logging tool,
useful both in small and large projects.

The main functionality of the package is accessed through a struct
called Logger.

Logger provices basic logging functionality through
shortcut functions, like Logger.Info(), Logger.Debug() and Logger.Error().
Each of these shortcut functions correspond to a specific
level of severity.


Levels of severity


There are 8 levels of severity, as defined by syslog:

 0: emerg
 1: alert
 2: crit
 3: error
 4: warn
 5: notice
 6: info
 7: debug

Each level of severity has a corresponding package-level constant of type LogLevel:

 LEVEL_EMERG
 LEVEL_ALERT
 LEVEL_CRIT
 LEVEL_ERROR
 LEVEL_WARN
 LEVEL_NOTICE
 LEVEL_INFO
 LEVEL_DEBUG


Handlers


A handler is a struct that implements the Handler interface.
A handler receives the final message and handles it. It might
write the message to a logfile, store the message into a database,
send the message over the network etc.

In practice, handlers are associated with a specific formatter
and a specific level of severity.

The subdirectory "handlers" contains a bunch of basic and useful
implementations of Handler.

    import "github.com/xrash/gol/v2/handlers"

    // ...

    stdoutHandler := handlers.NewStdoutHandler()
    fileHandler := handlers.NewFileHandler("/tmp/any.log")
    memoryHandler := handlers.NewMemoryHandler(100, 1000)

Formatters


A formatter is a struct that implements the Formatter interface.
The formatter receives the message as sent by the caller, and formats it
before it's passed to the handler.

The subdirectory "formatters" contains a bunch of basic and useful
implementations of Formatter.

    import "github.com/xrash/gol/v2/formatters"

    // ...

    basicFormatter := formatters.NewBasicFormatter()



Complete example


Complete example of package usage, with handlers and formatters:

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

    	logger.Emerg("Example %d, level %s", gol.LEVEL_EMERG, "emerg")
    	logger.Alert("Example %d, level %s", gol.LEVEL_ALERT, "alert")
    	logger.Crit("Example %d, level %s", gol.LEVEL_CRIT, "crit")
    	logger.Error("Example %d, level %s", gol.LEVEL_ERROR, "error")
    	logger.Warn("Example %d, level %s", gol.LEVEL_WARN, "warn")
    	logger.Info("Example %d, level %s", gol.LEVEL_INFO, "info")
    	logger.Notice("Example %d, level %s", gol.LEVEL_NOTICE, "notice")
    	logger.Debug("Example %d, level %s", gol.LEVEL_DEBUG, "debug")
    }

*/
package gol
