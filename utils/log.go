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

var maxLogCount = constant.Configs["log_max_count"]
var logDir = constant.Configs["log_dir"]

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
		defer func() {
			setupLogLock.Unlock()
			setupLogWorking = false
		}()
		logPath := filepath.Join(logDir, fmt.Sprintf("nexus-ai-%s.log", time.Now().Format("2006-01-02")))
		fd, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal("SetupLog failed to open log file:", err)
		}
		gin.DefaultWriter = io.MultiWriter(os.Stdout, fd)
		gin.DefaultErrorWriter = io.MultiWriter(os.Stderr, fd)
	}
}

func SysLog(s string) {
	t := time.Now().Format("2006-01-02 15:04:05")
	_, _ = fmt.Fprintf(gin.DefaultWriter, "[SYSLOG] %v | %s \n", t, s)
}

func SysError(s string) {
	t := time.Now().Format("2006-01-02 15:04:05")
	_, _ = fmt.Fprintf(gin.DefaultErrorWriter, "[SYSERR] %v | %s \n", t, s)
}

func SysWarn(s string) {
	t := time.Now().Format("2006-01-02 15:04:05")
	_, _ = fmt.Fprintf(gin.DefaultWriter, "[SYSWARN] %v | %s \n", t, s)
}

func UserLog(ctx context.Context, logType string, logContent string) {
	writer := gin.DefaultErrorWriter // 默认输出到错误输出
	if logType == logINFO {          // 日志类型为INFO时，输出到标准输出
		writer = gin.DefaultWriter
	}
	if ctx.Value(Request) != nil {
		
	t := time.Now().Format("2006-01-02 15:04:05")
	_, _ = fmt.Fprintf(writer, "[%s] %v | %s \n", logType, t, logContent)
}
