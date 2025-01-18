package model

import (
	"nexus-ai/utils"
)

// Usage DTO结构
type Usage struct {
	UsageID          string           `json:"usage_id"`             // 用量记录唯一标识
	TokenID          string           `json:"token_id"`             // 关联的令牌ID
	UserID           string           `json:"user_id"`              // 关联的用户ID
	ChannelID        string           `json:"channel_id"`           // 关联的渠道ID
	ModelID          string           `json:"model_id"`             // 关联的模型ID
	RequestID        string           `json:"request_id"`           // 关联的请求ID
	UsageType        string           `json:"usage_type"`           // 计费类型
	UnitPrice        float64          `json:"unit_price"`           // 单价
	TimesCount       int              `json:"times_count"`          // 请求次数
	TokensCount      int              `json:"tokens_count"`         // Token数量
	PriceTotalFactor float64          `json:"price_total_factor"`   // 价格总倍率
	TotalAmount      float64          `json:"total_amount"`         // 总金额
	CreatedAt        utils.MySQLTime  `json:"created_at"`           // 创建时间
	UpdatedAt        utils.MySQLTime  `json:"updated_at"`           // 更新时间
	DeletedAt        *utils.MySQLTime `json:"deleted_at,omitempty"` // 软删除时间
}
