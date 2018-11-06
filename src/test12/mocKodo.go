package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"
	"time"
)

func mocKodo(w http.ResponseWriter, r *http.Request) {

	fmt.Println("")
	fmt.Println("======================")
	r.ParseForm()
	fmt.Println("form:", r.Form)
	fmt.Println("method:", r.Method)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	fmt.Println("url:", r.URL)
	fmt.Println("path:", r.URL.Path)

	if r.Method == "POST" {
		var result interface{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&result)
		if err == nil {
			fmt.Println("post err:", err, " result:", result)
		}
	} else if r.Method == "GET" {
		fmt.Println("port:", r.URL.Port())
		query := r.URL.Query()
		fmt.Println("query:", query)
		fmt.Println("query:", query.Get("aa"))
	}

	w.Write([]byte("hello world curl"))
}

type item struct {
	str string
	num int
}

func testIsPrime(num int) {

	var sqrtValue = make([]int, 0, num*2)
	for i := 1; i < num*2; i++ {
		//sqrtValue[i] = int(math.Sqrt(float64(i)))
		sqrtValue = append(sqrtValue, int(math.Sqrt(float64(i))))
	}
	var slice = make([]int, 0, num)

	fmt.Println("start======================")

	start := time.Now()

	if num < 0 {
		return
	}
	if num == 0 {

	} else if num == 1 {

	} else if num == 2 {
		slice = append(slice, num)
	} else {

		slice = append(slice, 2)

		for i := 3; i <= num; i += 2 {

			var j = 0
			//var sqrtNum = sqrtValue[i]
			var sqrtNum = int(math.Sqrt(float64(i)))
			for j = 2; j <= sqrtNum; j++ {
				if i%j == 0 {
					break
				}
			}
			if j > sqrtNum {
				slice = append(slice, i)
			}
		}
	}

	fmt.Println("num:", num, " cost :", time.Since(start), " ", slice)
	fmt.Println("")
}

func outputPrime() {
	testIsPrime(1)
	testIsPrime(1)
	testIsPrime(0)
	testIsPrime(2)
	testIsPrime(3)
	testIsPrime(4)
	testIsPrime(10)
	testIsPrime(37)
	testIsPrime(38)
	testIsPrime(39)
	testIsPrime(100)
	testIsPrime(100000)

}

func testStringCmp() {

	var s1 = "1540535766267-0"
	var s2 = "1540535766267-0"

	var s3 = "1540535766268-0"
	var s4 = "1540535766368-0"
	var s5 = "1540535766368-10"
	var s6 = "1540535767008-10"
	var s7 = "1540635767008-10"

	fmt.Println(s1 >= s2, " ", s3 >= s2, " ", s4 >= s3, " ", s5 >= s4, " ", s6 >= s5, " ", s7 >= s6, " ", int(1540635767008*10000+10))
}

func main() {
	//outputPrime()
	testStringCmp()

	return

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	fmt.Println("slice2:", slice2, " len2:", len(slice2), " cap2:", cap(slice2), " len1:", len(slice1), " ", cap(slice1))
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	//copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置

	fmt.Println("slice2:", slice2, " len2:", len(slice2), " cap2:", cap(slice2), " len1:", len(slice1), " ", cap(slice1))

	var m = make(map[string]*item)
	m["a3"] = &item{str: "a3", num: 3}
	m["a1"] = &item{str: "a1", num: 1}
	m["a2"] = &item{str: "a2", num: 2}

	for key, conf := range m {
		conf.str = key + "111"
	}

	var str3 = "z1.xyjhub1.a1.1539345787286-1539345792168-2-p-wMAC6xnZrU2lwV.ts"

	var slice3 = strings.SplitN(str3, ".", 4)
	for k, v := range slice3 {
		fmt.Println("k:", k, " v:", v)
	}

	var startToEnd = slice3[3]

	fmt.Println("satrt-to-end:", startToEnd)
	var slice4 = strings.SplitN(startToEnd, "-", 4)

	for k, v := range slice4 {
		fmt.Println("k4:", k, " v4:", v)
	}

	v1, ok := m["a1"]
	fmt.Println("a1==>", v1, " ok:", ok, " direct:", m["a1"])
	fmt.Println("\n")
	for key, conf := range m {
		fmt.Println("key:", key, " conf:", conf)
	}

	str := "/jztest11/redis5022"
	slice := strings.SplitN(str, "/", 3)
	fmt.Println("0:", slice[0], " 1:", slice[1], " 2:", slice[2])

	http.HandleFunc("/v1/streamgate/upload2", mocKodo)
	err := http.ListenAndServe(":5000", nil)
	//var err error
	if err != nil {
		fmt.Println("err:", err)
	}
}
