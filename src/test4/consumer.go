package main

import (
	"flag"
	//"github.com/callmefisher/redis"
	"github.com/go-redis/redis"
	"github.com/qiniu/log"
	"strings"
	"time"
	"util"

	sha "crypto/sha1"
	"encoding/binary"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	//"strconv"
	"strconv"
)

var (
	redisAddr       string
	redisMasterName string
	consumer        string
	redisAddrSlice  []string
)

type Consumer struct {
	redisClient *redis.Client
}

func NewConsumer(redisAddr []string, master string) *Consumer {

	cli, err := util.NewRdsSentinelClient(redisAddr, master)
	log.Info("redisAddr:", redisAddr, " master:", master)
	if err != nil {
		log.Error("err1 ===>", err)
		return nil
	}
	return &Consumer{redisClient: cli}
}

func (consumer *Consumer) test1(ch chan bool) {

	key := "testa"
	var startTick time.Time
	var val string
	var err error

	for range time.Tick(5 * time.Second) {
		startTick = time.Now()
		val, err = consumer.redisClient.Get(key).Result()
		log.Info("get cost:", time.Since(startTick))
		if err != nil {
			if err.Error() == util.REDISNil {
				log.Info(key, " value is nil ")
			} else {
				log.Error("err ===>", err)
			}
			//ch <- false
			continue
		}
		//ch <- true
		log.Info(key, "==>2:", val)
	}
}

func init_1() {
	flag.StringVar(&redisAddr, "redisAddr", "", "example:127.0.0.1:6379")
	flag.StringVar(&redisMasterName, "master", "", "redis master name")
	flag.StringVar(&consumer, "consumer", "", "stream group consumer name")
	flag.Parse()

	if redisAddr == "" {
		log.Fatal("please input message queue addr")
	}

	if redisMasterName == "" {
		log.Fatal("please input redis master name")
	}

	redisAddrSlice = strings.Split(redisAddr, ",")
	if len(redisAddrSlice) == 0 {
		log.Fatal("please input message queue addr")
	}
	if consumer == "" {
		log.Fatal("please input consumer name")
	}

	log.Info(redisAddrSlice, " len:", len(redisAddrSlice))
}

func check(ch chan bool) {
	for {
		select {
		case <-ch:
			{
				//log.Info("finished")
				//close(ch)
			}
		default:
			//log.Info("waiting")
			time.Sleep(5 * time.Second)
		}

	}
}

func test2() {
	OneConsumer := NewConsumer(redisAddrSlice, redisMasterName)

	if OneConsumer == nil {
		return
	}

	ch := make(chan bool)

	go OneConsumer.test1(ch)
	go check(ch)
}

//func testConsumerV1() {
//	OneConsumer := NewConsumer(redisAddrSlice, redisMasterName)
//	if OneConsumer == nil {
//		return
//	}
//
//	var startTick time.Time
//	for range time.Tick(1 * time.Second) {
//
//		//log.Info(" prepare consume1")
//		startTick = time.Now()
//
//		//val, err := OneConsumer.redisClient.XReadBlock(10000 * time.Second, "stream1", "0").Result()
//		//val, err := OneConsumer.redisClient.XRange("stream1", "-", "+").Result()
//		//val, err := OneConsumer.redisClient.XReadGroupN("cg1", "zoned", 1, "stream1", ">").Result()
//		//val, err := OneConsumer.redisClient.XAdd("stream1", "-", "+").Result()
//		//log.Info("XRange Cost:", time.Since(startTick), " val:", val)
//		//for _, v := range val {
//		//	tmpMsg := *v
//		//	log.Info("msgId->", tmpMsg.ID)
//		//	for mskKey, msgValue := range tmpMsg.Values {
//		//		log.Info(mskKey, " -> ", msgValue)
//		//	}
//		//}
//		//	OneConsumer.redisClient.XAck("stream1", "cg1", )
//		val, err := OneConsumer.redisClient.XReadGroupN("cg1", consumer, 8, "stream1", ">").Result()
//		log.Info("XReadGroup Cost:", time.Since(startTick))
//		var ackArray []string
//		for k, v := range val {
//			log.Info()
//			log.Info("StreamId: ", k)
//			for _, msgValue := range v {
//				log.Info("     streamTimeSeq:", msgValue.Stream)
//				for _, v1 := range msgValue.Messages {
//					//log.Info( " id:", v1.ID )
//					for k2, v2 := range v1.Values {
//						log.Info("             msgField:", k2, " msgValue:", v2)
//					}
//				}
//				ackArray = append(ackArray, msgValue.Stream)
//
//			}
//		}
//		if err != nil && err != redis.Nil {
//			log.Info(" prepare consume3:", " err:", err)
//		}
//
//		if len(ackArray) > 0 {
//			cmd := OneConsumer.redisClient.XAck("stream1", "cg1", ackArray)
//			log.Info("ack:", cmd, " ->", ackArray)
//		}
//
//	}
//
//	for {
//
//	}
//
//}

