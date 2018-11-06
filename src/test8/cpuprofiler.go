package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("/Users/xiayanji/cpuload.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	bytesArray, err := ioutil.ReadAll(file)
	AllStr := strings.Fields(string(bytesArray))

	writeF, err := os.Create("/Users/xiayanji/cpubusy.txt")
	writeF2, err := os.Create("/Users/xiayanji/cpuidle.txt")
	writeF3, err := os.Create("/Users/xiayanji/cputime2.txt")
	defer writeF.Close()
	defer writeF2.Close()
	defer writeF3.Close()

	w := bufio.NewWriter(writeF)
	w2 := bufio.NewWriter(writeF2)
	w3 := bufio.NewWriter(writeF3)
	//使用 Flush 来确保所有缓存的操作已写入底层写入器。 w.Flush()

	var idleStr string
	var timeStr string
	var preFix string
	var busyStr string

	var timeSeq = make([]string, 10240)
	var useSeq = make([]string, 10240)

	for k, v := range AllStr {

		if k == 0 {
			idleStr = ""
			timeStr = ""
			preFix = ""
		} else {
			idleStr = AllStr[k-1]
			timeStr = AllStr[k-1]

		}

		if k > 1 {
			preFix = AllStr[k-2]
		}
		if v == "id," {
			//fmt.Println("v:", v, " pre:", idleStr )
			floatNum, err := strconv.ParseFloat(idleStr, 32)
			if err != nil {
				fmt.Println(err)
				continue
			}

			busyStr = fmt.Sprintf("%.1f\n", 100.0-floatNum)
			useSeq = append(useSeq, busyStr)
			w.WriteString(busyStr)
			w2.WriteString(idleStr + "\n")
		} else if v == "up" && preFix == "-" {
			timeSeq = append(timeSeq, timeStr)
		}
	}

	for k, v := range timeSeq {
		if v != "" {
			w3.WriteString("2018/08/15 " + timeSeq[k] + " " + useSeq[k])
			fmt.Println(k)
		}

	}

	w.Flush()
	w2.Flush()
	w3.Flush()
}
