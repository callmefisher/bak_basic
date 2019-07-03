package main

import (
	"flag"
	"fmt"
	"sync"

	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var (
	connectNum    int
	actualConnect int
	runSec        int
)

var f0_3 MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TIME = %s, TOPIC: %s, MSG: %s\n", time.Now(), msg.Topic(), msg.Payload())
}

func makeMttInstance3(wg *sync.WaitGroup, runInSec int, host string) {

	//dak := "genKey2"
	//dsk := "genKey2"
	//var p = func() (string, string) {
	//	username := fmt.Sprintf("dak=%s&timestamp=%d&version=v1", dak, time.Now().Unix())
	//	h := hmac.New(sha1.New, []byte(dsk))
	//	h.Write([]byte(username))
	//	password := h.Sum(nil)
	//	pwd := base64.URLEncoding.EncodeToString(password)
	//	return username, pwd
	//
	//}
	opts := MQTT.NewClientOptions().AddBroker(host)
	//opts.SetCredentialsProvider(p)
	//	opts.SetClientID(*dak)
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(time.Second)
	opts.SetKeepAlive(time.Second * 15)
	//opts.SetDefaultPublishHandler(f0_3)
	var onConnect = func(c MQTT.Client) {
		actualConnect++
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
	flag.IntVar(&connectNum, "conn", 2, "number of tcp connections")
	flag.IntVar(&runSec, "time", 10, "mqtt run time in minute")
	var host = flag.String("host", "10.200.20.26:1884", "host of mqtt server")
	var sleepMil = flag.Int("sleep", 100, "sleep miliseconds")
	flag.Parse()
	startSec := time.Now()
	var wg sync.WaitGroup

	go func() {
		sampleTick := time.NewTicker(time.Duration(5) * time.Second)
		for {
			select {
			case <-sampleTick.C:
				fmt.Println(" actual connections", actualConnect, " maxSec:", runSec, " wantConnect:", connectNum, " has elapse:",
					time.Now().Sub(startSec))
			default:
			}
		}
	}()
	for i := 0; i < connectNum; i++ {
		time.Sleep(time.Duration(*sleepMil) * time.Millisecond)
		wg.Add(1)
		go makeMttInstance3(&wg, runSec, *host)
	}

	go func() {

		tick := time.NewTicker(time.Duration(runSec) * time.Second)
		for {
			select {
			case <-tick.C:
				for i := 0; i < connectNum; i++ {
					wg.Done()
				}
				return
			default:
			}
		}
	}()

	wg.Wait()
	fmt.Println("done: actual connections", actualConnect, " runSec:", runSec, " wantConnect:", connectNum)
}
