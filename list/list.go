package list

import (
	"fmt"
)

// Node of list
type Node[V any] struct {
	Value      V
	Prev, Next *Node[V]
}

// List is double linked list
type List[V any] struct {
	length int

	Head, Tail *Node[V]
}

// Len returns length of List
func (l *List[V]) Len() int {
	return l.length
}

// New retruns a new an empty list
func New[V any]() *List[V] {
	return &List[V]{}
}

// PushHead push value of element into the head of list
func (l *List[V]) PushHead(v V) {
	l.PushHeadNode(&Node[V]{
		Value: v,
	})
}

// PushHeadNode push element into the head of list
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
}

// PushTail push value of element into the tail of list
func (l *List[V]) PushTail(v V) {
	l.PushTailNode(&Node[V]{
		Value: v,
	})
}

// PushTailNode push element into the tail of list
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

// Range will handle each node after n with fn
func (n *Node[V]) Range(fn func(val V)) {
	node := n
	for node != nil {
		fn(node.Value)
		node = node.Next
	}
}

// Range will handle each node before n with fn
func (n *Node[V]) ReverseRange(fn func(val V)) {
	node := n
	for node != nil {
		fn(node.Value)
		node = node.Prev
	}
}

// Remove a node in list
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

// String returns a string with all the message of list,
// you can print it out to know all about the list.
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
