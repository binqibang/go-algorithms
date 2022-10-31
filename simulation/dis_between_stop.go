package simulation

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// DistanceBetweenBusStops
// @summary: LeetCode #1184 (Easy)
// @author : binqibang
// @created: 2022-07-24 10:53
func DistanceBetweenBusStops(distance []int, start int, destination int) int {
	// 顺时针距离
	var cwDis = 0
	// 逆时针距离
	var ccwDis = 0
	if start > destination {
		start, destination = destination, start
	}
	// 遍历一次，顺向、逆向距离互补
	for i, d := range distance {
		if start <= i && i < destination {
			cwDis += d
		} else {
			ccwDis += d
		}
	}
	return min(cwDis, ccwDis)
}
