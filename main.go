package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"tidb-slowquery-lynx/collector"
	"tidb-slowquery-lynx/config"
	"time"
)

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:8080", nil)
	}()

	if len(os.Args) < 2 {
		fmt.Println("Usage: program <configFilePath>")
		os.Exit(1)
	}

	// 获取配置文件路径命令行参数
	config.LoadConfig(os.Args[1])
	//config.LoadConfig("/Users/renminghao/Desktop/pingcap/tidb-slowquery-lynx/config.toml")
	//fmt.Println(config.C)
	st, et := getPreviousWindowTimes(time.Duration(config.C.Global.TimeWindowMinutes) * time.Minute)
	collector.InitDB()
	for tdbname, tdb := range config.C.TargetDBs {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s&writeTimeout=3m&readTimeout=15m", tdb.User, tdb.Password, tdb.Host, tdb.Port, tdb.DBName)
		slowlogs, err := collector.GetSlowlog(tdbname, dsn, st, et)
		if err != nil {
			fmt.Printf("[ERROR] getslowlog error:%v\n", err)
			fmt.Printf("[ERROR] will skip %v\n", tdbname)
			continue
		}
		collector.Save(slowlogs, 100)
		//collector.SaveSlowlog(slowlogs)
		SlowlogDigests, err := collector.GetSlowlogDigest(tdbname, st, et)
		if err != nil {
			fmt.Printf("[ERROR] getslowlogdigest error:%v\n", err)
			fmt.Printf("[ERROR] will skip %v\n", tdbname)
			continue
		}
		if len(*SlowlogDigests) == 0 {
			continue
		}
		SlowlogQuerys, err := collector.GetSlowlogQuery(tdbname, dsn, SlowlogDigests)
		if err != nil {
			fmt.Printf("[ERROR] getslowlogquery error:%v\n", err)
			fmt.Printf("[ERROR] will skip %v\n", tdbname)
			continue
		}
		if len(*SlowlogDigests) == 0 {
			continue
		}
		collector.Save(SlowlogQuerys, 100)
	}
	time.Sleep(600 * time.Second)
}

func getPreviousWindowTimes(windowSize time.Duration) (startTime time.Time, endTime time.Time) {
	now := time.Now()
	endTime = now.Truncate(windowSize)
	startTime = endTime.Add(-windowSize)
	return
}
