package formatters

import (
	"github.com/xrash/gol/v2"
	"strings"
	"time"
)

const (
	TIMESTAMP_FORMAT = "2006-01-02T15:04:05.000Z07:00"
	MESSAGE_FORMAT   = "[%timestamp%] [%level%] %message%\n"
)

type BasicFormatter struct {
}

func NewBasicFormatter() *BasicFormatter {
	return &BasicFormatter{}
}

func (f *BasicFormatter) Format(message string, l gol.LogLevel) string {
	timestamp := time.Now().Format(TIMESTAMP_FORMAT)
	level := f.FormatLevel(l)

	params := map[string]string{
		"%timestamp%": timestamp,
		"%message%":   message,
		"%level%":     level,
	}

	line := MESSAGE_FORMAT

	for key, value := range params {
		line = strings.Replace(line, key, value, -1)
	}

	return line
}

func (f *BasicFormatter) FormatLevel(level gol.LogLevel) string {
	switch level {
	case gol.LEVEL_DEBUG:
		return "debug"
	case gol.LEVEL_INFO:
		return "info"
	case gol.LEVEL_WARN:
		return "warn"
	case gol.LEVEL_ERROR:
		return "error"
	case gol.LEVEL_NOTICE:
		return "notice"
	case gol.LEVEL_ALERT:
		return "alert"
	case gol.LEVEL_EMERG:
		return "emerg"
	}

	return "unknown"
}
