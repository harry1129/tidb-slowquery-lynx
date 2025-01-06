// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"tidb-slowquery-lynx/model"
)

func newAllClusterSlowQuery(db *gorm.DB, opts ...gen.DOOption) allClusterSlowQuery {
	_allClusterSlowQuery := allClusterSlowQuery{}

	_allClusterSlowQuery.allClusterSlowQueryDo.UseDB(db, opts...)
	_allClusterSlowQuery.allClusterSlowQueryDo.UseModel(&model.AllClusterSlowQuery{})

	tableName := _allClusterSlowQuery.allClusterSlowQueryDo.TableName()
	_allClusterSlowQuery.ALL = field.NewAsterisk(tableName)
	_allClusterSlowQuery.Cluster = field.NewString(tableName, "Cluster")
	_allClusterSlowQuery.SAMPLESTARTTIME = field.NewTime(tableName, "SAMPLE_START_TIME")
	_allClusterSlowQuery.SAMPLEENDTIME = field.NewTime(tableName, "SAMPLE_END_TIME")
	_allClusterSlowQuery.INSTANCE = field.NewString(tableName, "INSTANCE")
	_allClusterSlowQuery.Time = field.NewTime(tableName, "Time")
	_allClusterSlowQuery.TxnStartTs = field.NewUint64(tableName, "Txn_start_ts")
	_allClusterSlowQuery.User = field.NewString(tableName, "User")
	_allClusterSlowQuery.Host = field.NewString(tableName, "Host")
	_allClusterSlowQuery.ConnID = field.NewUint64(tableName, "Conn_ID")
	_allClusterSlowQuery.SessionAlias = field.NewString(tableName, "Session_alias")
	_allClusterSlowQuery.ExecRetryCount = field.NewUint64(tableName, "Exec_retry_count")
	_allClusterSlowQuery.ExecRetryTime = field.NewFloat64(tableName, "Exec_retry_time")
	_allClusterSlowQuery.QueryTime = field.NewFloat64(tableName, "Query_time")
	_allClusterSlowQuery.ParseTime = field.NewFloat64(tableName, "Parse_time")
	_allClusterSlowQuery.CompileTime = field.NewFloat64(tableName, "Compile_time")
	_allClusterSlowQuery.RewriteTime = field.NewFloat64(tableName, "Rewrite_time")
	_allClusterSlowQuery.PreprocSubqueries = field.NewUint64(tableName, "Preproc_subqueries")
	_allClusterSlowQuery.PreprocSubqueriesTime = field.NewFloat64(tableName, "Preproc_subqueries_time")
	_allClusterSlowQuery.OptimizeTime = field.NewFloat64(tableName, "Optimize_time")
	_allClusterSlowQuery.WaitTS = field.NewFloat64(tableName, "Wait_TS")
	_allClusterSlowQuery.PrewriteTime = field.NewFloat64(tableName, "Prewrite_time")
	_allClusterSlowQuery.WaitPrewriteBinlogTime = field.NewFloat64(tableName, "Wait_prewrite_binlog_time")
	_allClusterSlowQuery.CommitTime = field.NewFloat64(tableName, "Commit_time")
	_allClusterSlowQuery.GetCommitTsTime = field.NewFloat64(tableName, "Get_commit_ts_time")
	_allClusterSlowQuery.CommitBackoffTime = field.NewFloat64(tableName, "Commit_backoff_time")
	_allClusterSlowQuery.BackoffTypes = field.NewString(tableName, "Backoff_types")
	_allClusterSlowQuery.ResolveLockTime = field.NewFloat64(tableName, "Resolve_lock_time")
	_allClusterSlowQuery.LocalLatchWaitTime = field.NewFloat64(tableName, "Local_latch_wait_time")
	_allClusterSlowQuery.WriteKeys = field.NewUint64(tableName, "Write_keys")
	_allClusterSlowQuery.WriteSize = field.NewUint64(tableName, "Write_size")
	_allClusterSlowQuery.PrewriteRegion = field.NewUint64(tableName, "Prewrite_region")
	_allClusterSlowQuery.TxnRetry = field.NewUint64(tableName, "Txn_retry")
	_allClusterSlowQuery.CopTime = field.NewFloat64(tableName, "Cop_time")
	_allClusterSlowQuery.ProcessTime = field.NewFloat64(tableName, "Process_time")
	_allClusterSlowQuery.WaitTime = field.NewFloat64(tableName, "Wait_time")
	_allClusterSlowQuery.BackoffTime = field.NewFloat64(tableName, "Backoff_time")
	_allClusterSlowQuery.LockKeysTime = field.NewFloat64(tableName, "LockKeys_time")
	_allClusterSlowQuery.RequestCount = field.NewUint64(tableName, "Request_count")
	_allClusterSlowQuery.TotalKeys = field.NewUint64(tableName, "Total_keys")
	_allClusterSlowQuery.ProcessKeys = field.NewUint64(tableName, "Process_keys")
	_allClusterSlowQuery.RocksdbDeleteSkippedCount = field.NewUint64(tableName, "Rocksdb_delete_skipped_count")
	_allClusterSlowQuery.RocksdbKeySkippedCount = field.NewUint64(tableName, "Rocksdb_key_skipped_count")
	_allClusterSlowQuery.RocksdbBlockCacheHitCount = field.NewUint64(tableName, "Rocksdb_block_cache_hit_count")
	_allClusterSlowQuery.RocksdbBlockReadCount = field.NewUint64(tableName, "Rocksdb_block_read_count")
	_allClusterSlowQuery.RocksdbBlockReadByte = field.NewUint64(tableName, "Rocksdb_block_read_byte")
	_allClusterSlowQuery.DB = field.NewString(tableName, "DB")
	_allClusterSlowQuery.IndexNames = field.NewString(tableName, "Index_names")
	_allClusterSlowQuery.IsInternal = field.NewUint8(tableName, "Is_internal")
	_allClusterSlowQuery.Digest = field.NewString(tableName, "Digest")
	_allClusterSlowQuery.Stats = field.NewString(tableName, "Stats")
	_allClusterSlowQuery.CopProcAvg = field.NewFloat64(tableName, "Cop_proc_avg")
	_allClusterSlowQuery.CopProcP90 = field.NewFloat64(tableName, "Cop_proc_p90")
	_allClusterSlowQuery.CopProcMax = field.NewFloat64(tableName, "Cop_proc_max")
	_allClusterSlowQuery.CopProcAddr = field.NewString(tableName, "Cop_proc_addr")
	_allClusterSlowQuery.CopWaitAvg = field.NewFloat64(tableName, "Cop_wait_avg")
	_allClusterSlowQuery.CopWaitP90 = field.NewFloat64(tableName, "Cop_wait_p90")
	_allClusterSlowQuery.CopWaitMax = field.NewFloat64(tableName, "Cop_wait_max")
	_allClusterSlowQuery.CopWaitAddr = field.NewString(tableName, "Cop_wait_addr")
	_allClusterSlowQuery.MemMax = field.NewUint64(tableName, "Mem_max")
	_allClusterSlowQuery.DiskMax = field.NewUint64(tableName, "Disk_max")
	_allClusterSlowQuery.KVTotal = field.NewFloat64(tableName, "KV_total")
	_allClusterSlowQuery.PDTotal = field.NewFloat64(tableName, "PD_total")
	_allClusterSlowQuery.BackoffTotal = field.NewFloat64(tableName, "Backoff_total")
	_allClusterSlowQuery.WriteSqlResponseTotal = field.NewFloat64(tableName, "Write_sql_response_total")
	_allClusterSlowQuery.ResultRows = field.NewUint64(tableName, "Result_rows")
	_allClusterSlowQuery.Warnings = field.NewString(tableName, "Warnings")
	_allClusterSlowQuery.BackoffDetail = field.NewString(tableName, "Backoff_Detail")
	_allClusterSlowQuery.Prepared = field.NewUint8(tableName, "Prepared")
	_allClusterSlowQuery.Succ = field.NewUint8(tableName, "Succ")
	_allClusterSlowQuery.IsExplicitTxn = field.NewUint8(tableName, "IsExplicitTxn")
	_allClusterSlowQuery.IsWriteCacheTable = field.NewUint8(tableName, "IsWriteCacheTable")
	_allClusterSlowQuery.PlanFromCache = field.NewUint8(tableName, "Plan_from_cache")
	_allClusterSlowQuery.PlanFromBinding = field.NewUint8(tableName, "Plan_from_binding")
	_allClusterSlowQuery.HasMoreResults = field.NewUint8(tableName, "Has_more_results")
	_allClusterSlowQuery.ResourceGroup = field.NewString(tableName, "Resource_group")
	_allClusterSlowQuery.RequestUnitRead = field.NewFloat64(tableName, "Request_unit_read")
	_allClusterSlowQuery.RequestUnitWrite = field.NewFloat64(tableName, "Request_unit_write")
	_allClusterSlowQuery.TimeQueuedByRc = field.NewFloat64(tableName, "Time_queued_by_rc")
	_allClusterSlowQuery.TidbCPUTime = field.NewFloat64(tableName, "Tidb_cpu_time")
	_allClusterSlowQuery.TikvCPUTime = field.NewFloat64(tableName, "Tikv_cpu_time")
	_allClusterSlowQuery.Plan = field.NewString(tableName, "Plan")
	_allClusterSlowQuery.PlanDigest = field.NewString(tableName, "Plan_digest")
	_allClusterSlowQuery.BinaryPlan = field.NewString(tableName, "Binary_plan")
	_allClusterSlowQuery.PrevStmt = field.NewString(tableName, "Prev_stmt")
	_allClusterSlowQuery.Query = field.NewString(tableName, "Query")

	_allClusterSlowQuery.fillFieldMap()

	return _allClusterSlowQuery
}

