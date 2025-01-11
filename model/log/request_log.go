package log

import (
	"time"

	"nexus-ai/common"
)

// RequestLog 请求日志表
type RequestLog struct {
	RequestID       string      `gorm:"column:request_id;size:64;primaryKey" json:"request_id"`
	UserID          uint64      `gorm:"column:user_id;index;not null" json:"user_id"`
	TokenID         uint64      `gorm:"column:token_id;index;not null" json:"token_id"`
	ChannelID       uint64      `gorm:"column:channel_id;index;not null" json:"channel_id"`
	ModelID         uint64      `gorm:"column:model_id;index;not null" json:"model_id"`
	WorkerID        uint64      `gorm:"column:worker_id;index" json:"worker_id"`
	RequestType     string      `gorm:"column:request_type;size:50;index;not null" json:"request_type"`
	RequestPath     string      `gorm:"column:request_path;size:255;not null" json:"request_path"`
	RequestMethod   string      `gorm:"column:request_method;size:10;not null" json:"request_method"`
	RequestHeaders  common.JSON `gorm:"column:request_headers;type:json" json:"request_headers"`
	RequestParams   common.JSON `gorm:"column:request_params;type:json" json:"request_params"`
	RequestTokens   int         `gorm:"column:request_tokens" json:"request_tokens"`
	RequestStatus   int8        `gorm:"column:request_status;index;not null" json:"request_status"`
	RequestTime     time.Time   `gorm:"column:request_time;not null" json:"request_time"`
	ProcessTime     time.Time   `gorm:"column:process_time;not null" json:"process_time"`
	ResponseTime    time.Time   `gorm:"column:response_time;not null" json:"response_time"`
	UpstreamTime    int         `gorm:"column:upstream_time" json:"upstream_time"`
	TotalTime       int         `gorm:"column:total_time;not null" json:"total_time"`
	ResponseCode    int         `gorm:"column:response_code;index;not null" json:"response_code"`
	ResponseHeaders common.JSON `gorm:"column:response_headers;type:json" json:"response_headers"`
	ErrorMessage    string      `gorm:"column:error_message;type:text" json:"error_message"`
	ClientIP        string      `gorm:"column:client_ip;size:50;not null" json:"client_ip"`
	CreatedAt       time.Time   `gorm:"column:created_at;index;not null" json:"created_at"`
}

// TableName 表名
func (RequestLog) TableName() string {
	return "request_logs"
}
