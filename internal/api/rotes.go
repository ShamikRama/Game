package api

import (
	"Game/internal/service"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	services *service.Service
}

func NewHandlers(services *service.Service) *Handlers {
	return &Handlers{
		services: services,
	}
}

func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api", h.UserIdentity)
	{
		users := api.Group("/users")
		{
			users.GET("/:id/status", h.Status)
			users.GET("/leaderboard", h.Leaders)
			users.POST("/:id/task/complete", h.CompleteTaskTelegram)
			users.POST("/:id/referrer", h.ReferralCode)
		}
	}

	return router
}
