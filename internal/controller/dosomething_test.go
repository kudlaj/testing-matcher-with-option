package controller_test

import (
	"testing"
	"testingmatcher-with-option/internal/controller"
	"testingmatcher-with-option/internal/model"
	"testingmatcher-with-option/internal/service"
	"testingmatcher-with-option/internal/testutils"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestDoSomething(t *testing.T) {
	// Running test with any
	t.Run("should execute DoSomething call and MakeARequest return true", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockService := service.NewMockService(ctrl)
		mockService.EXPECT().MakeARequest(gomock.Any()).Return(true)
		resolver := controller.NewResolver(mockService)
		res := resolver.DoSomething()
		assert.Assert(t, res, true)
	})

	t.Run("should execute DoSomething call and MakeARequest and check Name and Id with testutils.CreateMyRequestMatcher and return true", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockService := service.NewMockService(ctrl)
		mockService.EXPECT().MakeARequest(testutils.CreateMyMatcher(model.MyRequest{},
			testutils.WithValue(model.MyRequest{}, "Name", "name1"),
			testutils.WithValue(model.MyRequest{}, "Id", "123"),
		)).Return(true)
		resolver := controller.NewResolver(mockService)
		res := resolver.DoSomething()
		assert.Assert(t, res, true)
	})
}
