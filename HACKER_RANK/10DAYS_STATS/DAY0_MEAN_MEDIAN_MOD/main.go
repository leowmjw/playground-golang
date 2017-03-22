package main

import (
	"fmt"
	"regexp"
	"sort"
)

func main() {
	// The first line contains an integer, , denoting the number of elements in the array.
	size_of_array := 0
	fmt.Scanf("%d", &size_of_array)
	// The second line contains  space-separated integers describing the array's elements.
	original_array := make([]int, size_of_array)
	for iter := 0; iter < size_of_array; iter++ {
		fmt.Scanf("%d", &original_array[iter])
	}
	/*
	Output Format

	Print  lines of output in the following order:

	Print the mean on a new line, to a scale of  decimal place (i.e., , ).
	Print the median on a new line, to a scale of  decimal place (i.e., , ).
	Print the mode on a new line; if more than one such value exists, print the numerically smallest one.

	 */
	fmt.Println(CalculateMean(original_array))
	fmt.Println(CalculateMedian(original_array))
	fmt.Println(CalculateMode(original_array))
}

func CalculateMean(array_to_calculate []int) string {
	total_sum := 0.0
	size_of_array := float64(len(array_to_calculate))
	for _, myval := range array_to_calculate {
		total_sum += float64(myval)
	}

	raw_mean := total_sum / size_of_array
	raw_mean_string := fmt.Sprintf("%.1f", raw_mean)
	match, _ := regexp.MatchString("^.*\\.0$", raw_mean_string)
	if match {
		return fmt.Sprintf("%d", int(raw_mean))
	}

	return raw_mean_string
}

func CalculateMedian(array_to_calculate []int) string {
	// First sort in order
	// fmt.Println("UNSORTED: ", array_to_calculate)
	sort.Ints(array_to_calculate)
	// fmt.Println("SORTED: ", array_to_calculate)
	// Test: Odd, Even??
	size_of_array := len(array_to_calculate)
	if size_of_array%2 != 0 {
		// if Odd; choose the one in the middle; next largest integer; roun dup so to speak
		// minus 1 in index as we start form 0
		mid_point := float64(size_of_array/2) + 0.5
		// fmt.Println("CHOSEN_INDEX: ", int(mid_point))
		return fmt.Sprintf("%d", array_to_calculate[int(mid_point)])
	}
	// If even
	// Take len/2 and len/2-1
	// fmt.Println("CHOSEN_INDEX: ", size_of_array/2, size_of_array/2-1)
	total_sum := array_to_calculate[size_of_array/2] + array_to_calculate[size_of_array/2-1]
	// DEBUG:
	// fmt.Println("SUM: ", total_sum)
	raw_median := float64(total_sum) / 2
	// DEBUG:
	// fmt.Println("RAW_MEDIAN: ", raw_median)
	raw_median_string := fmt.Sprintf("%.1f", raw_median)
	match, _ := regexp.MatchString("^.*\\.0$", raw_median_string)
	if match {
		return fmt.Sprintf("%d", int(raw_median))
	}

	return raw_median_string
}

// From Andrew Gerrand
// https://groups.google.com/forum/#!topic/golang-nuts/FT7cjmcL7gw
// A data structure to hold a key/value pair.
// a i++ was left out .. bug to be fixed :P
type Pair struct {
	Key   int
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int      { return len(p) }

// Use more than to reverse sort it ...
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[int]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		// DEBUG:
		// fmt.Println("K: ", k, " V: ", v)
		p[i] = Pair{k, v}
		i++
	}
	// DEBUG:
	// fmt.Println("PAIRLIST: ", p)
	sort.Sort(p)
	return p
}

func CalculateMode(array_to_calculate []int) string {
	// Use map to
	frequency_counter := make(map[int]int)
	lowest_value := 0
	// Count as we go through
	for _, myval := range array_to_calculate {
		frequency_counter[myval]++
		if lowest_value == 0 {
			lowest_value = myval
		} else {
			if myval < lowest_value {
				lowest_value = myval
			}
		}
	}
	// Used to break tie
	// DEBUG:
	// fmt.Println("LOWEST_VALUE: ", lowest_value)

	sorted_pairlist_byval := sortMapByValue(frequency_counter)
	// DEBUG:
	// fmt.Println("FREQ: ", sorted_pairlist_byval)
	if sorted_pairlist_byval[0].Value > sorted_pairlist_byval[1].Value {
		// Pick first
		// DEBUG:
		// fmt.Println("WINNER: ", sorted_pairlist_byval[0].Key)
		return fmt.Sprintf("%d", sorted_pairlist_byval[0].Key)
	}
	// DEBUG:
	// fmt.Println("WINNER_BY_LOWEST: ", lowest_value)

	return fmt.Sprintf("%d", lowest_value)
}
