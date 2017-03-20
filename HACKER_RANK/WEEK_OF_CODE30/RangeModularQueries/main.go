package main

import "fmt"

func main() {

	// The first line contains two space-separated integers describing the
	// respective values of  n(the size of A) and  q(the number of queries).
	size_of_array := 0
	number_of_queries := 0
	fmt.Scanf("%d %d", &size_of_array, &number_of_queries)
	// The second line has  space-separated integers describing the respective values of .
	original_array := make([]int, size_of_array)
	for iter := 0; iter < size_of_array; iter++ {
		fmt.Scanf("%d", &original_array[iter])
	}
	// Get all the subsequent lines; mentioned by number_of_queries
	for qiter := 0; qiter < number_of_queries; qiter++ {
		left_start := 0
		right_end := 0
		mod_value := 0
		remain_value := 0
		fmt.Scanf("%d %d %d %d", &left_start, &right_end, &mod_value, &remain_value)
		fmt.Println(
			CountModulo(
				ChooseArraySegment(original_array, left_start, right_end),
				mod_value,
				remain_value))
	}
}

func ChooseArraySegment(original_array []int, left int, right int) (array_to_check []int) {
	size_to_extract := right + 1
	return original_array[left:size_to_extract]
}

func CountModulo(array_to_check []int, mod_value int, remain_value int) int {
	number_of_matches := 0

	for _, myval := range array_to_check {
		// DEBUG:
		// fmt.Println("MOD: ", myval%mod_value)
		if myval%mod_value == remain_value {
			number_of_matches++
		}
	}
	// DEBUG:
	// fmt.Println("MATCHES: ", number_of_matches)
	return number_of_matches
}
