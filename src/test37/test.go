package main

import (
	"fmt"
	"math/rand"
	"os/exec"

	"github.com/globalsign/mgo"

	"time"
)


type PublicModel struct {
	Id            string    `json:"-" bson:"_id"`
	Appid         string    `json:"appid" bson:"appid"`
	DaySec        int64     `json:"daysec" bson:"daysec"`
	DeviceCount   int64     `json:"deviceCount" bson:"deviceCount"`
	GateWayCount  int64     `json:"gatewayCount" bson:"gatewayCount"`
	ExpireAt      time.Time `json:"-" bson:"expireAt"`
	LastUpdateSec int64     `json:"lastUpdate" bson:"lastUpdate"`
}

func main() {
	
	
	var p1 = PublicModel{
		Id:"p1",
	}
	var p2 = p1
	p2.Id = "p2"
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("==================")
	
	var c1 = make(chan  int , 1)
	
	close(c1)
	//c1<-1
	<-c1
	
	
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int31n(20))
	var mongoAddr = "localhost:26007"
	session, err := mgo.DialWithTimeout(mongoAddr, 5*time.Second)
	fmt.Println("uuuuuuuuuuuuuu1:", err)

	db := session.DB("linkingtest")
	db.DropDatabase()
	c := db.C("segments")
	//c2 := db.C("kodo")


	var shellAddr = "mongo "+ mongoAddr

	//var result interface{}

	var c2 = "echo \"sh.enableSharding('linkingtest')\" | " + shellAddr
	var c3 = "echo \"sh.shardCollection('linkingtest.segments', {_id:'hashed'})\" | " + shellAddr

	//var command1 = "mongo 'localhost:26007'/'admin' --eval 'sh.enableSharding('linkingtest')'"

	re, err2 := exec.Command("/bin/sh", "-c", c2).Output()
	re3, err3 := exec.Command("/bin/sh", "-c", c3).Output()

	//var err1 = session.Run(bson.D{{Name: "eval", Value: "sh.enableSharding('linkingtest')"}}, &result)
	//var err1 = session.Run(bson.D{{Name: "eval", Value: "sh.enableSharding('linkingtest')"}}, &result)

	fmt.Println(" aaaa:", 1, " \n err2:", err2, string(re), " \n", string(re3), " \n err3:", err3)
	//var err2 = db.Run(bson.D{{Name: "eval",
	//	Value: "sh.shardCollection('linkingtest.segments', { _id: 'hashed' })"}}, &result)
	//fmt.Println("bbb:", err2, result)
	//var err3 = db.Run(bson.D{{Name: "eval",
	//	Value: "sh.shardCollection('linkingtest.kodo', { _id: 'hashed' })"}}, &result)
	//fmt.Println("ccc:", err3, result)

	//c.Insert(bson.M{"_id":1, "expireAt":time.Now().Unix() + 30})

	c.EnsureIndex(mgo.Index{
		Key:         []string{"expireAt"},
		Sparse:      true,
		ExpireAfter: time.Second,
	})

}
