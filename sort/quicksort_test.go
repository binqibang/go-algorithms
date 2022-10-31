package sort

import (
	"testing"
)

func TestQuicksort(t *testing.T) {
	nums := []int{1, 4, 7, 9, 5, 3, 8}
	t.Log(nums)
	Quicksort(nums)
	t.Log(nums)
}
