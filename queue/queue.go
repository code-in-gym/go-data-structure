package queue

import "github.com/code-in-gym/go-data-structure/list"

type Queue[E any] struct {
	list *list.List[E]
}

// New returns a new empty queue
func New[E any]() *Queue[E] {
	return &Queue[E]{
		list: list.New[E](),
	}
}

func (q *Queue[E]) Add(v E) {
	q.list.PushTail(v)
}

func (q *Queue[E]) Peek() (v E) {
	return q.list.Head.Value
}

func (q *Queue[E]) Poll() (v E) {
	defer q.list.Remove(q.list.Head)
	return q.list.Head.Value
}

func (q *Queue[E]) Empty() bool {
	return q.list.Head == nil && q.list.Len() == 0
}

func (q *Queue[E]) Each(fn func(e E)) {
	q.list.Head.Each(fn)
}
