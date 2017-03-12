package main

import (
	"fmt"
)

func main() {
	mealCost:= 0.00
	tipPercent := 0
	taxPercent := 0
	fmt.Scanf("%f", &mealCost)
	fmt.Scanf("%d", &tipPercent)
	fmt.Scanf("%d", &taxPercent)
	// DEBUG:
	// fmt.Println("MEAL: ", mealCost, " TIP: ", tipPercent, " TAX: ", taxPercent)
	// fmt.Println("MEAL COST: ", CalculateTotalCost(mealCost, tipPercent, taxPercent))
	fmt.Println(fmt.Sprintf("The total meal cost is %d dollars.",CalculateTotalCost(mealCost, tipPercent, taxPercent) ))
}

func CalculateTip(mealCost float64, tipPercent int) float64 {
	return mealCost * float64(tipPercent) / 100
}

func CalculateTax(mealCost float64, taxPercent int) float64 {
	return mealCost * float64(taxPercent) / 100
}

func CalculateTotalCost(mealCost float64, tipPercent int, taxPercent int) int {
	return RoundNearestInteger(mealCost + CalculateTip(mealCost, tipPercent) + CalculateTax(mealCost, taxPercent))
}

func RoundNearestInteger(totalMealCost float64) int {
	return int(totalMealCost + 0.5)
}