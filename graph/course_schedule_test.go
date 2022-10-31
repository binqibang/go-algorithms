package graph

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCanFinish(t *testing.T) {
	numCourses := 2
	prerequisites := [][]int{{1, 0}}
	require.Equal(t, true, CanFinish(numCourses, prerequisites))
}
