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
var f0_0 MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	//fmt.Printf("TIME = %s, TOPIC: %s, MSG: %s\n", time.Now(), msg.Topic(), msg.Payload())
}

var f3 MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	//fmt.Println("sub msg callback:", string(msg.Payload()), " qos:", msg.Qos())
	subCount++
}

var (
	machineID0 []byte // 两个字节的机器ID
	subCount   int
	subCon     int
)

const subTopic string = "$share/1/linking/v1/${appid}/${device}/syslog/"

func init() {
	h, err := os.Hostname()
	if err != nil {
		//panic(err)
	}
	hash := sha1.Sum([]byte(h))
	machineID0 = hash[:2]
}

// same as qiniu access/secret key
func genKey2() string {
	var b [15]byte
	io.ReadFull(rand.Reader, b[:])
	return base64.URLEncoding.EncodeToString(b[:])
}

func makeMttInstance2(wg *sync.WaitGroup, runInSec int, host string) {

	dak := genKey2()
	dsk := genKey2()
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
	opts.SetDefaultPublishHandler(f0_0)
	var onConnect = func(c MQTT.Client) {
		c.Subscribe(subTopic, 1, f3)
		subCon++
	}
	opts.SetOnConnectHandler(onConnect)
	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		//panic(token.Error())

	}

}
func main() {
	var connections = flag.Int("conn", 18, "number of tcp connections")
	var runInSec = flag.Int("time", 60, "mqtt run time in minute")
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
				fmt.Println("sample sub:", subCount, " want connect", *connections, " actual connection:", subCon,
					" elapse:", time.Now().Sub(startSec), " maxRunTime:", *runInSec)
			default:
			}
		}

	}()

	for i := 0; i < *connections; i++ {
		time.Sleep(time.Duration(*sleepMil) * time.Millisecond)
		wg.Add(1)
		go makeMttInstance2(&wg, *runInSec, *host)
	}

	go func() {

		time.Sleep(time.Duration(*runInSec) * time.Second)
		for i := 0; i < *connections; i++ {
			wg.Done()
		}

	}()

	wg.Wait()
	fmt.Println("sub count:", subCount, " want connect", *connections, " actual connection:", subCon)
}
