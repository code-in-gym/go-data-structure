package list

import (
	"fmt"
)

type Node[V any] struct {
	Value      V
	Prev, Next *Node[V]
}

// Double linked list
type List[V any] struct {
	length int

	Head, Tail *Node[V]
}

func (l *List[V]) Len() int {
	return l.length
}

func New[V any](head *Node[V]) *List[V] {
	return &List[V]{
		length: 0,
	}
}

func (l *List[V]) PushHead(v V) {
	l.PushHeadNode(&Node[V]{
		Value: v,
	})
}

func (l *List[V]) PushHeadNode(n *Node[V]) {
	defer func() {
		l.length++
	}()
	n.Prev = nil
	n.Next = l.Head
	if l.Head != nil {
		l.Head.Prev = n
	} else {
		l.Tail = n
	}
	l.Head = n
	l.length++
}

func (l *List[V]) PushTail(v V) {
	l.PushTailNode(&Node[V]{
		Value: v,
	})
}

func (l *List[V]) PushTailNode(n *Node[V]) {
	defer func() {
		l.length++
	}()
	n.Next = nil
	n.Prev = l.Tail
	if l.Tail != nil {
		l.Tail.Next = n
	} else {
		l.Head = n
	}
	l.Tail = n
}

func (n *Node[V]) Range(fn func(val V)) {
	node := n
	for node != nil {
		fn(node.Value)
		node = node.Next
	}
}

func (n *Node[V]) ReverseRange(fn func(val V)) {
	node := n
	for node != nil {
		fn(node.Value)
		node = node.Prev
	}
}

func (l *List[V]) Remove(n *Node[V]) {
	defer func() {
		l.length--
	}()
	// 1. n in head
	if n.Prev == nil {
		n.Next.Prev = nil
		return
	}
	// 2. n in tail
	if n.Next == nil {
		n.Prev.Next = nil
		return
	}
	// 3. n in the middle of list
	n.Next.Prev = n.Prev
	n.Prev.Next = n.Next
}

func (l *List[V]) String() (str string) {
	if l == nil {
		return
	}
	l.Head.Range(func(val V) {
		str += fmt.Sprintf(" %v <=>", val)
	})
	str += fmt.Sprintf(" length: %d", l.length)
	return str
}
