package main

import (
	"github.com/go-redis/redis"
	"fmt"
)

func main()  {

	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:[]string{"127.0.0.1:6380", "127.0.0.1:6381", "127.0.0.1:6382"},
	})
	var pipe = clusterClient.Pipeline()
	pipe.Ping()
	pipe.MGet("{h1}.aaa", "{h1}.bbb", "{h1}.ccc")
	pipe.Get("h3")
	fmt.Println(pipe.Exec())
	
	
}
