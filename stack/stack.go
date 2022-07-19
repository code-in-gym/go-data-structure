package stack

import "github.com/code-in-gym/go-data-structure/list"

// Stack a LIFO stack
type Stack[N any] struct {
	list *list.List[N]
}

func New[N any]() *Stack[N] {
	return &Stack[N]{
		list: list.New[N](),
	}
}

// Empty returns true if there is no element in the stack,
// false otherwise.
func (s *Stack[N]) Empty() bool {
	return s.list.Tail == nil && s.list.Len() == 0
}

// Peek returns the top of this stack and true,
// if the stack is not empty.
// Otherwise returns default value of N and false.
func (s *Stack[N]) Peek() (v N, exist bool) {
	if s.Empty() {
		return
	}
	return s.list.Tail.Value, true
}

// Pop returns the top of this stack and true,
// if the stack is not empty.
// And remove this element from the stack.
// Otherwise returns default value of N and false.
func (s *Stack[N]) Pop() (v N, exist bool) {
	if s.Empty() {
		return
	}
	defer s.list.Remove(s.list.Tail)
	return s.list.Tail.Value, true
}

// Push an element into stack
func (s *Stack[N]) Push(v N) {
	s.list.PushTail(v)
}

// Size returns size of stack
func (s *Stack[N]) Size() int {
	return s.list.Len()
}
