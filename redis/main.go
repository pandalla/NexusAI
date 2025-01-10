package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	mu          sync.RWMutex
)

// 获取Redis客户端实例
func GetRedisClient() *redis.Client {
	mu.RLock()
	defer mu.RUnlock()
	return redisClient
}

// 健康检查
func HealthCheck(ctx context.Context) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("redis health check failed: can't get running redis client")
	}
	return client.Ping(ctx).Err()
}

// 初始化Redis连接池
func InitRedis(cfg *Config) error {
	if cfg == nil {
		cfg = DefaultConfig()
	}

	if err := cfg.Validate(); err != nil {
		return fmt.Errorf("invalid redis config: %v", err)
	}

	// 创建Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Basic.Host, cfg.Basic.Port),
		Password: cfg.Basic.Password,
		DB:       cfg.Basic.DB,

		// 超时配置
		DialTimeout:  cfg.Timeout.Dial,
		ReadTimeout:  cfg.Timeout.Read,
		WriteTimeout: cfg.Timeout.Write,

		// 连接池配置
		PoolSize:        cfg.Pool.Size,
		MinIdleConns:    cfg.Pool.MinIdleConns,
		ConnMaxLifetime: cfg.Pool.MaxLifetime,
		PoolTimeout:     cfg.Pool.Timeout,
		ConnMaxIdleTime: cfg.Pool.MaxIdleTime,
	})

	// 测试连接并使用指数退避重试
	var lastErr error
	backoff := cfg.Pool.RetryBackoff
	for i := 0; i < cfg.Pool.MaxConnAttempts; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout.Dial)
		_, err := client.Ping(ctx).Result()
		cancel()

		if err == nil {
			mu.Lock()
			redisClient = client
			mu.Unlock()
			return nil
		}
		lastErr = err
		sleepTime := time.Duration(1<<uint(i)) * backoff
		if sleepTime > 30*time.Second {
			sleepTime = 30 * time.Second
		}
		time.Sleep(sleepTime)
	}
	return fmt.Errorf("redis connection failed, retried %d times: %v", cfg.Pool.MaxConnAttempts, lastErr)
}

// 关闭Redis连接池
func GracefulClose(ctx context.Context) error {
	mu.Lock()
	defer mu.Unlock()

	if redisClient != nil {
		if err := redisClient.Close(); err != nil {
			return fmt.Errorf("error closing redis client: %v", err)
		}
		redisClient = nil
	}
	return nil
}
