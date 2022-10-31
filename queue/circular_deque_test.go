package queue

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMyCircularDeque(t *testing.T) {
	circularDeque := CircularDeque(3)

	require.Equal(t, circularDeque.InsertLast(1), true)
	require.Equal(t, circularDeque.InsertLast(2), true)
	require.Equal(t, circularDeque.InsertFront(3), true)
	require.Equal(t, circularDeque.InsertFront(4), false)

	require.Equal(t, circularDeque.GetRear(), 2)
	require.Equal(t, circularDeque.IsFull(), true)

	require.Equal(t, circularDeque.DeleteLast(), true)
	require.Equal(t, circularDeque.InsertFront(4), true)
	require.Equal(t, circularDeque.GetFront(), 4)
}
