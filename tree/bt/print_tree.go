package bt

import "strconv"

// PrintTree
// @summary: LeetCode #655 (Medium)
// @author : binqibang
// @created: 2022/8/22 16:35
func PrintTree(root *TreeNode) [][]string {
	m := getHeight(root)
	n := 1<<m - 1
	// Create multi-dimensional slice.
	ans := make([][]string, m)
	for i := range ans {
		ans[i] = make([]string, n)
	}
	fill(root, ans, 0, 0, n)
	return ans
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func fill(root *TreeNode, matrix [][]string, level, l, r int) {
	if root == nil {
		return
	}
	matrix[level][(l+r)/2] = strconv.Itoa(root.Val)
	fill(root.Left, matrix, level+1, l, (l+r)/2)
	fill(root.Right, matrix, level+1, (l+r+1)/2, r)
}
