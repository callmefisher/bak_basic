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
	"sync/atomic"
	"unsafe"

	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var (
	connectNum    int
	actualConnect int
	runSec        int
)

func genKey4() string {
	var b [15]byte
	io.ReadFull(rand.Reader, b[:])
	return base64.URLEncoding.EncodeToString(b[:])
}

type Test struct {
	str string
	num int
}

func makeMttInstance3(wg *sync.WaitGroup, runInSec int, host string) {

	dak := genKey4()
	dsk := genKey4()
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
	opts.SetClientID(dak)
	opts.SetAutoReconnect(false)
	opts.SetMaxReconnectInterval(time.Second)
	opts.SetKeepAlive(time.Second * 15)
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

type hmap struct {
	count       int
	flags       uint8
	B           uint8
	Yichu       uint16
	hash0       uint32
	buckets     unsafe.Pointer
	oldbuckets  unsafe.Pointer
	oldbuckets1 unsafe.Pointer
	oldbuckets2 unsafe.Pointer
}

func TestDefer(num int) {

	if num == 1 {
		fmt.Println("hello world! 0000 ")
		return
	}

	defer func() {

		fmt.Println("hello world!1111")
	}()
}

func getMapV(m map[int]*Test, num int) (*Test, bool) {
	var v, yes = m[num]
	return v, yes
}

func getMapALl(m map[int]*Test) {
	fmt.Println("\n")
	for k, v := range m {
		fmt.Print(k, "==>", v.num)
		fmt.Println("")
	}

	fmt.Println("\n")
	fmt.Println("\n")

}

func main() {

	TestDefer(1)

	var count int32

	atomic.AddInt32(&count, 1)

	fmt.Println(" count======> ", atomic.LoadInt32(&count))

	var testS []Test
	testS = append(testS, Test{num: 1, str: "a"})
	testS = append(testS, Test{num: 2, str: "b"})

	fmt.Println("before:", testS)

	for i := 0; i < len(testS); i++ {
		var tmpS = &testS[i]
		tmpS.num = 4444
	}
	fmt.Println("after:", testS)

	var t2 Test = Test{num: 3, str: "t2"}
	fmt.Println("before:", t2)
	t2.str = "t3"
	fmt.Println("after:", t2)

	var s1 []string
	s1 = append(s1, "wae")
	fmt.Println("s1:", s1)
	s2 := append(s1, "wdw")

	fmt.Println("s1:", s1, "\n", " s2:", s2)

	var slice2 = make([]int, 10, 10)

	fmt.Println("before ", slice2)
	slice2 = append(slice2, 8)
	fmt.Println("after ", slice2)

	var mTest = make(map[int]int)

	point1 := (**hmap)(unsafe.Pointer(&mTest))
	value := *point1

	fmt.Println("before map Test:", mTest, len(mTest), "  cap:", value)

	mTest[1] = 1
	mTest[2] = 2
	mTest[3] = 3
	mTest[30] = 3
	mTest[31] = 3
	mTest[32] = 3
	mTest[33] = 3
	mTest[34] = 3
	mTest[35] = 3
	mTest[36] = 3
	mTest[37] = 3
	mTest[38] = 3
	mTest[39] = 3

	point2 := (**hmap)(unsafe.Pointer(&mTest))
	value2 := *point2

	fmt.Println("after map Test:", mTest, len(mTest), "  cap:", value2)
	fmt.Println("======>", time.Now().Unix())

	fmt.Println("\n")
	fmt.Println("\n")

	var slice1 = make([]int, 10, 10)
	slice1[0] = 0
	slice1[1] = 1
	slice1[2] = 2
	slice1[3] = 3
	slice1[9] = 9

	var slice7 []int

	var slice3 = slice1[:]
	var slice4 = slice1[0:2]
	var slice5 = slice1[0:]
	var slice6 = slice1[10:]
	slice7 = slice1[1:3]

	fmt.Println("RRRRRRRRRRRRRRRRRRRRRRRR1", slice3)
	fmt.Println("RRRRRRRRRRRRRRRRRRRRRRRR2", slice4)
	fmt.Println("RRRRRRRRRRRRRRRRRRRRRRRR3", slice5)
	fmt.Println("RRRRRRRRRRRRRRRRRRRRRRRR3", slice5)
	fmt.Println("RRRRRRRRRRRRRRRRRRRRRRRR7", slice7)
	fmt.Println("RRRRRRRRRRRRRRRRRRRRRRRR4", slice6, len(slice1))

	fmt.Println("\n")
	fmt.Println("\n")
	type Meta struct {
		RecordType string
		Value      string
	}

	var mTest2 = make(map[int]*Meta)
	mTest2[1] = &Meta{Value: "1"}
	mTest2[2] = &Meta{Value: "2"}
	fmt.Println("map avail", mTest2[1] != nil)

	//var deviceMeta json.RawMessage
	//var deviceMeta Meta

	//data1 := Meta{"recordType", "ALARMS_RECORDING"}
	//meta1, err := json.Marshal(&data1)

	//PatchStr := "\"operations\":[{\"op\":\"replace\",\"key\":\"segmentExpireDays\",\"value\":\"" + "16" + "\" }]"
	//
	//var jsonStr = []byte(PatchStr)
	//
	//
	//
	//
	//
	//
	////fmt.Println("meta:", meta1, " err:", err)
	//
	//tr := &http.Transport{}
	//client := &http.Client{Transport: tr}
	//
	//req, err := http.NewRequest("PATCH",
	//	"http://47.105.118.51:8019/v1/apps/2xenzvf06ht5b/devices/eXVqaWEteDg2LXRlc3Q=",
	//	bytes.NewBuffer(jsonStr))
	//
	//req.Header.Set("Content-Type", "application/json")
	////req.Header.Set("Authorization","QiniuStub uid=1380518997")
	//
	//fmt.Println("request header==>", req.Header,  " req body==> ", req.Body ,err)
	//resp, err := client.Do(req)
	//
	//fmt.Println("resp:", resp, err)

	var t1 = &Test{
		num: 1,
	}
	var t4 = &Test{
		num: 4,
	}

	var t3 = &Test{
		num: 3,
	}

	var mapTest = make(map[int]*Test, 10)
	mapTest[1] = t1
	mapTest[4] = t4
	mapTest[3] = t3

	fmt.Println("before map:")
	getMapALl(mapTest)
	var v1, yes1 = getMapV(mapTest, 4)

	if yes1 {
		v1.num = 7981280
	}
	fmt.Println("after map:")

	getMapALl(mapTest)

	flag.IntVar(&connectNum, "conn", 1, "number of tcp connections")
	flag.IntVar(&runSec, "time", 1800, "mqtt run time in minute")
	var host = flag.String("host", "10.200.20.26:1883", "host of mqtt server")
	var sleepMil = flag.Int("sleep", 1, "sleep miliseconds")
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
