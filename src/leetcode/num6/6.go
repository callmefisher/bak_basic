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

// 3.旋转数组， 给定一个有序数组，如{1,2,3,4,5,6,7,8,9}，我们将对这个数组进行选择，位置旋转未知。
// 下面给出一个可能的旋转结果。 如{4,5,6.7.8,9.1.2.3}，我们可以理解为它从元素4位置开始旋转。之后给定一个指定的数字n，
// 让我们从{4,5,6,7,8,9,1,2,3}这个数组中找出它的
//位置，要求时间复杂度尽可能的低

//思路，先看哪边的数组是有序的，因为拐点至少一边是有序的
func searchIndexInRotateArr(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	var low = 0
	var high = len(arr) - 1
	for low <= high {
		var middleIndex = (low + high) / 2
		if arr[middleIndex] == target {
			return middleIndex
		}

		if arr[middleIndex] > arr[low] {
			//左侧有序
			if target >= arr[low] && target <= arr[middleIndex] {
				//在左半部分
				high = middleIndex - 1
			} else {
				//在右半部分
				low = middleIndex + 1
			}

		} else {
			//右侧有序
			if target >= arr[middleIndex] && target <= arr[high] {
				low = middleIndex + 1
			} else {
				high = middleIndex - 1
			}

		}
	}
	return -1
}

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


// 4. 有序数组，无重复元素
func BinarySearchUnique(arr[]int, target int) int  {
	var len = len(arr)
	if len == 0 {
		return -1
	}
	var low = 0
	var high = len - 1
	for low <= high  {
		var middle = low +  ( high - low ) / 2
		if target == arr[middle] {
			return middle
		} else if target > arr[middle] {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return -1
}

// 5. 有序数组， 有重复元素，寻找首个位置
func BinarySearchFirst(arr[]int, target int) int  {
	var len = len(arr)
	if len == 0 {
		return -1
	}
	var low = 0
	var high = len - 1
	for low <= high  {
		var middle = low +  ( high - low ) / 2
		if target == arr[middle] {
			return middle
		} else if target > arr[middle] {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return -1
}



// 6. 有序数组，有重复元素，寻找最后一个位置
func BinarySearchLast(arr[]int, target int) int  {
	var len = len(arr)
	if len == 0 {
		return -1
	}
	var low = 0
	var high = len - 1
	for low <= high  {
		var middle = low +  ( high - low ) / 2
		if target == arr[middle] {
			return middle
		} else if target > arr[middle] {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return -1
}




func main() {
	//MaxSubArray([]int{1, 8, -1, 0, 9, 18, -7, 8, 8})
	//MaxSubArray([]int{-1, -8, -2, 0, 9, -18, -7, -8, 8})
	//MaxSubArray([]int{-10, -8, -2, -1, -9, -18, -7, -8, -8})
	//longestSeqLen([]int{10, 2, 9})
	//longestSeqLen([]int{100, 4, 200, 1, 3, 2})
	//reorder([]int{2, 3, 4, 7, 6, 8, 9, 10})
	fmt.Println(searchIndexInRotateArr([]int{4, 5, 6, 7, 8, 9, 1, 2, 3}, 5))
	fmt.Println(BinarySearchUnique([]int{1, 2, 6, 9, 10, 11, 14}, 10))
}