type allClusterSlowQuery struct {
	allClusterSlowQueryDo

	ALL                       field.Asterisk
	Cluster                   field.String
	SAMPLESTARTTIME           field.Time
	SAMPLEENDTIME             field.Time
	INSTANCE                  field.String
	Time                      field.Time
	TxnStartTs                field.Uint64
	User                      field.String
	Host                      field.String
	ConnID                    field.Uint64
	SessionAlias              field.String
	ExecRetryCount            field.Uint64
	ExecRetryTime             field.Float64
	QueryTime                 field.Float64
	ParseTime                 field.Float64
	CompileTime               field.Float64
	RewriteTime               field.Float64
	PreprocSubqueries         field.Uint64
	PreprocSubqueriesTime     field.Float64
	OptimizeTime              field.Float64
	WaitTS                    field.Float64
	PrewriteTime              field.Float64
	WaitPrewriteBinlogTime    field.Float64
	CommitTime                field.Float64
	GetCommitTsTime           field.Float64
	CommitBackoffTime         field.Float64
	BackoffTypes              field.String
	ResolveLockTime           field.Float64
	LocalLatchWaitTime        field.Float64
	WriteKeys                 field.Uint64
	WriteSize                 field.Uint64
	PrewriteRegion            field.Uint64
	TxnRetry                  field.Uint64
	CopTime                   field.Float64
	ProcessTime               field.Float64
	WaitTime                  field.Float64
	BackoffTime               field.Float64
	LockKeysTime              field.Float64
	RequestCount              field.Uint64
	TotalKeys                 field.Uint64
	ProcessKeys               field.Uint64
	RocksdbDeleteSkippedCount field.Uint64
	RocksdbKeySkippedCount    field.Uint64
	RocksdbBlockCacheHitCount field.Uint64
	RocksdbBlockReadCount     field.Uint64
	RocksdbBlockReadByte      field.Uint64
	DB                        field.String
	IndexNames                field.String
	IsInternal                field.Uint8
	Digest                    field.String
	Stats                     field.String
	CopProcAvg                field.Float64
	CopProcP90                field.Float64
	CopProcMax                field.Float64
	CopProcAddr               field.String
	CopWaitAvg                field.Float64
	CopWaitP90                field.Float64
	CopWaitMax                field.Float64
	CopWaitAddr               field.String
	MemMax                    field.Uint64
	DiskMax                   field.Uint64
	KVTotal                   field.Float64
	PDTotal                   field.Float64
	BackoffTotal              field.Float64
	WriteSqlResponseTotal     field.Float64
	ResultRows                field.Uint64
	Warnings                  field.String
	BackoffDetail             field.String
	Prepared                  field.Uint8
	Succ                      field.Uint8
	IsExplicitTxn             field.Uint8
	IsWriteCacheTable         field.Uint8
	PlanFromCache             field.Uint8
	PlanFromBinding           field.Uint8
	HasMoreResults            field.Uint8
	ResourceGroup             field.String
	RequestUnitRead           field.Float64
	RequestUnitWrite          field.Float64
	TimeQueuedByRc            field.Float64
	TidbCPUTime               field.Float64
	TikvCPUTime               field.Float64
	Plan                      field.String
	PlanDigest                field.String
	BinaryPlan                field.String
	PrevStmt                  field.String
	Query                     field.String

	fieldMap map[string]field.Expr
}

