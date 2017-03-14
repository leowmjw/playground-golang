package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRecursiveCall(t *testing.T) {
	expectedStackCallSimple := "min(int, int)"
	assert.Equal(t, expectedStackCallSimple, RecursiveCallStack(2), "Simple Stack Call ..")

	expectedStackCallComplex := "min(int, min(int, min(int, min(int, int))))"
	assert.Equal(t, expectedStackCallComplex, RecursiveCallStack(5), "Complex Stack Call")
}
