package redis

import (
	"context"
	"fmt"
	"nexus-ai/utils"
	"time"
)

var stopChan chan struct{}

func Setup() error {
	cfg := DefaultConfig()
	if err := InitRedis(cfg); err != nil {
		return fmt.Errorf("failed to initialize redis: %v", err)
	}

	stopChan = make(chan struct{})
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-stopChan:
				return
			case <-ticker.C:
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				if err := HealthCheck(ctx); err != nil {
					utils.SysError("Redis | " + err.Error())
				}
				cancel()
			}
		}
	}()

	return nil
}

// Shutdown 关闭Redis连接
func Shutdown() error {
	// 停止健康检查
	if stopChan != nil {
		close(stopChan)
	}

	// 创建带超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return GracefulClose(ctx)
}
