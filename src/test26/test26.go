package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func runWorker(ch chan int, wg *sync.WaitGroup) {
	fmt.Println(" 随机数:", rand.Int())
	wg.Done()
	ch <- 1
}

func main() {

	var buf bytes.Buffer
	fmt.Println("bufff===>", buf.String())

	var num float32 = 0.0
	fmt.Println(num >= 0)

	var wg sync.WaitGroup
	var waitCount = 30
	wg.Add(waitCount)

	var ch = make(chan int)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < waitCount; i++ {
		go runWorker(ch, &wg)
	}

	wg.Wait()

	for {
		select {
		case num := <-ch:
			fmt.Println("read num from chain:", num)
		default:
			fmt.Println("return")
			return
		}
	}

}
