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
