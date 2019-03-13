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
package main


import (
	"fmt"
			"sync"
	"strings"
	)

var testM = make(map[int]int)
var testM2 = make(map[int]int)
var threadCount = 10
var lc sync.Mutex

//var slice = []int{1, 9, 10, 19, 20, 29, 30, 39, 40, 49, 50, 59, 60, 69, 70, 79, 80, 89, 90, 100}
var slice = []int{1, 9, 10, 19, 20, 29, 30, 39, 40, 49, 50, 59, 60, 69, 70, 79, 80, 89, 90, 100}

func insertToMap(start, end, goId int) {
	//lc.Lock()
	//defer lc.Unlock()
	//fmt.Println(goId, "range from", start, "  to:", end)
	//for i := start; i < end; i++ {
	//	testM[i] = i
	//}
	fmt.Println(goId, "range from", start, "  to:", end)
	testM2[2] = goId
}

func init() {

}

type S1 struct {
	num int
	str string
}


func testSyncGroup()  {
	
	var num = 1
	var s1 = &S1 {
		num: 1,
		str: "he",
	}
	var s2 = *s1
	
	fmt.Println("before: [s1 str]", s1.str  ," [s2 str]", s2.str ,)
	
	s1.str = "world"
	
	fmt.Println("after: [s1 str]", s1.str  ," [s2 str]", s2.str )
	
	fmt.Println("Done:", fmt.Sprint(num))
	
	
	m := make(map[int]int)
	
	//m[0] = 1
	m[0]++
	m[1] = 2
	
	for k, v:= range m {
		fmt.Println(k , ":", v)
	}
	
	var str1 = "z1.$(app).$(stream).$(startMs)-$(endMs)-$(seq)-$(connId).ts"
	var str2 = "$(app)/$(stream)/$(startMs)-$(endMs).ts"
	
	index1 := strings.Index(str1,"$(startMs1)")
	index2 := strings.Index(str2,"$(startMs)")
	
	
	
	fmt.Println(index1, " ->", index2)
	
	
	
}





func main() {
	
	for i := 0; i < threadCount; i++ {
		var index = i * 2
		//fmt.Println("index:", index)
		go insertToMap(slice[index], slice[index + 1], i)
	}
	
	for {
	
	}
	//testSyncGroup()
	//time.Sleep(1 * time.Second)

}
