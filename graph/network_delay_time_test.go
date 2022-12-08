package graph

import "testing"

func TestNetworkDelayTime(t *testing.T) {
	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	n, k := 4, 2
	t.Log(NetworkDelayTime(times, n, k))
}
