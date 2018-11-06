package main

import (
	"flag"
	//"github.com/callmefisher/redis"
	"github.com/go-redis/redis"
	"github.com/qiniu/log"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"util"
)

var (
	redisAddr       string
	redisMasterName string
	redisAddrSlice  []string
	msgInput        string
)

type Producer struct {
	redisClient *redis.Client
}

func NewProducer(redisAddr []string, master string) *Producer {

	cli, err := util.NewRdsSentinelClient(redisAddr, master)
	log.Info("redisAddr:", redisAddr, " master:", master)
	if err != nil {
		log.Error("err1 ===>", err)
		return nil
	}
	return &Producer{redisClient: cli}
}

func (consumer *Producer) test1(ch chan bool) {

	key := "testa"

	var startTick time.Time

	for range time.Tick(4 * time.Second) {
		val := strconv.Itoa(rand.Int())
		value := "2222222" + val
		startTick = time.Now()
		val, err := consumer.redisClient.Set(key, value, 0).Result()
		log.Info("set cost:", time.Since(startTick))
		if err != nil {
			if err.Error() == util.REDISNil {
				log.Info(key, " value is nil ")
			} else {
				log.Error("err ===>", err)
			}
			ch <- false
			return
		}
		ch <- true
		log.Info("set", key, ":", value, " result:", val)
	}
}

func init() {
	flag.StringVar(&redisAddr, "redisAddr", "", "example:127.0.0.1:6379")
	flag.StringVar(&redisMasterName, "master", "", "redis master name")
	flag.StringVar(&msgInput, "msg", "", "produce msg")
	flag.Parse()

	if redisAddr == "" {
		log.Fatal("please input message queue addr")
	}

	if redisMasterName == "" {
		log.Fatal("please input redis master name")
	}

	if msgInput == "" {
		log.Fatal("please input msgs")
	}

	redisAddrSlice = strings.Split(redisAddr, ",")
	if len(redisAddrSlice) == 0 {
		log.Fatal("please input message queue addr")
	}

	log.Info(redisAddrSlice, " len:", len(redisAddrSlice), " msgInput:", msgInput)
}

func check(ch chan bool) {

	for {
		select {
		case <-ch:
			{
				log.Info("finished")
				//close(ch)
				//goto Done
			}
		default:
			log.Info("waiting")
			time.Sleep(5 * time.Second)
		}
	}

}

func test2() {
	OneProducer := NewProducer(redisAddrSlice, redisMasterName)

	if OneProducer == nil {
		return
	}
	ch := make(chan bool)
	go OneProducer.test1(ch)
	go check(ch)
	for {

	}
}

func testProducer() {
	OneProducer := NewProducer(redisAddrSlice, redisMasterName)
	if OneProducer == nil {
		log.Info("error")
		return
	}
	var startTick time.Time
	pipe := OneProducer.redisClient.Pipeline()
	pipe.ClientSetName("producer")
	pipe.Exec()
	//		Msg := map[string]interface{}{strconv.Itoa(count): "this is first msg"}

	allKV := strings.Fields(msgInput)
	lenOfMsg := len(allKV)
	if lenOfMsg%2 != 0 {
		log.Info("Msg Format error")
		return
	}
	startTick = time.Now()
	for i := 0; i < 1; i++ {

		msg := map[string]interface{}{}
		for j := 0; j < lenOfMsg; j += 2 {
			//log.Info(j, " len:", len(allKV), " allKV[j]:", allKV[j], " allKV:", allKV)
			//msg[allKV[j] + strconv.Itoa(i)] = allKV[j+ 1]
			msg[allKV[j]] = allKV[j+1]

		}

		addArgs := &redis.XAddArgs{Stream: "stream1", MaxLenApprox: 200000, ID: "*", Values: msg}
		val, err := OneProducer.redisClient.XAdd(addArgs).Result()
		log.Info("XAdd Msg:", msg, " val:", val, " err:", err, " Cost:", time.Since(startTick))
	}
	OneProducer.redisClient.Close()
}

func main() {
	testProducer()
}
