package main

import (
	"fmt"
	"math/rand"
)

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (8 -> 7 -> 9)
//输出：0 -> 2 -> 3 -> 1
//原因：342 + 978 = 1320

type ListNode struct {
	Val  int
	Next *ListNode
}

func fillNode(p3 **ListNode, saveNum int)  {
	(*p3).Next = &ListNode{Val: saveNum, Next:nil}
	*p3 = (*p3).Next
}

func getRemainNum(sumNum int, jinwei *int) (saveNum int) {
	saveNum = (sumNum + *jinwei) % 10
	*jinwei = (sumNum + *jinwei) / 10
	return
}

func addTwoNumbers(p1 *ListNode, p2 *ListNode) *ListNode {
	var jinwei = 0
	var l3 = &ListNode{}
	var p3 = l3
	for ; p1 != nil && p2 != nil; p1 = p1.Next {

		var saveNum = getRemainNum(p1.Val+p2.Val, &jinwei)
		fillNode(&p3, saveNum)
		p2 = p2.Next
	
	}
	for ; p1 != nil; p1 = p1.Next {
		var saveNum = getRemainNum(p1.Val, &jinwei)
		fillNode(&p3, saveNum)
	}
	for ; p2 != nil; p2 = p2.Next {
		var saveNum = getRemainNum(p2.Val, &jinwei)
		fillNode(&p3, saveNum)
	}
	if jinwei != 0 {
		fillNode(&p3, jinwei)
	}
	
	
	PrintList("l3 ", l3.Next)

	return l3.Next
}

func createList(a []int) *ListNode {

	var q = &ListNode{}
	var p = q
	for i := 0; i < len(a); i++ {
		p.Next = &ListNode{Val: a[i], Next: nil}
		p = p.Next
	}
	return q.Next
}

func PrintList(name string, l *ListNode) {
	fmt.Print("\nname:", name)
	for ; l != nil; l = l.Next {
		if l.Next != nil {
			fmt.Print(l.Val, " --> ")
			continue
		}
		fmt.Print(l.Val)

	}
	fmt.Println("")
}




func main() {
	var p1 = createList([]int{2, 4, 3})
	PrintList("p1 ", p1)
	var p2 = createList([]int{5, 6, 4})
	PrintList("p2 ", p2)

	addTwoNumbers(p1, p2)
	
	

	fmt.Println(rand.Perm(5))

}
