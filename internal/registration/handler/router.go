package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/signup", h.createUser)
	router.POST("/login", h.loginUser)

	patientAPI := router.Group("/user")

	patientAPI.Use(h.authMiddleware())
	patientAPI.POST("/appointment/create", h.createAppointment)
	patientAPI.POST("/appointment/")

	return router
}
