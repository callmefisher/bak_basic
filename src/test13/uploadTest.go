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
	"io/ioutil"
	"os"
)

var (
	filePath string
)

func startLoad(ts []byte) error {

	//var err error
	//putPolicy := storage.PutPolicy{
	//	Scope:   priConf.Kodo.Bucket,
	//	Expires: priConf.Kodo.ExpireSec,
	//}
	//mac := qbox.NewMac(priConf.Kodo.KodoAccessKey, priConf.Kodo.KodoSeceretKey)
	//upToken := putPolicy.UploadToken(mac)
	//
	//cfg := storage.Config{}
	//// 空间对应的机房
	//cfg.Zone = &storage.Zone{SrcUpHosts: []string{priConf.Kodo.UploadAddr}, RsfHost: priConf.Kodo.RsAddr}
	//
	//// 是否使用https域名
	//cfg.UseHTTPS = false
	//// 上传是否使用CDN上传加速
	//cfg.UseCdnDomains = false
	//
	//formUploader := storage.NewFormUploader(&cfg)
	//ret := storage.PutRet{}
	//putExtra := storage.PutExtra{}
	//
	//filepath := fmt.Sprintf("%s.%d-%d-%d-%s.ts", req.ID, req.StartMs, req.EndMs, req.SeqId, req.ConnId)
	//data := req.Ts
	//dataLen := int64(len(data))
	//xl.Info("start upload ts path:", filepath, " bucket:", priConf.Kodo.Bucket)
	//err = formUploader.Put(context.Background(), &ret, upToken, filepath, bytes.NewReader(data), dataLen, &putExtra)
	//if err != nil {
	//	return err
	//}
	return nil

}

func main() {

	flag.StringVar(&filePath, "path", "", "file path")
	flag.Parse()
	if filePath == "" {
		fmt.Println("please input a valid file path")
		return
	}
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	bytesArray, err := ioutil.ReadAll(file)
	fmt.Println("====>", bytesArray)

}
