package log

import (
	"time"

	"nexus-ai/common"
)

// RequestLog 存储API请求日志，包括请求处理、响应结果等信息
type RequestLog struct {
	// 请求唯一标识
	RequestLogID string `gorm:"column:request_log_id;type:char(36);primaryKey;default:(UUID())" json:"request_log_id"`

	// 关联的用户ID
	UserID string `gorm:"column:user_id;type:char(36);index;not null;foreignKey:User(UserID)" json:"user_id"`

	// 关联的令牌ID
	TokenID string `gorm:"column:token_id;type:char(36);index;not null;foreignKey:Token(TokenID)" json:"token_id"`

	// 关联的渠道ID
	ChannelID string `gorm:"column:channel_id;type:char(36);index;not null;foreignKey:Channel(ChannelID)" json:"channel_id"`

	// 关联的模型ID
	ModelID string `gorm:"column:model_id;type:char(36);index;not null;foreignKey:Model(ModelID)" json:"model_id"`

	// 关联的主服务节点ID
	MasterID string `gorm:"column:master_id;type:char(36);index;foreignKey:Worker(WorkerID)" json:"master_id"`

	// 请求类型
	RequestType string `gorm:"column:request_type;size:50;index;not null" json:"request_type"`

	// 请求路径
	RequestPath string `gorm:"column:request_path;size:255;not null" json:"request_path"`

	// 请求方法(GET/POST等)
	RequestMethod string `gorm:"column:request_method;size:10;not null" json:"request_method"`

	// 请求头信息
	RequestHeaders common.JSON `gorm:"column:request_headers;type:json" json:"request_headers"`

	// 请求参数
	RequestParams common.JSON `gorm:"column:request_params;type:json" json:"request_params"`

	// 请求内容token数
	RequestTokens int `gorm:"column:request_tokens;not null;default:0" json:"request_tokens"`

	// 请求结果状态(1:成功 2:失败 3:超时)
	RequestStatus int8 `gorm:"column:request_status;index;not null;default:1" json:"request_status"`

	// 请求开始时间
	RequestTime time.Time `gorm:"column:request_time;not null" json:"request_time"`

	// 响应结束时间
	ResponseTime time.Time `gorm:"column:response_time;not null" json:"response_time"`

	// 总耗时(毫秒)
	TotalTime int `gorm:"column:total_time;not null;default:0" json:"total_time"`

	// 响应状态码
	ResponseCode int `gorm:"column:response_code;index;not null;default:0" json:"response_code"`

	// 响应头信息
	ResponseHeaders common.JSON `gorm:"column:response_headers;type:json" json:"response_headers"`

	// 错误信息
	ErrorMessage string `gorm:"column:error_message;type:text" json:"error_message"`

	// 客户端IP
	ClientIP string `gorm:"column:client_ip;size:50;not null" json:"client_ip"`

	// 记录创建时间
	CreatedAt time.Time `gorm:"column:created_at;index;not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`

	// 记录更新时间
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"`
}

// TableName 表名
func (RequestLog) TableName() string {
	return "request_logs"
}
