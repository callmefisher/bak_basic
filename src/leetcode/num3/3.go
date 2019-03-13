package main

import "fmt"

//在一个二维数组中（每个一维数组的长度相同），
//每一行都按照从左到右递增的顺序排序，
//每一列都按照从上到下递增的顺序排序。
//请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

func action(target int, array [][]int) bool {
	var row = len(array)
	var lines = len(array[0])
	var j = lines -1
	var i = 0
	for ; i < row && j >= 0;  {
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
	Print2DArray(array)
	fmt.Println("row:", row, " line:", lines, " target:", target, " in array?", false)
	return false
}





func main()  {
	fmt.Println("hello world")
	var arrary = [][]int {
		{1, 3, 5, 8},
		{2, 4, 8, 19},
		{7, 9, 10, 29},
	}
	action(2, arrary)
	action(11, arrary)
	
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


