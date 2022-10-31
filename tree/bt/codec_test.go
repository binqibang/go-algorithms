package bt

import (
	"testing"
)

func TestCodec(t *testing.T) {
	c := CodecConstructor()
	s := c.serialize(&bt)
	t.Log(s)
	res := c.deserialize(s)
	matrix := PrintTree(res)
	for _, row := range matrix {
		t.Log(row)
	}
}
