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
	"math/rand"
	"strings"
	"time"
)

type A interface {
	func1(a int) bool
}

type B struct {
	Num int
}

func (b *B) func1(a int) bool {
	return b.Num == a
}

func NewA() A {

	b := &B{Num: 1}
	return b
}

func SignatureNonceRandomString() string {
	return "2"
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%v", rand.Int())
}

func percentReplace(str string) string {
	defer func() {
		fmt.Println("Defer in percentReplace")
	}()
	str = strings.Replace(str, "+", "%20", -1)
	str = strings.Replace(str, "*", "%2A", -1)
	str = strings.Replace(str, "%7E", "~", -1)

	s1 := strings.TrimSpace(strings.SplitN("", ";", 2)[0])
	fmt.Println(s1)

	return SignatureNonceRandomString()
}

func main() {

	B := NewA()

	fmt.Println("interface test:", B.func1(2), "=》", SignatureNonceRandomString())

	fmt.Println(percentReplace(""))
}
