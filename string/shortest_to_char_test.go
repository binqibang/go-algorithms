package string

import "testing"

func TestShortestToChar(t *testing.T) {
	s := "iloveleetcode"
	c := byte('e')
	output := ShortestToChar(s, c)
	t.Log(output)
}
