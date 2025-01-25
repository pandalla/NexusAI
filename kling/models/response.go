package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// UnixTime 自定义时间类型，用于处理毫秒级时间戳
type UnixTime time.Time

// UnmarshalJSON 反序列化毫秒时间戳
func (ut *UnixTime) UnmarshalJSON(data []byte) error {
	var timestamp int64
	if err := json.Unmarshal(data, &timestamp); err != nil {
		return fmt.Errorf("invalid timestamp format: %w", err)
	}
	*ut = UnixTime(time.UnixMilli(timestamp))
	return nil
}

// MarshalJSON 序列化为毫秒时间戳
func (ut UnixTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(ut).UnixMilli())
}

// Time 转换为标准time.Time
func (ut UnixTime) Time() time.Time {
	return time.Time(ut)
}

// 基础响应结构
type BaseResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

// ------------------------- 创建任务响应 -------------------------
type CreateTaskResponse struct {
	BaseResponse
	Data CreateTaskData `json:"data"`
}

type CreateTaskData struct {
	TaskID     string `json:"task_id"`
	TaskStatus string `json:"task_status" validate:"oneof=submitted processing succeed failed"`
}

// ------------------------- 查询任务响应 -------------------------
type QueryTaskResponse struct {
	BaseResponse
	Data QueryTaskData `json:"data"`
}

type QueryTaskData struct {
	TaskID        string      `json:"task_id"`
	TaskStatus    string      `json:"task_status"`
	TaskStatusMsg string      `json:"task_status_msg,omitempty"`
	TaskInfo      TaskInfo    `json:"task_info"`
	TaskResult    *TaskResult `json:"task_result,omitempty"`
	CreatedAt     UnixTime    `json:"created_at"`
	UpdatedAt     UnixTime    `json:"updated_at"`
}

type TaskInfo struct {
	ExternalTaskID string                 `json:"external_task_id"`
	Parameters     map[string]interface{} `json:"parameters,omitempty"`
}

type TaskResult struct {
	Videos []VideoResult `json:"videos"`
}

type VideoResult struct {
	ID       string `json:"id"`
	URL      string `json:"url"`
	Duration string `json:"duration"`
}

// 状态枚举
const (
	TaskStatusSubmitted  = "submitted"
	TaskStatusProcessing = "processing"
	TaskStatusSucceed    = "succeed"
	TaskStatusFailed     = "failed"
)

// 错误码常量
const (
	ErrCodeSuccess           = 0
	ErrCodeAuthFailed        = 1000
	ErrCodeInvalidParams     = 1200
	ErrCodeRateLimit         = 1302
	ErrCodeContentModeration = 1301
)

// Error 实现error接口
func (r *BaseResponse) Error() string {
	return fmt.Sprintf("API Error %d: %s", r.Code, r.Message)
}

// IsSuccess 判断是否成功
func (r *BaseResponse) IsSuccess() bool {
	return r.Code == ErrCodeSuccess
}
