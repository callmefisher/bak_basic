//MIT License
//
//Copyright (c) 2018 XiaYanji
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.
package test2

import "fmt"
import "runtime"

const (
	Num int = 3
)

func init() {

	fmt.Println("pack 2")
}
func MyaddNum(num1, num2 int) int {
	return Num + num1 + num2
}

func sumInGorountine(s []int, result chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	result <- sum
}

func TestSum() {
	array := []int{1, 2, 3, -1, 0}
	result1 := make(chan int)
	result2 := make(chan int)
	fmt.Printf("0cur rontine num:%v\n", runtime.NumGoroutine())
	go sumInGorountine(array[:len(array)/2], result1)
	fmt.Printf("1cur rontine num:%v\n", runtime.NumGoroutine())
	go sumInGorountine(array[len(array)/2:], result2)
	fmt.Printf("2cur rontine num:%v\n", runtime.NumGoroutine())
	//r1 := <-result1
	//r2 := <-result2
	<-result1
	<-result2

	//fmt.Printf("%v %v %v\n", r1, r2, r1 + r2)
	fmt.Printf("4cur rontine num:%v\n", runtime.NumGoroutine())

}
