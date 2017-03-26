package main

import "fmt"

func main() {
	// The first line contains an integer, , denoting the number of elements in arrays  and .
	size_of_array := 0
	fmt.Scanf("%d", &size_of_array)
	// The second line contains  space-separated integers describing the respective elements of array .
	x := make([]int, size_of_array)
	for iter := 0; iter < size_of_array; iter++ {
		fmt.Scanf("%d", &x[iter])
	}
	// The third line contains  space-separated integers describing the respective elements of array .
	w := make([]int, size_of_array)
	for iter := 0; iter < size_of_array; iter++ {
		fmt.Scanf("%d", &w[iter])
	}

	// Output
	fmt.Println(CalculateWeightedMean(x, w))
}

func CalculateWeightedMean(x []int, w []int) string {
	total_weighted_sum := 0.0
	sum_of_weights := 0.0
	for i, single_data := range x {
		total_weighted_sum += float64(single_data * w[i])
		sum_of_weights += float64(w[i])
	}

	return fmt.Sprintf("%.1f", total_weighted_sum/sum_of_weights)
}
