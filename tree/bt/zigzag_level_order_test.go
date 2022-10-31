package bt

import "testing"

func TestZigzagLevelOrder(t *testing.T) {
	order := ZigzagLevelOrder(&bt)
	t.Log(order)
}
