package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"sync"

	//import the Paho Go MQTT library
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

//define a function for the default message handler

const pubString1 = "[2019-03-11 13:25:15][I][tsuploader]tsmuxuploader.c:1724 tsToken:ZLaj6xCX7-kN8FMz0AtPTmDuL4s6QyV_5i83JbbW:Q4GOSA43KMt7MBTuT69IeeqO2rs=:eyJzY29wZSI6InVlLXRlc3QiLCJkZWFkbGluZSI6MTU1MjI4MjA0NSwibWltZUxpbWl0IjoidmlkZW8vbXAydDt2aWRlby9tcDJ0cyIsImRlbGV0ZUFmdGVyR"
const pubTopic string = "linking/v1/${appid}/${device}/syslog/"

var (
	idmux1        sync.Mutex
	machineID1    []byte // 两个字节的机器ID
	pubCount      int
	pubConnection int
)
var f0_1 MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TIME = %s, TOPIC: %s, MSG: %s\n", time.Now(), msg.Topic(), msg.Payload())
}

func init() {
	h, err := os.Hostname()
	if err != nil {
		//panic(err)
	}
	hash := sha1.Sum([]byte(h))
	machineID1 = hash[:2]
}

// same as qiniu access/secret key
func genKey1() string {
	var b [15]byte
	io.ReadFull(rand.Reader, b[:])
	return base64.URLEncoding.EncodeToString(b[:])
}

func makeMttInstance1(wg *sync.WaitGroup, runInSec int, host string, wait int) {
	defer wg.Done()
	dak := genKey1()
	dsk := genKey1()
	var p = func() (string, string) {
		username := fmt.Sprintf("dak=%s&timestamp=%d&version=v1", dak, time.Now().Unix())
		h := hmac.New(sha1.New, []byte(dsk))
		h.Write([]byte(username))
		password := h.Sum(nil)
		pwd := base64.URLEncoding.EncodeToString(password)
		return username, pwd

	}
	opts := MQTT.NewClientOptions().AddBroker(host)
	opts.SetCredentialsProvider(p)
	//	opts.SetClientID(*dak)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(time.Second)
	opts.SetKeepAlive(time.Second * 15)
	opts.SetDefaultPublishHandler(f0_1)

	var onConnect = func(c MQTT.Client) {
		pubConnection++
	}
	opts.SetOnConnectHandler(onConnect)

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		//panic(token.Error())

	}

	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription

	time.Sleep(time.Duration(wait) * time.Second)
	for {
		select {
		default:
			if token := c.Publish(pubTopic, 1, false, pubString1); token.Error() == nil {
				pubCount++
			}
			time.Sleep(time.Millisecond * 500)
		}
	}

}
func main() {
	var connections = flag.Int("conn", 20, "number of tcp connections")
	var runInSec = flag.Int("time", 60, "mqtt run time in minute")
	var waitSec = flag.Int("wait", 1, "mqtt run time in minute")
	var host = flag.String("host", "10.200.20.26:1884", "host of mqtt server")
	var sleepMil = flag.Int("sleep", 100, "sleep miliseconds")
	flag.Parse()
	var wg sync.WaitGroup
	startSec := time.Now()
	go func() {

		sampleTick := time.NewTicker(time.Duration(5) * time.Second)
		for {
			select {
			case <-sampleTick.C:
				fmt.Println("sample publish", pubCount, "want connections:", *connections, " actual run:", pubConnection,
					" elapse:", time.Now().Sub(startSec), " maxSec:", *runInSec)
			default:
			}
		}
	}()

	for i := 0; i < *connections; i++ {
		time.Sleep(time.Duration(*sleepMil) * time.Millisecond)
		wg.Add(1)
		go makeMttInstance1(&wg, *runInSec, *host, *waitSec)
	}

	go func() {

		time.Sleep((time.Duration(*runInSec) * time.Second))
		for i := 0; i < *connections; i++ {
			wg.Done()
		}
	}()

	wg.Wait()
	fmt.Println("done publish", pubCount, "want connections:", *connections, " actual run:", pubConnection)
}