func (a allClusterSlowQuery) Table(newTableName string) *allClusterSlowQuery {
	a.allClusterSlowQueryDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a allClusterSlowQuery) As(alias string) *allClusterSlowQuery {
	a.allClusterSlowQueryDo.DO = *(a.allClusterSlowQueryDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *allClusterSlowQuery) updateTableName(table string) *allClusterSlowQuery {
	a.ALL = field.NewAsterisk(table)
	a.Cluster = field.NewString(table, "Cluster")
	a.SAMPLESTARTTIME = field.NewTime(table, "SAMPLE_START_TIME")
	a.SAMPLEENDTIME = field.NewTime(table, "SAMPLE_END_TIME")
	a.INSTANCE = field.NewString(table, "INSTANCE")
	a.Time = field.NewTime(table, "Time")
	a.TxnStartTs = field.NewUint64(table, "Txn_start_ts")
	a.User = field.NewString(table, "User")
	a.Host = field.NewString(table, "Host")
	a.ConnID = field.NewUint64(table, "Conn_ID")
	a.SessionAlias = field.NewString(table, "Session_alias")
	a.ExecRetryCount = field.NewUint64(table, "Exec_retry_count")
	a.ExecRetryTime = field.NewFloat64(table, "Exec_retry_time")
	a.QueryTime = field.NewFloat64(table, "Query_time")
	a.ParseTime = field.NewFloat64(table, "Parse_time")
	a.CompileTime = field.NewFloat64(table, "Compile_time")
	a.RewriteTime = field.NewFloat64(table, "Rewrite_time")
	a.PreprocSubqueries = field.NewUint64(table, "Preproc_subqueries")
	a.PreprocSubqueriesTime = field.NewFloat64(table, "Preproc_subqueries_time")
	a.OptimizeTime = field.NewFloat64(table, "Optimize_time")
	a.WaitTS = field.NewFloat64(table, "Wait_TS")
	a.PrewriteTime = field.NewFloat64(table, "Prewrite_time")
	a.WaitPrewriteBinlogTime = field.NewFloat64(table, "Wait_prewrite_binlog_time")
	a.CommitTime = field.NewFloat64(table, "Commit_time")
	a.GetCommitTsTime = field.NewFloat64(table, "Get_commit_ts_time")
	a.CommitBackoffTime = field.NewFloat64(table, "Commit_backoff_time")
	a.BackoffTypes = field.NewString(table, "Backoff_types")
	a.ResolveLockTime = field.NewFloat64(table, "Resolve_lock_time")
	a.LocalLatchWaitTime = field.NewFloat64(table, "Local_latch_wait_time")
	a.WriteKeys = field.NewUint64(table, "Write_keys")
	a.WriteSize = field.NewUint64(table, "Write_size")
	a.PrewriteRegion = field.NewUint64(table, "Prewrite_region")
	a.TxnRetry = field.NewUint64(table, "Txn_retry")
	a.CopTime = field.NewFloat64(table, "Cop_time")
	a.ProcessTime = field.NewFloat64(table, "Process_time")
	a.WaitTime = field.NewFloat64(table, "Wait_time")
	a.BackoffTime = field.NewFloat64(table, "Backoff_time")
	a.LockKeysTime = field.NewFloat64(table, "LockKeys_time")
	a.RequestCount = field.NewUint64(table, "Request_count")
	a.TotalKeys = field.NewUint64(table, "Total_keys")
	a.ProcessKeys = field.NewUint64(table, "Process_keys")
	a.RocksdbDeleteSkippedCount = field.NewUint64(table, "Rocksdb_delete_skipped_count")
	a.RocksdbKeySkippedCount = field.NewUint64(table, "Rocksdb_key_skipped_count")
	a.RocksdbBlockCacheHitCount = field.NewUint64(table, "Rocksdb_block_cache_hit_count")
	a.RocksdbBlockReadCount = field.NewUint64(table, "Rocksdb_block_read_count")
	a.RocksdbBlockReadByte = field.NewUint64(table, "Rocksdb_block_read_byte")
	a.DB = field.NewString(table, "DB")
	a.IndexNames = field.NewString(table, "Index_names")
	a.IsInternal = field.NewUint8(table, "Is_internal")
	a.Digest = field.NewString(table, "Digest")
	a.Stats = field.NewString(table, "Stats")
	a.CopProcAvg = field.NewFloat64(table, "Cop_proc_avg")
	a.CopProcP90 = field.NewFloat64(table, "Cop_proc_p90")
	a.CopProcMax = field.NewFloat64(table, "Cop_proc_max")
	a.CopProcAddr = field.NewString(table, "Cop_proc_addr")
	a.CopWaitAvg = field.NewFloat64(table, "Cop_wait_avg")
	a.CopWaitP90 = field.NewFloat64(table, "Cop_wait_p90")
	a.CopWaitMax = field.NewFloat64(table, "Cop_wait_max")
	a.CopWaitAddr = field.NewString(table, "Cop_wait_addr")
	a.MemMax = field.NewUint64(table, "Mem_max")
	a.DiskMax = field.NewUint64(table, "Disk_max")
	a.KVTotal = field.NewFloat64(table, "KV_total")
	a.PDTotal = field.NewFloat64(table, "PD_total")
	a.BackoffTotal = field.NewFloat64(table, "Backoff_total")
	a.WriteSqlResponseTotal = field.NewFloat64(table, "Write_sql_response_total")
	a.ResultRows = field.NewUint64(table, "Result_rows")
	a.Warnings = field.NewString(table, "Warnings")
	a.BackoffDetail = field.NewString(table, "Backoff_Detail")
	a.Prepared = field.NewUint8(table, "Prepared")
	a.Succ = field.NewUint8(table, "Succ")
	a.IsExplicitTxn = field.NewUint8(table, "IsExplicitTxn")
	a.IsWriteCacheTable = field.NewUint8(table, "IsWriteCacheTable")
	a.PlanFromCache = field.NewUint8(table, "Plan_from_cache")
	a.PlanFromBinding = field.NewUint8(table, "Plan_from_binding")
	a.HasMoreResults = field.NewUint8(table, "Has_more_results")
	a.ResourceGroup = field.NewString(table, "Resource_group")
	a.RequestUnitRead = field.NewFloat64(table, "Request_unit_read")
	a.RequestUnitWrite = field.NewFloat64(table, "Request_unit_write")
	a.TimeQueuedByRc = field.NewFloat64(table, "Time_queued_by_rc")
	a.TidbCPUTime = field.NewFloat64(table, "Tidb_cpu_time")
	a.TikvCPUTime = field.NewFloat64(table, "Tikv_cpu_time")
	a.Plan = field.NewString(table, "Plan")
	a.PlanDigest = field.NewString(table, "Plan_digest")
	a.BinaryPlan = field.NewString(table, "Binary_plan")
	a.PrevStmt = field.NewString(table, "Prev_stmt")
	a.Query = field.NewString(table, "Query")

	a.fillFieldMap()

	return a
}

func (a *allClusterSlowQuery) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *allClusterSlowQuery) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 85)
	a.fieldMap["Cluster"] = a.Cluster
	a.fieldMap["SAMPLE_START_TIME"] = a.SAMPLESTARTTIME
	a.fieldMap["SAMPLE_END_TIME"] = a.SAMPLEENDTIME
	a.fieldMap["INSTANCE"] = a.INSTANCE
	a.fieldMap["Time"] = a.Time
	a.fieldMap["Txn_start_ts"] = a.TxnStartTs
	a.fieldMap["User"] = a.User
	a.fieldMap["Host"] = a.Host
	a.fieldMap["Conn_ID"] = a.ConnID
	a.fieldMap["Session_alias"] = a.SessionAlias
	a.fieldMap["Exec_retry_count"] = a.ExecRetryCount
	a.fieldMap["Exec_retry_time"] = a.ExecRetryTime
	a.fieldMap["Query_time"] = a.QueryTime
	a.fieldMap["Parse_time"] = a.ParseTime
	a.fieldMap["Compile_time"] = a.CompileTime
	a.fieldMap["Rewrite_time"] = a.RewriteTime
	a.fieldMap["Preproc_subqueries"] = a.PreprocSubqueries
	a.fieldMap["Preproc_subqueries_time"] = a.PreprocSubqueriesTime
	a.fieldMap["Optimize_time"] = a.OptimizeTime
	a.fieldMap["Wait_TS"] = a.WaitTS
	a.fieldMap["Prewrite_time"] = a.PrewriteTime
	a.fieldMap["Wait_prewrite_binlog_time"] = a.WaitPrewriteBinlogTime
	a.fieldMap["Commit_time"] = a.CommitTime
	a.fieldMap["Get_commit_ts_time"] = a.GetCommitTsTime
	a.fieldMap["Commit_backoff_time"] = a.CommitBackoffTime
	a.fieldMap["Backoff_types"] = a.BackoffTypes
	a.fieldMap["Resolve_lock_time"] = a.ResolveLockTime
	a.fieldMap["Local_latch_wait_time"] = a.LocalLatchWaitTime
	a.fieldMap["Write_keys"] = a.WriteKeys
	a.fieldMap["Write_size"] = a.WriteSize
	a.fieldMap["Prewrite_region"] = a.PrewriteRegion
	a.fieldMap["Txn_retry"] = a.TxnRetry
	a.fieldMap["Cop_time"] = a.CopTime
	a.fieldMap["Process_time"] = a.ProcessTime
	a.fieldMap["Wait_time"] = a.WaitTime
	a.fieldMap["Backoff_time"] = a.BackoffTime
	a.fieldMap["LockKeys_time"] = a.LockKeysTime
	a.fieldMap["Request_count"] = a.RequestCount
	a.fieldMap["Total_keys"] = a.TotalKeys
	a.fieldMap["Process_keys"] = a.ProcessKeys
	a.fieldMap["Rocksdb_delete_skipped_count"] = a.RocksdbDeleteSkippedCount
	a.fieldMap["Rocksdb_key_skipped_count"] = a.RocksdbKeySkippedCount
	a.fieldMap["Rocksdb_block_cache_hit_count"] = a.RocksdbBlockCacheHitCount
	a.fieldMap["Rocksdb_block_read_count"] = a.RocksdbBlockReadCount
	a.fieldMap["Rocksdb_block_read_byte"] = a.RocksdbBlockReadByte
	a.fieldMap["DB"] = a.DB
	a.fieldMap["Index_names"] = a.IndexNames
	a.fieldMap["Is_internal"] = a.IsInternal
	a.fieldMap["Digest"] = a.Digest
	a.fieldMap["Stats"] = a.Stats
	a.fieldMap["Cop_proc_avg"] = a.CopProcAvg
	a.fieldMap["Cop_proc_p90"] = a.CopProcP90
	a.fieldMap["Cop_proc_max"] = a.CopProcMax
	a.fieldMap["Cop_proc_addr"] = a.CopProcAddr
	a.fieldMap["Cop_wait_avg"] = a.CopWaitAvg
	a.fieldMap["Cop_wait_p90"] = a.CopWaitP90
	a.fieldMap["Cop_wait_max"] = a.CopWaitMax
	a.fieldMap["Cop_wait_addr"] = a.CopWaitAddr
	a.fieldMap["Mem_max"] = a.MemMax
	a.fieldMap["Disk_max"] = a.DiskMax
	a.fieldMap["KV_total"] = a.KVTotal
	a.fieldMap["PD_total"] = a.PDTotal
	a.fieldMap["Backoff_total"] = a.BackoffTotal
	a.fieldMap["Write_sql_response_total"] = a.WriteSqlResponseTotal
	a.fieldMap["Result_rows"] = a.ResultRows
	a.fieldMap["Warnings"] = a.Warnings
	a.fieldMap["Backoff_Detail"] = a.BackoffDetail
	a.fieldMap["Prepared"] = a.Prepared
	a.fieldMap["Succ"] = a.Succ
	a.fieldMap["IsExplicitTxn"] = a.IsExplicitTxn
	a.fieldMap["IsWriteCacheTable"] = a.IsWriteCacheTable
	a.fieldMap["Plan_from_cache"] = a.PlanFromCache
	a.fieldMap["Plan_from_binding"] = a.PlanFromBinding
	a.fieldMap["Has_more_results"] = a.HasMoreResults
	a.fieldMap["Resource_group"] = a.ResourceGroup
	a.fieldMap["Request_unit_read"] = a.RequestUnitRead
	a.fieldMap["Request_unit_write"] = a.RequestUnitWrite
	a.fieldMap["Time_queued_by_rc"] = a.TimeQueuedByRc
	a.fieldMap["Tidb_cpu_time"] = a.TidbCPUTime
	a.fieldMap["Tikv_cpu_time"] = a.TikvCPUTime
	a.fieldMap["Plan"] = a.Plan
	a.fieldMap["Plan_digest"] = a.PlanDigest
	a.fieldMap["Binary_plan"] = a.BinaryPlan
	a.fieldMap["Prev_stmt"] = a.PrevStmt
	a.fieldMap["Query"] = a.Query
}

