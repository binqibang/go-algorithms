package bt

// DiameterOfBinaryTree
// @summary: LeetCode #543 (Easy); CodeTop MS
// @author : binqibang
// @created: 2022/10/6 10:29
func DiameterOfBinaryTree(root *TreeNode) int {
	var diameter int
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// depth of left or right subtree
		ld, rd := depth(node.Left), depth(node.Right)
		// pathLen = depth(lc) + depth(rc)
		diameter = max(ld+rd, diameter)
		// depth of current root
		return max(ld, rd) + 1
	}
	depth(root)
	return diameter
}
