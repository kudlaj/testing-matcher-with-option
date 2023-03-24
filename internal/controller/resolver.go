package controller

import "testingmatcher-with-option/internal/service"

type Resolver struct {
	service service.Service
}

func NewResolver(service service.Service) *Resolver {
	return &Resolver{
		service: service,
	}
}
