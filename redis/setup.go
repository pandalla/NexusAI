package redis

import (
	"context"
	"fmt"
	"nexus-ai/utils"
	"time"
)

var (
	stopChan               chan struct{}
	consecutiveFailures    int
	maxConsecutiveFailures = 3
)

func Setup() error {
	cfg := DefaultConfig()
	if err := InitRedis(cfg); err != nil {
		return fmt.Errorf("failed to initialize redis: %v", err)
	}

	stopChan = make(chan struct{})
	go func() {
		ticker := time.NewTicker(20 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-stopChan:
				return
			case <-ticker.C:
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				if err := HealthCheck(ctx); err != nil {
					consecutiveFailures++
					utils.SysError(fmt.Sprintf("Redis | health check failed (%d/%d): %v",
						consecutiveFailures, maxConsecutiveFailures, err))

					if consecutiveFailures >= maxConsecutiveFailures {
						utils.SysError("Redis | attempting to reconnect...")
						if err := InitRedis(cfg); err != nil {
							utils.SysError(fmt.Sprintf("Redis | reconnection failed: %v", err))
						} else {
							utils.SysInfo("Redis | reconnection successful")
							consecutiveFailures = 0
						}
					}
				} else {
					if consecutiveFailures > 0 {
						utils.SysInfo("Redis | health check recovered")
					} else {
						utils.SysInfo("Redis | health check passed")
					}
					consecutiveFailures = 0
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
