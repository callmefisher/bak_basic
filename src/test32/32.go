package main

import (
	"flag"
	"fmt"
	"github.com/globalsign/mgo"
	"strings"

	//"github.com/globalsign/mgo/bson"
	"time"
)

var (
	mgoAddr   string
	dbName    string
	table     string
	index     string
	unique    bool
	sparse    bool
	expireSec int
	dropIndex bool
)

type App struct {
	Id  string `json:"id" bson:"_id"`
	Uid uint32 `json:"-" bson:"uid"`
	App string `json:"app" bson:"app"`
}

func main() {

	flag.StringVar(&mgoAddr, "addr", "", "-addr='10.20.38.28:7001,10.20.36.39:7001,10.20.38.31:7001'")
	flag.StringVar(&dbName, "db", "", "-db=xxx")
	flag.StringVar(&table, "c", "", "-c=xxx")
	flag.StringVar(&index, "index", "", "-index=_id,...")
	flag.BoolVar(&unique, "unique", false, "-unique=true/false")
	flag.BoolVar(&sparse, "sparse", false, "-sparse=true/false")
	flag.BoolVar(&dropIndex, "dropindex", false, "-dropindex=true/false")
	flag.IntVar(&expireSec, "expire", -1, "-expire=123")
	flag.Parse()
	if dbName == "" || table == "" || mgoAddr == "" || index == "" {
		fmt.Println("error fmt , see   --help ")
		return
	}

	session, err := mgo.DialWithTimeout(mgoAddr, 3*time.Second)
	if err != nil {
		fmt.Println("mongo connect failed:", err)
		return
	}
	var readyIndexes = strings.Split(index, ",")
	if len(readyIndexes) == 0 {
		fmt.Println("please input index para")
		return
	}

	defer session.Close()
	var s1 = session.DB(dbName).C(table)

	if dropIndex {
		err = s1.DropAllIndexes()
		if err != nil {
			fmt.Println("step1:drop old index fail:", err)
			return
		}
	}

	var mgoIndex = mgo.Index{Key: readyIndexes}
	if unique {
		mgoIndex.Unique = true
	}
	if sparse {
		mgoIndex.Sparse = true
	}
	if expireSec >= 0 {
		mgoIndex.ExpireAfter = time.Duration(expireSec) * time.Second
	}
	fmt.Println("step2:start create index :", mgoIndex)
	err = s1.EnsureIndex(mgoIndex)
	if err != nil {
		fmt.Println("step2:create index error:", err)
		return
	}

	allCreateIndex, err := s1.Indexes()
	fmt.Println("step3:after create index len:", len(allCreateIndex), " error:", err)
	for k, v := range allCreateIndex {
		fmt.Println("k:", k, " v:", v.Key, "isUnique:", v.Unique, " isSparse", v.Sparse, " expire:", v.ExpireAfter)
	}

	fmt.Println("done!")

}
