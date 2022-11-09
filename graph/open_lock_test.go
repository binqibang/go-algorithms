package graph

import "testing"

func TestOpenLock(t *testing.T) {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	t.Log(OpenLock(deadends, target))
}

func TestGetNextCodes(t *testing.T) {
	code := "0000"
	t.Log(GetNextCodes(code))

	set := map[string]bool{}
	set["test"] = true
	t.Log(set["test"], set["0"])
}
