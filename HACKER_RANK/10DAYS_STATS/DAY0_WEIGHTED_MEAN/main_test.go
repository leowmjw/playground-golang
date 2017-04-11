package main

import "testing"

func TestWeightedMean(t *testing.T) {
	test_data := []int{10, 40, 30, 50, 20}
	weight_data := []int{1, 2, 3, 4, 5}
	result := CalculateWeightedMean(test_data, weight_data)
	if result != "32.0" {
		t.Fail()
	}
}
