//MIT License
//
//Copyright (c) 2018 XiaYanji
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.
package main

import (
	"github.com/qiniu/log"

	"flag"
	"fmt"
	lamgo "labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"qiniu.com/pili/common/mgo"
)

var (
	mongoAddr string
	timeSmall int64
	timeBig   int64
	keyId     string
	opType    string
	searchDB  string
)

type M map[string]interface{}

type Oplog struct {
	//TimeSpace       int64  `json:"ts" bson:"ts"`
	//Hash            int64  `json:"-" bson:"-"`
	//Version         int64  `json:"-" bson:"-"`
	OperationType   string `json:"op" bson:"op"`
	NameSpace       string `json:"ns" bson:"ns"`
	Operation       M      `json:"o" bson:"o"`
	UpdateOperation M      `json:"o2" bson:"o2"`
}

func oplogCloneSession(s mgo.Mongo) *lamgo.Collection {

	sess := s.Coll().Database.Session.Copy()
	return sess.DB("local").C("oplog.rs")
}

func dataCloneSession(s mgo.Mongo) *lamgo.Collection {

	sess := s.Coll().Database.Session.Copy()
	return sess.DB("pili").C("domaininfos")
}

func search() {

	opSession, err := mgo.New(mgo.Option{
		MgoAddr:     mongoAddr,
		MgoDB:       "local",
		MgoColl:     "oplog.rs",
		MgoPoolSize: 2,
	})
	if err != nil {
		log.Fatal("err:", err)
	}
	defer opSession.Close()

	if opType == "d" || opType == "i" || opType == "u" {

		opCollSession := oplogCloneSession(opSession)
		//opColl2 := opSession.Coll().Database.Session.Clone()
		//opCollSession := opColl2.DB("local").C("oplog.rs")
		//opQuery1 := M{"_id": bson.ObjectId("5b4eb925809860d11d003dd2")}
		//log.Info(opQuery1)
		//opQuery2 := M{"op": "d", "o":M{"_id":3} }
		//opQuery2 := M{"ts": M{"$gt": bson.MongoTimestamp(0)}}
		opQuery2 := M{"ts": M{"$gt": bson.MongoTimestamp(timeSmall), "$lt": bson.MongoTimestamp(timeBig)}}
		opQuery2["op"] = opType
		opQuery2["o._id"] = keyId
		//opQuery2["o"] = M{"_id":3}

		//log.Info("	bson.MongoTimestamp(m.TimeStamp)===>", 	bson.MongoTimestamp(1533181522 << 32 ))
		//log.Info(6584964495821504515 >> 32)

		var ret Oplog

		//err = opCollSession.Coll().Find(opQuery2).LogReplay().Sort("-$natural").One(&ret)
		queryRuslt := opCollSession.Find(opQuery2).LogReplay()

		count, err2 := queryRuslt.Count()
		iter := queryRuslt.Iter()
		for iter.Next(&ret) {
			fmt.Println("result:", ret, " err2:", err2, " count:", count)
		}
	} else {
		dbSess1 := opSession.Coll().Database.Session.Copy().DB(searchDB)
		allCollections, err := dbSess1.CollectionNames()
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		query := M{"_id": keyId}
		fmt.Println("searchDB:", searchDB, " collections:", allCollections)
		for _, collection := range allCollections {
			tmpTableSession := dbSess1.C(collection)
			queryResult := tmpTableSession.Find(query)
			iter := queryResult.Iter()
			var ret interface{}
			for iter.Next(&ret) {
				fmt.Println("currentTable", collection, "result:", ret, " query:", query)
			}
		}

	}
}
func init() {
	flag.StringVar(&mongoAddr, "addr", "", "example:127.0.0.1:27017")
	flag.Int64Var(&timeSmall, "timeSmall", -1, "search time range >= ")
	flag.Int64Var(&timeBig, "timeBig", -1, "search time range <= ")
	flag.StringVar(&keyId, "key", "", "search key")
	flag.StringVar(&opType, "op", "", "search key")
	flag.StringVar(&searchDB, "db", "", "search db")
	flag.Parse()

	if mongoAddr == "" || keyId == "" || opType == "" {
		log.Fatal("please input mongo server addr or search key")
	}

	if timeSmall == -1 || timeBig == -1 || timeSmall < 0 || timeBig < 0 || timeSmall > timeBig {
		log.Fatal("please input avail number range")
	}

}

func test2() {

	opSession, err := mgo.New(mgo.Option{
		MgoAddr:     mongoAddr,
		MgoDB:       "local",
		MgoColl:     "oplog.rs",
		MgoPoolSize: 2,
	})

	if err != nil {
		log.Fatal("err:", err)
	}
	defer opSession.Close()

	if opType == "d" || opType == "i" || opType == "u" {

		opCollSession := oplogCloneSession(opSession)
		var ret1 Oplog
		//先找到一条oplog
		searchCondition1 := M{"op": "d", "o._id": bson.ObjectIdHex("5b6d13b37a3b60eeb19d4a12")}
		err1 := opCollSession.Find(searchCondition1).One(&ret1)

		log.Info(err1, " result11111111 ==>", ret1, " begin  operation:", ret1.Operation)
		oldOperation := M{"_id": ret1.Operation["_id"]}

		//根据这条oplog的id反查
		dataSession := dataCloneSession(opSession)
		id, content := ret1.Operation["_id"], ret1.Operation
		err2 := dataSession.FindId(id).One(&content)
		log.Info(err2, " result3333333333333 ==>", content, " final oldOperation=====>:", ret1.Operation, " BBBBBak:", oldOperation)

		err3 := dataSession.Remove(oldOperation)

		log.Info(err3, " result4444444444 ==>")
		//M{"_id":bson.ObjectIdHex("5b6d13b37a3b60eeb19d4a11")}
		//err3 := dataSession.FindId(id).One(&ret2)
		//log.Info(err3, " result444444444444 ==>",ret2, " oldOperation:", ret1.Operation)

	}
}

func main() {
	//search()

	test2()

}
