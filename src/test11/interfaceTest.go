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
