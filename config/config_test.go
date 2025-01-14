package config

import (
	"os"
	"testing"
)

// 辅助函数，用于比较两个 Config 结构体是否相等
func compareConfig(c1, c2 Config) bool {
	if c1.Global.TimeWindowMinutes != c2.Global.TimeWindowMinutes {
		return false
	}
	if c1.Database.Host != c2.Database.Host ||
		c1.Database.Port != c2.Database.Port ||
		c1.Database.User != c2.Database.User ||
		c1.Database.Password != c2.Database.Password ||
		c1.Database.DBName != c2.Database.DBName ||
		c1.Database.MaxIdleConns != c2.Database.MaxIdleConns ||
		c1.Database.MaxOpenConns != c2.Database.MaxOpenConns ||
		c1.Database.ConnMaxLifetime != c2.Database.ConnMaxLifetime {
		return false
	}
	if len(c1.TargetDBs) != len(c2.TargetDBs) {
		return false
	}
	for key, db1 := range c1.TargetDBs {
		db2, exists := c2.TargetDBs[key]
		if !exists {
			return false
		}
		if db1.Host != db2.Host ||
			db1.Port != db2.Port ||
			db1.User != db2.User ||
			db1.Password != db2.Password {
			return false
		}
	}
	return true
}

// 测试正常情况，配置文件正确
func TestLoadConfig_Normal(t *testing.T) {
	// 创建一个临时的正确配置文件
	tmpFile, err := os.CreateTemp("", "config_test_*.toml")
	if err != nil {
		t.Fatalf("创建临时文件失败: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// 写入正确的配置内容
	configContent := `
[global]
time_window_minutes = 10

[database]
host = "localhost"
port = 3306
user = "root"
password = "123456"
db_name = "test_db"
#max_idle_conns = 10
max_open_conns = 100
conn_max_lifetime = 3600

[target_dbs.db1]
host = "192.168.1.1"
port = 3307
user = "user1"
password = "pwd1"

[target_dbs.db2]
host = "192.168.1.2"
port = 3308
user = "user2"
password = "pwd2"
`
	if _, err := tmpFile.WriteString(configContent); err != nil {
		t.Fatalf("写入临时文件失败: %v", err)
	}

	// 调用 LoadConfig 函数
	err = LoadConfig(tmpFile.Name())
	if err != nil {
		t.Errorf("LoadConfig 函数返回错误: %v", err)
	}

	// 检查配置是否正确加载
	expectedConfig := Config{
		Global: struct {
			TimeWindowMinutes int `toml:"time_window_minutes"`
		}{
			TimeWindowMinutes: 10,
		},
		Database: struct {
			Host            string `toml:"host"`
			Port            int    `toml:"port"`
			User            string `toml:"user"`
			Password        string `toml:"password"`
			DBName          string `toml:"db_name"`
			MaxIdleConns    int    `toml:"max_idle_conns"`
			MaxOpenConns    int    `toml:"max_open_conns"`
			ConnMaxLifetime int    `toml:"conn_max_lifetime"`
		}{
			Host:            "localhost",
			Port:            3306,
			User:            "root",
			Password:        "123456",
			DBName:          "test_db",
			MaxIdleConns:    10,
			MaxOpenConns:    100,
			ConnMaxLifetime: 3600,
		},
		TargetDBs: map[string]struct {
			Host     string `toml:"host"`
			Port     int    `toml:"port"`
			User     string `toml:"user"`
			Password string `toml:"password"`
		}{
			"db1": {
				Host:     "192.168.1.1",
				Port:     3307,
				User:     "user1",
				Password: "pwd1",
			},
			"db2": {
				Host:     "192.168.1.2",
				Port:     3308,
				User:     "user2",
				Password: "pwd2",
			},
		},
	}

	if !compareConfig(C, expectedConfig) {
		t.Errorf("配置加载不正确，期望值: %+v，实际值: %+v", expectedConfig, C)
	}
}

// 测试配置文件不存在的情况
func TestLoadConfig_FileNotExist(t *testing.T) {
	err := LoadConfig("non_existent_config.toml")
	if err == nil {
		t.Errorf("期望 LoadConfig 函数返回错误，但实际没有返回错误")
	}
}

// 测试配置文件格式错误的情况
func TestLoadConfig_FormatError(t *testing.T) {
	// 创建一个临时的格式错误的配置文件
	tmpFile, err := os.CreateTemp("", "config_test_*.toml")
	if err != nil {
		t.Fatalf("创建临时文件失败: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// 写入格式错误的配置内容
	configContent := `
[global]
time_window_minutes = "ten"  // 类型错误，应该是整数

[database]
host = localhost  // 缺少引号
port = 3306
user = root
password = 123456  // 类型错误，应该是字符串
db_name = test_db
max_idle_conns = 10
max_open_conns = 100
conn_max_lifetime = 3600
`
	if _, err := tmpFile.WriteString(configContent); err != nil {
		t.Fatalf("写入临时文件失败: %v", err)
	}

	// 调用 LoadConfig 函数
	err = LoadConfig(tmpFile.Name())
	if err == nil {
		t.Errorf("期望 LoadConfig 函数返回错误，但实际没有返回错误")
	}
}
