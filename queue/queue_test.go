package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	_one  = 1
	_zero = 0
)

func TestNew(t *testing.T) {
	t.Run("New int queue", func(t *testing.T) {
		q := New[int]()
		assert.IsType(t, Queue[int]{}, *q)
	})

	t.Run("New float64 queue", func(t *testing.T) {
		q := New[float64]()
		assert.IsType(t, Queue[float64]{}, *q)
	})
}

func mockAQueue(t *testing.T) *Queue[int] {
	t.Helper()
	testCase := New[int]()
	return testCase
}

func TestAddAndPeak(t *testing.T) {
	q := mockAQueue(t)
	assert.True(t, q.Empty())
	shouldBeZero, ok := q.Peek()
	assert.False(t, ok)
	assert.Equal(t, _zero, shouldBeZero)
	q.Add(_one)
	val, ok := q.Peek()
	assert.True(t, ok)
	assert.Equal(t, _one, val)
}

func TestPollAndEmpty(t *testing.T) {
	q := mockAQueue(t)
	assert.True(t, q.Empty())
	shouldBeZero, ok := q.Poll()
	assert.False(t, ok)
	assert.Equal(t, _zero, shouldBeZero)
	q.Add(_one)
	val, ok := q.Poll()
	assert.True(t, ok)
	assert.Equal(t, _one, val)
	assert.True(t, q.Empty())
}

func TestEach(t *testing.T) {
	t.Run("Only one element", func(t *testing.T) {
		q := mockAQueue(t)
		q.Add(_one)

		tmp := []int{}
		q.Each(func(e int) {
			tmp = append(tmp, e)
		})
		assert.ElementsMatch(t, []int{_one}, tmp)
	})

	t.Run("Add two elements", func(t *testing.T) {
		q := mockAQueue(t)
		q.Add(_one)
		q.Add(22)

		tmp := []int{}
		q.Each(func(e int) {
			tmp = append(tmp, e)
		})
		assert.ElementsMatch(t, []int{_one, 22}, tmp)

		tmpAfterPoll := []int{}
		val, ok := q.Poll()
		assert.True(t, ok)
		assert.Equal(t, _one, val)
		q.Each(func(e int) {
			t.Log(e)
			tmpAfterPoll = append(tmpAfterPoll, e)
		})
		assert.ElementsMatch(t, []int{22}, tmpAfterPoll)
	})
}
