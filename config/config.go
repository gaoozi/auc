package config

import (
	"fmt"
	"log/slog"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var (
  config Config
  once sync.Once
)

type ServerConfig struct {
	Name string
	Port int
}

type DatabaseConfig struct {
	Url         string
	User        string
	Password    string
	Name        string
  MaxIdleConn int `mapstructure:"max_idle_conn"`
	MaxOpenConn int `mapstructure:"max_open_conn"`
	MaxIdleTime int `mapstructure:"max_idle_time"`
}

type Config struct {
  Server ServerConfig
  Database DatabaseConfig
}

func initial() {
  viper.SetConfigType("yml")
  viper.SetConfigFile("config/config.yml")
  viper.AutomaticEnv()
  viper.SetEnvPrefix("auc")
  viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
  viper.BindEnv("database.user")
  viper.BindEnv("database.password")

  err := viper.ReadInConfig()
  if err != nil {
    panic("read config file err:" + err.Error())
  }

  err = viper.Unmarshal(&config)
  if err != nil {
    panic("config file unmarshal err:" + err.Error())
  }

  slog.Info("load config success.")
  slog.Info(fmt.Sprintf("config: %+v", config))
}

func GetConfig() *Config {
  once.Do(initial)
  return &config
}
