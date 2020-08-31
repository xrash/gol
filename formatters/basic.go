package formatters

import (
	"github.com/xrash/gol"
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

	params := map[string]string{
		"%timestamp%": timestamp,
		"%message%":   message,
		"%level%":     l.String(),
	}

	line := MESSAGE_FORMAT

	for key, value := range params {
		line = strings.Replace(line, key, value, -1)
	}

	return line
}
