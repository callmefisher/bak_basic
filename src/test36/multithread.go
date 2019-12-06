package main

import (
	"flag"
	"fmt"
	"github.com/globalsign/mgo"
	"math"
	"sync"

	"time"
)

var (
	SrcMgoAddr string
	SrcDbName  string
	SrcTable   string

	DstMgoAddr string
	DstDbName  string
	DstTable   string
	WaterMark  int
)

func main() {
	
	
	flag.StringVar(&SrcMgoAddr, "src", "10.20.38.28:7001", "-src='10.20.38.28:7001,'")
	flag.StringVar(&DstMgoAddr, "dst", "10.20.38.28:7011,10.20.36.39:7011", "-dst='10.20.38.28:7011,10.20.36.39:7011'")
	flag.StringVar(&SrcDbName, "srcdb", "", "-srcdb=xxx")
	flag.StringVar(&DstDbName, "dstdb", "", "-dstdb=xxx")
	flag.StringVar(&SrcTable, "srctable", "", "-srctable=xxx")
	flag.StringVar(&DstTable, "dstable", "", "-dstable=xxx")
	flag.IntVar(&WaterMark, "watermark", 40000, "-watermark=xxx")
	flag.Parse()
	if SrcMgoAddr == "" || DstMgoAddr == "" || SrcDbName == "" || DstDbName == "" || SrcTable == "" || DstTable == "" {
		fmt.Println("error fmt , see   --help ")
		return
	}

	srcSession, err := mgo.DialWithTimeout(SrcMgoAddr, 3*time.Second)
	if err != nil {
		fmt.Println("mongo connect failed:", err)
		return
	}

	defer srcSession.Close()

	dstSession, err := mgo.DialWithTimeout(DstMgoAddr, 3*time.Second)
	if err != nil {
		fmt.Println("mongo connect failed:", err)
		return
	}

	defer dstSession.Close()

	var srcT = srcSession.DB(SrcDbName).C(SrcTable)
	var dstT = dstSession.DB(DstDbName).C(DstTable)
	var change, err1 = dstT.RemoveAll(nil)
	if err1 != nil {
		fmt.Println(" drop old collection err:", DstDbName, DstTable, change, err1)
		fmt.Println("\n")
		return
	}

	var query interface{}
	var result []interface{}
	var err2 = srcT.Find(query).All(&result)

	if err2 != nil {

		fmt.Println("find data err:", SrcDbName, SrcTable, err2)
		fmt.Println("\n")
		return
	}

	var dataLen = len(result)
	var goRoutineCount = 1
	if dataLen > WaterMark {
		goRoutineCount = int(math.Ceil(float64(dataLen) / float64(WaterMark)))
	}
	fmt.Println("dataLen:", dataLen, " goroutine count:", goRoutineCount, " dstTable:", DstTable)
	//var lock sync.RWMutex
	var wg sync.WaitGroup
	wg.Add(goRoutineCount)

	var startTick = time.Now()

	for i := 0; i < goRoutineCount; i++ {

		go func(tag int) {
			defer wg.Done()

			var tmpRange = tag*WaterMark + WaterMark
			var rightRange = WaterMark
			if tmpRange > dataLen {
				rightRange = dataLen - tag*WaterMark
			}
			var lastIndex = tag*WaterMark + rightRange
			var dataSlice = result[tag*WaterMark : lastIndex]
			//lock.RLock()
			//fmt.Println("left:", tag*waterMark, " right:", rightRange, " table:", DstTable, " tag:", i, " ",
			//	len(dataSlice))
			//lock.RUnlock()

			for _, b := range dataSlice {
				var err3 = dstT.Insert(b)
				if err3 != nil {
					fmt.Println("insert data err:", DstDbName, DstTable, len(result), err3)
					fmt.Println("\n")
					return
				}
			}

			if lastIndex == dataLen {
				fmt.Println("insert last one:", result[dataLen-1])
			}

		}(i)
	}

	wg.Wait()

	srcDataCount, err := srcT.Count()
	if err != nil {
		fmt.Println("\n")
		return
	}
	fmt.Println("srcdb:", SrcDbName, " srctable:", SrcTable, " count:", srcDataCount)
	//dst

	fmt.Println("insert to db:", DstDbName, " table:", DstTable, " count:", len(result), " cost:", time.Now().Sub(startTick))
	fmt.Println("\n")

}
