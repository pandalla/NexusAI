package model

import (
	"fmt"
	"nexus-ai/model/log"
	"nexus-ai/model/worker"
	"nexus-ai/mysql"

	"time"

	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

// 初始化数据库连接和表结构
func InitGorm() error {
	sqlDB := mysql.GetDB()
	if sqlDB == nil {
		return fmt.Errorf("failed to get MySQL connection")
	}

	gormDB, err := gorm.Open(gmysql.New(gmysql.Config{
		Conn:                      sqlDB,
		DefaultStringSize:         256,  // 设置默认字符串长度
		DontSupportRenameIndex:    true, // 重命名索引时采用删除并新建的方式
		DontSupportRenameColumn:   true, // 用 `change` 重命名列，而不是 `rename`
		SkipInitializeWithVersion: true, // 跳过版本检查
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger:               logger.Default.LogMode(logger.Silent), // 添加详细日志
		DisableAutomaticPing: true,                                  // 禁用自动 ping
	})
	if err != nil {
		return fmt.Errorf("failed to initialize GORM: %v", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	db = gormDB

	// 初始化数据库表
	if err := Setup(gormDB); err != nil {
		return fmt.Errorf("failed to setup database tables: %v", err)
	}

	return nil
}

// GetDB 获取 GORM 数据库实例
func GetDB() *gorm.DB {
	return db
}

// Setup 初始化所有数据库表
func Setup(db *gorm.DB) error {
	// 先尝试创建数据库（如果不存在）
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci", mysql.GetConfig().Basic.Database)
	if err := db.Exec(sql).Error; err != nil {
		return fmt.Errorf("failed to create database: %v", err)
	}

	// 确保使用正确的数据库
	if err := db.Exec(fmt.Sprintf("USE %s", mysql.GetConfig().Basic.Database)).Error; err != nil {
		return fmt.Errorf("failed to use database: %v", err)
	}

	// 设置表选项
	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci")

	// 分批执行迁移，每个表单独处理
	tables := []interface{}{
		// 基础表
		&User{},
		&Model{},
		&Channel{},
		&Token{},
	}

	// 先尝试创建基础表
	for _, table := range tables {
		if err := db.AutoMigrate(table); err != nil {
			return fmt.Errorf("failed to migrate %T: %v", table, err)
		}
	}

	// 其他表
	otherTables := []interface{}{
		&worker.WorkerCluster{},
		&worker.Worker{},
		&worker.WorkerNode{},
		&log.GatewayLog{},
		&log.MasterLog{},
		&log.MasterMySQLLog{},
		&log.RedisLog{},
		&log.RedisPersistLog{},
		&log.RelayLog{},
		&log.RequestLog{},
		&log.WorkerLog{},
		&log.WorkerMySQLLog{},
		&Usage{},
		&Quota{},
		&Payment{},
		&MessageSave{},
		&Billing{},
	}

	// 分批创建其他表
	for _, table := range otherTables {
		if err := db.AutoMigrate(table); err != nil {
			return fmt.Errorf("failed to migrate %T: %v", table, err)
		}
	}

	return nil
}
