package handler

import "github.com/ymoldabe/Doodocs-Backend-Challenge/internal/service"

type HandlerType struct {
	service *service.Service
}

func New(service *service.Service) *HandlerType {
	return &HandlerType{
		service: service,
	}
}
