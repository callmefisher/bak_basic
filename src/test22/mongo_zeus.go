package main

import (
	"log"

	"flag"
	"fmt"
	lamgo "labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"qiniu.com/pili/common/mgo"
	"time"
	"context"
)

var (
	mongoAddr string
	parentIdc string
	targetIdc string
)

const (
	dbName   string = "rtc"
	collName string = "udpaccgateroute"
)

type OnePoint struct {
	IdcName string `json:"idc" bson:"idc"`
}

type RouteTable struct {
	Id     bson.ObjectId `json:"_id" bson:"_id"`
	SrcIdc string        `json:"srcIdc" bson:"srcIdc"`
	DstIdc string        `json:"destIdc" bson:"destIdc"`
	Level  int           `json:"level" bson:"level"`
	Path   [][]OnePoint  `json:"paths" bson:"paths"`
}

func dataCloneSession(s mgo.Mongo) *lamgo.Collection {

	sess := s.Coll().Database.Session.Copy()
	return sess.DB(dbName).C(collName)
}

func init() {
	flag.StringVar(&mongoAddr, "addr", "", "example:127.0.0.1:27017")
	flag.StringVar(&parentIdc, "parent", "", "src idc")
	flag.StringVar(&targetIdc, "dst", "", " dst idc")
	flag.Parse()

	if mongoAddr == "" || parentIdc == "" || targetIdc == "" {
		//		log.Fatal("please input mongo server addr or search key")
	}

}

func testPathAvail(points []OnePoint, srcIdc, targetIdc, parentIdc string) bool {

	var lastIdc = srcIdc
	var includeTarget = false
	for _, v := range points {
		if v.IdcName == targetIdc {
			includeTarget = true
			break
		}
		lastIdc = v.IdcName
	}

	if includeTarget {
		if lastIdc != parentIdc {
			return false
		}
	}

	return true
}

func test2() {

	opSession, err := mgo.New(mgo.Option{
		MgoAddr:     mongoAddr,
		MgoDB:       dbName,
		MgoColl:     collName,
		MgoPoolSize: 2,
	})

	if err != nil {
		log.Fatal("err:", err)
	}
	defer opSession.Close()

	dataSession := dataCloneSession(opSession)
	var query interface{}
	queryResult := dataSession.Find(query)
	iter := queryResult.Iter()
	var ret RouteTable
	var count = 0
	for iter.Next(&ret) {
		//fmt.Println("id:", ret.Id, "src:", ret.SrcIdc, " dst:", ret.DstIdc, " level:", ret.Level, " path:", ret.Path)
		if ret.DstIdc == targetIdc {
			if ret.SrcIdc != parentIdc {
				count++
				fmt.Println("direct id:", ret.Id, "src:", ret.SrcIdc, " dst:", ret.DstIdc, " level:", ret.Level)

			}
		} else {

			for _, v := range ret.Path {

				var flag = testPathAvail(v, ret.SrcIdc, targetIdc, parentIdc)

				if !flag {
					count++
					fmt.Println("not direct:", ret.Id, "src:", ret.SrcIdc, " dst:", ret.DstIdc, " level:", ret.Level, " path:", ret.Path)
				}
			}
		}
	}

	fmt.Println("target:", targetIdc, " parent:", parentIdc, " count:", count)

}

func test3() {
	var i int
	if true {

		if i == 1 {
			goto label
		}

		return
	}

label:
}

func test4()  {
	var m = make(map[string]int)
	m["1"]= 10
	m ["19"] = 1
	//var v = 111111
	if v, ok := m["19"]; !ok || v < 0 || v > 5 {
		fmt.Println("========================================:", v, " ok:", ok)
	}
	
	fmt.Println(m["0"], " : ", m["19"] )
}

type D struct {
   a * A;
}



type  A struct {
	//num int;
	str string;
}



type B struct {
	a string;
}

func testFunc(str string)  {
	
	fmt.Println(len(str))
}

func testF2()  {
	//fmt.Println("hello world2")
}

func (d *D )stack1(b * B, num int, str string )   {
	
	var newCtx = context.Background()
	_, cancel := context.WithCancel(newCtx)
	
	
	defer testF2()
	defer cancel()
	//testFunc(d.a.str)
	// panic("op")
	var chain = make(chan bool, 1)
	go func() {
		select {
		case <-chain:
		
		case <-time.After(1 * time.Second):
			testF2()
			cancel()
		}
	}()
	
	
	
	
	
	return
}





func main() {
	var d= &D{}
	//	a = nil
	d.stack1(&B{}, 1, "")
	
}