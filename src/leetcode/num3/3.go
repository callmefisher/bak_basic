package main

import (
	"fmt"
	"unsafe"
)

//在一个二维数组中（每个一维数组的长度相同），
//每一行都按照从左到右递增的顺序排序，
//每一列都按照从上到下递增的顺序排序。
//请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

//思路，右上角，或者左下角开始查找，因为这个地方的数字，一边是一直增长，另一边是一直下降

func action(target int, array [][]int) bool {
	var row = len(array)
	var lines = len(array[0])
	var j = lines - 1
	var i = 0
	for i < row && j >= 0 {
		if array[i][j] == target {
			fmt.Println(target, "yes in array")
			return true
		}
		if array[i][j] > target {
			j--
			continue
		}
		if array[i][j] < target {
			i++
			continue
		}

	}
	//Print2DArray(array)
	fmt.Println("row:", row, " line:", lines, " target:", target, " in array?", false)
	return false
}

// 二进制1的个数

func action2(n int) {
	//	n&(n-1) 操作相当于把二进制表示中最右边的1变成0
	var originNum = n
	var count = 0
	for n != 0 {
		n = n & (n - 1)
		count++

	}

	fmt.Printf(" way1 num:%d, binary:%b  include  %d个1\n", originNum, originNum, count)

	// way2
	var num = originNum
	var count2 = 0
	for num != 0 {
		count2 = count2 + num&1
		num = num >> 1
	}
	fmt.Printf(" way2 num:%d, binary:%b  include  %d个1\n", originNum, originNum, count2)
}

// 大端小端

func endian() {
	
	
	const INT_SIZE int = int(unsafe.Sizeof(0))
	var i int = 0x1   // 0000 0000 0000 0001(大端， 低地址存储高位)   1000 0000 0000 0000(小端, 低地址存储低位)
	bs := (*[INT_SIZE]byte)(unsafe.Pointer(&i))
	fmt.Println("INT_SIZE:", INT_SIZE)
	if bs[0] == 0 {
		fmt.Println("system edian is little endian")
	} else {
		fmt.Println("system edian is big endian")
	}
	

}

func main() {
	fmt.Println("hello world")
	var arrary = [][]int{
		{1, 3, 5, 8},
		{2, 4, 8, 19},
		{7, 9, 10, 29},
	}
	action(2, arrary)
	action(11, arrary)

	action2(301)

	endian()
}

func Print2DArray(array [][]int) {

	var rows = len(array)
	var lie = len(array[0])
	fmt.Println("row:", rows, " lie:", lie)

	for i := 0; i < rows; i++ {
		for j := 0; j < lie; j++ {
			fmt.Println("line:", i, "->", array[i][j])
		}
		fmt.Println("")
	}
}
