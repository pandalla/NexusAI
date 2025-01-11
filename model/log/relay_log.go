package log

import (
	"time"

	"nexus-ai/common"
)

// RelayLog 存储API转发日志，包括请求转发、响应处理等信息
type RelayLog struct {
	// 中继日志唯一标识
	RelayLogID string `gorm:"column:relay_log_id;type:char(36);primaryKey;default:(UUID())" json:"relay_log_id"`

	// 关联的用户ID
	UserID string `gorm:"column:user_id;type:char(36);index;not null;foreignKey:User(UserID)" json:"user_id"`

	// 关联的渠道ID
	ChannelID string `gorm:"column:channel_id;type:char(36);index;not null;foreignKey:Channel(ChannelID)" json:"channel_id"`

	// 关联的模型ID
	ModelID string `gorm:"column:model_id;type:char(36);index;not null;foreignKey:Model(ModelID)" json:"model_id"`

	// 关联的集群ID
	ClusterID string `gorm:"column:cluster_id;type:char(36);index;not null;foreignKey:WorkerCluster(ClusterID)" json:"cluster_id"`

	// 关联的工作节点ID
	WorkerID string `gorm:"column:worker_id;type:char(36);index;not null;foreignKey:Worker(WorkerID)" json:"worker_id"`

	// 关联的节点实例ID
	NodeID string `gorm:"column:node_id;type:char(36);index;not null;foreignKey:WorkerNode(NodeID)" json:"node_id"`

	// 转发状态(1:成功 2:失败 3:超时 4:限流 5:上游异常)
	RelayStatus int8 `gorm:"column:relay_status;index;not null;default:1" json:"relay_status"`

	// 上游返回状态码
	UpstreamStatus int `gorm:"column:upstream_status;not null;default:0" json:"upstream_status"`

	// 错误类型(rate_limit/server_error/timeout/network等)
	ErrorType string `gorm:"column:error_type;size:50;index" json:"error_type"`

	// 错误详细信息
	ErrorMessage string `gorm:"column:error_message;type:text" json:"error_message"`

	// 重试次数
	RetryCount int `gorm:"column:retry_count;not null;default:0" json:"retry_count"`

	// 重试结果记录
	RetryResult common.JSON `gorm:"column:retry_result;type:json" json:"retry_result"`

	// 请求头信息
	RequestHeaders common.JSON `gorm:"column:request_headers;type:json" json:"request_headers"`

	// 请求体信息
	RequestBody common.JSON `gorm:"column:request_body;type:json" json:"request_body"`

	// 响应头信息
	ResponseHeaders common.JSON `gorm:"column:response_headers;type:json" json:"response_headers"`

	// 响应体信息
	ResponseBody common.JSON `gorm:"column:response_body;type:json" json:"response_body"`

	// 上游服务延迟(ms)
	UpstreamLatency int `gorm:"column:upstream_latency;not null;default:0" json:"upstream_latency"`

	// 总延迟时间(ms)
	TotalLatency int `gorm:"column:total_latency;not null;default:0" json:"total_latency"`

	// 请求token数量
	RequestTokens int `gorm:"column:request_tokens;not null;default:0" json:"request_tokens"`

	// 响应token数量
	ResponseTokens int `gorm:"column:response_tokens;not null;default:0" json:"response_tokens"`

	// 消耗的配额数量
	QuotaConsumed float64 `gorm:"column:quota_consumed;type:decimal(10,6);not null;default:0.000000" json:"quota_consumed"`

	// 上游请求ID
	UpstreamRequestID string `gorm:"column:upstream_request_id;size:64" json:"upstream_request_id"`

	// 关联的请求ID
	RequestID string `gorm:"column:request_id;size:64;index;not null" json:"request_id"`

	// 记录创建时间
	CreatedAt time.Time `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`

	// 记录更新时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"`
}

// TableName 表名
func (RelayLog) TableName() string {
	return "relay_logs"
}
