package bt

var a = TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
var b = TreeNode{Val: 3, Left: &TreeNode{Val: 6}}
var bt = TreeNode{Val: 1, Left: &a, Right: &b}

var c = TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}
var d = TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 8}}
var bst = TreeNode{Val: 5, Left: &c, Right: &d}
