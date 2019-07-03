package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"path/filepath"
	"sync"
	"time"
)

func runWorker(ch chan int, wg *sync.WaitGroup) {
	fmt.Println(" 随机数:", rand.Int())
	wg.Done()
	ch <- 1
}

type Ability struct {
	Can    bool `json:"can" bson:"can"`
	Frozen bool `json:"frozen" bson:"frozen"`
}

func TestExists(v int, m map[int]Ability, str *string) bool {
	_, ok := m[v]
	*str = "afa"

	return ok
}

type Device struct {
	Id string `json:"-"`

	// meta data
	Meta json.RawMessage `json:"meta,omitempty" bson:"meta,omitempty"`
}

func main() {

	type Meta struct {
		Mode  int      `json:"mode"`
		Info  string   `json:"info"`
		Hosts []string `json:"hosts"`
	}

	data1 := Meta{Mode: 0, Info: "abcdef", Hosts: []string{"aas", "ddd", "ssssss", "ddddd"}}
	meta, _ := json.Marshal(&data1)
	_ = &Device{Meta: meta}

	fmt.Println("d ", len(meta))

	var slice = make([]int, 10, 10)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 4

	fmt.Println(filepath.Ext("aaa..."))
	fmt.Println(slice, " ", slice[1:2])

	var ability = Ability{
		Can:    true,
		Frozen: true,
	}
	var m = make(map[int]Ability)
	m[1] = ability
	m[2] = ability

	var str1 = "str1"
	var str2 = "str2"
	fmt.Println("default ability:", TestExists(1, m, &str1), TestExists(3, m, &str2))
	fmt.Println(str1, " ", str2)

	var streamID = "aa"
	var format = "22"

	var ft = fmt.Sprintf("psegments/%s/{{.start}}-{{.end}}.%s", streamID, format)
	var from = time.Now()

	fmt.Println(ft, " ", from.Unix())

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
