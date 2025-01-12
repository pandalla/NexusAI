package log

import (
	"nexus-ai/common"
	"nexus-ai/utils"

	"gorm.io/gorm"
)

// API转发日志，包括请求转发、响应处理等信息
type RelayLog struct {
	RelayLogID      string `gorm:"column:relay_log_id;type:char(36);primaryKey;default:(UUID())" json:"relay_log_id"`
	UserID          string `gorm:"column:user_id;type:char(36);index;not null;foreignKey:User(UserID)" json:"user_id"`                                       // 关联的用户ID
	ChannelID       string `gorm:"column:channel_id;type:char(36);index;not null;foreignKey:Channel(ChannelID)" json:"channel_id"`                           // 关联的渠道ID
	ModelID         string `gorm:"column:model_id;type:char(36);index;not null;foreignKey:Model(ModelID)" json:"model_id"`                                   // 关联的模型ID
	TokenID         string `gorm:"column:token_id;type:char(36);index;not null;foreignKey:Token(TokenID)" json:"token_id"`                                   // 关联的令牌ID
	WorkerClusterID string `gorm:"column:worker_cluster_id;type:char(36);index;not null;foreignKey:WorkerCluster(WorkerClusterID)" json:"worker_cluster_id"` // 关联的集群ID
	WorkerGroupID   string `gorm:"column:worker_group_id;type:char(36);index;not null;foreignKey:WorkerGroup(WorkerGroupID)" json:"worker_group_id"`         // 关联的工作节点组ID
	WorkerNodeID    string `gorm:"column:worker_node_id;type:char(36);index;not null;foreignKey:WorkerNode(WorkerNodeID)" json:"worker_node_id"`             // 关联的节点实例ID
	RequestID       string `gorm:"column:request_id;size:64;index;not null" json:"request_id"`                                                               // 关联的请求ID

	RelayStatus       int8        `gorm:"column:relay_status;index;not null;default:1" json:"relay_status"`   // 转发状态(1:成功 2:失败 3:超时 4:限流 5:上游异常)
	UpstreamRequestID string      `gorm:"column:upstream_request_id;size:64" json:"upstream_request_id"`      // 上游请求ID
	UpstreamStatus    int         `gorm:"column:upstream_status;not null;default:0" json:"upstream_status"`   // 上游返回状态码
	UpstreamLatency   int         `gorm:"column:upstream_latency;not null;default:0" json:"upstream_latency"` // 上游服务延迟(ms)
	TotalLatency      int         `gorm:"column:total_latency;not null;default:0" json:"total_latency"`       // 总延迟时间(ms)
	RetryCount        int         `gorm:"column:retry_count;not null;default:0" json:"retry_count"`           // 重试次数
	RetryResult       common.JSON `gorm:"column:retry_result;type:json" json:"retry_result"`                  // 重试结果记录
	ErrorType         string      `gorm:"column:error_type;size:50;index" json:"error_type"`                  // 错误类型(rate_limit/server_error/timeout/network等)
	ErrorMessage      string      `gorm:"column:error_message;type:text" json:"error_message"`                // 错误详细信息

	RequestHeaders  common.JSON `gorm:"column:request_headers;type:json" json:"request_headers"`                                  // 请求头信息
	RequestBody     common.JSON `gorm:"column:request_body;type:json" json:"request_body"`                                        // 请求体信息
	ResponseHeaders common.JSON `gorm:"column:response_headers;type:json" json:"response_headers"`                                // 响应头信息
	ResponseBody    common.JSON `gorm:"column:response_body;type:json" json:"response_body"`                                      // 响应体信息
	RequestTokens   int         `gorm:"column:request_tokens;not null;default:0" json:"request_tokens"`                           // 请求token数量
	ResponseTokens  int         `gorm:"column:response_tokens;not null;default:0" json:"response_tokens"`                         // 响应token数量
	QuotaConsumed   float64     `gorm:"column:quota_consumed;type:decimal(10,6);not null;default:0.000000" json:"quota_consumed"` // 消耗的配额数量

	CreatedAt utils.MySQLTime `gorm:"column:created_at;index;not null" json:"created_at"` // 记录创建时间
	UpdatedAt utils.MySQLTime `gorm:"column:updated_at;not null" json:"updated_at"`       // 记录更新时间
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName 表名
func (RelayLog) TableName() string {
	return "relay_logs"
}

// BeforeCreate 在创建记录前自动设置时间
func (relayLog *RelayLog) BeforeCreate(tx *gorm.DB) error {
	relayLog.CreatedAt = utils.MySQLTime(utils.GetTime())
	relayLog.UpdatedAt = utils.MySQLTime(utils.GetTime())
	return nil
}

// BeforeUpdate 在更新记录前自动设置更新时间
func (relayLog *RelayLog) BeforeUpdate(tx *gorm.DB) error {
	relayLog.UpdatedAt = utils.MySQLTime(utils.GetTime())
	return nil
}
