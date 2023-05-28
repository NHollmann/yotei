package controller

import (
	"github.com/NHollmann/yotei/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Yotei API
// @BasePath /api/v1
// @version v1
// @Accept json
// @Produce json
func (s *YoteiServer) initRoutes() {
	docs.SwaggerInfo.Title = "Yotei API"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Version = "v1"
	api := s.router.Group("/api/v1")
	{
		api.POST("/login", s.handleLogin)   // Login
		api.POST("/logout", s.handleLogout) // Logout

		api.GET("/user", s.handleUserList)              // Get all
		api.POST("/user", s.handleUserCreate)           // Create
		api.GET("/user/:userId", s.handleUserGet)       // Get one
		api.PUT("/user/:userId", s.handleUserUpdate)    // Update
		api.DELETE("/user/:userId", s.handleUserDelete) // Delete

		api.GET("/event", s.handleEventList)                           // Get all
		api.POST("/event", s.handleEventCreate)                        // Create
		api.GET("/event/:accessKey", s.handleEventGet)                 // Get one
		api.POST("/event/:accessKey", s.handleEventCreateParticipant)  // Add participant
		api.PATCH("/event/:accessKey", s.handleEventUpdateParticipant) // Update participant
		api.PUT("/event/:accessKey", s.handleEventUpdate)              // Update
		api.DELETE("/event/:accessKey", s.handleEventDelete)           // Delete
	}

	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
