package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
)


type T2 struct {

}



type T1 struct {
	Addrs     []string
	m1 map[int]int
	m2 map[int] int
}


func main() {
	
	
	var tmp [][]int
	
	fmt.Println(tmp[0])
	
	
	
	go func() {
		
		
		mux := http.NewServeMux()
		
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
		http.ListenAndServe("127.0.0.1:6070", nil)
		
		
	}()
	
	
	fmt.Println("hello world")
	
	for ; ;  {
		select {
		
		
		}
	}
}
