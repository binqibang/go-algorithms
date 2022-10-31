package bt

// @summary: LeetCode #623 (Medium)
// @author: binqibang
// @create: 2022-08-05

// AddOneRow DFS
func AddOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return nil
	}
	if depth == 1 {
		return &TreeNode{val, root, nil}
	}
	if depth == 2 {
		root.Left = &TreeNode{val, root.Left, nil}
		root.Right = &TreeNode{val, nil, root.Right}
	} else {
		root.Left = AddOneRow(root.Left, val, depth-1)
		root.Right = AddOneRow(root.Right, val, depth-1)
	}
	return root
}

// AddOneRow1 BFS
func AddOneRow1(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{val, root, nil}
	}
	queue := []*TreeNode{root}
	curDepth := 1
	for curDepth < depth-1 {
		n := len(queue)
		for i := 0; i < n; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		curDepth++
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		cur.Left = &TreeNode{val, cur.Left, nil}
		cur.Right = &TreeNode{val, nil, cur.Right}
	}
	return root
}
