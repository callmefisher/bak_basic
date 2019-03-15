package main

import "fmt"

// 链表
type LinkNode struct {
	Val interface{}
	Next * LinkNode
}

func CreateLinkList(arr[]int) *LinkNode  {
	var head = &LinkNode{Val:"Head"}
	var p = head
	for i:= 0; i < len(arr) ; i++ {
		var tmpNode = &LinkNode{Val:arr[i]}
		p.Next = tmpNode
		p = tmpNode
	}
	return  head
}

func PrintLinkList(h* LinkNode)  {
	for p := h; p != nil ; p = p.Next {
		if p.Next == nil {
			fmt.Print(p.Val, "  ")
		} else {
			fmt.Print(p.Val, " -> ")
		}
	}
	fmt.Println("")
}


//1. 反转单链表
func ReverseList(h *LinkNode)  *LinkNode{
	if h == nil {
		return nil
	}
	
	var p = h
	var q = h.Next
	p.Next = nil
	for ;q != nil ;  {
		var front = q.Next
		q.Next = p
		p = q
		q = front
	}
	h = p
	return h
}





func main() {
	var l1 = CreateLinkList([]int{1, 2, 3, 4})
	PrintLinkList(l1)
	var l2 = ReverseList(l1)
	PrintLinkList(l2)
}
