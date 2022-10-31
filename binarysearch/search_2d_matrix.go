package binarysearch

// SearchMatrix
// @summary: LeetCode #74 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/10/19 14:49
// @param matrix integers in each row are sorted from left to right, and the
//               first integer of each row is greater than the last integer
//               of the previous row.
func SearchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	low, high := 0, m*n-1
	for low <= high {
		mid := low + (high-low)/2
		// [i, j] -> i * n + j
		// x -> [x / n, x % n]
		cur := matrix[mid/n][mid%n]
		if cur == target {
			return true
		} else if cur < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}
