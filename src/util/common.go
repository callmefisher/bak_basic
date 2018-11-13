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
package util

import (
	//"github.com/callmefisher/redis"
	"encoding/json"
	"github.com/go-redis/redis"
	"time"
)

func NewRdsClusterClient(redisAddr []string) (redisClusterClient *redis.ClusterClient, err error) {

	redisClusterClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              redisAddr,
		PoolSize:           10,
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
		PoolSize:      10,
		ReadTimeout:   -1,
	})

	err = sentinelClient.Ping().Err()
	return
}

const REDISNil = string("redis: nil")

type Profile struct {
	//benchTime time.Duration
	start time.Time // Time test or benchmark started
}

func (p *Profile) StartAction() {
	p.start = time.Now()
}

func (p *Profile) resetAction() {
	//p.benchTime = 0
	p.start = time.Now()
}

func (p *Profile) EndAction() time.Duration {
	return time.Since(p.start)
}

type AccessLogEntry struct {
	Method       string  `json:"method"`
	Host         string  `json:"host"`
	Path         string  `json:"path"`
	IP           string  `json:"ip"`
	ResponseTime float64 `json:"response_time"`

	Encoded []byte
	Err     error
}

func (ale *AccessLogEntry) ensureEncoded() {
	if ale.Encoded == nil && ale.Err == nil {
		ale.Encoded, ale.Err = json.Marshal(ale)
	}
}

func (ale *AccessLogEntry) Length() int {
	ale.ensureEncoded()
	return len(ale.Encoded)
}

func (ale *AccessLogEntry) Encode() ([]byte, error) {
	ale.ensureEncoded()
	return ale.Encoded, ale.Err
}
