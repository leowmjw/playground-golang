package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.
	var it uint64 = 0
	var dt float64 = 0.0
	var st string = ""

	// Init the first to check and then onlyscan through
	for scanner.Scan() {
		var myinput = scanner.Text()
		// fmt.Println("INPUT: ", myinput)
		// DEBUG:
		// fmt.Println("In ProcessInput")
		myuint64, myerr := strconv.ParseUint(myinput, 10, 64)
		if myerr == nil {
			// return myuint64, "uint64"
			// IsInt
			it = myuint64
			fmt.Println(strconv.FormatUint(i+it, 10))
		} else {
			_, myerr2 := strconv.ParseInt(myinput, 10, 64)
			if myerr2 == nil {
				// return myint64, "int64"
				// Do Nothing
			} else {
				myfloat64, myerr3 := strconv.ParseFloat(myinput, 64)
				if myerr3 == nil {
					// return myfloat64, "float64"
					// IsFloat
					dt = myfloat64
					fmt.Println(strconv.FormatFloat(d+dt, 'f', 1, 64))
				} else {
					// IsString
					st = myinput
					// Print and break out!!!
					fmt.Println(fmt.Sprintf("%s%s", s, st))
					break
				}

			}

		}

	}

}

func main_alt() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.
	it := integertype{modifier: int64(i)}
	dt := floattype{modifier: d}
	st := stringtype{modifier: s}

	// DEBUG ..
	// fmt.Println(fmt.Sprintf("it: %v dt: %v st: %v", it, dt, st))
	// Read and save an integer, double, and String to your variables.
	var outer = 0
	checkFor := "uint64"
	// Init the first to check and then onlyscan through
	for scanner.Scan() {
		var inner = 0
		var myinput = scanner.Text()
		// fmt.Println("INPUT: ", myinput)
		myval, mytype := ProcessInput(myinput)
		// fmt.Println("TYPE: ", mytype, "VAL:", myval)
		if mytype == checkFor {
			// Advance the state
			if checkFor == "uint64" {
				it.input = int64(myval.(uint64))
				checkFor = "float64"
			} else if checkFor == "float64" {
				dt.input = myval.(float64)
				checkFor = "string"
			} else if checkFor == "string" {
				st.input = myval.(string)
				break
			}
		}
		inner++
		outer++
		// DEBUG:
		// fmt.Println("Inner:", inner, " Outer:", outer)
	}
	// DEBUG:
	// fmt.Println("IT: ", it, "DT: ", dt, "ST: ", st)
	// Print the sum of both integer variables on a new line.
	it.Add()
	// Print the sum of the double variables on a new line.
	dt.Add()
	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	st.Add()
}

type integertype struct {
	modifier int64
	input    int64
}

type floattype struct {
	modifier float64
	input    float64
}

type stringtype struct {
	modifier string
	input    string
}

// ProcessInput ...
// Input:
// Extract out either Integer,Float, String??
func ProcessInput(input string) (value interface{}, inputtype string) {

	// DEBUG:
	// fmt.Println("In ProcessInput")
	myuint64, myerr := strconv.ParseUint(input, 10, 64)
	if myerr == nil {
		return myuint64, "uint64"
	}
	myint64, myerr2 := strconv.ParseInt(input, 10, 64)
	if myerr2 == nil {
		return myint64, "int64"
	}
	myfloat64, myerr3 := strconv.ParseFloat(input, 64)
	if myerr3 == nil {
		return myfloat64, "float64"
	}

	return input, "string"

}

// Add ...
// called by type integertype
func (it integertype) Add() string {
	return strconv.FormatInt(it.modifier+it.input, 10)
}

// Add ...
// called by type floattype
func (ft floattype) Add() string {
	// DEBUG:
	// myresult := ft.modifier + ft.input
	// fmt.Println("FLOAT: ", strconv.FormatFloat(myresult, 'f', 1, 64))
	return strconv.FormatFloat(ft.modifier+ft.input, 'f', 1, 64)
}

// Add ...
// called by type stringtype
func (st stringtype) Add() string {
	return fmt.Sprintf("%s%s", st.modifier, st.input)
}
