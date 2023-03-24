package service

import "testingmatcher-with-option/internal/model"

type MyService struct {
}

func (s *MyService) MakeARequest(request model.MyRequest) bool {
	return true
}
