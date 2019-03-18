package main

import "fmt"
import "strings"

//1 左旋字符串, 字符序列S=”abcXYZdef”,要求输出循环左移3位后的结果，即“XYZdefabc

func LeftRotate(str string, shiftLen int) (string) {
	var len = len(str)
	if len  <= 1 {
		return str
	}
	var tmpStr = str + str
	var shitIndex = shiftLen % len
	return tmpStr[shitIndex :len + shitIndex]
	
	
}

// 2右旋字符串
func RightRotate(str string, shiftLen int) (string) {
	var len = len(str)
	if len  <= 1 || shiftLen == 0{
		return str
	}
	var tmpStr = make([]string, len, len)
	for i := 0; i < len; i++{
		tmpStr[(shiftLen + i ) % len] =str[i:i+1]
	}
	return  strings.Join(tmpStr, "")
}



//字符串
func main()  {

	fmt.Println(LeftRotate("abcXYZdef", 10))
	fmt.Println(RightRotate("abcXYZdef", 3))
	
}
