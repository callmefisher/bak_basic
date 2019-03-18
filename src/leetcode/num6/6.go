package main

import (
	"fmt"
)

// 数组

//1.连续子数组的最大值， 并返回子数组起始位置

//1 8 -1 0 9 18 -7 8 8

func PrintArray(str string, arr []int) {
	fmt.Print(str)
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func MaxSubArray(arr []int) []int {
	PrintArray("\n起始数组:", arr)
	var tmpPreMax, curMax = arr[0], arr[0]
	var startIndex, endIndex = 0, 0
	for i := 1; i < len(arr); i++ {

		var flag1, flag2 = false, false
		if tmpPreMax < 0 {
			tmpPreMax = arr[i]
			flag1 = true
		} else {
			tmpPreMax = arr[i] + tmpPreMax
		}

		if curMax < tmpPreMax {
			curMax = tmpPreMax
			endIndex = i
			flag2 = true
		}
		if flag1 && flag2 {
			startIndex = i
		}
	}

	var result = []int{curMax, startIndex, endIndex}
	PrintArray("子数组最大和：", result)
	return result
}

// 2. 寻找从上至下的一条值最大的路径
/*
            [ [2]]
[          [3, 4]]
[       [6, 5, 7]]
[    [4, 1, 8, 3]]
*/

// 3.旋转数组

//4. 给定一个存放整数的数组，重新排列数组使得数组左边为奇数，右边为偶数, 要求：空间复杂度 O(1)，时间复杂度为 O（n）

// 5. 最大公约数

func gcd1(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		var tmpOldb = b
		b = a % b
		a = tmpOldb
	}
	return a
}

func gcd2(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd2(b, a%b)
}

func main() {
	MaxSubArray([]int{1, 8, -1, 0, 9, 18, -7, 8, 8})
	MaxSubArray([]int{-1, -8, -2, 0, 9, -18, -7, -8, 8})
	MaxSubArray([]int{-10, -8, -2, -1, -9, -18, -7, -8, -8})
}
