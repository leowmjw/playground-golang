package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSortBalls(t *testing.T) {
	testCase := []int{6, 4, 1, 7, 10}
	expectedSorted := []int{1, 4, 6, 7, 10}
	myres := SortBalls(testCase)
	assert.Equal(t, expectedSorted, myres, "Must be same!!")
}

func TestFindContiguousSegments(t *testing.T) {
	expectedContiguousSegments :=
		[]SingleSegment{
			{ballPosition: 1, gapSize: 2, previousSegment: -1, nextSegment: -1},
			{ballPosition: 7, gapSize: 2, previousSegment: -1, nextSegment: -1},
			{ballPosition: 4, gapSize: 1, previousSegment: -1, nextSegment: -1},
		}

	continuousSegments := FindContiguousSegments([]int{1, 4, 6, 7, 10})
	assert.Equal(t, expectedContiguousSegments, continuousSegments, "Must be same!!!")
}

func TestSortedInsert(t *testing.T) {
	// Needs to match below
	expectedSortedSegments := []SingleSegment{
		{ballPosition: 1, gapSize: 2, previousSegment: -1, nextSegment: -1},
		{ballPosition: 7, gapSize: 2, previousSegment: -1, nextSegment: -1},
		{ballPosition: 4, gapSize: 1, previousSegment: -1, nextSegment: -1},
	}
	// Execute
	sortedSegments := SortedInsert([]SingleSegment{
		{ballPosition: 1, gapSize: 2, previousSegment: -1, nextSegment: -1},
		{ballPosition: 4, gapSize: 1, previousSegment: -1, nextSegment: -1},
	}, SingleSegment{7, 2, -1, -1})
	// Does it?
	assert.Equal(t, expectedSortedSegments, sortedSegments, "Sorted Segmments should match up!!")
}
