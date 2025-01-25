package service

import (
	"context"
	"fmt"
	"net/http"

	"nexus-ai/kling/api"
	"nexus-ai/kling/models"
)

// VideoService 视频服务接口定义
type VideoService interface {
	CreateVideoTask(ctx context.Context, req models.TextToVideoRequest) (*models.CreateTaskResponse, error)
	GetVideoTask(ctx context.Context, taskID, externalTaskID string) (*models.QueryTaskResponse, error)
}

// 私有实现结构体（小写开头表示不可导出）
type videoServiceImpl struct {
	apiClient *api.Client
}

// 编译时接口实现验证（确保实现所有接口方法）
var _ VideoService = (*videoServiceImpl)(nil)

// NewVideoService 视频服务工厂函数（返回接口类型）
func NewVideoService(apiClient *api.Client) VideoService {
	return &videoServiceImpl{
		apiClient: apiClient,
	}
}

// CreateVideoTask 创建视频任务实现
func (s *videoServiceImpl) CreateVideoTask(
	ctx context.Context,
	req models.TextToVideoRequest,
) (*models.CreateTaskResponse, error) {
	// 调用API客户端
	apiReq, err := s.apiClient.CreateRequest(ctx, "POST", "/v1/videos/text2video", req)
	if err != nil {
		return nil, fmt.Errorf("failed to create API request: %w", err)
	}

	var resp models.CreateTaskResponse
	if err := s.apiClient.DoRequest(apiReq, &resp); err != nil {
		return nil, wrapAPIError(err)
	}

	if !resp.IsSuccess() {
		return nil, NewServiceError(resp.Code, resp.Message)
	}

	return &resp, nil
}

// GetVideoTask 查询视频任务实现
func (s *videoServiceImpl) GetVideoTask(
	ctx context.Context,
	taskID,
	externalTaskID string,
) (*models.QueryTaskResponse, error) {
	// 构建请求路径
	path := buildTaskPath(taskID, externalTaskID)

	apiReq, err := s.apiClient.CreateRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create API request: %w", err)
	}

	var resp models.QueryTaskResponse
	if err := s.apiClient.DoRequest(apiReq, &resp); err != nil {
		return nil, wrapAPIError(err)
	}

	if !resp.IsSuccess() {
		return nil, NewServiceError(resp.Code, resp.Message)
	}

	return &resp, nil
}

/************************ 内部辅助函数 ************************/

// buildTaskPath 构建任务查询路径
func buildTaskPath(taskID, externalID string) string {
	if externalID != "" {
		return fmt.Sprintf("/v1/videos/text2video?external_task_id=%s", externalID)
	}
	return fmt.Sprintf("/v1/videos/text2video/%s", taskID)
}

// wrapAPIError 包装API错误
func wrapAPIError(err error) error {
	if apiErr, ok := err.(interface {
		GetCode() int
		GetMessage() string
	}); ok {
		return NewServiceError(apiErr.GetCode(), apiErr.GetMessage())
	}
	return fmt.Errorf("unknown API error: %w", err)
}

/************************ 错误处理 ************************/

// ServiceError 自定义服务错误
type ServiceError struct {
	Code    int
	Message string
}

func NewServiceError(code int, message string) *ServiceError {
	return &ServiceError{
		Code:    code,
		Message: message,
	}
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("service error [%d]: %s", e.Code, e.Message)
}

// BusinessErrorMapping 业务错误到HTTP状态码映射
func (e *ServiceError) HTTPStatus() int {
	switch {
	case e.Code >= 1000 && e.Code < 1100:
		return http.StatusUnauthorized
	case e.Code >= 1100 && e.Code < 1200:
		return http.StatusTooManyRequests
	case e.Code >= 1200 && e.Code < 1300:
		return http.StatusBadRequest
	case e.Code >= 1300 && e.Code < 1400:
		return http.StatusTooManyRequests
	default:
		return http.StatusInternalServerError
	}
}
