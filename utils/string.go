package utils

import (
	"math/rand"
	"nexus-ai/constant"
)

func GetRandomString(length int) string { // 获取随机字符串
	key := make([]byte, length)
	keyCharset := constant.KeyCharset
	for i := range key {
		key[i] = keyCharset[rand.Intn(len(keyCharset))]
	}
	return string(key)
}
