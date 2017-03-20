package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRangeModularQueries(t *testing.T) {

	Convey("Test ChooseArraySegment", t, func() {
		// Setup fixtures
		original_array := []int{250, 501, 5000, 5, 4}
		left_start := 0
		right_end := 4

		// Test Scenario #1
		Convey("Normal scenario ", func() {
			expected_array := []int{250, 501, 5000, 5, 4}
			array_to_check := ChooseArraySegment(original_array, left_start, right_end)
			So(array_to_check, ShouldResemble, expected_array)
		})
		// Test Scenario #2
		Convey("Edge scenario ", func() {
			left_start = 2
			right_end = 2
			expected_array := []int{5000}
			array_to_check := ChooseArraySegment(original_array, left_start, right_end)
			So(array_to_check, ShouldResemble, expected_array)
		})
	})

	Convey("Test CountModulo", t, func() {
		// Setup fixtures
		array_to_check := []int{250, 501, 5000, 5, 4}
		mod_value := 5
		remain_value := 0

		Convey("Normal scenario ", func() {
			number_of_matches := CountModulo(array_to_check, mod_value, remain_value)
			So(number_of_matches, ShouldEqual, 3)
		})

		Convey("Edge scenario ", func() {
			mod_value = 3
			remain_value = 2
			number_of_matches := CountModulo(array_to_check, mod_value, remain_value)
			So(number_of_matches, ShouldEqual, 2)
		})

		Convey("Wacky scenario ", func() {
			array_to_check = []int{5000}
			mod_value = 3
			remain_value = 2
			number_of_matches := CountModulo(array_to_check, mod_value, remain_value)
			So(number_of_matches, ShouldEqual, 1)
		})

	})

}
