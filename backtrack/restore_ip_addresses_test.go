package backtrack

import "testing"

func TestRestoreIpAddresses(t *testing.T) {
	s := "25525511135"
	res := RestoreIpAddresses(s)
	t.Log(res)
}
