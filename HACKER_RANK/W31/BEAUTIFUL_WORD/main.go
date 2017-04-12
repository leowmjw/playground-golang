package main

import (
	"strings"
	"fmt"
)

func main() {

	fmt.Println("Beautiful Word ...")
}

type CurrentState struct {
	has_seen_vowel bool
	previous_char  string
}

func NewCurrentState(has_seen_vowel bool, previous_char string) (*CurrentState) {
	return &CurrentState{has_seen_vowel, previous_char}
}

func (this *CurrentState) FoundUgly(current_char string) bool {
	is_ugly := false

	// Check if is vowel and update
	is_vowel := strings.ContainsAny(current_char, "aeiouy")
	// if previous is vowel and this is vowel; isUGlY
	if this.has_seen_vowel && is_vowel {
		is_ugly = true
	}
	this.has_seen_vowel = is_vowel
	// if previous_char is same as current char; isUGLY
	if this.previous_char == current_char {
		is_ugly = true
	}
	// Update previous_char
	this.previous_char = current_char

	// if not found ugly; can move on ..
	return is_ugly
}

func Beautiful_Word(w string) string {
	// Split the string given ...
	split_word := strings.Split(w, "")
	// fmt.Println("SPLIT: ", split_word)
	current_state := NewCurrentState(false, "")
	for i := 0; i < len(split_word); i++ {
		// DEBUG:
		// fmt.Println("CURRENT:", split_word[i])
		has_found_ugly := current_state.FoundUgly(split_word[i])
		// DEBUG:
		// fmt.Println("HAS_FOUND_UGLY: ", has_found_ugly)
		if has_found_ugly {
			return "No"
		}
	}
	// Got this far without harm; you are perfect!!
	return "Yes"
}
