package main

import (
	"fmt"
	"time"
	"os"
	"runtime/pprof"
)

var cpuProfile = "./cpu_profile"

func startCPUProfile() {
	if cpuProfile != "" {
		f, err := os.Create(cpuProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not create cpu profile output file: %s",
				err)
			return
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "Can not start cpu profile: %s", err)
			f.Close()
			return
		}
	}
}

func stopCPUProfile() {
	if cpuProfile != "" {
		pprof.StopCPUProfile() // 把记录的概要信息写到已指定的文件
	}
}



func main() {
	
	var now = time.Now()
	fmt.Println("sec:", now.IsZero())
	
	
	var str1 = fmt.Sprint("aaa", "bbbb", "cccc", "http://")
	fmt.Println(str1)
	go startCPUProfile()
	
	ch := make(chan int, 10)
	for {
		select {
		case <-time.After(4 * time.Second): {
			fmt.Println("===============>A")
			stopCPUProfile()
			fmt.Println("===============>")
			os.Exit(0)
		}
		case <- ch:
			fmt.Println("111")
		time.Sleep(1 * time.Second)
		}
	}
	


	
	
	//var arrar =[]int{1, 2, 3, 4}
	var startNano = time.Now().UnixNano()
	//_  = arrar[0]
	fmt.Println(time.Now().UnixNano() - startNano)
}
