package bt

// DeleteNode
// @summary: LeetCode #450 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/9/15 10:13
func DeleteNode(root *TreeNode, key int) *TreeNode {
	// empty bst
	if root == nil {
		return nil
	}

	// key node in the left subtree
	if key < root.Val {
		root.Left = DeleteNode(root.Left, key)
	}

	// key node in the right subtree
	if key > root.Val {
		root.Right = DeleteNode(root.Right, key)
	}

	// find the key node
	if key == root.Val {
		// leaf node, just delete it
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// left subtree is empty, return its right subtree as new subtree
		if root.Left == nil {
			return root.Right
		}
		// right subtree is empty, return its left subtree as new subtree
		if root.Right == nil {
			return root.Left
		}
		// if both left and right subtree is not empty, then find the max node of its
		// left subtree to replace (or min node of its right subtree)
		pre := root
		successor := root.Left
		for successor.Right != nil {
			pre = successor
			successor = successor.Right
		}
		root.Val = successor.Val
		if pre.Right.Val == successor.Val {
			pre.Right = successor.Left
		} else { // successor has only left subtree
			pre.Left = successor.Left
		}
	}
	return root
}
