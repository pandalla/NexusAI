package model

import (
	"nexus-ai/model/log"
	"nexus-ai/model/worker"

	"gorm.io/gorm"
)

// Setup 初始化所有数据库表
func Setup(db *gorm.DB) error {
	// 用户相关表
	if err := db.AutoMigrate(
		&Billing{},
		&Channel{},
		&MessageSave{},
		&Model{},
		&Payment{},
		&Quota{},
		&Token{},
		&Usage{},
		&User{},
	); err != nil {
		return err
	}

	// 工作节点相关表
	if err := db.AutoMigrate(
		&worker.WorkerCluster{},
		&worker.WorkerNode{},
		&worker.Worker{},
	); err != nil {
		return err
	}

	// 日志相关表
	if err := db.AutoMigrate(
		&log.GatewayLog{},
		&log.MasterLog{},
		&log.MasterMySQLLog{},
		&log.RedisLog{},
		&log.RedisPersistLog{},
		&log.RelayLog{},
		&log.RequestLog{},
		&log.WorkerLog{},
		&log.WorkerMySQLLog{},
	); err != nil {
		return err
	}

	return nil
}
