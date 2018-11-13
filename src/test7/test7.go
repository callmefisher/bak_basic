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
	"time"
)

var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	//use p
	fmt.Println(*p)
}

func test() {
	//var err error
	p, err := foo() //  bug
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
}

type tst struct {
	flag bool
}

func cycle() {

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(j)
			if j == 2 {
				goto RR
			}
		}
	}

RR:
	fmt.Println("rrrrrrrrrrrrrrr")

}

func main() {

	cycle()
	var t tst
	fmt.Println("=========>", t.flag)

	count := 0
	//for range time.After( 100 * time.Millisecond) {
	//	count ++
	//	fmt.Println("hello:", count)
	//}

	m := make(map[int]int)
	m[0] = 0
	m[2] = 2

	str1 := "hello"

	str2 := str1

	fmt.Println("1:", str1, " ->", str2)

	str2 = "world"
	fmt.Println("1:", str1, " ->", str2)

	for {
		timeout := time.Tick(100000 * time.Millisecond)
		select {
		// … do some stuff
		case <-timeout:
			count++
			if v, ok := m[3]; !ok {
				fmt.Println("hello:", count, " v:", v)
			}
		}
	}

}
