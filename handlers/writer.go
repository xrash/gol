package handlers

import (
	"io"
)

type WriterHandler struct {
	w io.Writer
}

func NewWriterHandler(w io.Writer) *WriterHandler {
	return &WriterHandler{
		w,
	}
}

func (h *WriterHandler) Handle(s string) error {
	_, err := h.w.Write([]byte(s))

	return err
}
