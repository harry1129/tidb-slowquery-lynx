// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameAllClusterSlowQueryInfo = "all_cluster_slow_query_info"

// AllClusterSlowQueryInfo mapped from table <all_cluster_slow_query_info>
type AllClusterSlowQueryInfo struct {
	Cluster    string  `gorm:"column:Cluster;type:varchar(64);primaryKey" json:"Cluster"`
	Digest     string  `gorm:"column:Digest;type:varchar(64);primaryKey" json:"Digest"`
	Plan       *string `gorm:"column:Plan;type:longtext" json:"Plan"`
	PlanDigest string  `gorm:"column:Plan_digest;type:varchar(128);primaryKey" json:"Plan_digest"`
	BinaryPlan *string `gorm:"column:Binary_plan;type:longtext" json:"Binary_plan"`
	Query      *string `gorm:"column:Query;type:longtext" json:"Query"`
}

// TableName AllClusterSlowQueryInfo's table name
func (*AllClusterSlowQueryInfo) TableName() string {
	return TableNameAllClusterSlowQueryInfo
}
