package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nexus-ai/controller"
	"nexus-ai/model"
	"nexus-ai/repository"
)

func SetupUserRouter(server *gin.Engine) {

	userRepo := repository.NewUserRepository(model.GetDB())
	userController := controller.NewUserController(userRepo)
	userRouter := server.Group("/user")
	{
		userRouter.POST("/login", userController.Login)
		userRouter.POST("/register", userController.Register)
	}
	userRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is user router!"})
	})

}
