# 简介
lynx 工具可以定时将 TiDB 集群的慢查询收集并持久化到后端数据库中，然后通过 grafana 查询展示出来，这可以帮助我们更好的分析慢查询日志。
# 背景
尽管 TiDB 提供了慢查询日志和 SQL 语句分析功能，并在 Dashboard 上提供了直观的查询界面，但在实际使用中仍存在一些局限性：
1. 早期版本 SQL 语句分析功能数据存储在内存中，当 TiDB Server 发生 Out-of-Memory (OOM) 时，数据会丢失，导致无法查询到相关数据。
2. SQL 语句分析功能存储的 SQL 类别有限，当 SQL 语句数量较多时，旧的记录可能会驱逐，从而无法查询到想要的数据。
3. 慢查询日志以单条语句的形式存在，当 TiDB 集群繁忙时，短时间会产生大量慢查询记录，界面返回的行数过多，不利于分析。
4. 慢查询日志底层存储在日志文件中，跨长时间段查询慢查询日志时，不仅速度较慢，还可能导致 TiDB Server OOM。
5. 慢查询日志界面无法自定义查询，无法满足复杂的分析需求。
6. 慢查询日志导出的 Excel 格式混乱，无法直接发送给应用研发团队。
7. 等等 。。。
# 特点
1. 支持多 TiDB 版本，适用于 TiDB v6.0 及以上版本。
2. 根据时间窗口采集慢查询数据，有效降低对目标 TiDB 集群的压力。
3. 同一集群中，具有相同 SQL 指纹和计划指纹的 SQL 语句和执行计划只收集一次，节约了存储空间，同时也降低了对采集目标集群的压力。
4. 可以灵活自定义 Grafana 面板，例如查找执行计划中包含隐式转换的 SQL 语句（计划中包含 cast 关键字）。
5. 通过 Grafana 可以方便地将慢查询数据导出为 Excel 格式，还可以开发脚本定期自动将数据通过邮件发送给相关人员。
6. lynx 收集的慢日志适合分析因执行计划不佳导致的查询慢的问题。然而，对于那些时快时慢的查询情况（可能是由于传入参数不同或集群运行状态变化引起的），TiDB Dashboard 的慢查询日志功能可能更为合适，因为同一执行计划 lynx 只采集一次执行计划。
7. 与慢查询日志一样，Lynx 无法获取低于慢查询阈值的语句信息。
# 部署lynx
## 1. 准备配置文件config.toml
```toml
[global]
time_window_minutes = 60  #采集窗口，单位（分钟），建议设置 30 ～ 60，与定时任务配置同样窗口

#后端存储库
[database]
host = "localhost"
port = 4000
user = "root"
password = ""
db_name = "test"
max_idle_conns = 2
max_open_conns = 10
conn_max_lifetime = 600    #单位（秒）

#采集目标库
[target_dbs]
cluster1.host = "196.128.1.1"
cluster1.port = 4000
cluster1.user = "root1"
cluster1.password = ""

cluster2.host = "196.128.1.2"
cluster2.port = 4000
cluster2.user = "root"
cluster2.password = ""


cluster3.host = "196.128.1.4"
cluster3.port = 4000
cluster3.user = "root"
cluster3.password = ""
```
目标数据库的采集用户最少需要以下权限
```sql
-- 查询 information_schema.schema 时需要看到所有的数据库名
GRANT SHOW DATABASES ON . TO 'xxxx'@'%';
-- 查询 information_schema.cluster_slow_query 时需要看到所有用户的慢查询
GRANT DASHBOARD_CLIENT ON . TO 'xxxx'@'%';
```
## 2. 添加定时执行任务
这里定时执行时间要与 time_window_minutes 一样
```shell
crontab -e
*/30 * * * * /your/path/lynx -c /your/path/config.toml -l /your/path/lynx.log
```
## 3. 添加 grafana 面板
1. 在 Grafana 上配置 mysql 数据源
Configuration =》 data sources =》Add data source =》Mysql
2. 将 https://github.com/harry1129/tidb-slowquery-lynx/tree/main/grafana 中带的面板导入到 Grafana
Create =》Import