func (a allClusterSlowQuery) clone(db *gorm.DB) allClusterSlowQuery {
	a.allClusterSlowQueryDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a allClusterSlowQuery) replaceDB(db *gorm.DB) allClusterSlowQuery {
	a.allClusterSlowQueryDo.ReplaceDB(db)
	return a
}

type allClusterSlowQueryDo struct{ gen.DO }

type IAllClusterSlowQueryDo interface {
	gen.SubQuery
	Debug() IAllClusterSlowQueryDo
	WithContext(ctx context.Context) IAllClusterSlowQueryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAllClusterSlowQueryDo
	WriteDB() IAllClusterSlowQueryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAllClusterSlowQueryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAllClusterSlowQueryDo
	Not(conds ...gen.Condition) IAllClusterSlowQueryDo
	Or(conds ...gen.Condition) IAllClusterSlowQueryDo
	Select(conds ...field.Expr) IAllClusterSlowQueryDo
	Where(conds ...gen.Condition) IAllClusterSlowQueryDo
	Order(conds ...field.Expr) IAllClusterSlowQueryDo
	Distinct(cols ...field.Expr) IAllClusterSlowQueryDo
	Omit(cols ...field.Expr) IAllClusterSlowQueryDo
	Join(table schema.Tabler, on ...field.Expr) IAllClusterSlowQueryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAllClusterSlowQueryDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAllClusterSlowQueryDo
	Group(cols ...field.Expr) IAllClusterSlowQueryDo
	Having(conds ...gen.Condition) IAllClusterSlowQueryDo
	Limit(limit int) IAllClusterSlowQueryDo
	Offset(offset int) IAllClusterSlowQueryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAllClusterSlowQueryDo
	Unscoped() IAllClusterSlowQueryDo
	Create(values ...*model.AllClusterSlowQuery) error
	CreateInBatches(values []*model.AllClusterSlowQuery, batchSize int) error
	Save(values ...*model.AllClusterSlowQuery) error
	First() (*model.AllClusterSlowQuery, error)
	Take() (*model.AllClusterSlowQuery, error)
	Last() (*model.AllClusterSlowQuery, error)
	Find() ([]*model.AllClusterSlowQuery, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AllClusterSlowQuery, err error)
	FindInBatches(result *[]*model.AllClusterSlowQuery, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AllClusterSlowQuery) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAllClusterSlowQueryDo
	Assign(attrs ...field.AssignExpr) IAllClusterSlowQueryDo
	Joins(fields ...field.RelationField) IAllClusterSlowQueryDo
	Preload(fields ...field.RelationField) IAllClusterSlowQueryDo
	FirstOrInit() (*model.AllClusterSlowQuery, error)
	FirstOrCreate() (*model.AllClusterSlowQuery, error)
	FindByPage(offset int, limit int) (result []*model.AllClusterSlowQuery, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAllClusterSlowQueryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a allClusterSlowQueryDo) Debug() IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Debug())
}

