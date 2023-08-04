package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/signup", h.createUser)
	router.POST("/login", h.loginUser)

	patientAPI := router.Group("/user")

	patientAPI.Use(h.authMiddleware())
	patientAPI.POST("/appointment/create", h.createAppointment)
	patientAPI.POST("/appointment/")

	return router
}
