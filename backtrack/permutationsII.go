package backtrack

import "sort"

// PermuteUnique
// @summary: LeetCode #47 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/11/1 10:00
func PermuteUnique(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums))
	n := len(nums)
	sort.Ints(nums)

	var backtrack func(depth int, perm []int)
	backtrack = func(depth int, perm []int) {
		if depth == n {
			res = append(res, append([]int{}, perm...))
			return
		}
		for i := 0; i < n; i++ {
			if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
				continue
			}
			perm = append(perm, nums[i])
			used[i] = true
			backtrack(depth+1, perm)
			used[i] = false
			perm = perm[:len(perm)-1]
		}
	}

	backtrack(0, []int{})
	return res
}
