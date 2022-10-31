// Package string
// @summary: LeetCode questions about "string"
// @author: binqibang
// @created: 2022-07-30 11:16
package string

// ShortestToChar LeetCode #821 (Easy)
func ShortestToChar(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)

	idx := -n
	for i := 0; i < n; i++ {
		if s[i] == c {
			idx = i
		}
		ans[i] = i - idx
	}

	idx = 2 * n
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			idx = i
		}
		ans[i] = min(ans[i], idx-i)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
