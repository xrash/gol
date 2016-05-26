package tests

import (
	"errors"
	"fmt"
	"github.com/xrash/gol"
	"testing"
)

// Simple Formatter
type FormatterMock struct{}

func (f *FormatterMock) Format(message string, level gol.LogLevel) string {
	return fmt.Sprintf("%v/%v", message, level)
}

// This Handler simply records the last logged message
type HandlerMock struct {
	LastLoggedLine string
}

func (h *HandlerMock) Handle(line string) error {
	h.LastLoggedLine = line

	return nil
}

// This Handler always return an error.
type ErrorHandlerMock struct {
	ErrorToReturn error
}

func (h *ErrorHandlerMock) Handle(line string) error {
	return h.ErrorToReturn
}

// Test different handlers with different levels.
func TestLevels(t *testing.T) {
	formatter := &FormatterMock{}

	handlerOne := &HandlerMock{}
	handlerTwo := &HandlerMock{}
	handlerThree := &HandlerMock{}

	logger := gol.NewLogger()
	logger.AddHandler(handlerOne, formatter, gol.LEVEL_DEBUG)
	logger.AddHandler(handlerTwo, formatter, gol.LEVEL_ERROR)
	logger.AddHandler(handlerThree, formatter, gol.LEVEL_INFO)

	executeForEveryHandlerEntry(logger, func(level gol.LogLevel, n int, handlerEntry *gol.HandlerEntry) {
		message := "wachacha"

		gol.FuncForLevel(logger, level)(message)
		expectedMessage := handlerEntry.Formatter.Format(message, level)
		handler := handlerEntry.Handler.(*HandlerMock)

		if handler.LastLoggedLine != expectedMessage {
			t.Fail()
			result := fmt.Sprintf("HandlerEntry [%v,%v]: [%v] should be equal [%v])", level, n, handler.LastLoggedLine, expectedMessage)
			fmt.Println(result)
		}
	})
}

// Test if errors are being correctly returned.
func TestErrors(t *testing.T) {
	formatter := &FormatterMock{}

	handlerOne := &ErrorHandlerMock{errors.New("one")}
	handlerTwo := &ErrorHandlerMock{errors.New("two")}
	handlerThree := &ErrorHandlerMock{errors.New("three")}
	handlerFour := &ErrorHandlerMock{errors.New("four")}
	handlerFive := &ErrorHandlerMock{errors.New("five")}

	logger := gol.NewLogger()
	logger.AddHandler(handlerOne, formatter, gol.LEVEL_INFO)
	logger.AddHandler(handlerTwo, formatter, gol.LEVEL_ERROR)
	logger.AddHandler(handlerThree, formatter, gol.LEVEL_WARN)
	logger.AddHandler(handlerFour, formatter, gol.LEVEL_INFO)
	logger.AddHandler(handlerFive, formatter, gol.LEVEL_DEBUG)

	executeForEveryHandlerEntry(logger, func(level gol.LogLevel, n int, handlerEntry *gol.HandlerEntry) {
		err := gol.FuncForLevel(logger, level)("test")
		handlerEntries := logger.HandlerEntries[int(level)]
		expectedErrors := make([]error, 0)

		for _, e := range handlerEntries {
			handler := e.Handler.(*ErrorHandlerMock)
			expectedErrors = append(expectedErrors, handler.ErrorToReturn)
		}

		if !areErrorSlicesEqual(err, expectedErrors) {
			t.Fail()
			result := fmt.Sprintf("%v should be equal %v", err, expectedErrors)
			fmt.Println(result)
		}
	})
}

// Executes function @fn for every HandlerEntry of @logger.
func executeForEveryHandlerEntry(logger *gol.Logger, fn func(gol.LogLevel, int, *gol.HandlerEntry)) {
	for _, level := range gol.Levels() {
		handlerEntries := logger.HandlerEntries[int(level)]
		for n, handlerEntry := range handlerEntries {
			fn(level, n, handlerEntry)
		}
	}
}

// Return true if both error slices are equal, false otherwise.
func areErrorSlicesEqual(a, b []error) bool {
	if len(a) != len(b) {
		return false
	}

	for k := range a {
		if a[k] != b[k] {
			return false
		}
	}

	return true
}
