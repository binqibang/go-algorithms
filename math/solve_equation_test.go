package math

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolveEquation(t *testing.T) {
	e := "x+5-3+x=6+x-2"
	res := SolveEquation(e)

	require.Equal(t, "x=2", res)
	t.Log(res)
}
