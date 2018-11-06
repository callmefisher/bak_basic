package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"qbox.us/errors"
	"sort"
	"strconv"
)

var (
	ERR1 = errors.New("hello")
)

func testErr() error {

	return ERR1
}

func hashTest(name string) {
	hash(name)
}

func randTest() {
	for i := 0; i < 100; i++ {
		fmt.Println(rand.Uint32() % 100)
	}
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func main() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	fmt.Println("Origin:", people)
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age:", people)

	var m1 = map[int]int{1: 1, 2: 2, 3: 3}
	var m2 = make(map[int]int)

	fmt.Println(m1, " ->", m2)
	copyMap(m1, m2)
	m1[1] = 10
	fmt.Println(m1, " -> ", m2)

	var n1 = 1
	var n2 = n1
	n1 = 3
	fmt.Println("n1:", n1, " ->", n2)

	err := testErr()
	fmt.Println(err.Error())
	fmt.Println(err == ERR1)

	hashNum := int(hash("abcwd"))
	strHash := strconv.Itoa(hashNum)

	fmt.Println(strHash[:4])

	randTest()

}

func copyMap(origin map[int]int, dest map[int]int) {
	for k, v := range origin {
		dest[k] = v
	}
}
