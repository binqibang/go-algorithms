package binarysearch

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{3, 1}
	target := 1
	res := Search(nums, target)
	t.Log(res)
}
