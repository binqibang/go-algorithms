package hashmap

import "testing"

func TestCountSubarrays(t *testing.T) {
	nums := []int{3, 2, 1, 4, 5}
	k := 4
	t.Log(CountSubarrays(nums, k))
}
