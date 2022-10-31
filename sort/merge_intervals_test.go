package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestMerge(t *testing.T) {
	intervals := [][]int{{1, 3}, {5, 10}, {2, 6}, {15, 18}}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := Merge(intervals)
	fmt.Println(res)
}
