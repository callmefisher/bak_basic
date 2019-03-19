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
//For example, Given {1, 2, 3, 4, 5, 6, 7, 8 }, reorder it to {1, 8, 2, 7, 3, 5, 4}.
func mergeLinkList(l1 *LinkNode, l2 *LinkNode) *LinkNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	var p1 = l1
	var p2 = l2

	for p1.Next != nil && p2.Next != nil {

		var tmpL1 = p1.Next
		var tmpL2 = p2.Next
		p1.Next = p2

		if tmpL1 == l2 {
			break
		}
		p2.Next = tmpL1
		p1 = tmpL1
		p2 = tmpL2

	}

	if p1.Next == l2 {
		p1.Next = p2
	}

	return l1
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

	var l1 = h.Next
	var l2 = afterReverse

	//PrintLinkList(h)

	return mergeLinkList(l1, l2)
}

func main() {
	//var l1 = CreateLinkList([]int{1, 2, 3, 4})
	//PrintLinkList(l1)
	//var l2 = ReverseList(l1)
	//PrintLinkList(l2)
	var l3 = CreateLinkList([]int{1, 2, 3, 4, 5, 6})
	var l4 = reorderLinkList(l3)
	PrintLinkList(l4)
}
