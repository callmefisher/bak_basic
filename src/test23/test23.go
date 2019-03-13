package main

import (
	"fmt"
	"net/http"
	"crypto/hmac"
	"crypto/sha1"
	"io"
	"encoding/base64"
	"math/rand"
	"time"
)


var roundNum = 0


type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	
	go func(bbbbbbbbb string) {
		fmt.Println("url_path:", bbbbbbbbb)
	}(req.URL.String())
	req = nil
	w.Write([]byte("Hello, world!"))
}

func GetRandomName() {
	var domains = []string{"a", "b", "c",}
	
	rand.Seed(time.Now().UnixNano())
	
	const randCount = 10
	//var randArray [randCount] int
	//
	for i:= 0; i < randCount; i++ {
		fmt.Println( domains[roundNum % len(domains)])
		roundNum++
	}
	
	
	//for i := 0; i < len(domains) / 2; i ++ {
	//
	//	domains[i], domains[roundNum] = domains[roundNum], domains[i]
	//}
	
	
	
	fmt.Println(domains)
	
	
	
}

func main() {
	http.Handle("/", &helloHandler{})
	
	
	var info = ""
	switch info {
	case "", "-":
		fmt.Println("111111111111")
	default:
		fmt.Println("22222222222222")
	}
	
	
	h := hmac.New(sha1.New, []byte("2"))
	io.WriteString(h, "http://www.baidu.com")
	
	fmt.Println( base64.URLEncoding.EncodeToString(h.Sum(nil)))
	
	
	GetRandomName()
	//http.ListenAndServe(":12345", nil)
}
