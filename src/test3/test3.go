//MIT License
//
//Copyright (c) 2018 XiaYanji
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.
package main

import (
	"flag"
	"fmt"
	"github.com/callmefisher/redis"
	"github.com/qiniu/log"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const N = 10
const MILIION = 1000000

//const
const LAYOUT = "2006/01/02 15:04:05.000000"
const LOCAL_TIME_SHIFT = 8 * 3600
const LOCAL_ZONE = "Asia/Shanghai"

func fn(m *map[int]int) {
	*m = make(map[int]int)
}

func test1() {

	m := make(map[int]*int, 99)

	for i := 0; i < N; i++ {
		j := i
		m[i] = &j //A
	}

	for i := range m {
		fmt.Println(*m[i])
	}
}

// 5
func funcA() int {
	x := 5
	defer func() {
		x += 1
	}()
	return x
}

// 示例代码二： 6
func funcB() (x int) {
	defer func() {
		x += 1
	}()
	return 5
}

// 示例代码三： 5
func funcC() (y int) {
	x := 5
	defer func() {
		x += 1
	}()
	return x
}

// 示例代码四：5, 解析代码四：返回x的值，传递x到匿名函数中执行时，传递的是x的拷贝，不影响外部x的值，最终返回值为5
func funcD() (x int) {
	defer func(x int) {
		x += 1
	}(x)
	return 5
}

func test2() {

	var m map[int]int

	fn(&m)
	fmt.Println(m == nil)
}

func MultiPanicRecover() {
	defer func() {
		if err := recover(); err != nil {
			log.Info("Panic info11 is: ", err)
		}
	}()
	defer func() {
		panic("222MultiPanicRecover defer inner panic")
	}()
	defer func() {
		if err := recover(); err != nil {
			log.Info("Panic info is33: ", err)
		}
	}()
	panic("MultiPanicRecover function panic-ed!")
}

func RecoverPlaceTest() {
	// 下面一行代码中 recover 函数会返回 nil，但也不影响程序运行
	defer recover()
	// recover 函数返回 nil
	defer log.Info("recover() is: ", recover())
	defer func() {
		func() {
			// 由于不是在 defer 调用函数中直接调用 recover 函数，recover 函数会返回 nil
			if err := recover(); err != nil {
				log.Info("Panic info is: ", err)
			}
		}()

	}()
	defer func() {
		if err := recover(); err != nil {
			log.Info("Panic info is55: ", err)
		}
	}()
	panic("RecoverPlaceTest function panic-ed!")
}

// 定义一个调用 recover 函数的函数
func CallRecover() {
	if err := recover(); err != nil {
		log.Info("Panic info is66 : ", err)
	}
}

// 定义个函数，在其中 defer 另一个调用了 recover 函数的函数
func RecoverInOutterFunc() {
	defer CallRecover()
	panic("RecoverInOutterFunc function panic-ed!")
}

func test4() {

	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	log.Info(len(m))

	for k, v := range m {
		log.Info("key:", k, " value:", v)
	}

}

func test5() {
	s := "123"
	ps := &s
	b := []byte(*ps)
	pb := &b

	s += "4"
	*ps += "5"
	b[1] = '0'

	println(*ps)
	println(string(*pb))
}
func test6(s string) {
	for i := 0; i < 2; i++ {

		runtime.Gosched()

		log.Info(s)
	}
}

func NewRdsClusterClient(redisAddr []string) (redisClusterClient *redis.ClusterClient, err error) {

	redisClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              redisAddr,
		PoolSize:           3,
		IdleTimeout:        2 * time.Minute,
		PoolTimeout:        1 * time.Second,
		IdleCheckFrequency: 1 * time.Minute,
	})

	err = redisClusterClient.Ping().Err()
	return
}

func NewRdsSentinelClient(redisAddr []string, master string) (sentinelClient *redis.Client, err error) {

	sentinelClient = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    master,
		SentinelAddrs: redisAddr,
	})

	err = sentinelClient.Ping().Err()
	return
}

const REDISNil = string("redis: nil")

func NewProduce1(redisAddr []string, master string) {

	client1, err := NewRdsSentinelClient(redisAddr, master)
	if err != nil {
		log.Error("err1 ===>", err)
		return
	}
	val, err := client1.Get("testa").Result()

	if err != nil {

		if err.Error() == REDISNil {
			log.Info("get key nil")

		} else {
			log.Error("err2===>", err)
		}
		return
	}
	log.Info(val, " ", err)
}

func NewCustomer(redisAddr []string) {

	client1, err := NewRdsClusterClient(redisAddr)
	if err != nil {
		log.Error("err1 ===>", err)
		return
	}

	val, err := client1.Get("testa").Result()

	if err != nil {

		if err.Error() == REDISNil {
			log.Info("get key nil")

		} else {
			log.Error("err2===>", err)
		}
		return
	}

	log.Info(val)
}

