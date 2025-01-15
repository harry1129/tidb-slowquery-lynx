package collector

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	"tidb-slowquery-lynx/config"
	"tidb-slowquery-lynx/model"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var gormLogger logger.Interface
var mylogger *zap.Logger

type slowlogDigest struct {
	Digest     *string   `json:"Digest"`
	PlanDigest *string   `json:"plan_digest"`
	MaxTime    time.Time `json:"Max_ime"`
}

func Initlogger(l *zap.Logger, gormL logger.Interface) {
	mylogger = l
	gormLogger = gormL
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
	dbtmp, err := gorm.Open(mysql.New(mysql.Config{Conn: pool}), &gorm.Config{SkipDefaultTransaction: true, Logger: gormLogger})
	if err != nil {
		fmt.Printf("dsn:%v,Open tidb error: %v\n", dsn, err)
		os.Exit(-1)
	}
	db = dbtmp
	db.AutoMigrate(&model.AllCluster{}, &model.AllClusterSlowQuery{}, &model.AllClusterSlowQueryInfo{})
}

/*func GetSlowlog(dbname string, dsn string, st time.Time, et time.Time) (*[]*model.AllClusterSlowQuery, error) {
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
}*/

func CollectClusterinfo(tdbname string, tdb *gorm.DB) {
	var dbs *[]*model.AllCluster
	err := tdb.Raw(fmt.Sprintf("SELECT '%s' AS Cluster,SCHEMA_NAME AS DB FROM INFORMATION_SCHEMA.SCHEMATA UNION ALL select '%s' AS Cluster,'' AS SCHEMA_NAME", tdbname, tdbname)).Scan(&dbs).Error
	if err != nil {
		mylogger.Error("GetClusterinfo", zap.String("database", tdbname), zap.Error(err))
		return
	}
	err = db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&dbs).Error
	if err != nil {
		mylogger.Error("GetClusterinfo", zap.String("database", tdbname), zap.Error(err))
		return
	}
}

func CollectSlowlog(tdbname string, tdb *gorm.DB, st time.Time, et time.Time) error {
	ch := make(chan interface{}, 2)
	var wg sync.WaitGroup
	defer wg.Wait()
	defer close(ch)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go saveDataCh(ch, &wg)
	}
	err := getSlowlog(tdbname, tdb, st, et, ch)
	if err != nil {
		return err
	}
	return nil
}

func CollectSlowlogQuery(tdbname string, tdb *gorm.DB, st time.Time, et time.Time) error {
	ch := make(chan interface{}, 2)
	var wg sync.WaitGroup
	defer wg.Wait()
	defer close(ch)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go saveDataCh(ch, &wg)
	}
	slowlogDigests, err := getSlowlogDigest(tdbname, st, et)
	if err != nil {
		return err
	}
	if slowlogDigests == nil {
		return nil
	}
	err = getSlowlogQuery(tdbname, tdb, slowlogDigests, ch)
	if err != nil {
		return err
	}
	return nil
}

func getSlowlogQuery(dbname string, tdb *gorm.DB, slowlogDigests *[]*slowlogDigest, ch chan interface{}) error {
	var slowLogQuerys []*model.AllClusterSlowQueryInfo
	var slowQuerySQL strings.Builder
	var count int
	if *slowlogDigests == nil {
		return nil
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
		}
		slowQuerySQL.WriteString(fmt.Sprintf("select '%s' Cluster,'%s' Digest,'%s' Plan_Digest,STR_TO_DATE('%s','%%Y-%%m-%%d %%H:%%i:%%s.%%f') max_time", dbname, digest, plan_Digest, (*slowlogDigests)[i].MaxTime.Format("2006-01-02 15:04:05.999999")))
		if i+1 < len(*slowlogDigests) {
			slowQuerySQL.WriteString(" union all ")
		}
	}
	slowQuerySQL.WriteString(") ss1) sss1")
	//fmt.Println(slowQuerySQL.String())
	sqlDB, err := tdb.DB()
	if err != nil {
		return err
	}
	rows, err := sqlDB.Query(slowQuerySQL.String())
	if err != nil {
		return err
	}
	for rows.Next() {
		var slowLogQuery *model.AllClusterSlowQueryInfo
		err = db.ScanRows(rows, &slowLogQuery)
		if err != nil {
			return err
		}
		slowLogQuerys = append(slowLogQuerys, slowLogQuery)
		count++
		if count == 10 {
			ch <- slowLogQuerys
			slowLogQuerys = []*model.AllClusterSlowQueryInfo{}
			count = 0
		}
	}
	if len(slowLogQuerys) > 0 {
		ch <- slowLogQuerys
	}
	return nil
}

