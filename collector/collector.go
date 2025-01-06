package collector

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"tidb-slowquery-lynx/config"
	"tidb-slowquery-lynx/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             10 * time.Second, // Slow SQL threshold
		LogLevel:                  logger.Silent,    // Log level
		IgnoreRecordNotFoundError: true,             // Ignore ErrRecordNotFound error for logger
		ParameterizedQueries:      true,             // Don't include params in the SQL log
		Colorful:                  false,            // Disable color
	},
)

type slowlogDigest struct {
	Digest     *string   `json:"Digest"`
	PlanDigest *string   `json:"plan_digest"`
	MaxTime    time.Time `json:"Max_ime"`
}

func InitDB() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s&writeTimeout=3m&readTimeout=10m", config.C.Database.User, config.C.Database.Password, config.C.Database.Host, config.C.Database.Port, config.C.Database.DBName)
	//fmt.Println(dsn)
	pool, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("dsn:%v,Open tidb error: %v\n", dsn, err)
		os.Exit(-1)
	}
	pool.SetMaxIdleConns(config.C.Database.MaxIdleConns)
	pool.SetMaxOpenConns(config.C.Database.MaxOpenConns)
	pool.SetConnMaxLifetime(time.Duration(config.C.Database.ConnMaxLifetime) * time.Second)
	dbtmp, err := gorm.Open(mysql.New(mysql.Config{Conn: pool}), &gorm.Config{SkipDefaultTransaction: true, Logger: newLogger})
	if err != nil {
		fmt.Printf("dsn:%v,Open tidb error: %v\n", dsn, err)
		os.Exit(-1)
	}
	db = dbtmp
}

func GetSlowlog(dbname string, dsn string, st time.Time, et time.Time) (*[]*model.AllClusterSlowQuery, error) {
	var slowLogs []*model.AllClusterSlowQuery
	var columns *string

	// 打开数据库连接
	sqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true, Logger: newLogger})
	if err != nil {
		return nil, err
	}
	err = sqlDB.Raw("SELECT GROUP_CONCAT(COLUMN_NAME ORDER BY ORDINAL_POSITION ASC SEPARATOR ',') FROM information_schema.columns WHERE table_name='cluster_slow_query' AND table_schema='information_schema' AND upper(COLUMN_NAME) NOT IN ('PLAN','BINARY_PLAN','QUERY','PREV_STMT')").Scan(&columns).Error
	if err != nil {
		return nil, err
	}
	SlowlogSQL := fmt.Sprintf("SELECT '%s' Cluster,STR_TO_DATE('%s','%%Y-%%m-%%d') SAMPLE_START_TIME,STR_TO_DATE('%s','%%Y-%%m-%%d') SAMPLE_END_TIME,%s FROM information_schema.cluster_slow_query WHERE time >= '%s' AND time < '%s'", dbname, st.Format("2006-01-02 15:04:05"), et.Format("2006-01-02 15:04:05"), *columns, st.Format("2006-01-02 15:04:05"), et.Format("2006-01-02 15:04:05"))
	//fmt.Println(SlowlogSQL)
	err = sqlDB.Raw(SlowlogSQL).Scan(&slowLogs).Error
	if err != nil {
		return nil, err
	}
	return &slowLogs, nil
}

func GetSlowlogQuery(dbname string, dsn string, slowlogDigests *[]*slowlogDigest) (*[]*model.AllClusterSlowQueryInfo, error) {
	var slowogQuerys []*model.AllClusterSlowQueryInfo
	var slowQuerySQL strings.Builder
	// 打开数据库连接
	sqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true, Logger: newLogger})
	if err != nil {
		return nil, err
	}
	slowQuerySQL.WriteString("SELECT sss1.Cluster,sss1.Digest,sss1.Plan_digest,sss1.info->>'$.query' Query,sss1.info->>'$.plan' Plan,sss1.info->>'$.binary_plan' Binary_plan FROM (SELECT ss1.*,(SELECT JSON_OBJECT('query',s2.query,'plan',s2.plan,'binary_plan',s2.binary_plan) FROM information_schema.cluster_slow_query s2 WHERE ss1.Digest=s2.Digest AND ss1.Plan_digest=s2.plan_digest AND s2.time =ss1.max_time  LIMIT 1) INFO FROM(")
	for i := 0; i < len(*slowlogDigests); i += 1 {

		var digest string
		var plan_Digest string
		if (*slowlogDigests)[i].Digest != nil {
			digest = *((*slowlogDigests)[i].Digest)
		}
		if (*slowlogDigests)[i].PlanDigest != nil {
			plan_Digest = *((*slowlogDigests)[i].PlanDigest)
			//fmt.Println(plan_Digest)
		}
		slowQuerySQL.WriteString(fmt.Sprintf("select '%s' Cluster,'%s' Digest,'%s' Plan_Digest,STR_TO_DATE('%s','%%Y-%%m-%%d %%H:%%i:%%s.%%f') max_time", dbname, digest, plan_Digest, (*slowlogDigests)[i].MaxTime.Format("2006-01-02 15:04:05.999999")))
		if i < len(*slowlogDigests)-1 {
			slowQuerySQL.WriteString(" union all ")
		}
	}
	slowQuerySQL.WriteString(") ss1) sss1")
	fmt.Println(slowQuerySQL.String())
	err = sqlDB.Raw(slowQuerySQL.String()).Scan(&slowogQuerys).Error
	if err != nil {
		return nil, err
	}
	return &slowogQuerys, nil
}

func GetSlowlogDigest(dbname string, st time.Time, et time.Time) (*[]*slowlogDigest, error) {
	var slowlogDigests []*slowlogDigest
	SlowlogDigestSQL := fmt.Sprintf("SELECT /*+ INL_JOIN(i@sel_2)*/  s.Digest,s.Plan_digest plan_digest,max(s.time) max_time FROM ALL_CLUSTER_SLOW_QUERY s WHERE s.cluster='%s' AND s.Time>='%s' AND s.Time<'%s' AND NOT EXISTS (SELECT 1 FROM ALL_CLUSTER_SLOW_QUERY_INFO i WHERE s.Cluster=i.Cluster AND s.digest=i.digest AND s.plan_digest=i.plan_digest) GROUP BY s.Digest,s.Plan_digest", dbname, st, et)
	//fmt.Println(SlowlogDigestSQL)
	err := db.Raw(SlowlogDigestSQL).Scan(&slowlogDigests).Error
	if err != nil {
		return nil, err
	}
	return &slowlogDigests, nil
}

func SaveSlowlog(slowLogs *[]*model.AllClusterSlowQuery) error {
	if slowLogs == nil {
		return nil
	}
	chunkSize := 500
	for i := 0; i < len(*slowLogs); i += chunkSize {
		end := i + chunkSize
		if end > len(*slowLogs) {
			end = len(*slowLogs)
		}
		if err := db.Create((*slowLogs)[i:end]).Error; err != nil {
			fmt.Printf("[ERROR] Error saving slow logs: %v", err)
			return err
		}
	}
	return nil
}

func SaveSlowlogQuery(slowLogQuerys *[]*model.AllClusterSlowQueryInfo) error {
	if slowLogQuerys == nil {
		return nil
	}
	chunkSize := 100
	for i := 0; i < len(*slowLogQuerys); i += chunkSize {
		end := i + chunkSize
		if end > len(*slowLogQuerys) {
			end = len(*slowLogQuerys)
		}
		if err := db.Create((*slowLogQuerys)[i:end]).Error; err != nil {
			fmt.Printf("[ERROR] Error saving slow log querys: %v", err)
			return err
		}
	}
	return nil
}
