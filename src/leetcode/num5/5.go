package main

import (
	"fmt"
)

// 链表
type LinkNode struct {
	Val  interface{}
	Next *LinkNode
}

func CreateLinkList(arr []int) *LinkNode {
	var head = &LinkNode{Val: "Head"}
	var p = head
	for i := 0; i < len(arr); i++ {
		var tmpNode = &LinkNode{Val: arr[i]}
		p.Next = tmpNode
		p = tmpNode
	}
	return head
}

func PrintLinkList(h *LinkNode) {
	for p := h; p != nil; p = p.Next {
		if p.Next == nil {
			fmt.Print(p.Val, "  ")
		} else {
			fmt.Print(p.Val, " -> ")
		}
	}
	fmt.Println("")
}

//1. 反转单链表
func ReverseList(h *LinkNode) *LinkNode {
	if h == nil {
		return nil
	}

	var p = h
	var q = h.Next
	p.Next = nil
	for q != nil {
		var front = q.Next
		q.Next = p
		p = q
		q = front
	}
	h = p
	return h
}

// 2. 快慢指针
//Given a singly linked list L: L0→L1→…→Ln-1→Ln,
//reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
//You must do this in-place without altering the nodes’ values.
//For example, Given {1, 2, 3, 4, 5, 6, 7}, reorder it to {1, 7, 2, 6, 3, 5, 4}.
func mergeLinkList(l1 *LinkNode, l2 *LinkNode) *LinkNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	var h1 = l1
	var h2 = l2

	for l1.Next != nil && l2.Next != nil {

		var tmpL1 = h1.Next
		var tmpL2 = h2.Next

		h1 = h1.Next
		h2 = h2.Next
	}

	return nil
}

func reorderLinkList(h *LinkNode) *LinkNode {
	if h == nil || h.Next == nil {
		return nil
	}
	PrintLinkList(h)
	var fast = h
	var slow = h
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var preHead = slow.Next

	var afterReverse = ReverseList(preHead)
	slow.Next = afterReverse
	PrintLinkList(h)
	return h
}

func main() {
	var l1 = CreateLinkList([]int{1, 2, 3, 4})
	PrintLinkList(l1)
	var l2 = ReverseList(l1)
	PrintLinkList(l2)
	var l3 = CreateLinkList([]int{1, 2, 3, 4, 5, 6, 7})
	reorderLinkList(l3)
}
