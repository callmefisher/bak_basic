package main

import "fmt"

// 链表
type LinkNode struct {
	Val interface{}
	Next * LinkNode
}

func CreateLinkList(arr[]int) *LinkNode  {
	var head = &LinkNode{}
	var p = head
	for i:= 0; i < len(arr) ; i++ {
		var tmpNode = &LinkNode{Val:arr[i]}
		p.Next = tmpNode
		p = tmpNode
	}
	return  head
}

func PrintLinkList(h* LinkNode)  {
	for p := h.Next; p != nil ; p = p.Next {
		if p.Next == nil {
			fmt.Print(p.Val, "  ")
		} else {
			fmt.Print(p.Val, " -> ")
		}
	}
}


//1. 反转单链表


func main() {
	var l1 = CreateLinkList([]int{1, 2, 3, 4})
	PrintLinkList(l1)
}
