package utils

import "testing"

func TestRandomString(t *testing.T) {
	t.Logf("%d %b %x", ' ', ' ', ' ')
	t.Logf("%d %b %x", '~', '~', '~')
	t.Log(RandomString(50))
}
