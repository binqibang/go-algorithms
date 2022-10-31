package bt

// CBTInserter
// @summary: LeetCode #919 (Medium)
// @author: binqibang
// @create: 2022-07-25 09:22
type CBTInserter struct {
	root      *TreeNode
	candidate []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	queue := []*TreeNode{root}
	var candidate []*TreeNode
	for len(queue) > 0 {
		// 模拟出队
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left == nil || node.Right == nil {
			candidate = append(candidate, node)
		}
	}
	return CBTInserter{root, candidate}
}

func (i *CBTInserter) Insert(val int) int {
	child := &TreeNode{val, nil, nil}
	cur := i.candidate[0]
	if cur.Left == nil {
		cur.Left = child
	} else {
		cur.Right = child
		// Go Slice 没有删除某元素的方法，只能用切片索引代替
		i.candidate = i.candidate[1:]
	}
	i.candidate = append(i.candidate, child)
	return cur.Val
}

func (i *CBTInserter) GetRoot() *TreeNode {
	return i.root
}
