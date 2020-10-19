package formatters

import (
	"github.com/franela/goblin"
	"github.com/xrash/gol/v2"
	"testing"
)

func TestNoopFormatter(t *testing.T) {
	f := NoopFormatter{}
	g := goblin.Goblin(t)
	g.Describe("NoopFormatter tests", func() {
		g.It("Result should be as expected", func() {
			r := f.Format("wachacha", gol.LEVEL_DEBUG)
			expected := "wachacha"
			g.Assert(r).Equal(expected)
		})
	})
}

func TestNoopLineFormatter(t *testing.T) {
	f := NoopLineFormatter{}
	g := goblin.Goblin(t)
	g.Describe("NoopLineFormatter tests", func() {
		g.It("Result should be as expected", func() {
			r := f.Format("wachacha", gol.LEVEL_DEBUG)
			expected := "wachacha\n"
			g.Assert(r).Equal(expected)
		})
	})
}
