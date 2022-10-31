package bt

// @summary: Lowest Common Ancestor problems.
// @author : binqibang
// @created: 2022/8/25 11:15

// LowestCommonAncestor LeetCode #236 (Medium)
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

var count = 0

// LowestCommonAncestorII LeetCode #1644 (Medium)
func LowestCommonAncestorII(root, p, q *TreeNode) *TreeNode {
	res := lca(root, p, q)
	if count == 2 {
		return res
	}
	return nil
}

func lca(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		count++
		return root
	}
	left := lca(root.Left, p, q)
	right := lca(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
