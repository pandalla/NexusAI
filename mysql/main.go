package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	mu   sync.RWMutex
	conf *Config
)

// GetDB 获取MySQL数据库连接实例
func GetDB() *sql.DB {
	mu.RLock()
	defer mu.RUnlock()
	return db
}

// HealthCheck 健康检查
func HealthCheck(ctx context.Context) error {
	client := GetDB()
	if client == nil {
		return fmt.Errorf("mysql health check failed: can't get running mysql connection")
	}
	return client.PingContext(ctx)
}

// GetConfig 获取MySQL配置
func GetConfig() *Config {
	return conf
}

// InitMySQL 初始化MySQL连接池
func InitMySQL(cfg *Config) error {
	if cfg == nil {
		cfg = DefaultConfig()
	}
	conf = cfg

	if err := cfg.Validate(); err != nil {
		return fmt.Errorf("invalid mysql config: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true&sql_mode='NO_ENGINE_SUBSTITUTION'",
		cfg.Basic.User,
		cfg.Basic.Password,
		cfg.Basic.Host,
		cfg.Basic.Port,
		cfg.Basic.Database,
	)

	var lastErr error
	var database *sql.DB

	// 使用指数退避重试连接
	for i := 0; i < cfg.Pool.MaxConnAttempts; i++ {
		database, lastErr = sql.Open("mysql", dsn)
		if lastErr == nil {
			// 测试连接
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			lastErr = database.PingContext(ctx)
			cancel()

			if lastErr == nil {
				break
			}
		}

		if i < cfg.Pool.MaxConnAttempts-1 {
			sleepTime := time.Duration(1<<uint(i)) * cfg.Pool.RetryBackoff
			if sleepTime > 30*time.Second {
				sleepTime = 30 * time.Second
			}
			time.Sleep(sleepTime)
		}
	}

	if lastErr != nil {
		return fmt.Errorf("mysql connection failed after %d attempts: %v", cfg.Pool.MaxConnAttempts, lastErr)
	}

	// 配置连接池
	database.SetMaxOpenConns(cfg.Pool.MaxOpenConns)
	database.SetMaxIdleConns(cfg.Pool.MaxIdleConns)
	database.SetConnMaxLifetime(cfg.Pool.ConnMaxLifetime)
	database.SetConnMaxIdleTime(cfg.Pool.ConnMaxIdleTime)

	mu.Lock()
	db = database
	mu.Unlock()

	return nil
}

// GracefulClose 优雅关闭MySQL连接
func GracefulClose(ctx context.Context) error {
	mu.Lock()
	defer mu.Unlock()

	if db != nil {
		if err := db.Close(); err != nil {
			return fmt.Errorf("error closing mysql connection: %v", err)
		}
		db = nil
	}
	return nil
}
