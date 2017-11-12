package main

import (
	. "github.com/pranavraja/zen"
	"testing"
)

func TestShelf(t *testing.T) {
	Desc(t, "shelf", func(it It) {
		shelf := NewShelf()
		shelf.AddProject("Known", "Known", 100)
		it("should find Room in projets when exist", func(expect Expect) {
			expect(shelf.FindProject("Known")).ToExist()
		})
		it("should know when things don't exist", func(expect Expect) {
			expect(shelf.FindProject("Unknown")).ToNotExist()
		})
	})
}
