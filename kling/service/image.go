package service

import (
	"context"
	"fmt"
	"net/url"

	"nexus-ai/kling/api"
	"nexus-ai/kling/models"
)

// ImageService 图像服务接口定义
type ImageService interface {
	CreateImageTask(ctx context.Context, req models.TextToImageRequest) (*models.CreateTaskResponse, error)
	GetImageTask(ctx context.Context, taskID string) (*models.QueryTaskResponse, error)
	ListImageTasks(ctx context.Context, pageNum, pageSize int) (*models.TaskListResponse, error)
}

type imageService struct {
	apiClient *api.Client
}

var _ ImageService = (*imageService)(nil)

// NewImageService 图像服务工厂函数
func NewImageService(apiClient *api.Client) ImageService {
	return &imageService{
		apiClient: apiClient,
	}
}

// CreateImageTask 创建图像任务实现
func (s *imageService) CreateImageTask(
	ctx context.Context,
	req models.TextToImageRequest,
) (*models.CreateTaskResponse, error) {
	apiReq, err := s.apiClient.CreateRequest(ctx, "POST", "/v1/images/generations", req)
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

// GetImageTask 查询图像任务实现
func (s *imageService) GetImageTask(
	ctx context.Context,
	taskID string,
) (*models.QueryTaskResponse, error) {
	path := buildImageTaskPath(taskID)

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

// ListImageTasks 查询图像任务列表实现
func (s *imageService) ListImageTasks(
	ctx context.Context,
	pageNum, pageSize int,
) (*models.TaskListResponse, error) {
	path := buildImageTaskListPath(pageNum, pageSize)

	apiReq, err := s.apiClient.CreateRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create API request: %w", err)
	}

	var resp models.TaskListResponse
	if err := s.apiClient.DoRequest(apiReq, &resp); err != nil {
		return nil, wrapAPIError(err)
	}

	if !resp.IsSuccess() {
		return nil, NewServiceError(resp.Code, resp.Message)
	}

	return &resp, nil
}

/************************ 内部辅助函数 ************************/

// buildImageTaskPath 构建图像任务路径
func buildImageTaskPath(taskID string) string {
	return fmt.Sprintf("/v1/images/generations/%s", url.PathEscape(taskID))
}

// buildImageTaskListPath 构建任务列表路径
func buildImageTaskListPath(pageNum, pageSize int) string {
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 || pageSize > 500 {
		pageSize = 30
	}
	return fmt.Sprintf("/v1/images/generations?pageNum=%d&pageSize=%d", pageNum, pageSize)
}
