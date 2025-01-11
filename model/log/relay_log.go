package log

import (
	"time"

	"nexus-ai/common"
)

// RelayLog 中继日志表
type RelayLog struct {
	RelayID           uint64      `gorm:"column:relay_id;primaryKey;autoIncrement" json:"relay_id"`
	UserID            uint64      `gorm:"column:user_id;index;not null" json:"user_id"`
	ChannelID         uint64      `gorm:"column:channel_id;index;not null" json:"channel_id"`
	ModelID           uint64      `gorm:"column:model_id;index;not null" json:"model_id"`
	UpstreamURL       string      `gorm:"column:upstream_url;size:255;not null" json:"upstream_url"`
	RelayStatus       int8        `gorm:"column:relay_status;index;not null" json:"relay_status"`
	UpstreamStatus    int         `gorm:"column:upstream_status" json:"upstream_status"`
	ErrorType         string      `gorm:"column:error_type;size:50;index" json:"error_type"`
	ErrorMessage      string      `gorm:"column:error_message;type:text" json:"error_message"`
	RetryCount        int         `gorm:"column:retry_count" json:"retry_count"`
	RetryResult       common.JSON `gorm:"column:retry_result;type:json" json:"retry_result"`
	RequestHeaders    common.JSON `gorm:"column:request_headers;type:json" json:"request_headers"`
	RequestBody       common.JSON `gorm:"column:request_body;type:json" json:"request_body"`
	ResponseHeaders   common.JSON `gorm:"column:response_headers;type:json" json:"response_headers"`
	ResponseBody      common.JSON `gorm:"column:response_body;type:json" json:"response_body"`
	UpstreamLatency   int         `gorm:"column:upstream_latency" json:"upstream_latency"`
	TotalLatency      int         `gorm:"column:total_latency" json:"total_latency"`
	RequestTokens     int         `gorm:"column:request_tokens" json:"request_tokens"`
	ResponseTokens    int         `gorm:"column:response_tokens" json:"response_tokens"`
	QuotaConsumed     float64     `gorm:"column:quota_consumed;type:decimal(10,6)" json:"quota_consumed"`
	UpstreamRequestID string      `gorm:"column:upstream_request_id;size:64" json:"upstream_request_id"`
	RequestID         string      `gorm:"column:request_id;size:64;index;not null" json:"request_id"`
	CreatedAt         time.Time   `gorm:"column:created_at;index;not null" json:"created_at"`
}

// TableName 表名
func (RelayLog) TableName() string {
	return "relay_logs"
}
