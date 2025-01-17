package model

import (
	"nexus-ai/common"
	"nexus-ai/utils"
)

// MessageOptions 消息配置选项
type MessageOptions struct {
	MaxTokens           int      `json:"max_tokens" binding:"required,gte=1"`                   // 最大token数
	MaxCompletionTokens int      `json:"max_completion_tokens" binding:"required,gte=1"`        // 最大补全token数
	Temperature         int      `json:"temperature" binding:"required,gte=0,lte=2"`            // 温度系数
	TopP                int      `json:"top_p" binding:"required,gte=0,lte=1"`                  // 采样阈值
	TopK                int      `json:"top_k" binding:"required,gte=0,lte=100"`                // 采样top-k
	Stop                []string `json:"stop" binding:"omitempty,dive,max=4"`                   // 停止词
	ResponseFormat      string   `json:"response_format" binding:"required,oneof=json text"`    // 响应格式
	EncodingFormat      string   `json:"encoding_format" binding:"required,oneof=utf-8 base64"` // 编码格式
	N                   int      `json:"n" binding:"required,gte=1,lte=10"`                     // 生成数量
	Size                int      `json:"size" binding:"required,gte=1,lte=10"`                  // 生成尺寸
	Seed                int      `json:"seed" binding:"required,gte=0,lte=1000000000"`          // 随机种子
	FunctionCall        int      `json:"function_call" binding:"required,gte=0,lte=1"`          // 是否启用函数调用
	PresencePenalty     int      `json:"presence_penalty" binding:"required,gte=0,lte=2"`       // 存在惩罚
	FrequencyPenalty    int      `json:"frequency_penalty" binding:"required,gte=0,lte=2"`      // 频率惩罚
	BestOf              int      `json:"best_of" binding:"required,gte=1,lte=20"`               // 最佳结果数
	Dimensions          int      `json:"dimensions" binding:"required,gte=1,lte=10"`            // 维度
}

// StreamOptions 流式输出配置
type StreamOptions common.JSON

// MessageExtra 消息额外信息
type MessageExtra common.JSON

// MessageSave DTO结构
type MessageSave struct {
	MessageID string `json:"message_id" binding:"required"` // 消息唯一标识
	RequestID string `json:"request_id" binding:"required"` // 关联的请求ID
	UserID    string `json:"user_id" binding:"required"`    // 关联的用户ID
	TokenID   string `json:"token_id" binding:"required"`   // 关联的令牌ID
	ModelID   string `json:"model_id" binding:"required"`   // 关联的模型ID
	ChannelID string `json:"channel_id" binding:"required"` // 关联的渠道ID

	RequestTitle     string         `json:"request_title" binding:"required"`    // 请求标题
	RequestContent   string         `json:"request_content" binding:"required"`  // 请求内容
	ResponseContent  string         `json:"response_content" binding:"required"` // 响应内容
	Input            string         `json:"input" binding:"required"`            // 输入
	Instruction      string         `json:"instruction" binding:"required"`      // 指令
	Prompt           string         `json:"prompt" binding:"required"`           // 提示词
	Stream           int8           `json:"stream" binding:"oneof=0 1"`          // 是否流式输出
	StreamOptions    StreamOptions  `json:"stream_options" binding:"required"`   // 流式输出配置
	RequestTokens    int            `json:"request_tokens" binding:"gte=0"`      // 请求token数
	PromptTokens     int            `json:"prompt_tokens" binding:"gte=0"`       // 提示词token数
	CompletionTokens int            `json:"completion_tokens" binding:"gte=0"`   // 补全token数
	Latency          int            `json:"latency" binding:"gte=0"`             // 响应延迟(ms)
	MessageOptions   MessageOptions `json:"message_options" binding:"required"`  // 消息配置
	MessageExtra     MessageExtra   `json:"message_extra" binding:"required"`    // 消息额外信息

	MessageStatus int8   `json:"message_status" binding:"oneof=0 1 2"` // 消息状态 0:处理中 1:成功 2:失败
	RetryCount    int    `json:"retry_count" binding:"gte=0,lte=5"`    // 重试次数
	ErrorType     string `json:"error_type"`                           // 错误类型
	ErrorInfo     string `json:"error_info"`                           // 错误信息

	CreatedAt utils.MySQLTime  `json:"created_at"`           // 创建时间
	UpdatedAt utils.MySQLTime  `json:"updated_at"`           // 更新时间
	DeletedAt *utils.MySQLTime `json:"deleted_at,omitempty"` // 删除时间
}
