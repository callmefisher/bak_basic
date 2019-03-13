package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {

}

func TestAction(t *testing.T) {
	var ast = assert.New(t)
	var r1 = action([]int{2, 7, 11, 15}, 9)
	ast.Equal(len(r1), 2)
	ast.Equal(r1[0], 0)
	ast.Equal(r1[1], 1)

	var r2 = action([]int{2, 7, 11, 15}, 14)
	ast.Equal(len(r2), 0)

	r1 = action([]int{2, 7, 11, 15}, 26)
	ast.Equal(len(r1), 2)
	ast.Equal(r1[0], 2)
	ast.Equal(r1[1], 3)

}

func BenchmarkAction(b *testing.B) {

	var arrayLen = 5
	var origin = make([]int, 0, arrayLen)

	for i := 0; i < arrayLen; i++ {
		origin = append(origin, rand.Int())
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		action(origin, rand.Int())
	}
}
