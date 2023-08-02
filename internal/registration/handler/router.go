package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/signup")

	doctorAPI := router.Group("/doctor")
	patientAPI := router.Group("/client")

	doctorAPI.GET("/list")
	doctorAPI.POST("/appointment/")

	patientAPI.POST("/orders")
	patientAPI.POST("/order/")

	return router
}
