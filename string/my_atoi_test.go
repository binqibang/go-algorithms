package string

import (
	"testing"
)

func TestMyAtoi(t *testing.T) {
	s := "  -429"
	t.Log(MyAtoi(s))
}
