package main

import "fmt"

func main() {
	number_of_integers := 0
	fmt.Scanf("%d", &number_of_integers)
	fmt.Println(RecursiveCallStack(number_of_integers))
}

func RecursiveCallStack(number_of_integers int) string {
	if number_of_integers == 2 {
		// End of stack; start to unwind ...
		return "min(int, int)"
	}

	if number_of_integers < 2 {
		// Soething wrong here ...
		return "bob"
	}
	// Get ready to be recursive; pop out the stack
	number_of_integers = number_of_integers - 1
	return fmt.Sprintf("min(int, %s)", RecursiveCallStack(number_of_integers))
}
