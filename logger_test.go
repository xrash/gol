package gol

import (
	"fmt"
	"github.com/franela/goblin"
	"testing"
)

// `formatterMock` implements the Formatter interface.
type formatterMock struct{}

func (f *formatterMock) Format(message string, level LogLevel) string {
	return fmt.Sprintf("%v/%v", message, level)
}

// `lastLineHandlerMock` implements the Handler interface,
// and records the last logged message.
type lastLineHandlerMock struct {
	lastLoggedLine string
}

func (h *lastLineHandlerMock) Handle(line string) error {
	h.lastLoggedLine = line
	return nil
}

// `errorHandlerMock` implements the Handler interface,
// and always returns an error.
type errorHandlerMock struct {
	errorToReturn error
}

func (h *errorHandlerMock) Handle(line string) error {
	return h.errorToReturn
}

// The Logger should call handlers to handle their specific LogLevel,
// and every LogLevel below it, and no LogLevel above it.
func TestIfLevelsAreWorkingCorrectly(t *testing.T) {
	logger := NewLogger()
	formatter := &formatterMock{}

	for _, level := range LogLevels {
		handler := &lastLineHandlerMock{}
		logger.AddHandler(handler, formatter, level)
	}

	for _, functionLevel := range LogLevels {
		message := "test"
		FuncForLevel(logger, functionLevel)(message)
		expectedLine := formatter.Format(message, functionLevel)

		g := goblin.Goblin(t)

		g.Describe(fmt.Sprintf("Testcase [%s]", expectedLine), func() {
			for _, handlerEntry := range logger.HandlerEntries[LEVEL_EMERG] {
				lastLineHandlerMock := handlerEntry.Handler.(*lastLineHandlerMock)
				handlerLevel := handlerEntry.Level
				lastLoggedLine := lastLineHandlerMock.lastLoggedLine

				if handlerLevel >= functionLevel {
					g.It("Last logged line should be equal", func() {
						g.Assert(lastLoggedLine).Equal(expectedLine)
					})
				} else {
					g.It("Last logged line should be different", func() {
						g.Assert(lastLoggedLine != expectedLine).IsTrue()
					})
				}
			}
		})

	}
}

// The Logger should correctly return the errors returned by handlers.
func TestErrors(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Test errors", func() {

		g.It("There must be one error", func() {
			logger := NewLogger()
			formatter := &formatterMock{}
			handler1 := &lastLineHandlerMock{}
			handler2 := &lastLineHandlerMock{}
			handler3 := &errorHandlerMock{
				errorToReturn: fmt.Errorf("any error"),
			}
			handler4 := &lastLineHandlerMock{}
			logger.AddHandler(handler1, formatter, LEVEL_DEBUG)
			logger.AddHandler(handler2, formatter, LEVEL_INFO)
			logger.AddHandler(handler3, formatter, LEVEL_NOTICE)
			logger.AddHandler(handler4, formatter, LEVEL_CRIT)
			err := logger.Alert("test")
			g.Assert(len(err)).Equal(1)
		})

		g.It("There must be two errors", func() {
			logger := NewLogger()
			formatter := &formatterMock{}
			handler1 := &lastLineHandlerMock{}
			handler2 := &errorHandlerMock{
				errorToReturn: fmt.Errorf("any error"),
			}
			handler3 := &errorHandlerMock{
				errorToReturn: fmt.Errorf("any error"),
			}
			handler4 := &lastLineHandlerMock{}
			logger.AddHandler(handler1, formatter, LEVEL_DEBUG)
			logger.AddHandler(handler2, formatter, LEVEL_INFO)
			logger.AddHandler(handler3, formatter, LEVEL_NOTICE)
			logger.AddHandler(handler4, formatter, LEVEL_CRIT)
			err := logger.Alert("test")
			g.Assert(len(err)).Equal(2)
		})

		g.It("There must be three errors", func() {
			logger := NewLogger()
			formatter := &formatterMock{}
			handler1 := &errorHandlerMock{
				errorToReturn: fmt.Errorf("any error"),
			}
			handler2 := &errorHandlerMock{
				errorToReturn: fmt.Errorf("any error"),
			}
			handler3 := &errorHandlerMock{
				errorToReturn: fmt.Errorf("any error"),
			}
			logger.AddHandler(handler1, formatter, LEVEL_DEBUG)
			logger.AddHandler(handler2, formatter, LEVEL_INFO)
			logger.AddHandler(handler3, formatter, LEVEL_NOTICE)
			err := logger.Alert("test")
			g.Assert(len(err)).Equal(3)
		})

		g.It("There must be no error", func() {
			logger := NewLogger()
			formatter := &formatterMock{}
			handler1 := &lastLineHandlerMock{}
			logger.AddHandler(handler1, formatter, LEVEL_DEBUG)
			err := logger.Alert("test")
			g.Assert(len(err)).Equal(0)
		})

	})

}
