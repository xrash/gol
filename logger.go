package gol

import (
	"fmt"
)

// Logger is the main struct.
type Logger struct {
	// HandlerEntries has the following layout:
	//
	// [
	//   LEVEL_EMERG:  [ A, B, C ],
	//   LEVEL_ALERT:  [ A, B ],
	//   LEVEL_CRIT:   [ A, B ],
	//   LEVEL_ERROR:  [ A, B ],
	//   LEVEL_WARN:   [ A, B ],
	//   LEVEL_NOTICE: [ A, B ],
	//   LEVEL_INFO:   [ A ],
	//   LEVEL_DEBUG:  [ A ],
	// ]
	//
	// A, B and C are instances of HandlerEntry.
	// 
	// Each key is a LogLevel, and each value stores a slice
	// of HandlerEntry instances, whose Handler will handle calls
	// for that specific LogLevel.
	HandlerEntries [][]*HandlerEntry
}

// HandlerEntry is how a handler is associated with a formatter
// and a specific level of severity.
type HandlerEntry struct {
	Handler   Handler
	Formatter Formatter
	Level     LogLevel
}

// Constructs a new Logger.
func NewLogger() *Logger {
	handlerEntries := make([][]*HandlerEntry, len(LogLevels))

	for e := range handlerEntries {
		handlerEntries[e] = make([]*HandlerEntry, 0)
	}

	return &Logger{
		handlerEntries,
	}
}

// Adds Handler h to Logger.
//
// The Handler's final message is formatted
// by Formatter f, and will be called for every message of LogLevel <= level.
//
// For example, a Handler for LEVEL_NOTICE will be called for messages
// of level LEVEL_NOTICE or LEVEL_ERROR, but will not be called for messages
// of level LEVEL_DEBUG or LEVEL_INFO, because LEVEL_DEBUG and LEVEL_INFO > LEVEL_NOTICE.
func (l *Logger) AddHandler(h Handler, f Formatter, level LogLevel) *HandlerEntry {
	handlerEntry := &HandlerEntry{
		Handler:   h,
		Formatter: f,
		Level:     level,
	}

	for i := 0; i <= int(level); i++ {
		l.HandlerEntries[i] = append(l.HandlerEntries[i], handlerEntry)
	}

	return handlerEntry
}

// Handle message according to level. Prefer to use shortcut functions instead of this one.
func (l *Logger) Handle(level LogLevel, format string, params ...interface{}) []error {
	message := fmt.Sprintf(format, params...)

	errors := make([]error, 0)

	for _, handlerEntry := range l.HandlerEntries[level] {
		err := handlerEntry.Handler.Handle(handlerEntry.Formatter.Format(message, level))
		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

type LogFunction func(format string, params ...interface{}) []error

// Shortcut function to handle message on the EMERG level.
func (l *Logger) Emerg(format string, params ...interface{}) []error {
	return l.Handle(LEVEL_EMERG, format, params...)
}

// Shortcut function to handle message on the ALERT level.
func (l *Logger) Alert(format string, params ...interface{}) []error {
	return l.Handle(LEVEL_ALERT, format, params...)
}

// Shortcut function to handle message on the CRIT level.
func (l *Logger) Crit(format string, params ...interface{}) []error {
	return l.Handle(LEVEL_CRIT, format, params...)
}

// Shortcut function to handle message on the ERROR level.
func (l *Logger) Error(format string, params ...interface{}) []error {
	return l.Handle(LEVEL_ERROR, format, params...)
}

// Shortcut function to handle message on the WARN level.
func (l *Logger) Warn(format string, params ...interface{}) []error {
	return l.Handle(LEVEL_WARN, format, params...)
}

// Shortcut function to handle message on the NOTICE level.
func (l *Logger) Notice(format string, params ...interface{}) []error {
	return l.Handle(LEVEL_NOTICE, format, params...)
}

// Shortcut function to handle message on the INFO level.
func (l *Logger) Info(format string, params ...interface{}) []error {
	return l.Handle(LEVEL_INFO, format, params...)
}

// Shortcut function to handle message on the DEBUG level.
func (l *Logger) Debug(format string, params ...interface{}) []error {
	return l.Handle(LEVEL_DEBUG, format, params...)
}
