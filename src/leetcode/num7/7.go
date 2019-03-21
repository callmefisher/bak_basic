package main

import "fmt"
import "strings"

//1 左旋字符串, 字符序列S=”abcXYZdef”,要求输出循环左移3位后的结果，即“XYZdefabc

func LeftRotate(str string, shiftLen int) string {
	var len = len(str)
	if len <= 1 {
		return str
	}
	var tmpStr = str + str
	var shitIndex = shiftLen % len
	return tmpStr[shitIndex : len+shitIndex]

}

// 2右旋字符串
func RightRotate(str string, shiftLen int) string {
	var len = len(str)
	if len <= 1 || shiftLen == 0 {
		return str
	}
	var tmpStr = make([]string, len, len)
	for i := 0; i < len; i++ {
		tmpStr[(shiftLen+i)%len] = str[i : i+1]
	}
	return strings.Join(tmpStr, "")
}

//3。 最长的不重复子串   例如"abcabcbb"--> abc 3         "pwwkew"->wke 3  "bbbbbb-->b" 1

//窗口的右边界就是当前遍历到的字符的位置，为了求出窗口的大小，我们需要一个变量left来指向滑动窗口的左边界，这样，如果当前遍历到的字符从未出现过，
//那么直接扩大右边界，如果之前出现过，那么就分两种情况，在或不在滑动窗口内，如果不在滑动窗口内，那么就没事，当前字符可以加进来，
//如果在的话，就需要先在滑动窗口内去掉这个已经出现过的字符了，去掉的方法并不需要将左边界left一位一位向右遍历查找，
//由于我们的HashMap已经保存了该重复字符最后出现的位置，所以直接移动left指针就可以了。我们维护一个结果res，每次用出现过的窗口大小来更新结果res

func longestUniquStr(str string) int {
	var len = len(str)
	var maxLen = 0
	if len == 0 {
		return maxLen
	}
	fmt.Print("str: ", str, " ")
	var m = make(map[byte]int)
	var tmpLeft = -1

	for i := 0; i < len; i++ {
		if lastPos, ok := m[str[i]]; ok {

			if lastPos > tmpLeft {
				tmpLeft = lastPos
			}
		}
		m[str[i]] = i

		if i-tmpLeft > maxLen {
			maxLen = i - tmpLeft
		}

	}
	return maxLen
}

//字符串
func main() {

	fmt.Println(LeftRotate("abcXYZdef", 10))
	fmt.Println(RightRotate("abcXYZdef", 3))
	fmt.Println(longestUniquStr("abcabcbb"))
	fmt.Println(longestUniquStr("pwwkew"))
	fmt.Println(longestUniquStr("bbbbbb"))
	fmt.Println(longestUniquStr("abcabcbdefg"))
}
