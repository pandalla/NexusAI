package router

import (
	"net/http"

	"nexus-ai/kling/controller"
)

func NewRouter(videoCtrl *controller.VideoController) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/videos", videoCtrl.CreateVideoTask)
	mux.HandleFunc("GET /v1/videos/{taskId}", videoCtrl.GetVideoTask)

	return mux
}
