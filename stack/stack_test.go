package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	_one  = 1
	_zero = 0
)

type testCase struct {
	val int
}

func TestNew(t *testing.T) {
	t.Run("New int stack", func(t *testing.T) {
		s := New[int]()
		assert.IsType(t, Stack[int]{}, *s)
	})

	t.Run("New float64 stack", func(t *testing.T) {
		s := New[float64]()
		assert.IsType(t, Stack[float64]{}, *s)
	})

	t.Run("New testCase stack", func(t *testing.T) {
		s := New[testCase]()
		assert.IsType(t, Stack[testCase]{}, *s)
	})
}

func mockAStack[N any](t *testing.T) *Stack[N] {
	return New[N]()
}

func TestEmpty(t *testing.T) {
	s := mockAStack[int](t)
	assert.True(t, s.Empty())
}

func TestSize(t *testing.T) {
	t.Run("szie == 0", func(t *testing.T) {
		s := mockAStack[int](t)
		assert.Equal(t, _zero, s.Size())
	})

	t.Run("szie != 0", func(t *testing.T) {
		s := mockAStack[int](t)
		s.Push(_one)
		s.Push(_one)
		assert.Equal(t, 2, s.Size())
	})
}

func TestPushAndPeek(t *testing.T) {
	s := mockAStack[int](t)
	_, shouldBefalse := s.Peek()
	assert.False(t, shouldBefalse)
	s.Push(_one)
	val, ok := s.Peek()
	assert.True(t, ok)
	assert.Equal(t, _one, val)
}

func TestPop(t *testing.T) {
	s := mockAStack[int](t)
	_, shouldBefalse := s.Pop()
	assert.False(t, shouldBefalse)
	s.Push(_one)
	val, ok := s.Pop()
	assert.True(t, ok)
	assert.Equal(t, _one, val)
	assert.True(t, s.Empty())
}
