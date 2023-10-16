package handler

import "github.com/ymoldabe/Doodocs-Backend-Challenge/internal/service"

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}
