package constant

type UserIDKeyType string    // 基于string的UserIDKeyType，可以避免与其他库的冲突
type RequestIDKeyType string // 基于string的RequestIDKeyType，可以避免与其他库的冲突

const (
	FrontendPort = "11000"
	BackendPort  = "10000"
	LogMaxCount  = 100000000
	LogDir       = "./logs"
	GitRepoURL   = "https://github.com/pandalla/NexusAI.git"
	UserIDKey    = UserIDKeyType("X-Nexus-AI-User-ID")
	RequestIDKey = RequestIDKeyType("X-Nexus-AI-Request-ID")
	KeyCharset   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)
