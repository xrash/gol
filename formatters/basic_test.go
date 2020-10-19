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

	g.Describe("BasicFormatter tests", func() {

		g.It("Result should be as expected", func() {
			r := f.Format("wachacha", gol.LEVEL_DEBUG)
			expected := "[1970-01-01T00:00:00.000Z] [debug] wachacha\n"
			g.Assert(r).Equal(expected)
		})

	})
}
