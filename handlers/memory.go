package handlers

import (
	"container/list"
)

type MemoryQueueHandler struct {
	minSize int
	maxSize int
	queue   *list.List
}

func NewMemoryQueueHandler(minSize, maxSize int) *MemoryQueueHandler {
	return &MemoryQueueHandler{
		minSize: minSize,
		maxSize: maxSize,
		queue:   list.New(),
	}
}

func (h *MemoryQueueHandler) Handle(s string) error {
	if h.queue.Len() >= h.maxSize {
		h.Adjust()
	}

	h.queue.PushBack(s)

	return nil
}

func (h *MemoryQueueHandler) Adjust() {
	for h.queue.Len() >= h.minSize {
		element := h.queue.Front()
		h.queue.Remove(element)
	}
}