func testConsumerV2() {
	OneConsumer := NewConsumer(redisAddrSlice, redisMasterName)
	if OneConsumer == nil {
		return
	}

	count := 0
	var startTick time.Time
	//for range time.Tick(3000 * time.Millisecond) {
	for {
		count++

		streams := []string{"stream1", ">"}
		//streams := []string{"stream1", "1533547324319-3"}
		startTick = time.Now()
		log.Info("consumer:", consumer, " count:", count)
		readArgs := &redis.XReadGroupArgs{Group: "cg1", Consumer: consumer, Streams: streams, Count: 8, Block: 0}
		val, err := OneConsumer.redisClient.XReadGroup(readArgs).Result()
		log.Info("XReadGroup Cost:", time.Since(startTick), " count:", count, " len result:", len(val))
		var ackArray []string
		//var msg  = make(map[string]interface{})

		var msg = make([]map[string]interface{}, 8, 16)
		for _, v := range val {
			for _, msgValue := range v.Messages {
				//log.Info("streamName:", v.Stream, "   streamTimeSeq:", msgValue.ID)
				//for k1, v1 := range msgValue.Values {
				//log.Info("streamName:", v.Stream, " streamTimeSeq:", msgValue.ID, "  msgField:", k1, " msgValue:", v1)
				//msg[k1] = v1
				//}
				msg = append(msg, msgValue.Values)
				ackArray = append(ackArray, msgValue.ID)
				//cmd := OneConsumer.redisClient.XAck("stream1", "cg1", msgValue.ID)
				//log.Info("ack:", cmd, " ->", msgValue.ID)
			}
		}

		if len(ackArray) > 0 {
			cmd := OneConsumer.redisClient.XAck("stream1", "cg1", ackArray...)
			log.Info("ack:", cmd, " ->", ackArray, " msg:", msg)
		}
		if err != nil && err != redis.Nil {
			log.Info(" prepare consume3:", " err:", err)
			if err != nil {
				if !strings.Contains(err.Error(), "i/o timeout") {
				}
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func test3() {
	c := 0
	for {
		c++
		log.Info("TEST3:", c)
		//time.Sleep(1 * time.Millisecond)
	}
}

func test5() {

	for {
	}
}

func test4() {
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGQUIT)
		for {
			<-sigs
			// 给buf 1MB的空间，最好判断下空间是否足够
			buf := make([]byte, 1<<10)
			runtime.Stack(buf, true)
			log.Infof("=== goroutine stack trace...\n%s\n=== end\n", buf)
			os.Exit(-1)
		}
	}()
}

func test6() {
	cli := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	cmd := cli.GetSet("test", 3)
	fmt.Println("cmd:", cmd, "  ==>", cmd.Err() == redis.Nil)
}

func EnableModelOrNot(s string, percent int, base int) bool {
	if s == "" || percent <= 0 || base <= 0 {
		return false
	}
	if percent >= base {
		return true
	}
	convertNum := int(binary.LittleEndian.Uint32(Hash([]byte(s))) % uint32(base))

	return convertNum >= 1 && convertNum <= percent
}

func main() {
	_, s0 := "eRw221wA-we", "nuLr32123--FYV"
	count1 := 0

	base := 100
	percent := 10

	for i := 0; i < 10000; i++ {

		s := s0 + strconv.Itoa(i) + "abvaRT"
		fl := EnableModelOrNot(s, percent, base)
		if fl {
			count1++
		}
	}

	fmt.Println(count1, "=>")

	test6()

	//go testConsumerV2()
	//go test5()
	//for {
	//runtime.Gosched()
	//time.Sleep(1 * time.Millisecond)
	//}
	//i := 3
	//for {
	//	time.Sleep(time.Second * 1)
	//	i--
	//	fmt.Println("I got scheduled!")
	//	if i == 0 {
	//		runtime.GC()
	//	}
	//}

	select {}
	//log.Info("main")
}

func Hash(val []byte) []byte {
	h := sha.New()
	h.Write(val)
	return h.Sum(nil)
}
