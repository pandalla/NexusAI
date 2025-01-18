package log

import (
	"nexus-ai/common"
	"nexus-ai/utils"
)

type RequestLogDetails struct {
	Content string `json:"content"`
}

// RequestLog DTO结构
type RequestLog struct {
	RequestLogID string `json:"request_log_id"` // 日志唯一标识
	UserID       string `json:"user_id"`        // 用户ID
	ChannelID    string `json:"channel_id"`     // 渠道ID
	ModelID      string `json:"model_id"`       // 模型ID
	TokenID      string `json:"token_id"`       // 令牌ID
	MasterID     string `json:"master_id"`      // 主服务节点ID

	RequestID      string          `json:"request_id"`      // 请求ID
	RequestType    string          `json:"request_type"`    // 请求类型
	RequestPath    string          `json:"request_path"`    // 请求路径
	RequestMethod  string          `json:"request_method"`  // 请求方法
	RequestHeaders common.JSON     `json:"request_headers"` // 请求头信息
	RequestParams  common.JSON     `json:"request_params"`  // 请求参数
	RequestTokens  int             `json:"request_tokens"`  // 请求内容token数
	RequestTime    utils.MySQLTime `json:"request_time"`    // 请求开始时间

	RequestStatus   int8            `json:"request_status"`   // 请求结果状态(1:成功 2:失败 3:超时)
	ResponseTime    utils.MySQLTime `json:"response_time"`    // 响应结束时间
	TotalTime       int             `json:"total_time"`       // 总耗时(毫秒)
	ResponseCode    int             `json:"response_code"`    // 响应状态码
	ResponseHeaders common.JSON     `json:"response_headers"` // 响应头信息
	ErrorMessage    string          `json:"error_message"`    // 错误信息
	ClientIP        string          `json:"client_ip"`        // 客户端IP

	EventType  string            `json:"event_type"`  // 事件类型(request/response)
	LogLevel   string            `json:"log_level"`   // 日志级别(info/warn/error)
	LogDetails RequestLogDetails `json:"log_details"` // 详细日志信息

	CreatedAt utils.MySQLTime  `json:"created_at"`           // 记录创建时间
	UpdatedAt utils.MySQLTime  `json:"updated_at"`           // 记录更新时间
	DeletedAt *utils.MySQLTime `json:"deleted_at,omitempty"` // 删除时间
}