func (a allClusterSlowQueryDo) WithContext(ctx context.Context) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a allClusterSlowQueryDo) ReadDB() IAllClusterSlowQueryDo {
	return a.Clauses(dbresolver.Read)
}

func (a allClusterSlowQueryDo) WriteDB() IAllClusterSlowQueryDo {
	return a.Clauses(dbresolver.Write)
}

func (a allClusterSlowQueryDo) Session(config *gorm.Session) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Session(config))
}

func (a allClusterSlowQueryDo) Clauses(conds ...clause.Expression) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a allClusterSlowQueryDo) Returning(value interface{}, columns ...string) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a allClusterSlowQueryDo) Not(conds ...gen.Condition) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a allClusterSlowQueryDo) Or(conds ...gen.Condition) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a allClusterSlowQueryDo) Select(conds ...field.Expr) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a allClusterSlowQueryDo) Where(conds ...gen.Condition) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a allClusterSlowQueryDo) Order(conds ...field.Expr) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a allClusterSlowQueryDo) Distinct(cols ...field.Expr) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a allClusterSlowQueryDo) Omit(cols ...field.Expr) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a allClusterSlowQueryDo) Join(table schema.Tabler, on ...field.Expr) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a allClusterSlowQueryDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a allClusterSlowQueryDo) RightJoin(table schema.Tabler, on ...field.Expr) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a allClusterSlowQueryDo) Group(cols ...field.Expr) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a allClusterSlowQueryDo) Having(conds ...gen.Condition) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a allClusterSlowQueryDo) Limit(limit int) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a allClusterSlowQueryDo) Offset(offset int) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a allClusterSlowQueryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a allClusterSlowQueryDo) Unscoped() IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Unscoped())
}