func getSlowlogDigest(dbname string, st time.Time, et time.Time) (*[]*slowlogDigest, error) {
	var slowlogDigests []*slowlogDigest
	SlowlogDigestSQL := fmt.Sprintf("SELECT /*+ INL_JOIN(i@sel_2)*/  s.Digest,s.Plan_digest plan_digest,max(s.time) max_time FROM ALL_CLUSTER_SLOW_QUERY s WHERE s.cluster='%s' AND s.Time>='%s' AND s.Time<'%s' AND NOT EXISTS (SELECT 1 FROM ALL_CLUSTER_SLOW_QUERY_INFO i WHERE s.Cluster=i.Cluster AND s.digest=i.digest AND s.plan_digest=i.plan_digest) GROUP BY s.Digest,s.Plan_digest", dbname, st, et)
	err := db.Raw(SlowlogDigestSQL).Scan(&slowlogDigests).Error
	if err != nil {
		return nil, err
	}
	return &slowlogDigests, nil
}

// func SaveData[T any](data *[]T, chunkSize int) error {
// 	if data == nil {
// 		return nil
// 	}
// 	for i := 0; i < len(*data); i += chunkSize {
// 		end := i + chunkSize
// 		if end > len(*data) {
// 			end = len(*data)
// 		}
// 		if err := db.Create((*data)[i:end]).Error; err != nil {
// 			fmt.Printf("[ERROR] Error saving %v ,err is %v", reflect.TypeOf(*data).Elem(), err)
// 			return err
// 		}
// 	}
// 	return nil
// }

func InitTargetDB(dsn string) (*gorm.DB, error) {
	targetdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true, Logger: gormLogger})
	if err != nil {
		return nil, err
	}
	return targetdb, nil
}

