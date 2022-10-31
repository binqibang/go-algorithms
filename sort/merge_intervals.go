package sort

import "sort"

// Merge
// @summary: LeetCode #56 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/9/13 11:14
func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var merged [][]int
	for i := 0; i < len(intervals); i++ {
		l, r, n := intervals[i][0], intervals[i][1], len(merged)
		if n == 0 || l > merged[n-1][1] {
			merged = append(merged, []int{l, r})
		} else {
			merged[n-1][1] = max(r, merged[n-1][1])
		}
	}
	return merged
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
