package handlers

import (
	"os"
)

type StdoutHandler struct {
	wh *WriterHandler
}

func NewStdoutHandler() *StdoutHandler {
	return &StdoutHandler{
		wh: NewWriterHandler(os.Stdout),
	}
}

func (h *StdoutHandler) Handle(s string) error {
	return h.wh.Handle(s)
}
