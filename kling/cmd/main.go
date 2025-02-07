package main

import (
	"fmt"
	"log"
	"net/http"
	"nexus-ai/kling/api"

	"nexus-ai/kling/config"
	"nexus-ai/kling/controller"
	"nexus-ai/kling/router"
	"nexus-ai/kling/service"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("kling/config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化依赖
	apiClient := api.NewClient(&cfg.KlingAI)
	videoService := service.NewVideoService(apiClient)
	controllerService := service.NewImageService(apiClient)

	// 注入控制器（自动类型匹配）
	videoCtrl := controller.NewVideoController(videoService)
	imageCtrl := controller.NewImageController(controllerService)

	// 创建路由
	router := router.NewRouter(videoCtrl, imageCtrl)

	// 启动服务器
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}

	log.Printf("Server starting on port %d", cfg.Server.Port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}
