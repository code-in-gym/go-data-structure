package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	_zero  = 0
	_one   = 1
	_two   = 2
	_three = 3
	_four  = 4
	str    = ` 0 <=> 1 <=> 2 <=> 3 <=> length: 4`
)

func TestNewAList(t *testing.T) {
	t.Run("New int list", func(t *testing.T) {
		l := New[int]()
		assert.IsType(t, List[int]{}, *l)
	})

	t.Run("New float64 list", func(t *testing.T) {
		l := New[float64]()
		assert.IsType(t, List[float64]{}, *l)
	})
}

func mockAList(t *testing.T) *List[int] {
	t.Helper()
	testCase := New[int]()
	return testCase
}

func TestPushHead(t *testing.T) {
	l := mockAList(t)
	const value1, value2 = -1, 2
	l.PushHead(value1)
	assert.Equal(t, value1, l.Head.Value)
	l.PushHead(value2)
	assert.Equal(t, value2, l.Head.Value)
}

func TestPushTail(t *testing.T) {
	l := mockAList(t)
	const value1, value2 = -1, 2
	l.PushTail(value1)
	assert.Equal(t, value1, l.Tail.Value)
	l.PushTail(value2)
	assert.Equal(t, value2, l.Tail.Value)
}

func mockPushTail(t *testing.T, testCase *List[int]) (*List[int], int) {
	t.Helper()
	testCase.PushTail(_zero)
	testCase.PushTail(_one)
	testCase.PushTail(_two)
	testCase.PushTail(_three)
	return testCase, _four
}

func TestEach(t *testing.T) {
	l := mockAList(t)
	l, length := mockPushTail(t, l)
	listNums := make([]int, length)
	i := 0
	l.Head.Each(func(val int) {
		listNums[i] = val
		i++
	})
	assert.ElementsMatch(
		t,
		[]int{_zero, _one, _two, _three},
		listNums,
	)
}

func TestReverseEach(t *testing.T) {
	l := mockAList(t)
	l, length := mockPushTail(t, l)
	listNums := make([]int, length)
	i := 0
	l.Tail.ReverseEach(func(val int) {
		listNums[i] = val
		i++
	})
	assert.ElementsMatch(
		t,
		[]int{_three, _two, _one, _zero},
		listNums,
	)
}

func TestLen(t *testing.T) {
	l := mockAList(t)
	l, length := mockPushTail(t, l)
	assert.Equal(t, length, l.Len())
}

func TestRemove(t *testing.T) {
	t.Run("Remove node in middle", func(t *testing.T) {
		l := mockAList(t)
		l, length := mockPushTail(t, l)
		nodeNext := l.Head.Next.Next
		l.Remove(l.Head.Next)
		assert.Equal(t, length-1, l.Len())
		assert.Same(t, nodeNext, l.Head.Next)
	})

	t.Run("Remove node in tail", func(t *testing.T) {
		l := mockAList(t)
		l, length := mockPushTail(t, l)
		l.Remove(l.Tail)
		assert.Equal(t, length-1, l.Len())
		assert.Equal(t, _two, l.Tail.Value)
	})

	t.Run("Remove node in head", func(t *testing.T) {
		l := mockAList(t)
		l, length := mockPushTail(t, l)
		l.Remove(l.Head)
		assert.Equal(t, length-1, l.Len())
		assert.Equal(t, _one, l.Head.Value)
	})

	t.Run(
		"Remove node in head with one element",
		func(t *testing.T) {
			l := mockAList(t)
			l.PushHead(1)
			assert.Equal(t, 1, l.Len())
			l.Remove(l.Head)
			assert.Equal(t, 0, l.Len())
		})
}

func TestString(t *testing.T) {
	t.Run("String list not nil", func(t *testing.T) {
		l := mockAList(t)
		l, _ = mockPushTail(t, l)
		assert.Equal(t, str, l.String())
	})

	t.Run("String list nil", func(t *testing.T) {
		l := mockAList(t)
		l = nil
		assert.Equal(t, l.String(), "")
	})
}
