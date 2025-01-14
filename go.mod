module tidb-slowquery-lynx

go 1.22.1

require gorm.io/driver/mysql v1.5.7

require go.uber.org/multierr v1.10.0 // indirect

require (
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d // indirect
	gorm.io/datatypes v1.1.1-0.20230130040222-c43177d3cf8c // indirect
	gorm.io/hints v1.1.0 // indirect
	gorm.io/plugin/dbresolver v1.5.3
)

require (
	github.com/BurntSushi/toml v1.4.0
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/things-go/gormzap v0.0.10
	go.uber.org/zap v1.27.0
	gorm.io/gen v0.3.26
	gorm.io/gorm v1.25.12
)