func convertUtcNanoToLocalTimeStr(timeStr string) string {
	handleStr := timeStr[0:10]
	handleMil := timeStr[10:13]
	handleMic := timeStr[13:17]
	convertInteger, err := strconv.Atoi(handleStr)
	convertInteger2, err := strconv.Atoi(handleMil)
	convertInteger3, err := strconv.Atoi(handleMic)
	if err != nil {
		log.Info(err)
		return ""
	}

	tm := time.Unix(int64(convertInteger), int64(convertInteger2*MILIION+convertInteger3*100))
	log.Info(timeStr, "===>", tm.Format(LAYOUT))
	return tm.Format(LAYOUT)
}

func timeZone(tick int64, nano int64) time.Time {
	loc, err := time.LoadLocation(LOCAL_ZONE)
	if err != nil {
		panic(err)
	}
	return time.Unix(tick, nano).In(loc)
}

func convertUtcSecToLocalTimeStr(timeStr string) string {
	convertInteger, err := strconv.Atoi(timeStr)
	if err != nil {
		log.Info(err)
		return ""
	}
	tmUnix := time.Unix(int64(convertInteger), 0)
	log.Info(timeStr, "===>", tmUnix.Format(LAYOUT))
	return tmUnix.Format("2006-01-02 15:04:05 00.000000")
}

func convertUtcMilSecToLocalTimeStr(timeStr string) string {

	handleMil1 := timeStr[0:10]
	handleMil2 := timeStr[10:]

	convertInteger1, err := strconv.Atoi(handleMil1)
	convertInteger2, err := strconv.Atoi(handleMil2)
	if err != nil {
		log.Info(err)
		return ""
	}

	tm := time.Unix(int64(convertInteger1), int64(convertInteger2*MILIION))
	log.Info(timeStr, "===>", tm.Format(LAYOUT))
	return tm.Format(LAYOUT)
}

func convertLocalStrTimeToUtcSec(timeStr string) {
	t, err := time.ParseInLocation(LAYOUT, timeStr, time.Local)
	if err != nil {
		log.Info(err)
	}
	log.Info(timeStr, "===>", "sec:", t.Unix(), " nano:", t.UnixNano(), " mil:", t.UnixNano()/10e6)
}

func openFile(path string) {

	file, err := os.Open(path)
	if err != nil {
		log.Error(err)
		return
	}
	defer file.Close()
	bytesArray, err := ioutil.ReadAll(file)

	/*

			按照行读取
			br := bufio.NewReader(fi)
		    for {
		        a, _, c := br.ReadLine()
		        if c == io.EOF {
		            break
		        }
		        fmt.Println(string(a))
		    }
			//字符串切割
			strings.Split

	*/
	if err != nil {

		log.Error(err)
		return
	}
	AllStr := strings.Fields(string(bytesArray))

	for k, v := range AllStr {
		log.Info(k, " ->", convertUtcNanoToLocalTimeStr(v))
	}
}

var (
	redisAddr       string
	redisMasterName string

	nano2str string
	mil2str  string
	sec2str  string

	str2sec string
)

func convertTime() {
	//openFile("/Users/xiayanji/time.log")
	convertUtcSecToLocalTimeStr("1533034488")
	convertLocalStrTimeToUtcSec("2018/07/31 18:54:48.000000")

	convertLocalStrTimeToUtcSec("2018/07/26 21:50:41.302794")
	convertUtcNanoToLocalTimeStr("1532613041302794000")
	convertUtcMilSecToLocalTimeStr("153261304130")

	convertUtcNanoToLocalTimeStr("15326472581986335")

}

func init() {

	flag.StringVar(&nano2str, "nano2str", "", "-nano2str  '1533034488000000000' ")
	flag.StringVar(&mil2str, "mil2str", "", "-mil2str '153303448800' ")
	flag.StringVar(&sec2str, "sec2str", "", "-sec2str '1533034488' ")
	flag.StringVar(&str2sec, "str2sec", "", "-str2sec '2018/07/27 07:20:58.198006' ")
	flag.Parse()
}

func main() {
	if nano2str != "" {
		convertUtcNanoToLocalTimeStr(nano2str)
		return
	}
	if mil2str != "" {
		convertUtcMilSecToLocalTimeStr(mil2str)
		return
	}
	if sec2str != "" {
		convertUtcSecToLocalTimeStr(sec2str)
		return
	}
	if str2sec != "" {
		convertLocalStrTimeToUtcSec(str2sec)
		return
	}

	log.Info("error fmt, use -help")
}