func (a allClusterSlowQueryDo) Create(values ...*model.AllClusterSlowQuery) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a allClusterSlowQueryDo) CreateInBatches(values []*model.AllClusterSlowQuery, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a allClusterSlowQueryDo) Save(values ...*model.AllClusterSlowQuery) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a allClusterSlowQueryDo) First() (*model.AllClusterSlowQuery, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AllClusterSlowQuery), nil
	}
}

func (a allClusterSlowQueryDo) Take() (*model.AllClusterSlowQuery, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AllClusterSlowQuery), nil
	}
}

func (a allClusterSlowQueryDo) Last() (*model.AllClusterSlowQuery, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AllClusterSlowQuery), nil
	}
}

func (a allClusterSlowQueryDo) Find() ([]*model.AllClusterSlowQuery, error) {
	result, err := a.DO.Find()
	return result.([]*model.AllClusterSlowQuery), err
}

func (a allClusterSlowQueryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AllClusterSlowQuery, err error) {
	buf := make([]*model.AllClusterSlowQuery, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a allClusterSlowQueryDo) FindInBatches(result *[]*model.AllClusterSlowQuery, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a allClusterSlowQueryDo) Attrs(attrs ...field.AssignExpr) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a allClusterSlowQueryDo) Assign(attrs ...field.AssignExpr) IAllClusterSlowQueryDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a allClusterSlowQueryDo) Joins(fields ...field.RelationField) IAllClusterSlowQueryDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a allClusterSlowQueryDo) Preload(fields ...field.RelationField) IAllClusterSlowQueryDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a allClusterSlowQueryDo) FirstOrInit() (*model.AllClusterSlowQuery, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AllClusterSlowQuery), nil
	}
}

func (a allClusterSlowQueryDo) FirstOrCreate() (*model.AllClusterSlowQuery, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AllClusterSlowQuery), nil
	}
}

func (a allClusterSlowQueryDo) FindByPage(offset int, limit int) (result []*model.AllClusterSlowQuery, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a allClusterSlowQueryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a allClusterSlowQueryDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a allClusterSlowQueryDo) Delete(models ...*model.AllClusterSlowQuery) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *allClusterSlowQueryDo) withDO(do gen.Dao) *allClusterSlowQueryDo {
	a.DO = *do.(*gen.DO)
	return a
}
