package main

import (
	"fmt"
	"time"
)



type s1 interface {
	id() int
}

type conn struct {
	m map[int]int
	num int
}

func (s* conn) id()int{
	return s.num
}


func StartTimer(s s1) *time.Timer {
	
	timer := time.AfterFunc(2 * time.Second, func() {
		//fmt.Println("timer callback ======", s.id(), " ", globalT.Stop())
		
	})
	return timer
}




func StartRetryTimer(ttl int64) * time.Timer{
	
	return time.AfterFunc(4*time.Second, func() {
		if time.Now().Unix() >= ttl {
			fmt.Println("timer ttl delete")
			return
		}
		fmt.Println("time not ttl, prepare:", ttl)
		if ttl == 0 {
			fmt.Println("time not ttl, continue")
			StartRetryTimer(ttl)
		}

	})
}



type RegionCount struct {
	Zone        string `json:"zone" bson:"zone"`
	DeviceCount int    `json:"deviceCount" bson:"deviceCount"`
	GateCount   int    `json:"gateCount" bson:"gateCount"`
}


func main()  {
	
	var m = []int{0, 1, 2, 3, 4, }
	var i = 0
	for ; i < len(m); i ++ {
	
	}
	
	
	fmt.Println("i:", i, "len:", len(m),  "\n",m [:0], "\n", m [:1], "\n",  m[:i])
	
	
	var appRegionInfo  = make(map[string] *RegionCount)
	
	appRegionInfo["t"] =  &RegionCount{
		DeviceCount:100,
		GateCount: 101,
	}
	
	fmt.Println("before:", )
	fmt.Println("aftrer:", )
	for k, v := range appRegionInfo {
		fmt.Println(k, v)
	}
	if tmpRegion, ok := appRegionInfo["t"]; ok {
		tmpRegion.DeviceCount = 1
		tmpRegion.GateCount = tmpRegion.GateCount + 7
	} else {
	
	}
	
	
	fmt.Println("aftrer:", )
	for k, v := range appRegionInfo {
		fmt.Println(k, v)
	}
	
	
	StartRetryTimer(time.Now().Unix() + 19)
	
	
	fmt.Println(time.Now().Second(), " ", time.Now().Unix())
	
	
	var slice1 = [] int{1, 2, 3}
	var slice2 = slice1
	
	fmt.Println(slice1)
	
	fmt.Println("after ")
	slice2[0] = 111
	fmt.Println(slice1)
	
	
	
	
	var s = &conn{
		num : 1,
	}


	var _ = StartTimer(s)
	fmt.Println("hello world", )
	
	time.Sleep(18 * time.Second)
}
