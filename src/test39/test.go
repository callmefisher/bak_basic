package main

import (
	"fmt"
	"strings"

	"github.com/qiniu/rpc.v1/lb.v2.1"
	"github.com/qiniu/xlog.v1"
	"qiniu.com/auth/qiniumac.v1"
	"qiniu.com/pili/api/themisd.v1"
)

type liveArgs struct {
	Appid     string `json:"appid"`
	Device    string `json:"device"`
	PublishIP string `json:"publishIP"`
	PlayIP    string `json:"playIP"`
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
	return lb.New(&lb.Config{
		Hosts:           hosts,
		ClientTimeoutMS: timeoutMs,
		TryTimes:        3,
	},
		qiniumac.NewTransport(mac, nil))
}

func main() {
	xl := xlog.NewDummy()

	//uid := flag.Uint("u", 1381694418, "`uid`")
	//stream := flag.String("s", "3nm4x0knux28v:xyj_test1_device2", "`stream`")
	//publishIP := flag.String("pub", "127.0.0.1", "`publishIP`")
	//playIP := flag.String("play", "127.0.0.1", "`playIP`")
	//flushRoute := flag.Bool("flush", false, "`flushRoute`")

	//if !flag.Parsed() {
	//	flag.Parse()
	//}

	mac := &qiniumac.Mac{
		AccessKey: "ASs6_77Km5cDlZt7_K9eI1-P2z-_WMWwiPhT9fwk",
		SecretKey: []byte("X6devmDVzYaiJGCi8BF46FU2bGdSZJN31oUV4BkA"),
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

	hosts := []string{"http://linking.qiniuapi.com"}

	var httpClient = NewHttpQiniuAuthClient(hosts, mac, 5000)

	args := liveArgs{
		Appid:     "3nm4x0knux28v",
		Device:    "xyj_test1_device2",
		PublishIP: "140.206.66.42",
		PlayIP:    "140.206.66.42",
	}

	ret := themisd.RouteRet{}
	err := httpClient.CallWithJson(xl, &ret, "/v1/startlive", args)
	if err != nil {
		xl.Error(err)
		return
	}

	fmt.Printf("ffplay '%s'\n", ret.PlayUrls.Rtmp)
	fmt.Printf("ffplay '%s'\n", ret.PlayUrls.Flv)
	fmt.Printf("ffplay '%s'\n", ret.PlayUrls.Hls)
	fmt.Printf("Push '%s'\n", ret.PublishUrl)
}
