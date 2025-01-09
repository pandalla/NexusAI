package utils

import (
	"fmt"
	"time"
)

func GetTimeStamp() int64 { // 获取当前时间戳
	return time.Now().Unix()
}

func GetTimeString() string { // 获取当前时间字符串
	return fmt.Sprintf("%s%d", time.Now().Format("20060102150405"), time.Now().UnixNano()%1e9)
}
