package main

import (
	"flag"
	"fmt"
	
	"github.com/qiniu/rpc.v1/lb.v2.1"
	xlog "github.com/qiniu/xlog.v1"
	authutil "qiniu.com/auth/authutil.v1"
	qiniumac "qiniu.com/auth/qiniumac.v1"
	themisd "qiniu.com/pili/api/themisd.v1"
)





func main() {
	xl := xlog.NewDummy()
	
	uid := flag.Uint("u", 1380668373, "`uid`")
	stream := flag.String("s", "stream1", "`stream`")
	publishIP := flag.String("pub", "127.0.0.1", "`publishIP`")
	playIP := flag.String("play", "127.0.0.1", "`playIP`")
	flushRoute := flag.Bool("flush", false, "`flushRoute`")
	
	if !flag.Parsed() {
		flag.Parse()
	}
	
	mac := qiniumac.Mac{
		AccessKey: "",
		SecretKey: []byte(""),
	}
	suInfo := authutil.FormatSuInfo(uint32(*uid), 0)
	tr := qiniumac.NewAdminTransport(&mac, suInfo, nil)
	
	client := themisd.Client{
		Conn: lb.New(&lb.Config{
			Hosts:           []string{"http://pili-themis.qiniuapi.com"},
			ClientTimeoutMS: 5000,
			TryTimes:        3,
		}, tr),
	}
	
	args := themisd.Route{
		Uid:        uint32(*uid),
		Stream:     *stream,
		PublishIP:  *publishIP,
		PlayIP:     *playIP,
		FlushRoute: *flushRoute,
	}
	
	ret, err := client.Startlive(xl, args)
	if err != nil {
		xl.Error(err)
		return
	}
	
	fmt.Printf("ffmpeg -re -i qiniu2.mp4 -f flv -acodec copy -vcodec libx264 '%s'\n", ret.PublishUrl)
	fmt.Printf("ffplay '%s'\n", ret.PlayUrls.Rtmp)
	fmt.Printf("ffplay '%s'\n", ret.PlayUrls.Flv)
	fmt.Printf("ffplay '%s'\n", ret.PlayUrls.Hls)
}