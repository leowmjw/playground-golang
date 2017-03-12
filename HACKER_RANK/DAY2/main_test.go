package main

import (
	"testing"
	"fmt"
)

func TestCalculateTip(t *testing.T) {
	// Tip is 20%
	mealCost := 12.00
	tipPercent := 20
	tipCost := CalculateTip(mealCost, tipPercent)
	if tipCost != 2.4 {
		fmt.Println("tipCost: ", tipCost)
		t.Fail()
	}
	fmt.Println("OK!")
}

func TestCalculateTax(t *testing.T) {
	// Tax is 8%
	mealCost := 12.00
	taxPercent := 8
	taxCost := CalculateTax(mealCost, taxPercent)
	if taxCost != 0.96 {
		fmt.Println("taxCost: ", taxCost)
		t.Fail()
	}
	fmt.Println("OK!")
}

func TestRoundNearestInteger(t *testing.T) {
	totalMealCost := 15.36
	roundedTotalMealCost := RoundNearestInteger(totalMealCost)
	if roundedTotalMealCost != 15 {
		fmt.Println("EXPECTED 15, got ", roundedTotalMealCost)
		t.Fail()
	}
	totalMealCost = 99.88
	roundedTotalMealCost = RoundNearestInteger(totalMealCost)
	if roundedTotalMealCost != 100 {
		fmt.Println("EXPECTED 100, got ", roundedTotalMealCost)
		t.Fail()
	}
	fmt.Println("OK!")
}

func TestCalculateTotal(t *testing.T) {
	mealCost := 12.00
	tipPercent := 20
	taxPercent := 8
	totalCost := CalculateTotalCost(mealCost, tipPercent, taxPercent)
	if totalCost != 15 {
		t.Fail()
	}

	fmt.Println("OK!")
}
