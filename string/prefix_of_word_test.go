package string

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsPrefixOfWord(t *testing.T) {
	sentence := "i love eating burger"
	searchWord := "burg"
	res := IsPrefixOfWord(sentence, searchWord)
	require.Equal(t, res, 4)
	t.Log(res)
}
