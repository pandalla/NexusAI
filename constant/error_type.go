package constant

const (
	ErrorTypeInternalServerError     = "nexus_ai_panic_error"               // 内部服务器错误
	ErrorTypeIPVerifyFailed          = "nexus_ai_ip_verify_failed"          // IP验证失败
	ErrorTypeTokenVerifyFailed       = "nexus_ai_token_verify_failed"       // 令牌验证失败
	ErrorTypeUserVerifyFailed        = "nexus_ai_user_verify_failed"        // 用户验证失败
	ErrorTypeQuotaVerifyFailed       = "nexus_ai_quota_verify_failed"       // 配额验证失败
	ErrorTypeModelVerifyFailed       = "nexus_ai_model_verify_failed"       // 模型验证失败
	ErrorTypeChannelDistributeFailed = "nexus_ai_channel_distribute_failed" // 渠道分配失败
)
