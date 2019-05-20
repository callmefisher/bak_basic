package main

import (
	"github.com/qiniu/xlog.v1"
	"qiniu.com/pili/common/redisutilv5"
)

func test() {

	var conf = redisutilv5.RedisClusterCfg{
		Addrs: []string{"127.0.0.1:6380", "127.0.0.1:6381", "127.0.0.1:6382", "127.0.0.1:6383", "127.0.0.1:6384",
			"127.0.0.1:6385"},
	}

	cluster, _ := redisutilv5.NewRdsClusterClient(conf)

	var xl = xlog.NewDummy()
	// test1
	//pipeline := cluster.PipelineWithTrack()

	//var r1, e1 = pipeline.Set("h1", "w1", 0 ).Result()
	//fmt.Println("r1===> ", r1,  " error1:", e1)
	//var r2, e2 = pipeline.Set("h2", "v2", 0 ).Result()
	//
	//fmt.Println("r2===> ", r2,  " error2:", e2)
	//
	//var r3, e3 = pipeline.Set("h3", "v3", 0).Result()
	//fmt.Println("r3===> ", r3,  " error3:", e3)

	//var script1 = "\"redis.call('set',KEYS[1],'ARGV[1]')\" 1 h1 w1"
	//var script2 = "\"redis.call('set',KEYS[1],'ARGV[1]')\" 1 h2 w2"

	//var script1 = "for index, key in pairs(KEYS) do redis.call('set',key, ARGV[index]) end"
	//fmt.Println(cluster.ClusterKeySlot("key3"))
	//v1, err := cluster.Eval(script1,  []string{"key3", "key4", "key5"}, "v3", "v4", "v5").Result()
	//v1, err := cluster.Eval(script1,  []string{"key3", }, "v3", ).Result()

	//fmt.Println(v1, " err:", err)

	//test2

	//var script1 = " local v1 = redis.call('get',KEYS[1]) if v1 == false or v1 == '' then redis.call('set',KEYS[1],ARGV[1]) end"
	////var script2 = " local v1 = redis.call('get',KEYS[1]) if v1 ~= false and v1 ~= '' then redis.call('del',KEYS[1]) end"
	//var m = make(map[string]string)
	//m["key3"] = "v3"
	//m["key4"] = "v4"
	//m["key5"] = "v5"
	//m["key6"] = "v6"
	//m["key7"] = "v7"
	//m["key8"] = "v8"
	//m["key9"] = "v9"
	//m["key10"] = "v10"
	//pipeline := cluster.PipelineWithTrack()
	//for k, v := range m {
	//	pipeline.Eval(script1, []string{k}, v)
	//	//pipeline.Eval(script2, []string{k, },  v)
	//}
	//_, e1 := pipeline.ExecWithTrack(xl, "script")
	//fmt.Println(" ====> ", e1)
	//pipeline.Close()

	cluster.Set("k1", "v1", 0)
	v1, e1 := cluster.Get("k1").Result()
	xl.Info("k1's value:", v1, "   error等于redis.Nil？:", e1)

	// test3

}

func main() {

	test()

}
