package string

import "testing"

func TestMatchString(t *testing.T) {
	words := []string{"mass", "as", "hero", "superhero"}
	ans := MatchString(words)
	t.Log(ans)
}
