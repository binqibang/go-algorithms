package bt

// ZigzagLevelOrder
// @summary: LeetCode #103 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/9/9 10:55
func ZigzagLevelOrder(root *TreeNode) [][]int {
	var order [][]int
	if root == nil {
		return order
	}
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		var row []int
		length := len(queue)
		for i := 0; i < length; i++ {
			cur := queue[0]
			queue = queue[1:]
			row = append(row, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		// 奇数层翻转顺序
		if level%2 == 1 {
			for i, n := 0, len(row); i < n/2; i++ {
				row[i], row[n-1-i] = row[n-1-i], row[i]
			}
		}
		order = append(order, row)
	}
	return order
}
