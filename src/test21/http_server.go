package main

import (
	"fmt"
	"net/http"
	"net/url"
)

/*
一、Client-Get



package main

import (
"fmt"
"net/url"
"net/http"
"io/ioutil"
"log"
)

func main() {
	u, _ := url.Parse("http://localhost:9001/xiaoyue")
	q := u.Query()
	q.Set("username", "user")
	q.Set("password", "passwd")
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String());
	 if err != nil {
		log.Fatal(err) return
	}
	 result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err) return
	}
	fmt.Printf("%s", result)
}


二、Client-Post

package main

import (
"fmt"
"net/url"
"net/http"
"io/ioutil"
"log"
"bytes"
"encoding/json"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
	ServersID  string
}


func main() {

	var s Serverslice

	 var newServer Server;
	newServer.ServerName = "Guangzhou_VPN";
	newServer.ServerIP = "127.0.0.1"
	s.Servers = append(s.Servers, newServer)

	        s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.2"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.3"})

	        s.ServersID = "team1"

	        b, err := json.Marshal(s)
	        if err != nil {
		                fmt.Println("json err:", err)
		        }

	        body := bytes.NewBuffer([]byte(b))
	        res,err := http.Post("http://localhost:9001/xiaoyue", "application/json;charset=utf-8", body)
	        if err != nil {
		                log.Fatal(err)
		                return
		        }
	        result, err := ioutil.ReadAll(res.Body)
	        res.Body.Close()
	        if err != nil {
		                log.Fatal(err)
		                return
		        }
	        fmt.Printf("%s", result)
}

三、Server



package main

import (
"fmt"
"net/http"
"strings"
"html"
"io/ioutil"
"encoding/json"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
	ServersID  string
}

func main() {
	      http.HandleFunc("/", handler)
	http.ListenAndServe(":9001", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Fprintf(w, "Hi, I love you %s", html.EscapeString(r.URL.Path[1:]))
	 if r.Method == "GET" {
		 fmt.Println("method:", r.Method) //获取请求的方法

		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])

		for k, v := range r.Form {
			 fmt.Print("key:", k, "; ")
			 fmt.Println("val:", strings.Join(v, ""))
			 }
	} else if r.Method == "POST" {
		  result, _:= ioutil.ReadAll(r.Body)
		 r.Body.Close()
		 fmt.Printf("%s\n", result)

		//未知类型的推荐处理方法

		 var f interface{}
		 json.Unmarshal(result, &f)
		m := f.(map[string]interface{})
		 for k, v := range m {
			 switch vv := v.(type) {
				 case string:
				 fmt.Println(k, "is string", vv)
				 case int:
				 fmt.Println(k, "is int", vv)
				 case float64:
				 fmt.Println(k,"is float64",vv)
				 case []interface{}:
				 fmt.Println(k, "is an array:")
				 for i, u := range vv {
					 fmt.Println(i, u)
					 }
				 default:
				 fmt.Println(k, "is of a type I don't know how to handle")
			}
			 }

		//结构已知，解析到结构体

		 var s Serverslice;
		                 json.Unmarshal([]byte(result), &s)

		                 fmt.Println(s.ServersID);

		                 for i:=0; i<len(s.Servers); i++ {
			                         fmt.Println(s.Servers[i].ServerName)
			                         fmt.Println(s.Servers[i].ServerIP)
			                 }
		        }
}

*/

type A struct {
	Host string
}

func main() {
	Url := "rtmp://pili-publish.qiniu1.wawazhua.com:1250/javava-online-002/j001-stream-00085-1405128A?e=1564446614\u0026token=QEYBv8GhD5xuzczGFpRA-Tusn8E6rXl6mE_U8dQt:NtHUb0YKQouYXEb6i6-Us4tkaDU=\u0026noforward__=true"
	u2, err := url.Parse(Url)
	//var a *A
	fmt.Println(u2, err)

	fmt.Println(AppendQuery(Url, "nosegmenter__", "true"))

	preHost, preUrl := "req.Host", "req.URL.String()"

	fmt.Println("========>", preHost, preUrl)

	//
	http.HandleFunc("/foo", testHandler)
	http.ListenAndServe(":9999", nil)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("11111111111")
}

func AppendQuery(in, k, v string) string {

	uu, err := url.Parse(in)
	if err == nil {
		if uu.RawQuery != "" {
			uu.RawQuery += "&"
		}
		k, v = url.QueryEscape(k), url.QueryEscape(v)
		uu.RawQuery += fmt.Sprintf("%s=%s", k, v)
		return uu.String()
	}

	return in
}
