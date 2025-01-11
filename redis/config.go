package redis

import (
	"fmt"
	"nexus-ai/constant"
	"nexus-ai/utils"
	"strconv"
	"time"
)

// Config Redis配置结构体
type Config struct {
	// 基础配置组
	Basic struct {
		Host     string `json:"host" yaml:"host"`
		Password string `json:"password" yaml:"password"`
		Port     int    `json:"port" yaml:"port"`
		DB       int    `json:"db" yaml:"db"`
	}

	// 超时配置组
	Timeout struct {
		Dial  time.Duration `json:"dial" yaml:"dial"`
		Read  time.Duration `json:"read" yaml:"read"`
		Write time.Duration `json:"write" yaml:"write"`
	}

	// 连接池配置组
	Pool struct {
		Size            int           `json:"size" yaml:"size"`
		MinIdleConns    int           `json:"minIdleConns" yaml:"minIdleConns"`
		MaxConnAttempts int           `json:"maxConnAttempts" yaml:"maxConnAttempts"`
		MaxLifetime     time.Duration `json:"maxLifetime" yaml:"maxLifetime"`
		Timeout         time.Duration `json:"timeout" yaml:"timeout"`
		MaxIdleTime     time.Duration `json:"maxIdleTime" yaml:"maxIdleTime"`
		RetryBackoff    time.Duration `json:"retryBackoff" yaml:"retryBackoff"`
	}
}

// Validate 验证配置的合法性
func (c *Config) Validate() error {
	// 验证基础配置
	if c.Basic.Host == "" {
		return fmt.Errorf("redis host cannot be empty")
	}
	if c.Basic.Port < 1 || c.Basic.Port > 65535 {
		return fmt.Errorf("invalid port number: %d", c.Basic.Port)
	}
	if c.Basic.DB < 0 {
		return fmt.Errorf("invalid DB number: %d", c.Basic.DB)
	}

	// 验证连接池配置
	if c.Pool.Size <= 0 {
		return fmt.Errorf("pool size must be positive")
	}
	if c.Pool.MinIdleConns < 0 {
		return fmt.Errorf("min idle connections cannot be negative")
	}
	if c.Pool.Size < c.Pool.MinIdleConns {
		return fmt.Errorf("pool size (%d) cannot be smaller than min idle connections (%d)",
			c.Pool.Size, c.Pool.MinIdleConns)
	}
	if c.Pool.MaxConnAttempts <= 0 {
		return fmt.Errorf("max connection attempts must be positive")
	}
	if c.Pool.Timeout <= 0 || c.Pool.MaxLifetime <= 0 || c.Pool.MaxIdleTime <= 0 {
		return fmt.Errorf("pool timeout values must be positive")
	}
	if c.Pool.Timeout > 24*time.Hour || c.Pool.MaxLifetime > 24*time.Hour || c.Pool.MaxIdleTime > 24*time.Hour {
		return fmt.Errorf("pool timeout values too large")
	}
	if c.Pool.RetryBackoff <= 0 {
		return fmt.Errorf("retry backoff must be positive")
	}
	// 验证超时配置
	if c.Timeout.Dial <= 0 || c.Timeout.Read <= 0 || c.Timeout.Write <= 0 {
		return fmt.Errorf("timeout values must be positive")
	}

	return nil
}

// DefaultConfig 返回默认的Redis配置
func DefaultConfig() *Config {
	cfg := &Config{}

	// 解析基础配置
	port, err := strconv.Atoi(utils.GetEnv("REDIS_PORT", constant.RedisDefaultPort))
	if err != nil {
		port, _ = strconv.Atoi(constant.RedisDefaultPort)
	}

	db, err := strconv.Atoi(utils.GetEnv("REDIS_DB", constant.RedisDefaultDB))
	if err != nil {
		db, _ = strconv.Atoi(constant.RedisDefaultDB)
	}

	cfg.Basic.Host = utils.GetEnv("REDIS_HOST", constant.RedisDefaultHost)
	cfg.Basic.Password = utils.GetEnv("REDIS_PASSWORD", constant.RedisDefaultPassword)
	cfg.Basic.Port = port
	cfg.Basic.DB = db
	utils.SysInfo("Redis | host: " + cfg.Basic.Host + " | port: " + strconv.Itoa(cfg.Basic.Port) + " | password: " + cfg.Basic.Password + " | db: " + strconv.Itoa(cfg.Basic.DB))
	// 设置超时配置
	cfg.Timeout.Dial = 10 * time.Second
	cfg.Timeout.Read = 10 * time.Second
	cfg.Timeout.Write = 10 * time.Second

	// 设置连接池配置
	poolSize, _ := strconv.Atoi(utils.GetEnv("REDIS_MAX_POOL_SIZE", constant.RedisDefaultMaxPoolSize))
	minIdleConns, _ := strconv.Atoi(utils.GetEnv("REDIS_MIN_IDLE_CONNS", constant.RedisDefaultMinIdleConns))

	cfg.Pool.Size = poolSize
	cfg.Pool.MinIdleConns = minIdleConns
	cfg.Pool.MaxConnAttempts = 10
	cfg.Pool.MaxLifetime = 10 * time.Minute
	cfg.Pool.Timeout = 10 * time.Second
	cfg.Pool.MaxIdleTime = 10 * time.Minute
	cfg.Pool.RetryBackoff = 1 * time.Second

	return cfg
}
