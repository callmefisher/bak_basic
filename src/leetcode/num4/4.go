package main

import "fmt"
import "sync"
import (
	"container/list"
	"errors"
	)

//二叉搜索树第k个结点
// 栈
//队列

// 前序递归遍历，中序递归遍历, 前/中/后序非递归， 层次遍历, 输出二叉树每一层的第一个节点,树的高度, zigzag遍历

type Tree struct {
	Val    int
	LChild *Tree
	RChild *Tree
	Level  int
}

/*          1
2                  3
  5                  9
6   7            10
      8             12
                  -1
*/

type stack struct {
	lock sync.Mutex // you don't have to do this if you don't want thread safety
	s    []interface{}
}

func NewStack() *stack {
	return &stack{sync.Mutex{}, make([]interface{}, 0)}
}

func (s *stack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}
func (s *stack) IsEmpty() bool {
	return len(s.s) == 0
}

func (s *stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.s[len(s.s)-1]
}

func (s *stack) Pop() (interface{}, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return 0, errors.New("Empty Stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

type Queue struct {
	L    *list.List
	lock sync.Mutex
}

func NewQueue() *Queue {
	q := &Queue{L: list.New(), lock: sync.Mutex{}}
	return q
}

func (q *Queue) Enqueue(v interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.L.PushBack(v) // Enqueue
}

func (q *Queue) Dequeue() interface{}{
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.L.Len() > 0 {
		e := q.L.Front() // First element
		q.L.Remove(e) // Dequeue
		return e.Value
	}
	return nil
}
func (q* Queue) IsEmpty() bool{
	return q.L.Len() == 0
}
func (q* Queue)Len() int {
	return q.L.Len()
}
func (q* Queue)Front() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.L.Len() > 0 {
		return q.L.Front().Value
	}
	return nil
}


func test() {
	fmt.Println("\nStack=================")
	s := NewStack()
	s.Push("a1")
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	fmt.Println("\nQueue=================")
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue("试试2")
	q.Enqueue(3)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
}

func createTree() *Tree {
	var root = &Tree{Val: 1}

	var p1 = &Tree{Val: 2}
	var p2 = &Tree{Val: 3}

	var p3 = &Tree{Val: 5}
	var p4 = &Tree{Val: 9}
	var p5 = &Tree{Val: 6}
	var p6 = &Tree{Val: 7}
	var p7 = &Tree{Val: 10}
	var p8 = &Tree{Val: 8}
	var p9 = &Tree{Val: 12}
	var p10 = &Tree{Val: -1}

	root.LChild = p1
	root.RChild = p2
	p1.RChild = p3
	p2.RChild = p4
	p3.LChild = p5
	p3.RChild = p6
	p4.LChild = p7
	p6.RChild = p8
	p7.RChild = p9
	p9.LChild = p10

	return root

}


func createTree2() *Tree {
	var p1 = &Tree{Val: 1}
	
	var p2 = &Tree{Val: 2}
	var p3 = &Tree{Val: 3}
	var p4 = &Tree{Val:4}
	var p6 = &Tree{Val: 6}
	var p7 = &Tree{Val: 7}
	var p8 = &Tree{Val: 8}
	var p9 = &Tree{Val: 9}
	var p10 = &Tree{Val: 10}
	
	p1.LChild = p2
	p1.RChild = p3
	p2.RChild = p4
	p3.LChild = p6
	p3.RChild = p7
	
	
	p7.RChild = p8
	p8.LChild = p9
	p8.RChild = p10
	
	return p1
	
}

func PreOrderTraverseTree(t *Tree) {
	if t == nil {
		return
	}
	fmt.Println(t.Val)
	if t.LChild != nil {
		PreOrderTraverseTree(t.LChild)
	}
	if t.RChild != nil {
		PreOrderTraverseTree(t.RChild)
	}
}

func MiddleOrderTraverseTree(t *Tree) {
	if t == nil {
		return
	}
	if t.LChild != nil {
		MiddleOrderTraverseTree(t.LChild)
	}
	fmt.Println(t.Val)
	if t.RChild != nil {
		MiddleOrderTraverseTree(t.RChild)
	}
}
func AfterOrderTraverseTree(t *Tree) {
	if t == nil {
		return
	}
	if t.LChild != nil {
		AfterOrderTraverseTree(t.LChild)
	}

	if t.RChild != nil {
		AfterOrderTraverseTree(t.RChild)
	}
	fmt.Println(t.Val)
}

func PreOrderTraverseTreeWithStack(root *Tree) {
	if root == nil {
		return
	}
	stack := NewStack()
	stack.Push(root)
	for stack.IsEmpty() == false {
		v, _ := stack.Pop()
		var tmpNode = v.(*Tree)
		fmt.Print(tmpNode.Val, " ")
		if tmpNode.RChild != nil {
			stack.Push(tmpNode.RChild)
		}
		if tmpNode.LChild != nil {
			stack.Push(tmpNode.LChild)
		}
	}
	fmt.Println("")
}

func MiddleTraverseTreeWithStack(root *Tree) {
	if root == nil {
		return
	}
	stack := NewStack()

	/*           1
	2                  3
	  5                  9
	6   7            10
	      8             12
	                  -1
	*/
	// middle 2 6 5 7 8 1 3 10 -1 12 9
	// pre    1 2 5 6 7 8 3 9 10 12 -1
	// after  6 8 7 5 2 -1 12  10 9 3 1
	var p = root
	for p != nil {
		stack.Push(p)
		p = p.LChild
	}

	for stack.IsEmpty() == false {
		v, _ := stack.Pop()
		var tmpNode = v.(*Tree)
		fmt.Print(tmpNode.Val, " ")
		var p = tmpNode.RChild
		for p != nil {
			stack.Push(p)
			p = p.LChild
		}
	}

	fmt.Println("")
}

func EnablePop(t* Tree, m map[*Tree]bool) bool {
	if t == nil {
		return false
	}
	if t.RChild == nil && t.LChild == nil {
		return true
	}
	var flag1, flag2 = false, false
	if t.RChild != nil {
		_, flag1 = m[t.RChild]
	} else {
		flag1 = true
	}
	if t.LChild != nil {
		_, flag2 = m[t.LChild]
	} else {
		flag2 = true
	}
	return flag2 && flag1
}


func AfterTraverseWithStack(root *Tree) {

	/*        1
	2                  3
	  5                  9
	6   7            10
		  8             12
					  -1
	*/
	// middle 2 6 5 7 8 1 3 10 -1 12 9
	// pre    1 2 5 6 7 8 3 9 10 12 -1
	// after  6 8 7 5 2 -1 12  10 9 3 1
	if root == nil {
		return
	}
	stack := NewStack()
	var p = root
	for ;p != nil; {
		stack.Push(p)
		p = p.LChild
	}
	var tags = make(map[*Tree]bool)

	for ;stack.IsEmpty() == false; {
		var tmpTopNode = stack.Top().(*Tree)
		if EnablePop(tmpTopNode, tags) {
			stack.Pop()
			tags[tmpTopNode] = true
			fmt.Print(tmpTopNode.Val, " ")
			continue
		}
		if tmpTopNode.RChild == nil {
			continue
		}
		var p = tmpTopNode.RChild
		for ;p != nil ;{
			stack.Push(p)
			p = p.LChild
		}
	}
	fmt.Println("")

}


// 层次遍历
func LevelTraverseTree(root* Tree) {
	if root == nil {
		return
	}
	q := NewQueue()
	q.Enqueue(root)
	for ;q.IsEmpty() == false ;  {
		var tmpNode = q.Dequeue().(*Tree)
		fmt.Print(tmpNode.Val, " ")
		if tmpNode.LChild != nil {
			q.Enqueue(tmpNode.LChild)
		}
		if tmpNode.RChild != nil {
			q.Enqueue(tmpNode.RChild)
		}
	}
	fmt.Println("")
}

//输出每层中的第一个节点
func LevelFirstNodeWay2(root *Tree) (high int) {
	if root == nil {
		return 0
	}
	var q = NewQueue()
	q.Enqueue(root)
	
	for  q.IsEmpty() != true  {
		var levelLen = q.Len()
		for i:= 0; i < levelLen; i++ {
			var tmpNode = q.Dequeue().(*Tree)
			if i == 0{
				fmt.Print(tmpNode.Val, " ")
			}
			if tmpNode.LChild != nil {
				q.Enqueue(tmpNode.LChild)
			}
			if tmpNode.RChild != nil {
				q.Enqueue(tmpNode.RChild)
			}
		}
		high = high + 1
	}
	fmt.Println(" high:", high)
	return
}
func LevelFirstNode(root* Tree)  (high int){
	if root ==nil {
		return 0
	}
	/*           1
	2                  3
	  5                  9
	6   7            10
	      8             12
	                  -1
	*/
	var m = make(map[int] bool)
	var q = NewQueue()
	root.Level = 1
	high = 1
	q.Enqueue(root)
	m[root.Level] = true
	for ; q.IsEmpty() == false ;  {
		var tmpNode = q.Dequeue().(*Tree)
		if _, ok := m[tmpNode.Level]; ok {
			fmt.Print(tmpNode.Val, " ")
			delete(m, tmpNode.Level)
		}
		if tmpNode.LChild != nil {
			tmpNode.LChild.Level = tmpNode.Level + 1
			q.Enqueue(tmpNode.LChild)
			m[tmpNode.LChild.Level] = true
			if tmpNode.LChild.Level > high {
				high = tmpNode.LChild.Level
			}
		}
		
		if tmpNode.RChild != nil {
			tmpNode.RChild.Level = tmpNode.Level + 1
			q.Enqueue(tmpNode.RChild)
			m[tmpNode.RChild.Level] = true
			if tmpNode.RChild.Level > high {
				high = tmpNode.RChild.Level
			}
		}
	}
	fmt.Println(" high:", high)
	return
	
}

func GetTreeHigh(t * Tree)  (high int){
	if t == nil {
		return
	}
	
	var leftHigh = GetTreeHigh(t.LChild)
	var rightHigh = GetTreeHigh(t.RChild)
	if leftHigh > rightHigh {
		return leftHigh + 1
	}
	return rightHigh + 1
}

func ZigZagTree(t * Tree)  {
	if t == nil {
		return
	}
	var s1 = NewStack()
	var s2 = NewStack()
	t.Level = 1
	s1.Push(t)
	
	for {
		if s1.IsEmpty() && s2.IsEmpty() {
			break
		}
		
		// s1
		for ; s1.IsEmpty() == false; {
			
			var v1, _= s1.Pop()
			var tmpNode1= v1.(*Tree)
			fmt.Print(tmpNode1.Val, " ", )
			
			if tmpNode1.LChild != nil {
				s2.Push(tmpNode1.LChild)
			}
			if tmpNode1.RChild != nil {
				s2.Push(tmpNode1.RChild)
			}
		}
		
		// s2
		for ; s2.IsEmpty() == false; {
			
			var v2, _= s2.Pop()
			var tmpNode2= v2.(*Tree)
			fmt.Print(tmpNode2.Val, " ", )
			if tmpNode2.RChild != nil {
				s1.Push(tmpNode2.RChild)
			}
			if tmpNode2.LChild != nil {
				s1.Push(tmpNode2.LChild)
			}
			
		}
	}
	
	fmt.Println("")
}

func ReverseTree1(root* Tree)  {
	if root == nil {
		return
	}
	var tmpNode = root.LChild
	root.LChild = root.RChild
	root.RChild = tmpNode
	ReverseTree1(root.LChild)
	ReverseTree1(root.RChild)
}

func ReverseTree2(root* Tree)  {
	if root == nil {
		return
	}
	var q = NewQueue()
	q.Enqueue(root)
	
	for ;q.IsEmpty() != true ;  {
		var tmpNode = q.Dequeue().(*Tree)
		var canSwap1,canSwap2  = false, false
		if tmpNode.LChild != nil {
			canSwap1 = true
			q.Enqueue(tmpNode.LChild)
		}
		if tmpNode.RChild != nil {
			canSwap2 = true
			q.Enqueue(tmpNode.RChild)
		}
		if canSwap1 && canSwap2 {
			fmt.Println("swap root:", tmpNode.Val, " left:", tmpNode.LChild.Val, " right:", tmpNode.RChild.Val)
			var tmpLChild = tmpNode.LChild
			tmpNode.LChild = tmpNode.RChild
			tmpNode.RChild = tmpLChild
		}
		
	}
}

func getIndex(arr[]int, target int)  int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func testAvail(preOrder []int, middleOrder []int) bool {
	var m = make(map[int]bool)
	for i:=0 ; i < len(preOrder); i++ {
		m[preOrder[i]] = true
	}
	
	for j := 0; j < len(middleOrder); j++ {
		delete(m, middleOrder[j])
	}
	
	return len(m) == 0
}


// 递归，根据前序确认根，根据中序确认左右节点
func CreateTreeByPreANdMiddle(preOrder []int, middleOrder[]int) (root* Tree) {
	if len(preOrder) == 0 || len(middleOrder) == 0 || len(middleOrder) != len(preOrder) || !testAvail(preOrder, middleOrder) {
		return nil
	}
	var index = getIndex(middleOrder, preOrder[0])
	if index == -1 {
		return nil
	}
	
	root = &Tree{Val:preOrder[0]}
	// 左子树
	root.LChild = CreateTreeByPreANdMiddle(preOrder[1:1+index], middleOrder[0:index])
	//右子树
	root.RChild = CreateTreeByPreANdMiddle(preOrder[index+1:], middleOrder[index + 1:])
	return
}

func CreateTreeByOrderedArray(arr[]int )(root* Tree) {
	var arrLen = len(arr)
	if arrLen == 0 {
		return nil
	}
	if arrLen == 1 {
		root = &Tree{Val:arr[0]}
		return
	}
	var middleValue = arr[arrLen / 2]
	root = &Tree{Val:middleValue}
	root.LChild = CreateTreeByOrderedArray(arr[0:arrLen / 2 ])
	root.RChild = CreateTreeByOrderedArray(arr[arrLen / 2 + 1 :])
	
	return
}



func main() {
	var root = createTree()
	// 1. 前序非递归
	//PreOrderTraverseTreeWithStack(root)
	//2.中序非递归
	//MiddleTraverseTreeWithStack(root)
	//3.后续非递归
	//AfterTraverseWithStack(root)
	
	var root2 = createTree2()
	// 4. 层次遍历
	//LevelTraverseTree(root2)
	//LevelTraverseTree(root)
	
	// 5. 层次中的第一个节点输出

	//LevelFirstNode(root2)
	//LevelFirstNode(root)
	
	LevelFirstNodeWay2(root2)
	LevelFirstNodeWay2(root)
	// 6. 树的高度
	//fmt.Println("root2 high: ", GetTreeHigh(root2))
	//fmt.Println("root high:", GetTreeHigh(root))
	
	
	// 7. z字遍历
	//ZigZagTree(root2)
	//ZigZagTree(root)
	
	// 8. 树的子结构， 判断B树是不是A树的一颗子树
	// 双层递归
	// 1. 第一层, A树的节点等于B树的节点，调用第二层递归，否则，则继续第一层 递归，分别向父树的左右节点扩展
	// a->lchild, b || a->right, b
	// 2. 第二层， A树的子节点等于B树的子节点，则继续第二层的递归，扩展到A, B树的左右节点。
	//否则返回false, 既B树不是A树的子树
	
	// 9.前序和中序重建二叉树，
	var _ = CreateTreeByPreANdMiddle(
		[]int{1, 2, 4, 3 , 6, 7, 8, 9 , 10, },
		[]int{2, 4, 1, 6,  3, 7, 9, 8, 10,})
	
	//PrintTreeDetail(root3)
	
	// 10. 二叉树镜像，反转二叉树， 借助递归或者队列，顺次交换node的左右节点
	//ReverseTree1(root2)
	ReverseTree2(root2)
	//LevelTraverseTree(root2)
	
	// 11. 有序数组重建二叉树
	var root4 = CreateTreeByOrderedArray([]int{1, 2, 3, 4, 5, 6, 8, 9})
	fmt.Println("root4:", root4)
	PrintTreeDetail(root4)
}

func PrintTreeDetail(root *Tree)  {
	if root == nil {
		return
	}
	var  q = NewQueue()
	q.Enqueue(root)
	for ;q.IsEmpty() != true ;  {
		var tmpNode = q.Dequeue().(*Tree)
		fmt.Println("\n\nroot:", tmpNode.Val)
		if tmpNode.LChild != nil {
			fmt.Print(" lchild: ", tmpNode.LChild.Val)
			q.Enqueue(tmpNode.LChild)
		}
		if tmpNode.RChild != nil {
			fmt.Print(" rchild: ", tmpNode.RChild.Val)
			q.Enqueue(tmpNode.RChild)
		}
		
		
	}
}