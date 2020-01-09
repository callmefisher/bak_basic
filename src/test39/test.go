package main

import (
	"encoding/json"
	"fmt"
	"github.com/qiniu/xlog.v1"
	"io/ioutil"
	"net/http"
	"strings"
	
	"github.com/qiniu/rpc.v1/lb.v2.1"
	"qiniu.com/auth/qiniumac.v1"
)

type liveArgs struct {
	Appid     string `json:"appid"`
	Device    string `json:"device"`
	PublishIP string `json:"publishIP"`
	PlayIP    string `json:"playIP"`
	Model     int    `json:"model"`
}

type RouteRet struct {
	PublishUrl string        `json:"publishUrl"` // 推流URL
	PlayUrls   RoutePlayUrls `json:"playUrls"`   // 拉流URLs
}

type RoutePlayUrls struct {
	Rtmp string `json:"rtmp"` // RTMP拉流URL
	Flv  string `json:"flv"`  // FLV拉流URL
	Hls  string `json:"hls"`  // HLS拉流URL
}

var appUnit = make(map[string]AppInfo)

const (
	PublishIP = "127.0.0.1"
	PlayIP    = "127.0.0.1"
	Model     = 0
)

type AppInfo struct {
	Id     string
	hosts  []string
	Ak     string
	SK     string
	Device string
}

var (
	appInfo AppInfo
	args    = liveArgs{}
)

func init() {
	
	appUnit["2xenzvm26zx2b"] = AppInfo{
		Id:    "2xenzvm26zx2b",
		hosts: []string{"http://linking.qiniuapi.com"},
		//Ak:     "ASs6_77Km5cDlZt7_K9eI1-P2z-_WMWwiPhT9fwk",
		//SK:     "X6devmDVzYaiJGCi8BF46FU2bGdSZJN31oUV4BkA",
		Ak:     "JAwTPb8dmrbiwt89Eaxa4VsL4_xSIYJoJh4rQfOQ",
		SK:     "G5mtjT3QzG4Lf7jpCAN5PZHrGeoSH9jRdC96ecYS",
		Device: "anjia_test_rtmp",
	}
	//AppID:2xenzvm26zx2b	Uid:1381539624， XYJ_Application1
	//device:anjia_test_rtmp
	//bucket: linking-vdn-test
	
	appUnit["2akrarewgurpw"] = AppInfo{
		Id:     "2akrarewgurpw",
		hosts:  []string{"http://10.200.20.26:5276/v1/startlive"},
		Ak:     "Ves3WTXC8XnEHT0I_vacEQQz-9jrJZxNExcmarzQ",
		SK:     "eNFrLXKG3R8TJ-DJA9YiMjLwuEfQnw8krrDuZzoy",
		Device: "device_2akrar9djdptk",
	}
	// 1380310120, stor-in-qbox-07
	//2akrarewgurpw
	// bucket-test1
	//{ "_id" : "2akrarewgurpw", "createdAt" : NumberLong(1558017626), "playbackStatus" : 0,
	//	"bucket" : "bucket-test1", "deviceCounter" : 10219, "state" : 0, "comment" : "remark",
	//	"publishDomain" : "www.qiniu.com", "activedDeviceCounter" : 3, "fileType" : 1,
	//	"updatedAt" : NumberLong(1577085666), "bucketDownloadDomain" : "prlghu509.test.bkt.clouddn.com",
	//	"uploadMode" : 0, "playLimit" : 36, "playDomain" : "www.hao123.com",
	//	"segmentExpireDays" : 14, "app" : "aaa123111", "uid" : 1810757928, "liveStatus" : 0, "logLevel" : 2 }
	appInfo = appUnit["2xenzvm26zx2b"]
	args = liveArgs{
		Appid:     appInfo.Id,
		Device:    appInfo.Device,
		PublishIP: PublishIP,
		PlayIP:    PlayIP,
		Model:     Model,
	}
}

func NewHttpQiniuAuthClient(hosts []string, mac *qiniumac.Mac, timeoutMs int) *lb.Client {
	
	for i, host := range hosts {
		if !strings.HasPrefix(host, "http") {
			host = "http://" + host
			hosts[i] = host
		}
	}
	if timeoutMs <= 0 {
		timeoutMs = 3000
	}
	fmt.Println(hosts)
	return lb.New(&lb.Config{
		Hosts:           hosts,
		ClientTimeoutMS: timeoutMs,
		TryTimes:        3,
	},
		qiniumac.NewTransport(mac, nil))
}

func test1() {
	
	xl := xlog.NewDummy()
	mac := &qiniumac.Mac{
		AccessKey: appInfo.Ak,
		SecretKey: []byte(appInfo.SK),
	}
	//tp := qiniumac.NewTransport(mac, nil)
	//admin auth
	//suInfo := authutil.FormatSuInfo(uid, 0)
	//tr := qiniumac.NewAdminTransport(&mac, suInfo, nil)
	//
	//client := themisd.Client{
	//	Conn: lb.New(&lb.Config{
	//		//Hosts:           []string{"http://pili-themis.qiniuapi.com"},
	//		Hosts:           []string{"http://linking.qiniuapi.com"},
	//		ClientTimeoutMS: 5000,
	//		TryTimes:        3,
	//	}, tp),
	//}
	
	var httpClient = NewHttpQiniuAuthClient(appInfo.hosts, mac, 5000)
	ret := RouteRet{}
	err := httpClient.CallWithJson(xl, &ret, "/v1/startlive", args)
	if err != nil {
		xl.Error(err)
		return
	}
	
	fmt.Printf("ffplay '%s'\n", ret.PlayUrls.Rtmp)
	fmt.Printf("ffplay '%s'\n", ret.PlayUrls.Flv)
	fmt.Printf("ffplay '%s'\n", ret.PlayUrls.Hls)
	fmt.Printf("'%s'\n", ret.PublishUrl)
	
}

func test2() {
	
	tr2 := &http.Transport{}
	client2 := &http.Client{Transport: tr2}
	
	b2, _ := json.Marshal(args)
	//发起请求
	fmt.Println(string(b2))
	
	req2, _ := http.NewRequest("POST", appInfo.hosts[0], strings.NewReader(string(b2)))
	req2.Header.Add("Authorization", "QiniuStub uid=1810757928")
	req2.Header.Add("Content-Type", "application/json")
	
	resp2, err := client2.Do(req2)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	bodyBytes2, _ := ioutil.ReadAll(resp2.Body)
	
	//fmt.Println(string(bodyBytes2))
	ret := RouteRet{}
	json.Unmarshal(bodyBytes2, &ret)
	fmt.Println("status:", resp2.StatusCode, resp2.Header, resp2.Status)
	fmt.Printf("ffplay '%s'\n", ret.PlayUrls.Rtmp)
	fmt.Printf("ffplay '%s'\n", ret.PlayUrls.Flv)
	fmt.Printf("ffplay '%s'\n", ret.PlayUrls.Hls)
	fmt.Printf("'%s'\n", ret.PublishUrl)
}

func main() {
	
	test1()
	test2()
	
}
