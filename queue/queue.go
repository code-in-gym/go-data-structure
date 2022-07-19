package queue

import "github.com/code-in-gym/go-data-structure/list"

type Queue[E any] struct {
	list *list.List[E]
}

// New returns a new empty queue.
func New[E any]() *Queue[E] {
	return &Queue[E]{
		list: list.New[E](),
	}
}

// Add an element into queue.
func (q *Queue[E]) Add(v E) {
	q.list.PushTail(v)
}

// Peek returns the value of element in the head of queue.
func (q *Queue[E]) Peek() (v E, exist bool) {
	if q.Empty() {
		return
	}
	return q.list.Head.Value, true
}

// Poll returns the value of element in the head of queue,
// And remove this element from queue.
func (q *Queue[E]) Poll() (v E, exist bool) {
	if q.Empty() {
		return
	}
	defer q.list.Remove(q.list.Head)
	return q.list.Head.Value, true
}

// Empty returns true if there is no element in the queue,
// false otherwise.
func (q *Queue[E]) Empty() bool {
	return q.list.Head == nil && q.list.Len() == 0
}

// Each calls `fn` on each in the queue,
// starting with the first pushed element.
func (q *Queue[E]) Each(fn func(e E)) {
	q.list.Head.Each(fn)
}
