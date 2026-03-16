package database

import (
	"database/sql"
	"fmt"
	"log"
	"sysmonitor-web/config"

	// _ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/glebarez/sqlite"
)

var DB *sql.DB

func InitDB() error {
	var err error
	cfg := config.GlobalConfig.Database

	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
	// 	cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	dsn := cfg.Path

	// DB, err = sql.Open("pgx", dsn)
	DB, err = sql.Open("sqlite", dsn)
	if err != nil {
		return fmt.Errorf("连接数据库失败: %w", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("数据库连接测试失败: %w", err)
	}

	log.Printf("数据库连接成功: %s", dsn)

	DB.SetMaxOpenConns(1)
	DB.SetMaxIdleConns(1)

	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
