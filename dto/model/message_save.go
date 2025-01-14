package model

import (
	"nexus-ai/utils"
)

// MessageOptions 消息配置选项
type MessageOptions struct {
	Temperature      float64  `json:"temperature"`       // 温度系数
	TopP             float64  `json:"top_p"`             // 采样阈值
	N                int      `json:"n"`                 // 生成数量
	MaxTokens        int      `json:"max_tokens"`        // 最大token数
	PresencePenalty  float64  `json:"presence_penalty"`  // 存在惩罚
	FrequencyPenalty float64  `json:"frequency_penalty"` // 频率惩罚
	Stop             []string `json:"stop"`              // 停止词
	Echo             bool     `json:"echo"`              // 是否回显输入
	Stream           bool     `json:"stream"`            // 是否流式输出
	BestOf           int      `json:"best_of"`           // 最佳结果数
}

// PromptTemplate 提示词模板
type PromptTemplate struct {
	TemplateID      string            `json:"template_id"`      // 模板ID
	TemplateName    string            `json:"template_name"`    // 模板名称
	TemplateContent string            `json:"template_content"` // 模板内容
	Variables       map[string]string `json:"variables"`        // 变量映射
	DefaultValues   map[string]string `json:"default_values"`   // 默认值
}

// MessageExtra 消息额外信息
type MessageExtra struct {
	Source    string            `json:"source"`     // 消息来源
	IP        string            `json:"ip"`         // 请求IP
	UserAgent string            `json:"user_agent"` // 用户代理
	Referer   string            `json:"referer"`    // 来源页面
	TraceID   string            `json:"trace_id"`   // 追踪ID
	SessionID string            `json:"session_id"` // 会话ID
	BatchID   string            `json:"batch_id"`   // 批次ID
	Tags      []string          `json:"tags"`       // 标签列表
	Metadata  map[string]string `json:"metadata"`   // 元数据
}

// MessageSave DTO结构
type MessageSave struct {
	MessageID        string           `json:"message_id"`           // 消息唯一标识
	UserID           string           `json:"user_id"`              // 关联的用户ID
	TokenID          string           `json:"token_id"`             // 关联的令牌ID
	ModelID          string           `json:"model_id"`             // 关联的模型ID
	ChannelID        string           `json:"channel_id"`           // 关联的渠道ID
	RequestID        string           `json:"request_id"`           // 关联的请求ID
	MessageTitle     string           `json:"message_title"`        // 消息标题
	MessageContent   string           `json:"message_content"`      // 消息内容
	MessageTokens    int              `json:"message_tokens"`       // 消息token数
	PromptTokens     int              `json:"prompt_tokens"`        // 提示词token数
	CompletionTokens int              `json:"completion_tokens"`    // 补全token数
	Latency          int              `json:"latency"`              // 响应延迟(ms)
	MessageStatus    int8             `json:"message_status"`       // 消息状态
	RetryCount       int              `json:"retry_count"`          // 重试次数
	ErrorType        string           `json:"error_type"`           // 错误类型
	ErrorInfo        string           `json:"error_info"`           // 错误信息
	MessageOptions   MessageOptions   `json:"message_options"`      // 消息配置
	PromptTemplate   PromptTemplate   `json:"prompt_template"`      // 提示词模板
	MessageExtra     MessageExtra     `json:"message_extra"`        // 消息额外信息
	CreatedAt        utils.MySQLTime  `json:"created_at"`           // 创建时间
	UpdatedAt        utils.MySQLTime  `json:"updated_at"`           // 更新时间
	DeletedAt        *utils.MySQLTime `json:"deleted_at,omitempty"` // 删除时间
}
