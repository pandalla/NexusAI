package model

import (
	"nexus-ai/utils"
)

// ModelPrice 模型价格配置
type ModelPrice struct {
	RequestPrice    float64 `json:"request_price"`    // 请求价格
	CompletionPrice float64 `json:"completion_price"` // 补全价格
	ResponsePrice   float64 `json:"response_price"`   // 响应价格
}

// ModelAlias 模型映射配置
type ModelAlias struct {
	ProviderAlias []string `json:"provider_alias"` // 提供商别名列表
	DisplayName   string   `json:"display_name"`   // 显示名称
	ShortName     string   `json:"short_name"`     // 短名称
}

// ModelOptions 模型配置选项
type ModelOptions struct {
	MaxTokens         int     `json:"max_tokens"`          // 最大token数
	Temperature       float64 `json:"temperature"`         // 温度
	TopP              float64 `json:"top_p"`               // Top P
	FrequencyPenalty  float64 `json:"frequency_penalty"`   // 频率惩罚
	PresencePenalty   float64 `json:"presence_penalty"`    // 存在惩罚
	MaxContextLength  int     `json:"max_context_length"`  // 最大上下文长度
	DefaultSystemRole string  `json:"default_system_role"` // 默认系统角色
	DefaultUserRole   string  `json:"default_user_role"`   // 默认用户角色
	DefaultAssistRole string  `json:"default_assist_role"` // 默认助手角色
	RequireAuth       bool    `json:"require_auth"`        // 是否需要认证
	DisableRateLimit  bool    `json:"disable_rate_limit"`  // 是否禁用频率限制
}

// Model DTO结构
type Model struct {
	ModelID          string           `json:"model_id"`             // 模型ID
	ModelGroupID     string           `json:"model_group_id"`       // 模型组ID
	ModelName        string           `json:"model_name"`           // 模型名称
	ModelDescription string           `json:"model_description"`    // 模型描述
	ModelType        string           `json:"model_type"`           // 模型类型
	Provider         string           `json:"provider"`             // 模型提供商
	PriceType        string           `json:"price_type"`           // 计费类型
	ModelPrice       ModelPrice       `json:"model_price"`          // 模型价格配置
	Status           int8             `json:"status"`               // 模型状态
	ModelAlias       ModelAlias       `json:"model_alias"`          // 模型映射
	ModelOptions     ModelOptions     `json:"model_options"`        // 模型配置
	CreatedAt        utils.MySQLTime  `json:"created_at"`           // 创建时间
	UpdatedAt        utils.MySQLTime  `json:"updated_at"`           // 更新时间
	DeletedAt        *utils.MySQLTime `json:"deleted_at,omitempty"` // 删除时间
}
