package main

import (
	"flag"
	"fmt"
	_ "net/http/pprof"
	"os"
	"tidb-slowquery-lynx/collector"
	"tidb-slowquery-lynx/config"
	"tidb-slowquery-lynx/log"
	"time"

	"go.uber.org/zap"
)

func main() {
	// go func() {
	// 	http.ListenAndServe("0.0.0.0:8080", nil)
	// }()

	var configPath string
	var logPath string

	flag.StringVar(&configPath, "c", "", "Path to the configuration file")
	flag.StringVar(&logPath, "l", "", "Path to the log file")

	flag.Usage = func() {
		fmt.Println("Usage:")
		flag.PrintDefaults()
	}
	flag.Parse()
	if configPath == "" || logPath == "" {
		flag.Usage()
		os.Exit(1)
	}
	config.LoadConfig(configPath)
	log.InitLogger(logPath)
	mylogger := log.GetLogger()
	mylogger.Info("Welcome to lynx")
	collector.Initlogger(mylogger, log.GetgormLogger())
	collector.InitDB()
	st, et := getPreviousWindowTimes(time.Duration(config.C.Global.TimeWindowMinutes) * time.Minute)
	for tdbname, tdb := range config.C.TargetDBs {
		start := time.Now()
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s&writeTimeout=3m&readTimeout=15m&group_concat_max_len=4096", tdb.User, tdb.Password, tdb.Host, tdb.Port)
		mylogger.Info("Start collect slowlog", zap.String("database", tdbname))
		db, err := collector.InitTargetDB(dsn)
		if err != nil {
			mylogger.Error("Open database error", zap.Error(err), zap.String("database", tdbname))
			mylogger.Error("Skip this database", zap.String("database", tdbname))
			continue
		}
		collector.CollectClusterinfo(tdbname, db)
		err = collector.CollectSlowlog(tdbname, db, st, et)
		if err != nil {
			mylogger.Error("CollectSlowlog error", zap.Error(err), zap.String("database", tdbname))
			mylogger.Error("Skip this database", zap.String("database", tdbname))
			continue
		}
		err = collector.CollectSlowlogQuery(tdbname, db, st, et)
		if err != nil {
			mylogger.Error("CollectSlowlog error", zap.Error(err), zap.String("database", tdbname))
			mylogger.Error("Skip this database", zap.String("database", tdbname))
			continue
		}
		mylogger.Info("Complete collect slowlog", zap.String("database", tdbname), zap.String("time", time.Since(start).String()))

	}
	mylogger.Info("Lynx has completed. Goodbye!")
}

func getPreviousWindowTimes(windowSize time.Duration) (startTime time.Time, endTime time.Time) {
	now := time.Now()
	endTime = now.Truncate(windowSize)
	startTime = endTime.Add(-windowSize)
	return
}