func getSlowlog(dbname string, tdb *gorm.DB, st time.Time, et time.Time, ch chan interface{}) error {
	var slowLogs []*model.AllClusterSlowQuery
	var columns *string
	var count int
	//err := tdb.Raw("WITH aa AS ( select 'INSTANCE' AS COLUMN_NAME,4 AS ORDINAL_POSITION UNION ALL select 'Time' AS COLUMN_NAME,5 AS ORDINAL_POSITION UNION ALL select 'Txn_start_ts' AS COLUMN_NAME,6 AS ORDINAL_POSITION UNION ALL select 'User' AS COLUMN_NAME,7 AS ORDINAL_POSITION UNION ALL select 'Host' AS COLUMN_NAME,8 AS ORDINAL_POSITION UNION ALL select 'Conn_ID' AS COLUMN_NAME,9 AS ORDINAL_POSITION UNION ALL select 'Session_alias' AS COLUMN_NAME,10 AS ORDINAL_POSITION UNION ALL select 'Exec_retry_count' AS COLUMN_NAME,11 AS ORDINAL_POSITION UNION ALL select 'Exec_retry_time' AS COLUMN_NAME,12 AS ORDINAL_POSITION UNION ALL select 'Query_time' AS COLUMN_NAME,13 AS ORDINAL_POSITION UNION ALL select 'Parse_time' AS COLUMN_NAME,14 AS ORDINAL_POSITION UNION ALL select 'Compile_time' AS COLUMN_NAME,15 AS ORDINAL_POSITION UNION ALL select 'Rewrite_time' AS COLUMN_NAME,16 AS ORDINAL_POSITION UNION ALL select 'Preproc_subqueries' AS COLUMN_NAME,17 AS ORDINAL_POSITION UNION ALL select 'Preproc_subqueries_time' AS COLUMN_NAME,18 AS ORDINAL_POSITION UNION ALL select 'Optimize_time' AS COLUMN_NAME,19 AS ORDINAL_POSITION UNION ALL select 'Wait_TS' AS COLUMN_NAME,20 AS ORDINAL_POSITION UNION ALL select 'Prewrite_time' AS COLUMN_NAME,21 AS ORDINAL_POSITION UNION ALL select 'Wait_prewrite_binlog_time' AS COLUMN_NAME,22 AS ORDINAL_POSITION UNION ALL select 'Commit_time' AS COLUMN_NAME,23 AS ORDINAL_POSITION UNION ALL select 'Get_commit_ts_time' AS COLUMN_NAME,24 AS ORDINAL_POSITION UNION ALL select 'Commit_backoff_time' AS COLUMN_NAME,25 AS ORDINAL_POSITION UNION ALL select 'Backoff_types' AS COLUMN_NAME,26 AS ORDINAL_POSITION UNION ALL select 'Resolve_lock_time' AS COLUMN_NAME,27 AS ORDINAL_POSITION UNION ALL select 'Local_latch_wait_time' AS COLUMN_NAME,28 AS ORDINAL_POSITION UNION ALL select 'Write_keys' AS COLUMN_NAME,29 AS ORDINAL_POSITION UNION ALL select 'Write_size' AS COLUMN_NAME,30 AS ORDINAL_POSITION UNION ALL select 'Prewrite_region' AS COLUMN_NAME,31 AS ORDINAL_POSITION UNION ALL select 'Txn_retry' AS COLUMN_NAME,32 AS ORDINAL_POSITION UNION ALL select 'Cop_time' AS COLUMN_NAME,33 AS ORDINAL_POSITION UNION ALL select 'Process_time' AS COLUMN_NAME,34 AS ORDINAL_POSITION UNION ALL select 'Wait_time' AS COLUMN_NAME,35 AS ORDINAL_POSITION UNION ALL select 'Backoff_time' AS COLUMN_NAME,36 AS ORDINAL_POSITION UNION ALL select 'LockKeys_time' AS COLUMN_NAME,37 AS ORDINAL_POSITION UNION ALL select 'Request_count' AS COLUMN_NAME,38 AS ORDINAL_POSITION UNION ALL select 'Total_keys' AS COLUMN_NAME,39 AS ORDINAL_POSITION UNION ALL select 'Process_keys' AS COLUMN_NAME,40 AS ORDINAL_POSITION UNION ALL select 'Rocksdb_delete_skipped_count' AS COLUMN_NAME,41 AS ORDINAL_POSITION UNION ALL select 'Rocksdb_key_skipped_count' AS COLUMN_NAME,42 AS ORDINAL_POSITION UNION ALL select 'Rocksdb_block_cache_hit_count' AS COLUMN_NAME,43 AS ORDINAL_POSITION UNION ALL select 'Rocksdb_block_read_count' AS COLUMN_NAME,44 AS ORDINAL_POSITION UNION ALL select 'Rocksdb_block_read_byte' AS COLUMN_NAME,45 AS ORDINAL_POSITION UNION ALL select 'DB' AS COLUMN_NAME,46 AS ORDINAL_POSITION UNION ALL select 'Index_names' AS COLUMN_NAME,47 AS ORDINAL_POSITION UNION ALL select 'Is_internal' AS COLUMN_NAME,48 AS ORDINAL_POSITION UNION ALL select 'Digest' AS COLUMN_NAME,49 AS ORDINAL_POSITION UNION ALL select 'Stats' AS COLUMN_NAME,50 AS ORDINAL_POSITION UNION ALL select 'Cop_proc_avg' AS COLUMN_NAME,51 AS ORDINAL_POSITION UNION ALL select 'Cop_proc_p90' AS COLUMN_NAME,52 AS ORDINAL_POSITION UNION ALL select 'Cop_proc_max' AS COLUMN_NAME,53 AS ORDINAL_POSITION UNION ALL select 'Cop_proc_addr' AS COLUMN_NAME,54 AS ORDINAL_POSITION UNION ALL select 'Cop_wait_avg' AS COLUMN_NAME,55 AS ORDINAL_POSITION UNION ALL select 'Cop_wait_p90' AS COLUMN_NAME,56 AS ORDINAL_POSITION UNION ALL select 'Cop_wait_max' AS COLUMN_NAME,57 AS ORDINAL_POSITION UNION ALL select 'Cop_wait_addr' AS COLUMN_NAME,58 AS ORDINAL_POSITION UNION ALL select 'Mem_max' AS COLUMN_NAME,59 AS ORDINAL_POSITION UNION ALL select 'Disk_max' AS COLUMN_NAME,60 AS ORDINAL_POSITION UNION ALL select 'KV_total' AS COLUMN_NAME,61 AS ORDINAL_POSITION UNION ALL select 'PD_total' AS COLUMN_NAME,62 AS ORDINAL_POSITION UNION ALL select 'Backoff_total' AS COLUMN_NAME,63 AS ORDINAL_POSITION UNION ALL select 'Write_sql_response_total' AS COLUMN_NAME,64 AS ORDINAL_POSITION UNION ALL select 'Result_rows' AS COLUMN_NAME,65 AS ORDINAL_POSITION UNION ALL select 'Warnings' AS COLUMN_NAME,66 AS ORDINAL_POSITION UNION ALL select 'Backoff_Detail' AS COLUMN_NAME,67 AS ORDINAL_POSITION UNION ALL select 'Prepared' AS COLUMN_NAME,68 AS ORDINAL_POSITION UNION ALL select 'Succ' AS COLUMN_NAME,69 AS ORDINAL_POSITION UNION ALL select 'IsExplicitTxn' AS COLUMN_NAME,70 AS ORDINAL_POSITION UNION ALL select 'IsWriteCacheTable' AS COLUMN_NAME,71 AS ORDINAL_POSITION UNION ALL select 'Plan_from_cache' AS COLUMN_NAME,72 AS ORDINAL_POSITION UNION ALL select 'Plan_from_binding' AS COLUMN_NAME,73 AS ORDINAL_POSITION UNION ALL select 'Has_more_results' AS COLUMN_NAME,74 AS ORDINAL_POSITION UNION ALL select 'Resource_group' AS COLUMN_NAME,75 AS ORDINAL_POSITION UNION ALL select 'Request_unit_read' AS COLUMN_NAME,76 AS ORDINAL_POSITION UNION ALL select 'Request_unit_write' AS COLUMN_NAME,77 AS ORDINAL_POSITION UNION ALL select 'Time_queued_by_rc' AS COLUMN_NAME,78 AS ORDINAL_POSITION UNION ALL select 'Tidb_cpu_time' AS COLUMN_NAME,79 AS ORDINAL_POSITION UNION ALL select 'Tikv_cpu_time' AS COLUMN_NAME,80 AS ORDINAL_POSITION UNION ALL select 'Plan' AS COLUMN_NAME,81 AS ORDINAL_POSITION UNION ALL select 'Plan_digest' AS COLUMN_NAME,82 AS ORDINAL_POSITION UNION ALL select 'Binary_plan' AS COLUMN_NAME,83 AS ORDINAL_POSITION UNION ALL select 'Prev_stmt' AS COLUMN_NAME,84 AS ORDINAL_POSITION UNION ALL select 'Query' AS COLUMN_NAME,85 AS ORDINAL_POSITION) SELECT GROUP_CONCAT(concat_ws('',IFNULL(c.COLUMN_NAME,'NULL'),' as ',aa.COLUMN_NAME) ORDER BY aa.ORDINAL_POSITION SEPARATOR ',') FROM aa LEFT JOIN information_schema.COLUMNS c on lower(aa.COLUMN_NAME)=lower(c.COLUMN_NAME) AND c.TABLE_SCHEMA='information_schema' AND c.TABLE_NAME='cluster_slow_query' AND upper(c.COLUMN_NAME) NOT IN ('PLAN','BINARY_PLAN','QUERY','PREV_STMT')").Scan(&columns).Error
	err := tdb.Raw("SELECT GROUP_CONCAT(COLUMN_NAME ORDER BY ORDINAL_POSITION ASC SEPARATOR ',') FROM information_schema.columns WHERE table_name='cluster_slow_query' AND table_schema='information_schema' AND upper(COLUMN_NAME) NOT IN ('PLAN','BINARY_PLAN','QUERY','PREV_STMT')").Scan(&columns).Error
	if err != nil {
		return err
	}
	slowlogSQL := fmt.Sprintf("SELECT '%s' Cluster,STR_TO_DATE('%s','%%Y-%%m-%%d %%H:%%i:%%s') SAMPLE_START_TIME,STR_TO_DATE('%s','%%Y-%%m-%%d %%H:%%i:%%s') SAMPLE_END_TIME,%s FROM information_schema.cluster_slow_query WHERE time >= '%s' AND time < '%s'", dbname, st.Format("2006-01-02 15:04:05"), et.Format("2006-01-02 15:04:05"), *columns, st.Format("2006-01-02 15:04:05"), et.Format("2006-01-02 15:04:05"))

	sqlDB, err := tdb.DB()
	if err != nil {
		return err
	}
	rows, err := sqlDB.Query(slowlogSQL)
	if err != nil {
		return err
	}
	for rows.Next() {
		var slowLog *model.AllClusterSlowQuery
		err = db.ScanRows(rows, &slowLog)
		if err != nil {
			return err
		}
		slowLogs = append(slowLogs, slowLog)
		count++
		if count == 100 {
			ch <- slowLogs
			slowLogs = []*model.AllClusterSlowQuery{}
			count = 0
		}
	}
	if len(slowLogs) > 0 {
		ch <- slowLogs
	}
	/*err = tdb.Raw(slowlogSQL).Scan(&slowLogs).Error
	if err != nil {
		return err
	}
	for i := 0; i < len(slowLogs); i += chunkSize {
		end := i + chunkSize
		if end > len(slowLogs) {
			end = len(slowLogs)
		}
		ch <- slowLogs[i:end]
	} */
	return nil
}

func saveDataCh(ch chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		data, ok := <-ch
		if !ok { // 通道关闭
			mylogger.Debug("Channel close")
			return
		}
		result := db.Create(data)
		if result.Error != nil {
			mylogger.Error("Save data error", zap.String("datatype", reflect.TypeOf(data).String()), zap.Error(result.Error))
		} else {
			mylogger.Info("Save data succeed", zap.String("datatype", reflect.TypeOf(data).String()), zap.Int64("rows", result.RowsAffected))
		}
	}
}
