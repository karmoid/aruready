package main

import (
	. "github.com/karmoid/aruready/karmostate"
	. "github.com/pranavraja/zen"
	"testing"
)

func TestZen(t *testing.T) {
	Desc(t, "zen", func(it It) {
		it("should know when things exist", func(expect Expect) {
			expect("tree").ToExist()
		})
		it("should know when things don't exist", func(expect Expect) {
			expect(nil).ToNotExist()
		})
		it("should know that a thing is equal to itself", func(expect Expect) {
			expect(1).ToEqual(1)
		})
		it("should know when things are different", func(expect Expect) {
			expect(1).ToNotEqual(2)
		})
		it("should be able to learn about new things", func(expect Expect) {
			divisibleBy := func(a, b interface{}) bool {
				return a.(int)%b.(int) == 0
			}
			expect(1).To("be divisible by", divisibleBy, 1)
		})
		it("should return 1 when we send One", func(expect Expect) {
			expect(GetState("un")).ToEqual("1")
		})
	})
}

func TestZen2(t *testing.T) {
	Desc(t, "before and after", func(it It) {
		count := 0

		before := func() {
			count++
		}

		after := func() {
			count--
		}

		setup := Setup(before, after)

		it("should execute before and after functions", setup(func(expect Expect) {
			expect(count).ToEqual(1)
		}))

		it("should execute before and after functions", setup(func(expect Expect) {
			expect(count).ToEqual(1)
		}))
	})
}

func TestCli(t *testing.T) {
	Desc(t, "CLI", func(it It) {
		submitData := "anything #tag1 #tag2 #tag3"
		unknownTag := "tag4"
		knownTag := "tag3"
		analyzedData, _ := CliAnalyze(submitData)
		it("should return 3 tags when submit 3 #", func(expect Expect) {
			expect(len(analyzedData.Tag())).ToEqual(3)
		})
		it("should not return unknown tag when don't exist", func(expect Expect) {
			expect(analyzedData.FindTag(unknownTag)).ToNotExist()
		})
		it("should return tag exists when exist", func(expect Expect) {
			expect(analyzedData.FindTag(knownTag)).ToExist()
		})
		// it("should return known tag when exist", func(expect Expect) {
		// 	expect(analyzedData.FindTag(knownTag)).ToEqual(knownTag)
		// })
	})
}
