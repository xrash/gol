package gol

type LogLevel int

func LevelsFuncMap(logger *Logger) map[LogLevel]LogFunction {
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

func FuncForLevel(logger *Logger, level LogLevel) LogFunction {
	return LevelsFuncMap(logger)[level]
}

func Levels() []LogLevel {
	return []LogLevel{
		LEVEL_EMERG,
		LEVEL_ALERT,
		LEVEL_CRIT,
		LEVEL_ERROR,
		LEVEL_WARN,
		LEVEL_NOTICE,
		LEVEL_INFO,
		LEVEL_DEBUG,
	}
}

const (
	LEVELS_QUANTITY = 8
)

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

func NewLogLevel(level string) LogLevel {
	switch level {
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
