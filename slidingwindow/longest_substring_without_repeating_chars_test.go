package slidingwindow

import (
	"go-algorithms/utils"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := utils.RandomString(50)
	t.Log(s)
	t.Log(LengthOfLongestSubstring(s))
}
