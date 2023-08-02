package handler

import "xamss.onelab.final/internal/registration/service"

type Handler struct {
	srvs service.Service
}

func NewHandler(srvs service.Service) *Handler {
	return &Handler{
		srvs: srvs,
	}
}
