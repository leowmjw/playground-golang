package main

import (
	"fmt"
	"github.com/runningwild/sorts/smooth"
	"github.com/davecgh/go-spew/spew"
	"sort"
)

type SingleSegment struct {
	ballPosition    int
	gapSize         int
	previousSegment int
	nextSegment     int
}

type ContiguousSegments struct {
	segments []SingleSegment
}

func main() {
	fmt.Println("Welcome to Day3!!!")
}

func SortBalls(arrayBalls []int) []int {
	sortedBalls := arrayBalls
	a := smooth.IntSlice(sortedBalls[0:])
	smooth.Sort(a)

	// DEBUG:
	// fmt.Println("SORTED: ", sortedBalls)
	return sortedBalls
	// return []int{1, 4, 6, 7, 10}
}

func FindContiguousSegments(sortedBalls []int) []SingleSegment {
	// currentSegmentSliceStart := 0
	currentSegmentPosition := -1
	segments := []SingleSegment{}
	for i := 0; i < len(sortedBalls); i++ {
		if currentSegmentPosition == -1 {
			// at the start; set it to current value
			currentSegmentPosition = sortedBalls[i]
		} else {
			// Now test if the previousSegmentPosition is adjacent to current Ball
			// if yes; move up previousSegmentPosition
			if (currentSegmentPosition+1 == sortedBalls[i]) {
				currentSegmentPosition++
			} else {
				singleSegment := SingleSegment{previousSegment: -1, nextSegment: -1}
				singleSegment.ballPosition = currentSegmentPosition
				// if no; create a new Segment; fill with slice of the start of the previous Segment?
				// calculate Gap!
				singleSegment.gapSize = sortedBalls[i] - 1 - currentSegmentPosition
				// DEBUG
				// fmt.Println("INDEX: ", i, " DATA: ", sortedBalls[i])
				// if contiguous; extend the edgePosition
				// else calculate Gap and append a new item
				// FoundGap?
				// append if gap is equal/higher current
				// segments = append(segments, singleSegment)
				// prepend
				// segments = append([]SingleSegment{SingleSegment{previousSegment: -1, nextSegment: -1}}, segments ...)
				// update the currentSegmentPosition
				segments = SortedInsert(segments, singleSegment)
				currentSegmentPosition = sortedBalls[i]
			}

		}
	}

	// DEBUG:
	// fmt.Println(fmt.Sprintf("SEGMENTS: %#v", segments))
	// fmt.Println(fmt.Sprintf("SEGMENTS: %v", segments))
	spew.Dump(segments)

	return segments
}

func CalculateGapBetweenSegments(sortedBalls *[]int) int {
	return 0
}

func (s *SingleSegment) Less(f SingleSegment) bool {
	fmt.Println("NODE_GAP: ", s.gapSize, " INSERT_GAP: ", f.gapSize)
	if s.gapSize < f.gapSize {
		return true
	}
	return false
}

func SortedInsert(s []SingleSegment, f SingleSegment) []SingleSegment {
	l := len(s)
	if l == 0 {
		return append([]SingleSegment{}, f)
	}
	i := 0
	// Iterate over until find a node where Gap is mor than it!
	for ; i < len(s); i++ {
		if s[i].Less(f) {
			break
		}
	}
	fmt.Println("LENGTH: ", l, " FOUND_NODEINDEX: ", i)
	if i == l { // not found = new value is the smallest
		return append(s[0:l], f)
	}
	if i == 0 { // new value is the biggest
		return append([]SingleSegment{f}, s...)
	}
	sam := append([]SingleSegment{}, s[i:]...)
	fmt.Println(fmt.Sprintf("SLICE: %v", sam))
	// What's up with the belo; why not working??
	// return append(s[0:l], f, s[l+1:]...)
	bob := append(s[0:i], f)
	// DEBUG:
	fmt.Println(fmt.Sprintf("BOB: %v", bob))
	fmt.Println(fmt.Sprintf("SAM_SPREAD: %v", sam))
	return append(bob, sam...)
}

// https://gist.github.com/zhum/57cb45d8bbea86d87490
func OldSortedInsert(s []SingleSegment, f SingleSegment) []SingleSegment {
	l := len(s)
	if l == 0 {
		segments := []SingleSegment{}
		return append(segments, f)
	}

	i := sort.Search(l,
		func(i int) bool {
			return s[i].Less(f)
		})
	fmt.Println("LENGTH: ", l, " INDEX: ", i)
	if i == l { // not found = new value is the smallest
		return append(s[0:l], f)
	}
	if i == 0 { // new value is the biggest
		return append([]SingleSegment{f}, s...)
	}
	// What's up with the belo; why not working??
	// return append(s[0:l], f, s[l+1:]...)
	bob := append(s[0:i], f)
	// DEBUG:
	fmt.Println(fmt.Sprintf("BOB: %v", bob))
	fmt.Println(fmt.Sprintf("SLICE: %v", s[l:]))
	return append(bob, s[i+1:]...)
}
