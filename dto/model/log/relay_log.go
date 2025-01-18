package log

import (
	"nexus-ai/common"
	"nexus-ai/utils"
)

type RelayLogDetails struct {
	Content string `json:"content"`
}

// RelayLog DTO结构
type RelayLog struct {
	RelayLogID      string `json:"relay_log_id"`      // 日志唯一标识
	UserID          string `json:"user_id"`           // 关联的用户ID
	ChannelID       string `json:"channel_id"`        // 关联的渠道ID
	ModelID         string `json:"model_id"`          // 关联的模型ID
	TokenID         string `json:"token_id"`          // 关联的令牌ID
	WorkerClusterID string `json:"worker_cluster_id"` // 关联的集群ID
	WorkerGroupID   string `json:"worker_group_id"`   // 关联的工作节点组ID
	WorkerNodeID    string `json:"worker_node_id"`    // 关联的节点实例ID
	RequestID       string `json:"request_id"`        // 关联的请求ID

	RelayStatus       int8        `json:"relay_status"`        // 转发状态(1:成功 2:失败 3:超时 4:限流 5:上游异常)
	UpstreamRequestID string      `json:"upstream_request_id"` // 上游请求ID
	UpstreamStatus    int         `json:"upstream_status"`     // 上游返回状态码
	UpstreamLatency   int         `json:"upstream_latency"`    // 上游服务延迟(ms)
	TotalLatency      int         `json:"total_latency"`       // 总延迟时间(ms)
	RetryCount        int         `json:"retry_count"`         // 重试次数
	RetryResult       common.JSON `json:"retry_result"`        // 重试结果记录
	ErrorType         string      `json:"error_type"`          // 错误类型(rate_limit/server_error/timeout/network等)
	ErrorMessage      string      `json:"error_message"`       // 错误详细信息

	RequestHeaders  common.JSON `json:"request_headers"`  // 请求头信息
	RequestBody     common.JSON `json:"request_body"`     // 请求体信息
	ResponseHeaders common.JSON `json:"response_headers"` // 响应头信息
	ResponseBody    common.JSON `json:"response_body"`    // 响应体信息
	RequestTokens   int         `json:"request_tokens"`   // 请求token数量
	ResponseTokens  int         `json:"response_tokens"`  // 响应token数量
	QuotaConsumed   float64     `json:"quota_consumed"`   // 消耗的配额数量

	EventType  string          `json:"event_type"`  // 事件类型(request/response)
	LogLevel   string          `json:"log_level"`   // 日志级别(info/warn/error)
	LogDetails RelayLogDetails `json:"log_details"` // 详细日志信息

	CreatedAt utils.MySQLTime  `json:"created_at"`           // 记录创建时间
	UpdatedAt utils.MySQLTime  `json:"updated_at"`           // 记录更新时间
	DeletedAt *utils.MySQLTime `json:"deleted_at,omitempty"` // 删除时间
}
