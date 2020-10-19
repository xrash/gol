package formatters

import (
	"github.com/xrash/gol/v2"
)

// NoopFormatter just forwards the raw message as it was
// passed to the methods Info(), Error() etc.
type NoopFormatter struct{}

func NewNoopFormatter() *NoopFormatter {
	return &NoopFormatter{}
}

func (f *NoopFormatter) Format(message string, level gol.LogLevel) string {
	return message
}

// NoopLineFormatter is the same as NoopFormatter, but
// adds a newline at the end of the message.
type NoopLineFormatter struct{}

func NewNoopLineFormatter() *NoopLineFormatter {
	return &NoopLineFormatter{}
}

func (f *NoopLineFormatter) Format(message string, level gol.LogLevel) string {
	return message + "\n"
}
