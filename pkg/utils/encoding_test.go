package utils

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestCharmapEqual(t *testing.T) {
	g := Goblin(t)
	g.Describe("Check DecodeCharmap()", func() {
		g.It("check cp866", func() {
			g.Assert(DecodeCharmap("\x92\xa5\xe1\xe2", "CP866")).Equal("Тест")
		})
		g.It("check cp895", func() {
			g.Assert(DecodeCharmap("\x80\x81\x82\x83", "CP895")).Equal("Čěšý")
		})
		g.It("check utf-8", func() {
			g.Assert(DecodeCharmap("Тест", "UTF-8")).Equal("Тест")
		})
	})
	g.Describe("Check EncodeCharmap()", func() {
		g.It("check cp866", func() {
			g.Assert(EncodeCharmap("Тест", "CP866")).Equal("\x92\xa5\xe1\xe2")
		})
		g.It("check cp895", func() {
			g.Assert(EncodeCharmap("Čěšý", "CP895")).Equal("\x80\x81\x82\x83")
		})
		g.It("check utf-8", func() {
			g.Assert(EncodeCharmap("Тест", "UTF-8")).Equal("Тест")
		})
	})
}
