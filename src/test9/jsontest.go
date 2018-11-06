package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

//type StreamMsg struct {
//	Msg string `json:"msg"`
//	Flag bool     `json:"flag"`
//	Num int  `json:"num"`
//}

type StreamMsg struct {
	Msg  string `json:"msg"`
	Flag bool   `json:"flag"`
	Num2 int    `json:"num"`
}

func test(ch chan int, num int) {
	ch <- num
	fmt.Println("write num==========>", num)
}
func t2(ch chan int) {
	fmt.Println("read from chain:====>", <-ch)
}

type StreamMsg2 struct {
	Msg  string `json:"msg"`
	Flag bool   `json:"flag"`
	Num2 int    `json:"num"`
}

func main() {

	s1 := &StreamMsg{Msg: "hello", Flag: true, Num2: 3}
	bytes, _ := json.Marshal(s1)

	m := make(map[string]interface{})
	m["test"] = bytes

	var s2 StreamMsg2
	b2, _ := m["test"]
	json.Unmarshal(b2.([]byte), &s2)

	//	s2 := m["test"]

	fmt.Println(reflect.TypeOf(s1))
	fmt.Println(reflect.TypeOf(s2))

	ch := make(chan int)

	go test(ch, 20)
	//go test(ch, 2)
	//	close(ch)

	go t2(ch)
	//r1 := <-ch
	//fmt.Println("--------->1  ", r1)

	//r2 := <-ch
	//fmt.Println("---------> 2  ", r2)

	signals := make(chan os.Signal, 16)
	signal.Notify(signals, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT)
	msg := map[int]string{1: "1", 2: "2"}
	delete(msg, 10)
	//panic("")

	for i := 0; i < 100; i++ {
		fmt.Println(rand.Uint32()%uint32(100) < 1)
	}

	defer func() {
		fmt.Println("Defer===============>")
	}()

	for {
		select {
		case sig := <-signals:
			fmt.Println("--------->", sig.String())
			return
		}
	}

}
