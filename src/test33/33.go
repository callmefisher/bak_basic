package main

import (
	"flag"
	"fmt"
	"github.com/globalsign/mgo"
	jsoniter "github.com/json-iterator/go"
	"strconv"

	"time"
)

var (
	mgoAddr string
	dbName  string
	table   string
)

type App struct {
	Id  string `json:"id" bson:"_id"`
	Uid uint32 `json:"-" bson:"uid"`
	App string `json:"app" bson:"app"`
}

//1. history

type DeviceHistoryElement struct {
	Id           string    `bson:"_id" json:"-"`
	Date         time.Time `bson:"date,omitempty" json:"-"`
	Appid        string    `bson:"appid" json:"appid"`
	Device       string    `bson:"device" json:"device"`
	LoginAt      int64     `bson:"loginAt" json:"loginAt"`
	LogoutAt     int64     `bson:"logoutAt" json:"logoutAt"`
	LogoutReason string    `bson:"reason" json:"reason"`
	RemoteIp     string    `bson:"remoteIp" json:"remoteIp"`
}

const deviceHistoryElementFmt = "%s:%s:%014d"

func makeid(appid, device string, loginAt int64) string {
	return fmt.Sprintf(deviceHistoryElementFmt, appid, device, loginAt)
}

type DeviceKey struct {
	AccessKey string `json:"accessKey" bson:"accessKey"`
	SecretKey string `json:"secretKey" bson:"secretKey"`
	State     int    `json:"state" bson:"state"`
	CreatedAt int64  `json:"createdAt" bson:"createdAt"`
}

type Channel struct {
	Channelid int    `json:"channelid" bson:"channelid"`
	Comment   string `json:"comment" bson:"comment"`
}
type Device struct {
	Id       string `json:"-" bson:"_id"`
	Device   string `json:"device" bson:"device"`
	LoginAt  int64  `json:"loginAt,omitempty" bson:"-"`  // 设备在线时才会出现该字段
	RemoteIp string `json:"remoteIp,omitempty" bson:"-"` // 设备在线时才会出现该字段
	// 0 不录制
	// -1 永久
	// -2 继承app配置
	SegmentExpireDays int `json:"segmentExpireDays" bson:"segmentExpireDays"`

	// 0 客户端配置上传
	// -1 继承app配置
	// 1 强制持续上传
	// 2 强制关闭上传
	UploadMode int `json:"uploadMode,omitempty" bson:"uploadMode,omitempty"`

	CreatedAt int64 `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64 `json:"updatedAt" bson:"updatedAt"`
	State     int   `json:"state" bson:"state"`

	ActivedAt int64  `json:"activedAt,omitempty" bson:"activedAt"`
	ActivedSn string `json:"activedSn,omitempty" bson:"activedSn,omitempty"`

	Keys []DeviceKey `json:"-" bson:"keys"`

	// 0 免费使用
	// 1 正常收费
	LicenseMode int `json:"licenseMode,omitempty" bson:"licenseMode,omitempty"`

	// batchId
	Batch string `json:"batch,omitempty" bson:"batch,omitempty"`

	// meta data
	Meta jsoniter.RawMessage `json:"meta,omitempty" bson:"meta,omitempty"`

	// 0: 最大存多少天(default)
	// 1: 最大能占用多少内存
	SdcardRotatePolicy int `json:"sdcardRotatePolicy,omitempty" bson:"sdcardRotatePolicy,omitempty"`

	// 上面policy 对应的值(默认存7天)
	SdcardRotateValue int `json:"sdcardRotateValue,omitempty" bson:"sdcardRotateValue,omitempty"`

	// device type 0:normal type, 1:gateway
	Type int `json:"type" bson:"type,omitempty"`
	// max channel of gateway [1,64]
	MaxChannel int       `json:"maxChannel,omitempty" bson:"maxChannel,omitempty"`
	Channels   []Channel `json:"channels,omitempty" bson:"channels,omitempty"`
}

func main() {

	flag.StringVar(&mgoAddr, "addr", "", "-addr='10.200.20.59:2810,10.200.20.57:2810'")
	flag.StringVar(&dbName, "db", "", "-db=xxx")
	flag.StringVar(&table, "c", "", "-c=xxx")
	flag.Parse()
	if dbName == "" || table == "" || mgoAddr == "" {
		fmt.Println("error fmt , see   --help ")
		return
	}
	var curData = time.Now()
	var curMil = time.Now().UnixNano() / 1e6
	fmt.Println(curMil)
	session, err := mgo.DialWithTimeout(mgoAddr, 3*time.Second)
	if err != nil {
		fmt.Println("mongo connect failed:", err)
		return
	}

	defer session.Close()

	var curSec = time.Now().Unix()
	var s1 = session.DB(dbName).C(table)

	//write to historyactivity
	if table == "historyactivity" {

		var oneItem1 = DeviceHistoryElement{
			Date:         curData,
			Appid:        "2akrar93dqurh",
			Device:       "yujiatest",
			LoginAt:      curMil,
			LogoutAt:     curMil,
			LogoutReason: "xxxxx",
			RemoteIp:     "199.199.222.222",
		}

		for i := 0; i < 1e6; i++ {
			oneItem1.Id = makeid(oneItem1.Appid, oneItem1.Device, curMil+int64(i))
			var err = s1.Insert(oneItem1)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	}

	if table == "device" {

		var app = "2akrar999jqpb:"
		var deviceId = "Test_Device"

		var ak = "bW2YOdEaDRhjd-5ev7p2bbBxRccRgWFi5cR2DpGI"

		var oneItem1 = Device{
			SegmentExpireDays: 0,
			CreatedAt:         curSec,
			Batch:             "001",
		}

		for i := 0; i < 1e6; i++ {

			var keyInfo = []DeviceKey{
				{
					CreatedAt: curSec,
					AccessKey: ak + strconv.Itoa(i),
				},
			}

			oneItem1.Id = app + deviceId + strconv.Itoa(i)
			oneItem1.Keys = keyInfo
			var err = s1.Insert(oneItem1)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(i)
		}

	}

	if table == "app" {

		var Uid = 1810757928
		var app = ""
		var id = "2akrara1qgkuu"

		var count = 0
		for i := 0; i < 1e4; i++ {

			var oneItem1 = App{
				Uid: uint32(Uid) + uint32(i),
				App:app,
			}
			for j := 0; j < 1e2; j++ {

				
				oneItem1.Id = id + strconv.Itoa(count)
				var err = s1.Insert(oneItem1)
				if err != nil {
					fmt.Println(err)
					break
				}
				count++
			}
			fmt.Println(count)
		}
	}

	var insertCount, err2 = s1.Count()
	fmt.Println("insert table count:", table, insertCount, err2)
}
