package mysql

import (
	"fmt"
	"nexus-ai/constant"
	"nexus-ai/utils"
	"strconv"
	"time"
)

// Config MySQL配置结构体
type Config struct {
	// 基础配置组
	Basic struct {
		Host     string `json:"host" yaml:"host"`
		Port     int    `json:"port" yaml:"port"`
		User     string `json:"user" yaml:"user"`
		Password string `json:"password" yaml:"password"`
		Database string `json:"database" yaml:"database"`
	}

	// 连接池配置组
	Pool struct {
		MaxOpenConns    int           `json:"maxOpenConns" yaml:"maxOpenConns"`
		MaxIdleConns    int           `json:"maxIdleConns" yaml:"maxIdleConns"`
		ConnMaxLifetime time.Duration `json:"connMaxLifetime" yaml:"connMaxLifetime"`
		ConnMaxIdleTime time.Duration `json:"connMaxIdleTime" yaml:"connMaxIdleTime"`
		MaxConnAttempts int           `json:"maxConnAttempts" yaml:"maxConnAttempts"`
		RetryBackoff    time.Duration `json:"retryBackoff" yaml:"retryBackoff"`
	}
}

// Validate 验证配置的合法性
func (c *Config) Validate() error {
	// 验证基础配置
	if c.Basic.Host == "" {
		return fmt.Errorf("database host cannot be empty")
	}
	if c.Basic.Port < 1 || c.Basic.Port > 65535 {
		return fmt.Errorf("invalid port number: %d", c.Basic.Port)
	}
	if c.Basic.User == "" {
		return fmt.Errorf("database user cannot be empty")
	}
	if c.Basic.Database == "" {
		return fmt.Errorf("database name cannot be empty")
	}

	// 验证连接池配置
	if c.Pool.MaxOpenConns <= 0 {
		return fmt.Errorf("max open connections must be positive")
	}
	if c.Pool.MaxIdleConns <= 0 {
		return fmt.Errorf("max idle connections must be positive")
	}
	if c.Pool.MaxOpenConns < c.Pool.MaxIdleConns {
		return fmt.Errorf("max open connections (%d) cannot be smaller than max idle connections (%d)",
			c.Pool.MaxOpenConns, c.Pool.MaxIdleConns)
	}
	if c.Pool.ConnMaxLifetime <= 0 || c.Pool.ConnMaxIdleTime <= 0 {
		return fmt.Errorf("connection lifetime values must be positive")
	}
	if c.Pool.MaxConnAttempts <= 0 {
		return fmt.Errorf("max connection attempts must be positive")
	}
	if c.Pool.RetryBackoff <= 0 {
		return fmt.Errorf("retry backoff must be positive")
	}
	return nil
}

// DefaultConfig 返回默认的MySQL配置
func DefaultConfig() *Config {
	cfg := &Config{}

	// 解析基础配置
	port, err := strconv.Atoi(utils.GetEnv("MYSQL_PORT", constant.MysqlDefaultPort))
	if err != nil {
		port, _ = strconv.Atoi(constant.MysqlDefaultPort)
	}

	cfg.Basic.Host = utils.GetEnv("MYSQL_HOST", constant.MysqlDefaultHost)
	cfg.Basic.Port = port
	cfg.Basic.User = utils.GetEnv("MYSQL_USER", constant.MysqlDefaultUser)
	cfg.Basic.Password = utils.GetEnv("MYSQL_PASSWORD", constant.MysqlDefaultPassword)
	cfg.Basic.Database = utils.GetEnv("MYSQL_DATABASE", constant.MysqlDefaultDatabase)

	utils.SysInfo(fmt.Sprintf("MySQL | host: %s | port: %d | user: %s | database: %s",
		cfg.Basic.Host, cfg.Basic.Port, cfg.Basic.User, cfg.Basic.Database))

	// 设置连接池配置
	cfg.Pool.MaxOpenConns = 100
	cfg.Pool.MaxIdleConns = 10
	cfg.Pool.ConnMaxLifetime = 10 * time.Minute
	cfg.Pool.ConnMaxIdleTime = 5 * time.Minute
	cfg.Pool.MaxConnAttempts = 5
	cfg.Pool.RetryBackoff = 1 * time.Second

	return cfg
}
