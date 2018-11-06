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
