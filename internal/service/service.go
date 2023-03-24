package service

import "testingmatcher-with-option/internal/model"

type Service interface {
	MakeARequest(model.MyRequest) bool
}

func GetServiceImplemetation() Service {
	return &MyService{}
}
