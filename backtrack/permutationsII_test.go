package backtrack

import "testing"

func TestPermuteUnique(t *testing.T) {
	nums := []int{3, 3, 0, 3}
	t.Log(PermuteUnique(nums))
}
