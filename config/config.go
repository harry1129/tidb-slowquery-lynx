package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Global struct {
		TimeWindowMinutes int `toml:"time_window_minutes"`
	} `toml:"global"`
	Database struct {
		Host            string `toml:"host"`
		Port            int    `toml:"port"`
		User            string `toml:"user"`
		Password        string `toml:"password"`
		DBName          string `toml:"db_name"`
		MaxIdleConns    int    `toml:"max_idle_conns"`
		MaxOpenConns    int    `toml:"max_open_conns"`
		ConnMaxLifetime int    `toml:"conn_max_lifetime"`
	} `toml:"database"`
	TargetDBs map[string]struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		User     string `toml:"user"`
		Password string `toml:"password"`
	} `toml:"target_dbs"`
}

var C = Config{}

func LoadConfig(filename string) error {
	if _, err := toml.DecodeFile(filename, &C); err != nil {
		fmt.Printf("error parsing config file: %v", err)
		os.Exit(1)
	}
	return nil
}
