package main_test

import (
	"testing"
	"github.com/smartystreets/goconvey/convey"
	bw "github.com/leowmjw/playground-golang/HACKER_RANK/W31/BEAUTIFUL_WORD"
)

func TestFindingUgly(t *testing.T) {
	convey.Convey("Confirming the Beauty", t, func() {
		split_word := []string{"a", "b", "a", "c", "a", "b", "a", }
		current_state := bw.NewCurrentState(false, "")

		convey.Convey("Scenario #1: Start up with first char", func() {
			current_state.FoundUgly(split_word[0])
			// Test state change
			convey.So(current_state, convey.ShouldResemble, bw.NewCurrentState(true, "a"))
		})

		convey.Convey("Scenario #2: Ending char .. happy ending!", func() {
			current_state := bw.NewCurrentState(false, "b")
			current_state.FoundUgly(split_word[6])
			// Test state change
			convey.So(current_state, convey.ShouldResemble, bw.NewCurrentState(true, "a"))
		})
	})

	convey.Convey("Pilloring the Wretch", t, func() {
		split_word := []string{"b", "a", "d", "d", }
		current_state := bw.NewCurrentState(false, "")

		convey.Convey("Scenario #1: Start up with first char", func() {
			has_found_ugly := current_state.FoundUgly(split_word[0])
			// Test finding ugly
			convey.So(has_found_ugly, convey.ShouldEqual, false)
			// Test state change
			convey.So(current_state, convey.ShouldResemble, bw.NewCurrentState(false, "b"))
		})

		convey.Convey("Scenario #2: Ending char hitting consecutive same char", func() {
			current_state := bw.NewCurrentState(false, "d")
			has_found_ugly := current_state.FoundUgly(split_word[3])
			// Test finding ugly
			convey.So(has_found_ugly, convey.ShouldEqual, true)
			// Test state change
			convey.So(current_state, convey.ShouldResemble, bw.NewCurrentState(false, "d"))
		})

		convey.Convey("Scenario #3: Hitting consecutive vowels", func() {
			current_state := bw.NewCurrentState(true, "a")
			has_found_ugly := current_state.FoundUgly("i")
			// Test finding ugly
			convey.So(has_found_ugly, convey.ShouldEqual, true)
			// Test state change
			convey.So(current_state, convey.ShouldResemble, bw.NewCurrentState(true, "i"))
		})

		convey.Convey("Scenario #4: False alarm; saw vowel but now is consonents", func() {
			current_state := bw.NewCurrentState(true, "a")
			has_found_ugly := current_state.FoundUgly(split_word[2])
			// Test finding ugly
			convey.So(has_found_ugly, convey.ShouldEqual, false)
			// Test state change
			convey.So(current_state, convey.ShouldResemble, bw.NewCurrentState(false, "d"))
		})

	})

}

func TestBeautifulWordDetection(t *testing.T) {
	convey.Convey("Full Detection in Words", t, func() {

		convey.Convey("Scenario #1: Beautiful", func() {
			word := "batman"
			is_beautiful := bw.Beautiful_Word(word)
			convey.So(is_beautiful, convey.ShouldEqual, "Yes")
		})

		convey.Convey("Scenario #2: Not Beautiful - Consecutive chars", func() {
			word := "badd"
			is_beautiful := bw.Beautiful_Word(word)
			convey.So(is_beautiful, convey.ShouldEqual, "No")
		})

		convey.Convey("Scenario #3: Not Beautiful - Consecutive vowels", func() {
			word := "beauty"
			is_beautiful := bw.Beautiful_Word(word)
			convey.So(is_beautiful, convey.ShouldEqual, "No")
		})

		convey.Convey("Scenario #4: Not Beautiful - Mixed", func() {
			word := "sonniebond"
			is_beautiful := bw.Beautiful_Word(word)
			convey.So(is_beautiful, convey.ShouldEqual, "No")
		})

	})
}
