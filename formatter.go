package gol

type Formatter interface {
	// Format returns the final, formatted message, as it will be passed to Handler.
	Format(message string, level LogLevel) string
}
