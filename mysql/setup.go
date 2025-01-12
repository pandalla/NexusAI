package mysql

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
	if err := InitMySQL(cfg); err != nil {
		return fmt.Errorf("failed to initialize mysql: %v", err)
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
					utils.SysError(fmt.Sprintf("MySQL | health check failed (%d/%d): %v",
						consecutiveFailures, maxConsecutiveFailures, err))

					if consecutiveFailures >= maxConsecutiveFailures {
						utils.SysError("MySQL | attempting to reconnect...")
						if err := InitMySQL(cfg); err != nil {
							utils.SysError(fmt.Sprintf("MySQL | reconnection failed: %v", err))
						} else {
							utils.SysInfo("MySQL | reconnection successful")
							consecutiveFailures = 0
						}
					}
				} else {
					if consecutiveFailures > 0 {
						utils.SysInfo("MySQL | health check recovered")
					} else {
						utils.SysInfo("MySQL | health check passed")
					}
					consecutiveFailures = 0
				}
				cancel()
			}
		}
	}()

	return nil
}

// Shutdown 关闭MySQL连接
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
