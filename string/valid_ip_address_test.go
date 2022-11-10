package string

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidIPAddress(t *testing.T) {
	ip := "192.168.10.12."
	ip1 := "2001:0db8:85a3:0:0:8A2E:0370:7334:"
	t.Log(ValidIPAddress(ip))
	require.Equal(t, "Neither", ValidIPAddress(ip1))
}
