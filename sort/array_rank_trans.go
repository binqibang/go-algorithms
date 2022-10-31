package sort

import "sort"

// ArrayRankTransform
// @summary: LeetCode #1331 (Easy)
// @author : binqibang
// @create : 2022-07-28 11:03
func ArrayRankTransform(arr []int) []int {
	// 复制 arr 到 sorted，且改变 sorted 的值不会影响原切片
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)
	ranks := map[int]int{}
	for _, num := range sorted {
		_, contains := ranks[num]
		if !contains {
			ranks[num] = len(ranks) + 1
		}
	}
	for i, v := range arr {
		arr[i] = ranks[v]
	}
	return arr
}
