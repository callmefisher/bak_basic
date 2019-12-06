package main

import (
	"flag"
	"fmt"
	"github.com/globalsign/mgo"

	"time"
)

var (
	SrcMgoAddr string
	SrcDbName  string
	SrcTable   string

	DstMgoAddr string
	DstDbName  string
	DstTable   string
)

func main() {

	flag.StringVar(&SrcMgoAddr, "src", "10.20.38.28:7001", "-src='10.20.38.28:7001,'")
	flag.StringVar(&DstMgoAddr, "dst", "10.20.38.28:7011", "-dst='10.20.38.28:7011,'")
	flag.StringVar(&SrcDbName, "srcdb", "", "-srcdb=xxx")
	flag.StringVar(&DstDbName, "dstdb", "", "-dstdb=xxx")
	flag.StringVar(&SrcTable, "srctable", "", "-srctable=xxx")
	flag.StringVar(&DstTable, "dstable", "", "-dstable=xxx")
	flag.Parse()
	if SrcMgoAddr == "" || DstMgoAddr == "" || SrcDbName == "" || DstDbName == "" || SrcTable == "" || DstTable == "" {
		fmt.Println("error fmt , see   --help ")
		return
	}

	var query interface{}
	var result []interface{}

	var i = 0
	for {

		i++
		srcSession, err := mgo.DialWithTimeout(SrcMgoAddr, 3*time.Second)
		if err != nil {
			fmt.Println("mongo connect failed:", err)
			break
		}
		srcSession.SetSocketTimeout(300 * time.Second)

		dstSession, err := mgo.DialWithTimeout(DstMgoAddr, 3*time.Second)
		if err != nil {
			fmt.Println("mongo connect failed:", err)
			break
		}
		dstSession.SetSocketTimeout(300 * time.Second)
		var srcT = srcSession.DB(SrcDbName).C(SrcTable)
		var err2 = srcT.Find(query).Limit(10000)
		srcSession.Close()
		fmt.Println("step :", i)

		if err2 != nil {

			fmt.Println("find data err:", SrcDbName, SrcTable, err2)
			fmt.Println("\n")
			continue
		}

		var dstT = dstSession.DB(DstDbName).C(DstTable)

		for k, b := range result {
			var err3 = dstT.Insert(b)
			if err3 != nil {
				if mgo.IsDup(err3) {
					continue
				}
				fmt.Println("insert data err:", DstDbName, DstTable, len(result), err3)
				fmt.Println("\n")
				return
			}
			if k == len(result)-1 {
				fmt.Println("insert last one:", k, b)
			}
		}

		dstSession.Close()

		if i > 200 {
			break
		}

	}

	srcSession, err := mgo.DialWithTimeout(SrcMgoAddr, 3*time.Second)
	if err != nil {
		fmt.Println("mongo connect failed:", err)
		return
	}
	srcSession.SetSocketTimeout(300 * time.Second)
	defer srcSession.Close()
	var srcT = srcSession.DB(SrcDbName).C(SrcTable)
	srcDataCount, err := srcT.Count()
	if err != nil {
		fmt.Println(" err:", err)
		return
	}

	dstSession, err := mgo.DialWithTimeout(DstMgoAddr, 3*time.Second)
	if err != nil {
		fmt.Println("mongo connect failed:", err)
		return
	}
	defer dstSession.Close()
	dstSession.SetSocketTimeout(300 * time.Second)
	var dstT = dstSession.DB(DstDbName).C(DstTable)
	dstDataCount, err := dstT.Count()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("srcdb:", SrcDbName, " srctable:", SrcTable, " count:", srcDataCount)
	//dst

	fmt.Println("insert to db:", DstDbName, " table:", DstTable, " count:", dstDataCount)

	fmt.Println("\n")

}
