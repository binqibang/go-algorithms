package binarysearch

import "testing"

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 4
	t.Log(SearchMatrix(matrix, target))
}
