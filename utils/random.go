package utils

import (
	"math/rand"
	"nexus-ai/constant"

	"github.com/google/uuid"
)

func GenerateRandomString(length int) string { // 获取随机字符串
	key := make([]byte, length)
	keyCharset := constant.KeyCharset
	for i := range key {
		key[i] = keyCharset[rand.Intn(len(keyCharset))]
	}
	return string(key)
}

// GenerateRandomNumber 生成指定长度的随机数字字符串
func GenerateRandomNumber(n int) string {
	b := make([]byte, n)
	numberCharset := constant.NumberCharset
	for i := range b {
		b[i] = numberCharset[rand.Intn(len(numberCharset))]
	}
	return string(b)
}

// GenerateRandomUUID 生成随机UUID
func GenerateRandomUUID(n int) string {
	return uuid.New().String()[:n]
}
