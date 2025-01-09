package utils

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"nexus-ai/constant"

	"github.com/gin-gonic/gin"
)

const (
	logINFO  = "INFO"
	logWARN  = "WARN"
	logERROR = "ERROR"
)

var maxLogCount = constant.LogMaxCount
var logDir = constant.LogDir

var logCount int
var setupLogWorking bool
var setupLogLock sync.Mutex

func SetupLog() {
	if logDir != "" {
		ok := setupLogLock.TryLock()
		if !ok {
			log.Println("SetupLog is working")
			return
		}
		logPath := filepath.Join(logDir, fmt.Sprintf("nexus-ai-%s.log", time.Now().Format("2006-01-02 15-04-05")))
		fd, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal("SetupLog failed to open log file:", err)
		}
		gin.DefaultWriter = io.MultiWriter(os.Stdout, fd)
		gin.DefaultErrorWriter = io.MultiWriter(os.Stderr, fd)
		defer func() {
			setupLogLock.Unlock()
			setupLogWorking = false
			SysInfo("SetupLog Successfully")
		}()
	}
}

func logCommon(ctx context.Context, logType string, logContent string) {
	writer := gin.DefaultErrorWriter // 默认输出到错误输出
	if logType == logINFO {          // 日志类型为INFO时，输出到标准输出
		writer = gin.DefaultWriter
	}
	userID := GetContextValue(ctx.Value(constant.UserIDKey), "test_user_id")
	requestID := GetContextValue(ctx.Value(constant.RequestIDKey), "test_request_id")
	t := time.Now().Format("2006-01-02 15:04:05")
	switch logType {
	case logINFO:
		_, _ = fmt.Fprintf(writer, "[%s]    %v | %s | %s | %s \n", logType, t, userID, requestID, logContent)
	case logWARN:
		_, _ = fmt.Fprintf(writer, "[%s]    %v | %s | %s | %s \n", logType, t, userID, requestID, logContent)
	case logERROR:
		_, _ = fmt.Fprintf(writer, "[%s]   %v | %s | %s | %s \n", logType, t, userID, requestID, logContent)
	}
	logCount++
	if logCount > maxLogCount && !setupLogWorking { // 如果日志数量超过最大值，并且没有正在设置日志，则设置日志
		logCount = 0
		setupLogWorking = true
		go func() {
			SetupLog()
		}()
	}
}

func LogInfo(ctx context.Context, logContent string) { // 信息日志
	logCommon(ctx, logINFO, logContent)
}

func LogWarn(ctx context.Context, logContent string) { // 警告日志
	logCommon(ctx, logWARN, logContent)
}

func LogError(ctx context.Context, logContent string) { // 错误日志
	logCommon(ctx, logERROR, logContent)
}

func FatalLog(s string) { // 致命错误日志 系统退出
	t := time.Now().Format("2006-01-02 15:04:05")
	_, _ = fmt.Fprintf(gin.DefaultWriter, "[SYSLOG]  %v | %s \n", t, s)
	os.Exit(1)
}

func SysInfo(s string) { // 系统信息日志
	t := time.Now().Format("2006-01-02 15:04:05")
	_, _ = fmt.Fprintf(gin.DefaultWriter, "[SYSINFO] %v | %s \n", t, s)
}

func SysError(s string) { // 系统错误日志
	t := time.Now().Format("2006-01-02 15:04:05")
	_, _ = fmt.Fprintf(gin.DefaultErrorWriter, "[SYSERR]  %v | %s \n", t, s)
}

func SysWarn(s string) { // 系统警告日志
	t := time.Now().Format("2006-01-02 15:04:05")
	_, _ = fmt.Fprintf(gin.DefaultWriter, "[SYSWARN] %v | %s \n", t, s)
}
