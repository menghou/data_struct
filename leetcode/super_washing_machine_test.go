package leetcode

import (
	"testing"
)

func TestFindMinMoves(t *testing.T) {
	input := []int{4,0,0,4}
	result := findMinMoves(input)
	t.Log(result)
}
