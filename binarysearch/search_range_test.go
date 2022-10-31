package binarysearch

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	ans := []int{3, 4}
	require.EqualValues(t, ans, SearchRange(nums, target))
}
