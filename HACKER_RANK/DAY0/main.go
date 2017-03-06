package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	readText := ""
	for scanner.Scan() {
		readText = scanner.Text()
		// fmt.Println(readText)
		break
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error", err)
	} else {
		inputString := "Welcome to 30 Days of Code!"
		if readText != "" {
			inputString = readText
		}
		fmt.Println("Hello, World.")
		fmt.Println(inputString)
		//	println("Hello, World.")
		//	println(inputString)
	}
}
