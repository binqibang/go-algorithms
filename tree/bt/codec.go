package bt

import (
	"container/list"
	"strconv"
	"strings"
)

// Codec
// @summary: LeetCode #297 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/9/2 10:45
type Codec struct {
}

func CodecConstructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "nil"
	}
	return strconv.Itoa(root.Val) + "," + c.serialize(root.Left) + "," + c.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	arr := strings.Split(data, ",")
	nodes := list.New()
	for _, s := range arr {
		nodes.PushBack(s)
	}
	return dfs(nodes)
}

func dfs(nodes *list.List) *TreeNode {
	cur := nodes.Front()
	if cur.Value.(string) == "nil" {
		nodes.Remove(cur)
		return nil
	}
	val, _ := strconv.Atoi(cur.Value.(string))
	node := &TreeNode{Val: val}
	nodes.Remove(cur)
	node.Left = dfs(nodes)
	node.Right = dfs(nodes)
	return node
}
