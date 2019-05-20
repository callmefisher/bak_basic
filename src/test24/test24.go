package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"qiniu.com/pili-zeus/themisd.v1/model"
	"regexp"
	"runtime/pprof"
	"time"
)

var cpuProfile = "./cpu_profile"

func startCPUProfile() {
	if cpuProfile != "" {
		f, err := os.Create(cpuProfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not create cpu profile output file: %s",
				err)
			return
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "Can not start cpu profile: %s", err)
			f.Close()
			return
		}
	}
}

func stopCPUProfile() {
	if cpuProfile != "" {
		pprof.StopCPUProfile() // 把记录的概要信息写到已指定的文件
	}
}

func MergeStreamReports(args *model.Report) {
	tableMap := make(map[string]int)
	skip := 0
	oldSize := len(args.Streams)
	for index, stream := range args.Streams {
		if oldIndex, ok := tableMap[stream.ID]; ok {
			fmt.Println("merge stream report", stream.ID, args.Streams[oldIndex].Conn, stream.Conn, stream.Status)
			if stream.Status != "disconnected" {
				args.Streams[oldIndex] = stream
			}
			skip++
		} else {
			tableMap[stream.ID] = index
			args.Streams[index-skip] = stream
		}
	}

	newSize := oldSize - skip
	args.Streams = args.Streams[:newSize]

	for k, v := range args.Streams {
		fmt.Println("k:", k, " Id:", v.ID, " status", v.Status)
	}
}

type Report struct {
	NodeID  string         `json:"nodeID"`  // 节点ID
	Streams []ReportStream `json:"streams"` // 流列表
}

type ReportStream struct {
	ID     string `json:"id"`     // ID: uid/stream
	Stream string `json:"stream"` // 流名称
}

func getError() (err error) {

	err = errors.New("all image test fail")
	return nil

}

func main() {

	fmt.Println("\n============================ ip test start\n")
	var slice = make([]int, 10, 10)
	slice[0] = 1
	slice[1] = 2
	slice[9] = 3

	fmt.Println(slice[:10])

	var args = &model.Report{
		NodeID: "a",
		Streams: []model.ReportStream{

			{
				ID:     "id1",
				Status: "disconnected",
				Stream: "stream1",
			},

			{
				ID:     "id1",
				Status: "connect",
				Stream: "stream1",
			},

			{
				ID:     "id1",
				Status: "disconnected",
				Stream: "stream1",
			},

			//{
			//    ID:"id1",
			//    Status:"connect",
			//    Stream:"stream1",
			//
			//},
			//{
			//    ID:"id1",
			//    Status:"disconnected",
			//    Stream:"stream1",
			//
			//},

		},
	}
	MergeStreamReports(args)

	streamPattern, _ := regexp.Compile("^([a-zA-Z0-9_-]{1}|[a-zA-Z0-9_-][@a-zA-Z0-9_/-]{0,278}[a-zA-Z0-9_-])$")
	fmt.Println(" match:", streamPattern.MatchString("11111111111111"))

	fmt.Println("============================ ip test done\n")

	var arrar = []int{1, 2, 3, 4}
	fmt.Println(arrar[1:3])

	var now = time.Now()
	fmt.Println("sec:", now.IsZero())

	var str1 = fmt.Sprint("aaa", "bbbb", "cccc", "http://")
	fmt.Println(str1)
	go startCPUProfile()

	ch := make(chan int, 10)
	for {
		select {
		case <-time.After(4 * time.Second):
			{
				fmt.Println("===============>A")
				stopCPUProfile()
				fmt.Println("===============>")
				os.Exit(0)
			}
		case <-ch:
			fmt.Println("111")
			time.Sleep(1 * time.Second)
		}
	}

	var startNano = time.Now().UnixNano()
	//_  = arrar[0]
	fmt.Println(time.Now().UnixNano() - startNano)

	fmt.Println("===> ", getError())

}
