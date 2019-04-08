package main

import (
	"fmt"
)

// map, hash， btree, trie tree等

//单调栈，栈中的元素都是单调递增或者单调递减，
// 全排列

// 最小栈，
//Example:
//
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> Returns -3.
//minStack.pop();
//minStack.top();      --> Returns 0.
//minStack.getMin();   --> Returns -2.

//使用两个栈来实现，一个栈来按顺序存储push进来的数据，另一个用来存出现过的最小值
/*
class MinStack {
public:
MinStack() {}

void push(int x) {
s1.push(x);
if (s2.empty() || x <= s2.top()) s2.push(x);
}

void pop() {
if (s1.top() == s2.top()) s2.pop();
s1.pop();
}

int top() {
return s1.top();
}

int getMin() {
return s2.top();
}

private:
stack<int> s1, s2;
};
*/

func main() {
	var array = make([]int, 10)
	array[0] = 1
	fmt.Println(array[2])
	fmt.Println(array[0])
	//fmt.Println(array[20])

}
