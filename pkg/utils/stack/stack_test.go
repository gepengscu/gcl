package stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPush(t *testing.T) {
	s := &stack[int]{}
	require.Nil(t, s.root)
	s.Push(1)
	require.NotNil(t, s.root)
}

func TestPop(t *testing.T) {
	s := &stack[int]{}
	require.Nil(t, s.root)
	s.Push(1)
	s.Push(2)
	require.Equal(t, 2, s.Pop())
	require.Equal(t, 1, s.Pop())
}

func TestIsEmpty(t *testing.T) {
	s := &stack[int]{}
	require.True(t, s.IsEmpty())
	s.Push(1)
	require.False(t, s.IsEmpty())
	s.Push(2)
	require.False(t, s.IsEmpty())
	s.Pop()
	require.False(t, s.IsEmpty())
	s.Pop()
	require.True(t, s.IsEmpty())
}

func TestIsClear(t *testing.T) {
	s := &stack[int]{}
	s.Push(1)
	s.Push(2)
	require.False(t, s.IsEmpty())
	s.Clear()
	require.True(t, s.IsEmpty())
}
