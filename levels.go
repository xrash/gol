package gol

import (
	"strings"
)

type LogLevel int

const (
	LEVEL_EMERG LogLevel = iota
	LEVEL_ALERT
	LEVEL_CRIT
	LEVEL_ERROR
	LEVEL_WARN
	LEVEL_NOTICE
	LEVEL_INFO
	LEVEL_DEBUG
)

var LogLevels = []LogLevel {
	LEVEL_EMERG,
	LEVEL_ALERT,
	LEVEL_CRIT,
	LEVEL_ERROR,
	LEVEL_WARN,
	LEVEL_NOTICE,
	LEVEL_INFO,
	LEVEL_DEBUG,
}

func levelsFuncMap(logger *Logger) map[LogLevel]LogFunction {
	return map[LogLevel]LogFunction{
		LEVEL_EMERG:  logger.Emerg,
		LEVEL_ALERT:  logger.Alert,
		LEVEL_CRIT:   logger.Crit,
		LEVEL_ERROR:  logger.Error,
		LEVEL_WARN:   logger.Warn,
		LEVEL_NOTICE: logger.Notice,
		LEVEL_INFO:   logger.Info,
		LEVEL_DEBUG:  logger.Debug,
	}
}

// Returns the specific LogFunction of Logger for LogLevel.
func FuncForLevel(logger *Logger, level LogLevel) LogFunction {
	return levelsFuncMap(logger)[level]
}

// Returns LogLevel by its string representation. Valid strings
// are: "emerg", "alert", "crit", "error", "warn",
// "notice", "info" and "debug".
//
// If the string is invalid, LEVEL_EMERG is returned.
func NewLogLevel(level string) LogLevel {
	switch strings.ToLower(level) {
	case "emerg":
		return LEVEL_EMERG
	case "alert":
		return LEVEL_ALERT
	case "crit":
		return LEVEL_CRIT
	case "error":
		return LEVEL_ERROR
	case "warn":
		return LEVEL_WARN
	case "notice":
		return LEVEL_NOTICE
	case "info":
		return LEVEL_INFO
	case "debug":
		return LEVEL_DEBUG
	}

	return LEVEL_EMERG
}

// Returns a string representation of LogLevel.
func (l LogLevel) String() string {
	switch l {
	case LEVEL_DEBUG:
		return "debug"
	case LEVEL_NOTICE:
		return "notice"
	case LEVEL_INFO:
		return "info"
	case LEVEL_WARN:
		return "warn"
	case LEVEL_ERROR:
		return "error"
	case LEVEL_CRIT:
		return "crit"
	case LEVEL_ALERT:
		return "alert"
	case LEVEL_EMERG:
		return "emerg"
	}

	return "unknown"
}
