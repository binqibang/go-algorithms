package binarysearch

// SearchMatrixII
// @summary: LeetCode #240 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/10/19 14:49
// @param matrix Integers in each row are sorted in ascending from
//               left to right. Integers in each column are sorted
//               in ascending from top to bottom.
func SearchMatrixII(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1
	for x < m && y >= 0 {
		cur := matrix[x][y]
		if cur == target {
			return true
		} else if cur < target {
			x++
		} else {
			y--
		}
	}
	return false
}
