package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	)

func main() {
	resp, err := http.Get("http://www.sina.com")
	if resp != nil {
		fmt.Println("close ")
		defer resp.Body.Close()
	}
//	defer resp.Body.Close()//not ok
	if err != nil {
		fmt.Println(err)
		return
	}
	
	body, err := ioutil.ReadAll(resp.Body)
	var newBody = ioutil.NopCloser(bytes.NewReader((body)))
	//fmt.Println(string(body), len(body))
	fmt.Println("11==>", err, len(body))
	
	
	body2, err2 := ioutil.ReadAll(newBody)
	
	//fmt.Println("22222222:", string(body2),err2,  len(body2))
	fmt.Println("22222222:",err2 , len(body2))
}
