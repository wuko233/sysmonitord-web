package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// 服务端配置
type ServerConfig struct {
	ServerPort int    `mapstructure:"server_port"`
	Mode       string `mapstructure:"mode"`
}

type Databaseconfig struct {
	Type     string `mapstructure:"type"`
	Path     string `mapstructure:"path"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"db_name"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database Databaseconfig `mapstructure:"database"`
}

var GlobalConfig *Config

func LoadConfig(configPath string) error {
	v := viper.New()

	v.SetConfigFile(configPath)

	if configPath == "" {
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		v.AddConfigPath("./config")
	}

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("解析配置文件失败: %w", err)
	}

	GlobalConfig = &cfg

	log.Printf("配置加载成功: %+v", GlobalConfig)

	return nil
}
