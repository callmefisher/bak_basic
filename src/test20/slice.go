package main

import (
	"fmt"
	"strings"
	"errors"
)

func test2()  (int,error) {
	
	return 1, errors.New("11")
}

func test3()  (int,error) {
	
	return 1, errors.New("11")
}

func test4()  (int,error) {
	
	return 1, errors.New("11")
}

func test5()  error {
	
	return errors.New("11")
}


func test1() (err error){
	
	if true {
		
		
		num , err := test4()
		
		 //if _, err = test4(); err != nil {
		 	fmt.Println(num)
		 //	return
		 //}
		
		fmt.Println( err)
		// err = nil
		 return err
	}
	return
}


type BiliStreamSpec string

const (
	BiliStreamSpecNormal BiliStreamSpec = ""
	BiliStreamSpecH265   BiliStreamSpec = "h265"
	BiliStreamSpecMmask  BiliStreamSpec = "mmask"
	BiliStreamSpecWmask  BiliStreamSpec = "wmask"
)


func cutMount(strmKey string) string {
	if i := strings.IndexByte(strmKey, '/'); i != -1 {
		return strmKey[i+1:]
	}
	return strmKey
}


func BiliStreamKeyFormat(strmKey string) (streamTitle, profile string, spec BiliStreamSpec) {
	split := strings.Split(cutMount(strmKey), "@")
	streamTitle = split[0]
	if len(split) == 2 {
		profile = split[1]
	}
	if strings.HasSuffix(streamTitle, "_1500") {
		profile = "ddddddddddddddddddddd"
		streamTitle = strings.TrimSuffix(streamTitle, "_1500")
	}
	
	split = strings.Split(streamTitle, "_")
	switch split[len(split)-1] {
	case "hevc":
		spec = BiliStreamSpecH265
		streamTitle = strings.TrimSuffix(streamTitle, "_hevc")
	case "mmask":
		spec = BiliStreamSpecMmask
		streamTitle = strings.TrimSuffix(streamTitle, "_mmask")
	case "wmask":
		spec = BiliStreamSpecWmask
		streamTitle = strings.TrimSuffix(streamTitle, "_wmask")
	}
	
	return streamTitle, profile, spec
}


func testStreamFormat() {
	fmt.Println("=======================================")
	title1, profile1, spec1 := BiliStreamKeyFormat("s1")
	fmt.Println("ttle:", title1, " profile:", profile1, " spec:", spec1)
	
	
	title2, profile2, spec2 := BiliStreamKeyFormat("stream@b_720p")
	fmt.Println("title:", title2, " profile:", profile2, " spec:", spec2)
	
	
	
	title3, profile3, spec3 := BiliStreamKeyFormat("stream_hevc")
	fmt.Println("title:", title3, " profile:", profile3, " spec:", spec3)
	
	
	title4, profile4, spec4 := BiliStreamKeyFormat("stream_1500")
	fmt.Println("title:", title4, " profile:", profile4, " spec:", spec4)
	
	
	title5, profile5, spec5 := BiliStreamKeyFormat("stream_mmask")
	fmt.Println("title:", title5, " profile:", profile5, " spec:", spec5)
	
	title6, profile6, spec6 := BiliStreamKeyFormat("stream_wmask")
	fmt.Println("title:", title6, " profile:", profile6, " spec:", spec6)
	
	
	title7, profile7, spec7 := BiliStreamKeyFormat("stream_34324")
	fmt.Println("title:", title7, " profile:", profile7, " spec:", spec7)
	
}





func main()  {
	
	s1 := make([]int, 3)
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, )
	fmt.Println(len(s1), cap(s1))
	
	s1 = append(s1, 1)
	fmt.Println(len(s1), cap(s1))
	
	s1 = append(s1, 1)
	fmt.Println(len(s1), cap(s1))
	
	s1 = append(s1, 1)
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, 1)
	
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, 1)
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, 1)
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, 1)
	fmt.Println(len(s1), cap(s1))
	
	
	s2 := make([]int, 0, 3)
	s2 = append(s2, 2)
	fmt.Println(s2, cap(s2))
	
	s2 = append(s2, 3)
	fmt.Println(s2, cap(s2))
	
	s2 = append(s2, 4)
	fmt.Println(s2, cap(s2))
	
	s2 = append(s2, -2)
	fmt.Println(s2, cap(s2))
	
	fmt.Println("hello world")
	
	streamTitle := "stream_1500"
	if strings.HasSuffix(streamTitle, "_1500") {
		streamTitle = strings.TrimSuffix(streamTitle, "_1500")
	}
	fmt.Println(streamTitle)
	
	
	var hubName = "iblued"
	var domainId = "ph3dtijs3.sabkt.gdipper.com"
	var url  = "http://www.baidu.com/hehe"
	if hubName == "iblued" && domainId != "" {
	
		urlWithoutScheme := url[len("http://"):]
		fmt.Println("url no scheme:", urlWithoutScheme, " index:", strings.Index(urlWithoutScheme, "/"))
		domain := urlWithoutScheme[:strings.Index(urlWithoutScheme, "/")]
		fmt.Println("domain:", domain)
		url = strings.Replace(url, domain, domainId, 1)
		fmt.Println("result url:", url)
	}
	
	s5 := []int{1, 2, 3, 4}
	fmt.Println(s5[1:3])
	
	mapTest := make(map[string] string)
	mapTest["1"] = "10"
	fmt.Println(mapTest["1"], " ===>", mapTest["2"])
	
	
	s6 := []int{4}
	s7 := []int{1, 2}
	
	copy(s6, s7)
	fmt.Println(s6)
	
	
	
	s8 := make([]int, 0, 8)
	s8 = append(s8, 1, 1, 2,)
	s9 := s8[2:]
	
	fmt.Println("test====================")
	
	fmt.Println(s8, " ", len(s8), " ",  cap(s8), " ", s9)
	s8 = append(s8, 1, 1,)
	s9[0] = 10
	fmt.Println(s8, " ", len(s8), " ",  cap(s8), " ", s9, cap(s9), " ", )
	
	s10 := 1025
	fmt.Println(byte(s10))
	
	
	//copy(s7, s6)
	//fmt.Println(s7)
	test1()
	
	
	testStreamFormat()
	
}
