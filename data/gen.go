package main

import (
	"strings"

	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// 更新会覆盖原有文件，所以通过 g.GenerateModel("oss", fieldOpts...) 指定需要更新的表，不要全部覆盖

const DBDSN = "root:@(localhost:4000)/test?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	cfg := gen.Config{
		OutPath: "../dal",                                                           // 输出的文件夹
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: true, // generate pointer when field is nullable

		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false, // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values

		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: false, // detect integer field's unsigned type, adjust generated data type
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: false, // generate with gorm index tag
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true, // generate with gorm column type tag
	}
	// 处理表名
	cfg.WithTableNameStrategy(func(tableName string) (targetTableName string) {
		// 需要忽略的表
		if strings.EqualFold(tableName, "pay_credit") ||
			strings.EqualFold(tableName, "pay_log") ||
			strings.EqualFold(tableName, "space") ||
			strings.EqualFold(tableName, "space_account") {
			return ""
		}
		return tableName
	})
	// 处理 model名
	cfg.WithModelNameStrategy(func(tableName string) (targetTableName string) {
		s := tableName
		if strings.HasPrefix(tableName, "conf_") {
			s = strings.TrimPrefix(tableName, "conf_")
		}
		ns := schema.NamingStrategy{}
		return ns.SchemaName(s)
	})
	// 处理文件名
	cfg.WithFileNameStrategy(func(tableName string) (targetTableName string) {
		if strings.HasPrefix(tableName, "conf_") {
			return strings.TrimPrefix(tableName, "conf_")
		}
		return tableName
	})
	g := gen.NewGenerator(cfg)

	gormdb, _ := gorm.Open(mysql.Open(DBDSN))

	g.UseDB(gormdb) // reuse your gorm db

	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf
	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "uint8" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int32" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int32" },
		"bigint": func(detailType gorm.ColumnType) (dataType string) {
			if detailType.Name() == "size" {
				return "uint64"
			}
			return "uint64"
		},
		"int": func(detailType gorm.ColumnType) (dataType string) { return "uint32" },
	}
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)
	// 自定义模型结体字段的标签
	// 将特定字段名的 json 标签加上`string`属性,即 MarshalJSON 时该字段由数字类型转成字符串类型
	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
		toStringField := `balance, `
		if strings.Contains(toStringField, columnName) {
			return columnName + ",string"
		}
		return columnName
	})
	delField := gen.FieldType("deleted_at", "time.Time") // 不生成 默认的类型
	//sizeField := gen.FieldType("size", "uint64")         // 不生成 默认的类型
	// 将非默认字段名的字段定义为自动时间戳和软删除字段;
	// 自动时间戳默认字段名为:`updated_at`、`created_at, 表字段数据类型为: INT 或 DATETIME
	// 软删除默认字段名为:`deleted_at`, 表字段数据类型为: DATETIME
	//autoUpdateTimeField := gen.FieldGORMTag("update_time", "column:update_time;type:int unsigned;autoUpdateTime")
	//autoCreateTimeField := gen.FieldGORMTag("create_time", "column:create_time;type:int unsigned;autoCreateTime")
	//softDeleteField := gen.FieldType("delete_time", "soft_delete.DeletedAt")
	// 模型自定义选项组
	fieldOpts := []gen.ModelOpt{jsonField, delField}

	// 创建模型的结构体,生成文件在 model 目录; 先创建的结果会被后面创建的覆盖
	// 这里创建个别模型仅仅是为了拿到`*generate.QueryStructMeta`类型对象用于后面的模型关联操作中
	// Address := g.GenerateModel("address")

	// 创建 全部模型文件 , 并覆盖前面创建的同名模型
	allModel := g.GenerateAllTable(fieldOpts...)

	// 指定特定的表名
	//models := []interface{}{
	//  g.GenerateModel("workspace_subscribe_log", fieldOpts...),
	//  g.GenerateModel("workspace_bill_log", fieldOpts...),
	//  g.GenerateModel("workspace_version", fieldOpts...),
	//  g.GenerateModel("storage_usage", fieldOpts...),
	//  g.GenerateModel("inbox", fieldOpts...),
	//}

	// 创建模型的方法,生成文件在 query 目录; 先创建结果不会被后创建的覆盖
	g.ApplyBasic()
	g.ApplyBasic(allModel...)

	g.Execute()
}
