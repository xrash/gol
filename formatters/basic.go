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

type nowProvider func() time.Time

type BasicFormatter struct {
	nowProvider nowProvider
}

func NewBasicFormatter() *BasicFormatter {
	return &BasicFormatter{
		nowProvider: time.Now,
	}
}

func (f *BasicFormatter) Format(message string, l gol.LogLevel) string {
	timestamp := f.nowProvider().UTC().Format(TIMESTAMP_FORMAT)

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
