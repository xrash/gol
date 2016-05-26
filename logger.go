package gol

import (
	"fmt"
)

type Logger struct {
	HandlerEntries [][]*HandlerEntry
}

type HandlerEntry struct {
	Handler   Handler
	Formatter Formatter
}

func NewLogger() *Logger {
	handlerEntries := make([][]*HandlerEntry, LEVELS_QUANTITY)

	for e := range handlerEntries {
		handlerEntries[e] = make([]*HandlerEntry, 0)
	}

	return &Logger{
		handlerEntries,
	}
}

func (l *Logger) AddHandler(h Handler, f Formatter, level LogLevel) {
	handlerEntry := &HandlerEntry{
		h,
		f,
	}

	for i := 0; i <= int(level); i++ {
		l.HandlerEntries[i] = append(l.HandlerEntries[i], handlerEntry)
	}
}

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

func (l *Logger) Emerg(format string, params ...interface{}) []error {
	return l.Handle(LEVEL_EMERG, format, params...)
}

func (l *Logger) Alert(format string, params ...interface{}) []error {
	return l.Handle(LEVEL_ALERT, format, params...)
}

func (l *Logger) Crit(format string, params ...interface{}) []error {
	return l.Handle(LEVEL_CRIT, format, params...)
}

func (l *Logger) Error(format string, params ...interface{}) []error {
	return l.Handle(LEVEL_ERROR, format, params...)
}

func (l *Logger) Warn(format string, params ...interface{}) []error {
	return l.Handle(LEVEL_WARN, format, params...)
}

func (l *Logger) Notice(format string, params ...interface{}) []error {
	return l.Handle(LEVEL_NOTICE, format, params...)
}

func (l *Logger) Info(format string, params ...interface{}) []error {
	return l.Handle(LEVEL_INFO, format, params...)
}

func (l *Logger) Debug(format string, params ...interface{}) []error {
	return l.Handle(LEVEL_DEBUG, format, params...)
}
