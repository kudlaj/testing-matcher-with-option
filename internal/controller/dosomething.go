package controller

import (
	"testingmatcher-with-option/internal/model"
)

func (r *Resolver) DoSomething() bool {
	res := r.service.MakeARequest(model.MyRequest{
		Id:       "123",
		Name:     "name",
		Datetime: "2020-01-01 00:00:00",
		Value1:   1,
		Value2:   2,
		Value3:   3,
		Value4:   4,
		Value5:   5,
		Value6:   6,
	})
	return res
}
