package utils

import (
	"fmt"
	"strings"
)

type Node struct {
	Value      int
	Next, Prev *Node
}

// NewCircularLinkedList creates a new circular linked list with a single node
// of the given value. The node's next and prev pointers point to itself.
func NewCircularLinkedList(value int) *Node {
	n := &Node{Value: value}
	n.Next = n
	n.Prev = n
	return n
}

// InsertAfter inserts a new node between n and n.next, and returns this node
func (n *Node) InsertAfter(value int) *Node {
	m := &Node{Value: value, Next: n.Next, Prev: n}
	n.Next.Prev = m
	n.Next = m
	return m
}

// Remove removes a node. The node is invalid after this operation.
func (n *Node) Remove() {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev

	n.Next = nil
	n.Prev = nil
	n.Value = -1
	n = nil
}

// Skip returns the node s steps away from this node
func (n *Node) Skip(s int) *Node {
	ret := n
	for i := 0; i < Abs(s); i++ {
		if s > 0 {
			ret = ret.Next
		} else {
			ret = ret.Prev
		}
	}
	return ret
}

// Debug string
func (n *Node) StringMarkCurrent(cur *Node) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(" %d ", n.Value))
	for p := n.Next; p != n; p = p.Next {
		if p == cur {
			sb.WriteString(fmt.Sprintf("(%d)", p.Value))
		} else {
			sb.WriteString(fmt.Sprintf(" %d ", p.Value))
		}
	}
	return sb.String()
}
