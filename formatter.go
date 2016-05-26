package gol

type Formatter interface {
	Format(message string, level LogLevel) string
}
