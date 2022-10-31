package bt

import "testing"

func TestPrintTree(t *testing.T) {
	matrix := PrintTree(&bt)
	for _, row := range matrix {
		t.Log(row)
	}
}
