package main

import (
	"fmt"
	"github.com/qiniu/log.v1"
	"math/rand"
	"runtime"
	"strconv"
	t "test2"
	"time"
	//"os"
	//	"log"
	"flag"
	"qiniu.com/pili/common/mgo"
	"reflect"
	//"gopkg.in/mgo.v2/bson"
)

func test1() {

	for i := 1; i < 5; i++ {
		fmt.Println("num is:", rand.Intn(100))
	}

	fmt.Println(1 << 31)
	fmt.Println(1 << 32)

}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	c, python, java bool
)
var t2 rune = 3222832

func test2() {
	var i int = 2
	//i = 3
	fmt.Println(i, c, python, java, t2)
}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func test3(slice []int) {

	slice = slice[2:]
	fmt.Println("after len:", len(slice), " cap:", cap(slice), " first:", slice[0])
}

type s1 struct {
	namee string
	num   int
}

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (i *IPAddr) string() string {
	lastIndex := len(i) - 1
	var s string
	for i, v := range i {
		s += strconv.Itoa(int(v)) // 数字文字间互转用strconv类
		if i != lastIndex {
			s += "."
		}
	}

	return s
}

func say(s string) {
	fmt.Println("start")
	for i := 0; i < 5; i++ {
		fmt.Println("i")
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func testDeferf() int {
	//return 0
	//panic("404")
	defer func() {
		fmt.Println("tttttttttDefer===>")
	}()

	return 0
}

func testCoutine(c chan int) {

	for {
		time.Sleep(time.Second * 1)
		c <- 1
		log.Println("channel start sleep")

	}
}

type Oplog struct {
	TimeSpace       int64  `json:"ts" bson:"ts"`
	Hash            int64  `json:"-" bson:"-"`
	Version         int64  `json:"-" bson:"-"`
	OperationType   string `json:"op" bson:"op"`
	NameSpace       string `json:"ns" bson:"ns"`
	Operation       M      `json:"o" bson:"o"`
	UpdateOperation M      `json:"o2" bson:"o2"`
}

type M map[string]interface{}

var (
	intFlag int
)

func main() {

	flag.IntVar(&intFlag, "id", 230, "id")
	flag.Parse()
	log.Info("1111111111111111===>", intFlag)

	ss := []int{1, 2, 3}

	for _, v := range ss {
		v += 10
	}

	for i := range ss {
		ss[i] += 10
	}

	log.Info(ss)

	dataColl, err := mgo.New(mgo.Option{
		MgoAddr:     "127.0.0.1:27017",
		MgoDB:       "testsync",
		MgoColl:     "base",
		MgoPoolSize: 2,
	})

	defer func() {
		dataColl.Close()
	}()

	if err != nil {
		return
	}

	var ret Oplog

	dataColl.Coll().Remove(M{"host": "27017"})
	dataColl.Coll().Find(M{"host": "27018"}).One(&ret)
	log.Info("-============>", ret)

	opQuery := M{"op": "d"}
	tmpSession := dataColl.Coll().Database.Session.Copy()
	opColl := tmpSession.DB("local").C("oplog.rs")

	err = opColl.Find(opQuery).Sort("-$natural").One(&ret)

	log.Info("-2322 ============>", ret)

	tmpSession2 := dataColl.Coll().Database.Session.Copy()
	opColl2 := tmpSession2.DB("test").C("test")

	opColl2.Find(M{"host": "27017"}).One(&ret)

	log.Info("-============>", ret)

	var synModule interface{} = "str34242"

	log.Info("synModule:", synModule.(string))

	stmp1 := s1{namee: "hello", num: 1000}
	rflecctResultOfs1 := reflect.ValueOf(stmp1)
	log.Info("rflecctResultValueOfs1:", rflecctResultOfs1, " Kind:", rflecctResultOfs1.Kind())
	log.Info("Type===>", rflecctResultOfs1.Type())
	log.Info("TypeOf===>", reflect.TypeOf(stmp1))
	log.Info("NumIn===>", rflecctResultOfs1.Type().NumField())
	log.Info("StringIn===>", rflecctResultOfs1.Type().Field(0).Name)
	log.Info("FiledByName===>", rflecctResultOfs1.FieldByName("namee"))
	log.Info("FiledByName===>", rflecctResultOfs1.FieldByName("num"))
	//log.Info("Elem===>",     rflecctResultOfs1.Type().Elem())

	tickTime := 3 * time.Second
	count := 0
	log.Error("Hi=>", count)
	for range time.Tick(tickTime) {

		log.Error("Hi=>", count)
		count++

	}

	//f, err := os.OpenFile("./log_test.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	//defer f.Close()
	//if nil != err {
	//	panic("9999")
	//}
	//log.SetOutput(f)
	//log.SetOutputLevel(log.Ldebug)

	defer func() {
		fmt.Println("DDDDDDDDDDDDD")
	}()
	testC := make(chan int, 1)
	go testCoutine(testC)

	go func() {
		for {
			select {
			case readResult := <-testC:
				log.Println("read from channale::%v ============================\n\n", readResult)
			default:

				//log.Warn("gg===========ggggo")
				// time.Sleep(time.Second * 2)
			}
		}
	}()

	//设置随机种子
	rand.Seed(time.Now().UnixNano())
	//test1()
	arr := []int{1, 2, 3}

	sliceT := make([]int, 3, 5)

	fmt.Println("capTest:", cap(sliceT))

	fmt.Println("before len:", len(arr), " cap:", cap(arr), " first:", arr[0])
	test3(arr)
	//var DefaultTransport = NewTransport()

	fmt.Println(t.MyaddNum(1, 3), " ->", t.Num)

	var i interface{} = "dadsd"
	switch v := i.(type) {
	case int:
		fmt.Println(v)
	default:
		fmt.Println("fail switch")
	}
	tests1 := &s1{namee: "heh"}

	fmt.Printf("%v ", tests1)
	fmt.Printf("%v", tests1)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip.string())
	}
	//go say("xiaolezi")
	//
	//time.Sleep(time.Second * 4)
	//fmt.Printf("%v", "end")

	t.TestSum()

	testDeferf()
	fmt.Printf("\n5cur rontine num:%d\n", runtime.NumGoroutine())
	for {

	}

}
