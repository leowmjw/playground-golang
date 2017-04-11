package main

import (
	"fmt"
	"testing"
)

func TestIntegerInputs(t *testing.T) {
	myval, mytype := ProcessInput("123")
	// fmt.Println("VAL: ", myval, " TYPE: ", mytype)
	if mytype == "uint64" {
		myintval := myval.(uint64)
		// DEBUG
		fmt.Println("VALUE: ", myintval)
	}

}

func TestFloatInputs(t *testing.T) {
	myval, mytype := ProcessInput("123.76")
	if mytype == "float64" {
		myfloatval := myval.(float64)
		// DEBUG
		fmt.Println("VALUE: ", myfloatval)
	}
}

func TestStringInputs(t *testing.T) {
	myval, mytype := ProcessInput("bob")
	if mytype == "string" {
		mystringval := myval.(string)
		// DEBUG
		fmt.Println("VALUE: ", mystringval)
	}
}

func TestAdd(t *testing.T) {
	// Scenario #1
	it := integertype{modifier: 4, input: 2}
	result := it.Add()
	// DEBUG:
	fmt.Println("RESULT:", result)
	if result == "6" {
		fmt.Println("Scenario #1 OK!")
	} else {
		t.Fail()
	}
	// Scenario #2: Negatives
	it2 := integertype{modifier: -5, input: 3}
	result2 := it2.Add()
	//	fmt.Println("RESULT:", result2)
	if result2 == "-2" {
		fmt.Println("Scenario #2 OK!")
	} else {
		t.Fail()
	}
	// Scenario #3: Float
	ft := floattype{modifier: 4.0, input: 2.456}
	result3 := ft.Add()
	//	fmt.Println("RESULT:", result3)
	if result3 == "6.5" {
		fmt.Println("Scenario #3 OK!")
	} else {
		t.Fail()
	}
	// Scenario #4: String
	st := stringtype{modifier: "Goodbye", input: "    Cruel Wor  ld!!!"}
	result4 := st.Add()
	//	fmt.Println("RESULT:", result4)
	if result4 == "Goodbye    Cruel Wor  ld!!!" {
		fmt.Println("Scenario #4 OK!")
	} else {
		t.Fail()
	}
}
