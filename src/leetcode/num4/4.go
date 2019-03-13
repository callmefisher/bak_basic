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

// 前序递归遍历，中序递归遍历, 前/中/后序非递归， 层次遍历, 输出二叉树每一层的第一个节点,树的高度

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
func LevelFirstNode(root* Tree)  {
	if root ==nil {
		return
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
		}
		
		if tmpNode.RChild != nil {
			tmpNode.RChild.Level = tmpNode.Level + 1
			q.Enqueue(tmpNode.RChild)
			m[tmpNode.RChild.Level] = true
		}
	}
	fmt.Println("")
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

	LevelFirstNode(root2)
	LevelFirstNode(root)
}
