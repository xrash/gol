package formatters

import (
	"github.com/franela/goblin"
	"github.com/xrash/gol/v2"
	"testing"
	"time"
)

func testNow() time.Time {
	return time.Unix(0, 0)
}

func TestBasicFormatter(t *testing.T) {
	f := BasicFormatter{
		nowProvider: testNow,
	}

	g := goblin.Goblin(t)

	g.Describe("Basic Formatter tests", func() {

		g.It("Result should be as expected", func() {
			r := f.Format("wachacha", gol.LEVEL_DEBUG)
			expected := "[1969-12-31T21:00:00.000-03:00] [debug] wachacha\n"
			g.Assert(r).Equal(expected)
		})

	})
}
