package handlers

import (
	"os"
)

type FileHandler struct {
	filename string
	file     *os.File
}

func NewFileHandler(filename string) *FileHandler {
	return &FileHandler{
		filename,
		nil,
	}
}

func (h *FileHandler) Open() error {
	var file *os.File
	var err error

	if fileExists(h.filename) {
		file, err = os.OpenFile(h.filename, os.O_WRONLY|os.O_APPEND, 0755)
	} else {
		file, err = os.OpenFile(h.filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	}

	if err != nil {
		return err
	}

	h.file = file

	return nil
}

func (h *FileHandler) Close() error {
	return h.file.Close()
}

func (h *FileHandler) Handle(s string) error {
	_, err := h.file.WriteString(s)

	return err
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}
