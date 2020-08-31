package gol

type Handler interface {
	// Handle receives the final, formatted message, and handles it.
	// Examples of handling are: outputting to STDOUT, writing on a file,
	// storing in memory, writing on a database, sending over the network etc.
	Handle(message string) error
}
