package main

import (
	"flag"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"strconv"

	"time"
)

type M bson.M

var (
	SrcMgoAddr string
	SrcDbName  string
	SrcTable   string

	DstMgoAddr  string
	DstDbName   string
	DstTable    string
	GreaterThan int
)

var LAYOUT = "2006-01-02T15:04:05.000Z"
var LAYOUT1 = "2006/01/02 15:04:05.000000"
var convertInteger = 1565767236058

func convertTime(milSec int) (time.Time, string) {

	var timeStr = strconv.Itoa(int(milSec))
	handleMil1 := timeStr[0:10]
	handleMil2 := timeStr[10:]

	convertInteger1, _ := strconv.Atoi(handleMil1)
	convertInteger2, _ := strconv.Atoi(handleMil2)

	tm := time.Unix(int64(convertInteger1), int64(convertInteger2*1e6))
	fmt.Println(timeStr, "===>", tm.Format(LAYOUT))
	return tm, tm.Format(LAYOUT)
}

type Segment struct {
	Id string `bson:"_id"` // did:start

	Did         string            `json:"-" bson:"-"`
	Start       int               `json:"start" bson:"start"`     // 毫秒
	End         int               `json:"end" bson:"end"`         // 毫秒
	Session     string            `json:"session" bson:"session"` // 切片会话
	VD          int               `json:"vd" bson:"vd"`           // 切片视频内容时长, 毫秒
	AD          int               `json:"ad" bson:"ad"`           // 切片音频内容时长, 毫秒
	Meta        map[string]string `json:"meta" bson:"meta"`
	EndReason   string            `json:"endReason" bson:"endReason"`
	Sequence    int               `json:"sequence" bson:"sequence"` // 最大的切片序列号
	ExpireAt    time.Time         `json:"-" bson:"expireAt"`
	FrameStatus int               `json:"frameStatus" bson:"frameStatus"` // 0 表示成功， 1 表示timeout导致失败， 2 客户端强制停止上传， 3 客户端异常导致停止上传
}

func main() {

	convertTime(convertInteger)

	flag.StringVar(&SrcMgoAddr, "src", "10.20.38.28:7001", "-src='10.20.38.28:7001,'")
	flag.StringVar(&DstMgoAddr, "dst", "10.20.38.28:7011", "-dst='10.20.38.28:7011,'")
	flag.StringVar(&SrcDbName, "srcdb", "linking-segment", "-srcdb=linking-segment")
	flag.StringVar(&DstDbName, "dstdb", "linking-segment", "-dstdb=linking-segment")
	flag.StringVar(&SrcTable, "srctable", "segment", "-srctable=segment")
	flag.StringVar(&DstTable, "dstable", "segment", "-dstable=segment")
	flag.IntVar(&GreaterThan, "gt", 1565933228408, "-gt=1565933228408")
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

	srcSession.SetSocketTimeout(300 * time.Second)
	dstSession, err := mgo.DialWithTimeout(DstMgoAddr, 3*time.Second)
	if err != nil {
		fmt.Println("mongo connect failed:", err)
		return
	}

	defer dstSession.Close()
	var srcT = srcSession.DB(SrcDbName).C(SrcTable)
	var dstT = dstSession.DB(DstDbName).C(DstTable)
	//dst
	dstDataCount, err4 := dstT.Count()
	if err4 != nil {
		fmt.Println("done insert to new table count:", DstDbName, DstTable, err4)
		fmt.Println("\n")
		return
	}
	var query interface{}
	var result []interface{}
	var allSegresult []Segment

	if GreaterThan > 0 {

		if DstTable == "historyactivity" {
			tm, _ := convertTime(GreaterThan)
			query = M{"date": M{"$gt": tm}}
			srcT.Find(query).All(&result)
		} else if DstTable == "segment" {
			query = M{"start": M{"$gt": GreaterThan}}
			srcT.Find(query).All(&allSegresult)
		}

	}

	var validInsert int
	var sameIdCount int
	if DstTable == "historyactivity" {
		validInsert = len(result)
		for k, b := range result {
			var err3 = dstT.Insert(b)
			if err3 != nil {
				if mgo.IsDup(err3) {
					validInsert--
					continue
				}
				fmt.Println("insert data err:", DstDbName, DstTable, len(result), err3)
				fmt.Println("\n")
				return
			}
			if k == len(result)-1 {
				fmt.Println("last item:", b)
			}
		}
	} else {

		validInsert = len(allSegresult)
		for k, OldSeg := range allSegresult {
			var err3 = dstT.Insert(OldSeg)
			if err3 != nil {
				if mgo.IsDup(err3) {
					var tmpNewSegresult Segment
					dstT.FindId(OldSeg.Id).One(&tmpNewSegresult)
					if OldSeg.End > tmpNewSegresult.End {
						dstT.UpsertId(OldSeg.Id, OldSeg)
						sameIdCount++
						fmt.Println("update old :", OldSeg, " new:", tmpNewSegresult)
					} else {
						validInsert--
					}

					continue
				}
				fmt.Println("insert data err:", DstDbName, DstTable, len(result), err3)
				fmt.Println("\n")
				return
			}
			if k == len(result)-1 {
				fmt.Println("last item:", OldSeg)
			}
		}
	}

	srcDataCount, err := srcT.Count()
	if err != nil {
		fmt.Println("\n")
		return
	}
	fmt.Println("srcdb:", SrcDbName, " srctable:", SrcTable, " count:", srcDataCount)

	fmt.Println("insert to db:", DstDbName, " table:", DstTable, " count:", dstDataCount,
		" remidy:", validInsert, " result len:", len(result), " finally:", dstDataCount+validInsert, " sameId:", sameIdCount)
	fmt.Println("\n")

}
