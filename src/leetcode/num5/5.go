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

	var afterReverse = reverseListWay2(preHead, nil)
	slow.Next = afterReverse

	var l1 = h.Next
	var l2 = afterReverse

	//PrintLinkList(h)

	return mergeLinkList(l1, l2)
}

//. 给定一个不含有重复数字的链表，和一个数组，其中数组的元素全部来自于链表，返回相连子链表的数量
// 例如 链表A: 10->12->2->4->5->1->null, 数组:[12, 5, 4], 链表被分为[12,] [4，5]，2个子链表
//                                          [ 2, 1, 5,12] --> [12,2]和[ 5,1]2个

func getConnectListNum(h *LinkNode, arr []int) {

	fmt.Print("arr:", arr, " ")
	var arrLen = len(arr)
	if arrLen == 0 || h == nil || h.Next == nil {
		return
	}
	var m = make(map[int]bool)
	for i := 0; i < arrLen; i++ {
		m[arr[i]] = true
	}

	var p = h.Next
	var count = 0

	for p != nil {

		for {

			if _, ok1 := m[p.Val.(int)]; !ok1 {
				p = p.Next
			} else {
				break
			}
		}

		count++

		for {
			if p == nil {
				break
			}
			if _, ok1 := m[p.Val.(int)]; ok1 {
				p = p.Next
			} else {
				break
			}

		}
	}
	fmt.Print(" num of sub link:", count, " \n")
}

// 链表排序
//Example 1:
//Input: 4->2->1->3
//Output: 1->2->3->4
//时间复杂度必须为O(nlgn)
//Example 2:
//Input: -1->5->3->4->0
//Output: -1->0->3->4->5

func listSort(l *LinkNode) {

}

// 链表局部反转
//Example 1:
//Input: 4->2->1->3->5->0,  m = 3, n = 5, 反转[m , n]之间的部分
//Output: 4->2->5->3->1->0
func reversePartList(head *LinkNode, m, n int) *LinkNode {
	var fast = head
	var slow = head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	fmt.Println("slowNode:", slow.Val)

	if m >= n || m <= 0 || n <= 0 {
		return head
	}

	var tmpNode1 = head
	var tmpNode2 = head
	var count1, count2 = 1, 1

	for count2 < n {

		if count1 < m {
			tmpNode1 = tmpNode1.Next
			if tmpNode1 == nil {
				return head
			}
		}
		tmpNode2 = tmpNode2.Next
		if tmpNode2 == nil {
			return head
		}
		count1++
		count2++
	}
	if tmpNode1.Next == nil || tmpNode2.Next == nil {
		return head
	}

	var firstTail = tmpNode1
	var secondTail = tmpNode2.Next.Next
	tmpNode2.Next.Next = nil
	var newPartHead = reverseListWay2(tmpNode1.Next, secondTail)
	firstTail.Next = newPartHead

	return head
}

// 反转链表方法2
func reverseListWay2(head *LinkNode, secondTail *LinkNode) *LinkNode {
	var preNode *LinkNode = secondTail
	for head != nil {
		var tmpNode = head.Next
		head.Next = preNode
		preNode = head
		head = tmpNode
	}
	return preNode
}

func main() {
	//var l1 = CreateLinkList([]int{1, 2, 3, 4})
	//PrintLinkList(l1)
	//var l2 = ReverseList(l1)
	//PrintLinkList(l2)
	var l3 = CreateLinkList([]int{1, 2, 3, 4, 5, 6})
	var l4 = reorderLinkList(l3)
	PrintLinkList(l4)
	var l5 = CreateLinkList([]int{10, 12, 2, 4, 5, 1})
	PrintLinkList(l5)
	//getConnectListNum(l5, []int{2, 1, 5, 12})
	//getConnectListNum(l5, []int{4, 1, 5, 12})
	//getConnectListNum(l5, []int{4, 1, 12})
	//PrintLinkList(reverseListWay2(l5))

	//PrintLinkList(reversePartList(l5, 3, 5))
	//PrintLinkList(reversePartList(l5, 3, 6))
	PrintLinkList(reversePartList(l5, 3, 4))
}
