package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Result struct {
	r   *http.Response
	err error
}

func process() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//释放资源
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	resultChan := make(chan Result, 1)
	//发起请求
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println("http request failed, err:", err)
		return
	}
	/*
	   func (c *Client) Do(req *Request) (*Response, error)
	*/
	go func() {
		resp, err := client.Do(req)

		fmt.Println("============>", resp)
		if resp != nil {
			pack := Result{r: resp, err: err}
			//将返回信息写入管道(正确或者错误的)
			resultChan <- pack
		} else {
			cancel()
		}

	}()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		er := <-resultChan
		fmt.Println("Timeout!", er.err)
	case res := <-resultChan:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Server Response: %s", out)
	}
	return
}
func main() {

	var num uint = 1

	fmt.Println("num:", byte(num))

	process()
}
