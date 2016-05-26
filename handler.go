package gol

type Handler interface {
	Handle(line string) error
}
