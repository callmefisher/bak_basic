package main

import (
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
