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
//例如 2 3 4 7 6 8 9 10 ===> 3 7 9 2 4 6 8 10
func reorder(arr []int) {
	PrintArray("before order:", arr)
	var availeOushuTag = 0 // 记录一个可以交换的偶数位置
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 1 {
			arr[i], arr[availeOushuTag] = arr[availeOushuTag], arr[i]
			availeOushuTag++
		}
	}
	PrintArray("after order:", arr)
}

// 5. 最长连续字串长度
//Given[100, 4, 200, 1, 3, 2],
//The longest consecutive elements sequence is[1, 2, 3, 4]. Return its length:4.
// 思路采用hash表
func longestSeqLen(arr []int) int {
	var lenOfArr = len(arr)
	if lenOfArr == 0 {
		return 0
	}
	var m = make(map[int]bool)
	for i := 0; i < lenOfArr; i++ {
		m[arr[i]] = true
	}
	var maxSeqLen = 1
	for i := 0; i < lenOfArr; i++ {
		var tmpNum = arr[i]
		if _, ok1 := m[tmpNum]; !ok1 {
			continue
		}
		delete(m, tmpNum)
		var preNum = tmpNum - 1
		var postNum = tmpNum + 1

		for {
			if _, ok := m[preNum]; !ok {
				break
			}
			delete(m, preNum)
			preNum--
		}

		for {
			if _, ok := m[postNum]; !ok {
				break
			}
			delete(m, postNum)
			postNum++
		}
		var tmpLen = postNum - preNum - 1
		if maxSeqLen < tmpLen {
			maxSeqLen = tmpLen
		}
	}

	fmt.Println(arr, "max seq len:", maxSeqLen)

	return maxSeqLen
}

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
	//MaxSubArray([]int{1, 8, -1, 0, 9, 18, -7, 8, 8})
	//MaxSubArray([]int{-1, -8, -2, 0, 9, -18, -7, -8, 8})
	//MaxSubArray([]int{-10, -8, -2, -1, -9, -18, -7, -8, -8})
	//longestSeqLen([]int{10, 2, 9})
	//longestSeqLen([]int{100, 4, 200, 1, 3, 2})
	reorder([]int{2, 3, 4, 7, 6, 8, 9, 10})
}
