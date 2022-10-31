package math

// ValidSquare
// @summary: LeetCode #593 (Medium)
// @author : binqibang
// @created: 2022-07-29 10:58
func ValidSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	if isPointEqual(p1, p2) || isPointEqual(p1, p3) || isPointEqual(p1, p4) ||
		isPointEqual(p2, p3) || isPointEqual(p2, p4) || isPointEqual(p3, p4) {
		return false
	}
	normalCount := isNormal(p1, p2, p3) + isNormal(p2, p3, p4) +
		isNormal(p1, p3, p4) + isNormal(p1, p2, p4)
	if normalCount < 4 {
		return false
	}
	return isAdjEqual(p1, p2, p3)
}

func isNormal(p1, p2, p3 []int) int {
	v1 := []int{p2[0] - p1[0], p2[1] - p1[1]}
	v2 := []int{p3[0] - p1[0], p3[1] - p1[1]}
	v3 := []int{p3[0] - p2[0], p3[1] - p2[1]}
	if v1[0]*v2[0]+v1[1]*v2[1] == 0 {
		return 1
	} else if v1[0]*v3[0]+v1[1]*v3[1] == 0 {
		return 1
	} else if v2[0]*v3[0]+v2[1]*v3[1] == 0 {
		return 1
	} else {
		return 0
	}
}

func isAdjEqual(p1, p2, p3 []int) bool {
	l1 := (p2[0]-p1[0])*(p2[0]-p1[0]) + (p2[1]-p1[1])*(p2[1]-p1[1])
	l2 := (p3[0]-p1[0])*(p3[0]-p1[0]) + (p3[1]-p1[1])*(p3[1]-p1[1])
	l3 := (p3[0]-p2[0])*(p3[0]-p2[0]) + (p3[1]-p2[1])*(p3[1]-p2[1])
	maxLen := max(l1, max(l2, l3))
	if l1 == maxLen {
		return l2 == l3
	} else if l2 == maxLen {
		return l1 == l3
	} else {
		return l1 == l2
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func isPointEqual(p1, p2 []int) bool {
	return p1[0] == p2[0] && p1[1] == p2[1]
}
