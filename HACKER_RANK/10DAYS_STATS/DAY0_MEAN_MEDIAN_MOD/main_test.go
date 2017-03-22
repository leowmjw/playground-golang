package main

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCalculateMean(t *testing.T) {
	Convey("Overall ..", t, func() {

		test_data := []int{64630, 11735, 14216, 99233, 14470, 4978, 73429, 38120, 51135, 67060}

		Convey("Scenario: Normal, nice round value", func() {
			test_data = []int{4, 11, 1, 2, 7}
			result_mean := CalculateMean(test_data)
			So(result_mean, ShouldEqual, "5")
		})

		Convey("Scenario: Average, round to 1 decimal", func() {
			result_mean := CalculateMean(test_data)
			So(result_mean, ShouldEqual, "43900.6")
		})

	})
}

func TestCalculateMedian(t *testing.T) {
	Convey("Calculate Mean", t, func() {

		test_data := []int{64630, 11735, 14216, 99233, 14470, 4978, 73429, 38120, 51135, 67060}

		Convey("Scenario: Even elements in set (average it)", func() {
			test_data = []int{50, 44, 4, 2}
			result_median := CalculateMedian(test_data)
			So(result_median, ShouldEqual, "24")
		})

		Convey("Scenario: Even elements in set (average it, round to 1 decimal)", func() {
			result_median := CalculateMedian(test_data)
			So(result_median, ShouldEqual, "44627.5")
		})

		Convey("Scenario: Odd elements in set (pick middle)", func() {
			test_data = []int{50, 1, 44, 3, 2}
			result_median := CalculateMedian(test_data)
			So(result_median, ShouldEqual, "3")
		})

	})
}

func TestCalculateMode(t *testing.T) {
	Convey("Calculate Mode", t, func() {

		test_data := []int{64630, 11735, 14216, 99233, 14470, 4978, 73429, 38120, 51135, 67060}

		Convey("Scenario: Normal Mod (pick most frequent)", func() {
			test_data = []int{1, 3, 3, 5, 6, 10000, 3, 2, 1}
			result_mode := CalculateMode(test_data)
			So(result_mode, ShouldEqual, "3")
		})

		Convey("Scenario: Multimodal Mod (pick smallest numerical)", func() {
			result_mode := CalculateMode(test_data)
			So(result_mode, ShouldEqual, "4978")

		})

	})
}

func TestPairSort(t *testing.T) {
	SkipConvey("Test Len", t, func() {
		// test_case := []Pair{Pair{10, 1, }, Pair{1, 2, }, Pair{3, 5, }}
		// _ := test_case
		Convey("Normal", func() {

		})
	})

	SkipConvey("Test Less", t, func() {
		Convey("Normal", func() {

		})
	})

	SkipConvey("Test Swap", t, func() {
		Convey("Normal", func() {

		})
	})

	SkipConvey("Test Actual Sorting byVal", t, func() {
		Convey("Normal", func() {

		})
	})
}
